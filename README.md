# 6.857Coin

This is the blockchain server for 6.857Coin, a coin made for the
[6.857](http://courses.csail.mit.edu/6.857/2015/) security class.

The server is (was) running at: http://6857coin.csail.mit.edu

## Usage

1. Install (assuming `GOPATH=~/go`):

        $ go get github.com/davidlazar/6.857coin/...

2. Create required directories:

        $ mkdir logs blocks

3. Create the genesis block:

        $ cat blocks/genesis.block
        {
          "Contents": "Genesis",
          "Nonce": 0,
          "Length": 0
        }

4. Run the blockchain server:

        $ ~/go/bin/coin-server

5. Build a miner using the API described at http://localhost:8080
