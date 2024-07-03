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
{"from":"0x67ad0b6ab4cf735670ebb48e9c7d63f498ec2f55","to":"0x245db945c485b68fdc429e4f7085a1761aa4d45d","value":500000000000000000,"block":36004745}
{"from":"0x1277ebe8c161802c273c666d0884c1879f36b2d7","to":"0x245db945c485b68fdc429e4f7085a1761aa4d45d","value":84357669465561926,"block":36004746}
{"from":"0x67ad0b6ab4cf735670ebb48e9c7d63f498ec2f55","to":"0x245db945c485b68fdc429e4f7085a1761aa4d45d","value":500000000000000000,"block":36004748}
{"from":"0xfff9ce5f71ca6178d3beecedb61e7eff1602950e","to":"0x245db945c485b68fdc429e4f7085a1761aa4d45d","value":50150000000000,"block":36004749}
{"from":"0x67ad0b6ab4cf735670ebb48e9c7d63f498ec2f55","to":"0x245db945c485b68fdc429e4f7085a1761aa4d45d","value":500000000000000000,"block":36004751}
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

