package coin

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"
)

type Hash [sha256.Size]byte

func NewHash(hexstr string) (Hash, error) {
	var h Hash
	b, err := hex.DecodeString(hexstr)
	if err != nil {
		return h, fmt.Errorf("hex decode error: %s", err)
	}
	copy(h[:], b)
	return h, nil
}

func (h Hash) String() string {
	return hex.EncodeToString(h[:])
}

// TODO speed up
func (h Hash) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf(`"%s"`, h.String())
	return []byte(s), nil
}

func (h *Hash) UnmarshalJSON(b []byte) error {
	if b[0] != '"' || b[len(b)-1] != '"' {
		return fmt.Errorf("expecting string for hash value")
	}
	n, err := hex.Decode(h[:], b[1:len(b)-1])
	if err != nil {
		return fmt.Errorf("hex decode error: %s", err)
	}
	if n != sha256.Size {
		return fmt.Errorf("short hash value")
	}
	return nil
}

type Block struct {
	PrevHash  Hash
	Contents  string
	Nonce     uint64
	Length    uint32
	Timestamp time.Time
}

func (b *Block) Sum() Hash {
	w := sha256.New()
	w.Write(b.PrevHash[:])
	w.Write([]byte(b.Contents))
	binary.Write(w, binary.BigEndian, b.Nonce)
	binary.Write(w, binary.BigEndian, b.Length)
	sum := w.Sum(nil)
	var h Hash
	copy(h[:], sum[:])
	return h
}

func (b *Block) Verify() (Hash, bool) {
	h := b.Sum()
	d := int(b.Length/100 + 24)
	return h, fastCheck(d, &h)
}

func fastCheck(d int, h *Hash) bool {
	for i := 0; i < d/8; i++ {
		if h[i] != 0 {
			return false
		}
	}
	if m := d % 8; m != 0 {
		if h[d/8]>>uint(8-m) != 0 {
			return false
		}
	}
	return true
}

func slowCheck(d int, h *Hash) bool {
	target := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(256-d)), nil)
	n := new(big.Int).SetBytes(h[:])
	return n.Cmp(target) == -1
}
