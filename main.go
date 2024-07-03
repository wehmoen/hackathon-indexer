package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"math/big"
	"nmyk.io/cowsay"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	providerURL             = "https://api-gateway.skymavis.com/rpc/archive"
	treasuryDeploymentBlock = int64(16377111)
	walletAddress           = "0x245db945c485b68fdc429e4f7085a1761aa4d45d"
)

var erc20ABI = `[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"Transfer","type":"event"}]`

var contractAddresses = []string{
	"0xc99a6a985ed2cac1ef41640596c5a5f9f4e19ef5",
	"0x97a9107c1793bc407d6f527b77e7fff4d812bece",
	"0xa8754b9fa15fc18bb59458815510e40a12cd2014",
}

func printError(message string, exit ...bool) {
	data := struct {
		Message string `json:"message"`
	}{
		Message: message,
	}
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Printf(`{"message": "%s"}\n`, message)
	} else {
		fmt.Println(string(b))
	}

	if len(exit) == 0 || exit[0] {
		os.Exit(1)
	}

}

func main() {

	startBlockFlag := flag.Int64("start", treasuryDeploymentBlock, "Start block number")
	apiKeyFlag := flag.String("apikey", "", "SkyMavis RPC API key. Required if using default RPC URL")
	noCowFlag := flag.Bool("disable-cow", false, "Disable ASCII cow")
	customRPCFlag := flag.String("rpc", providerURL, "RPC URL")
	flag.Parse()

	if !*noCowFlag {
		cowsay.Cowsay("AxieGov Data Hackathon")
	}

	if *startBlockFlag < treasuryDeploymentBlock {
		printError(fmt.Sprintf("Start block number must be greater than or equal to %d", treasuryDeploymentBlock))
	}

	var provider string

	if *customRPCFlag == providerURL {
		if *apiKeyFlag == "" {
			printError("API key is required. Get one at https://developers.skymavis.com/")
		}
		provider = fmt.Sprintf("%s?apikey=%s", providerURL, *apiKeyFlag)
	} else {
		provider = *customRPCFlag
	}

	client, err := ethclient.Dial(fmt.Sprintf("%s?apikey=%s", provider, *apiKeyFlag))
	if err != nil {
		printError(fmt.Sprintf("Failed to connect to Ronin client: %v", err))
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		printError(fmt.Sprintf("Failed to get latest block number: %v", err))
	}

	latestBlock := header.Number.Int64()

	if *startBlockFlag > latestBlock {
		printError(fmt.Sprintf("Start block number must be less than or equal to the latest block number: %d", latestBlock))
	}

	startBlock := *startBlockFlag

	erc20ABIParsed, err := abi.JSON(strings.NewReader(erc20ABI))
	if err != nil {
		printError(fmt.Sprintf("Failed to parse ERC20 ABI: %v", err))
	}

	addresses := make([]common.Address, len(contractAddresses))
	for i, addr := range contractAddresses {
		addresses[i] = common.HexToAddress(addr)
	}

	wallet := common.HexToAddress(walletAddress)

	for blockNumber := startBlock; blockNumber <= latestBlock; blockNumber++ {
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(blockNumber),
			ToBlock:   big.NewInt(blockNumber),
			Addresses: addresses,
		}

		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			printError(fmt.Sprintf("Failed to filter logs for block %d: %v", blockNumber, err))
		}

		for _, vLog := range logs {
			if vLog.Topics[0] == erc20ABIParsed.Events["Transfer"].ID {
				var transferEvent struct {
					From  common.Address
					To    common.Address
					Value *big.Int
				}

				err := erc20ABIParsed.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
				if err != nil {
					printError(fmt.Sprintf("Failed to unpack log: %v", err), false)
					continue
				}

				transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
				transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())
				if transferEvent.From == wallet || transferEvent.To == wallet {

					output := struct {
						From  common.Address `json:"from"`
						To    common.Address `json:"to"`
						Value *big.Int       `json:"value"`
						Block int64          `json:"block"`
					}{
						From:  transferEvent.From,
						To:    transferEvent.To,
						Value: transferEvent.Value,
						Block: blockNumber,
					}

					jsonBytes, err := json.Marshal(output)

					if err != nil {
						printError(fmt.Sprintf("Failed to marshal output: %v", err))
					}

					fmt.Println(string(jsonBytes))
				}
			}
		}
	}
}
