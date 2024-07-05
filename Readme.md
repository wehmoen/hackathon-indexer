# AxieGov Data Hackathon Indexer

```
 ------------------------
< AxieGov Data Hackathon >
 ------------------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||

```


This project provides a simple golang cli tool that scans the blockchain for all AXS, WETH and SLP transfers in and out the Axie Treasury.

### Example

```json
  {
    "from": "0xb1b165d6c53a69cc903109247da54045af7c08ab",
    "to": "0x245db945c485b68fdc429e4f7085a1761aa4d45d",
    "tokenAddress": "0x97a9107c1793bc407d6f527b77e7fff4d812bece",
    "transactionHash": "0x29b2600116f30e280f1c6fd8ae6ac97124b04e102880d36d8570ad7a872cc158",
    "value": 288197335793115211,
    "block": 36097704
  }
```

### Prerequisites

You need an API key from https://developer.skymavis.com/ to run this tool as well as access to the "Ronin Archive Node" service. Alternatively, you can use your own RPC URL.

### How to run

First clone the repository. 

### Build on Linux/Mac

- `go build -o hackathon main.go`

### Build on Windows

- `go build -o hackathon.exe main.go`

### Run

- `./hackathon -help or ./hackathon.exe -help`

```
Usage of hackathon.exe:
  -apikey string
        SkyMavis RPC API key. Required if using default RPC URL
  -disable-cow
        Disable ASCII cow
  -rpc string
        RPC URL (default "https://api-gateway.skymavis.com/rpc/archive")
  -start int
        Start block number (default 16377111)
```


### Output 

The tool will output a json file with all the transactions in the Axie Treasury. The filename will be `transfers_STARTBLOCK_ENDBLOCK.json`
