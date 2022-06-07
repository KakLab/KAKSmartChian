package clique

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/abi"
	"github.com/umbracle/ethgo/contract"
	"github.com/umbracle/ethgo/jsonrpc"
)

// http url for contract call
var httpUrl string

func SetStake(httpPort string) {
	httpUrl = "http://localhost:" + httpPort
}

// get stake list for contract
func getStake() []*common.Address {

	waddr := getWtake()
	if waddr == nil{
		return nil
	}
	if waddr.String() == "0x0000000000000000000000000000000000000000"{
		return nil
	}

	// function list
	var functions = []string{
		"function validators() public view returns (address[])",
	}

	// new contract
	abiContract, err := abi.NewABIFromList(functions)
	if err != nil {
		return nil
	}

	// grpc
	addr := ethgo.HexToAddress(waddr.String()) // 001 contract address for miner1
	client, err := jsonrpc.NewClient(httpUrl)
	if err != nil {
		return nil
	}

	// call contract
	c := contract.NewContract(addr, abiContract, contract.WithJsonRPC(client.Eth()))
	res, err := c.Call("validators", ethgo.Latest)
	if err != nil {
		return nil
	}

	addrs := make([]*common.Address, 0)
	for _, v := range res {
		if gethAddr, ok := v.([]ethgo.Address); ok {
			for _, v := range gethAddr {
				addr := common.BytesToAddress(v.Bytes())
				addrs = append(addrs, &addr)
			}
		}
	}
	log.Info(fmt.Sprintf("stake signer list:%s", addrs))
	return addrs
}

func getWtake() *common.Address {
	// function list
	var functions = []string{
		"function getStakeAddr() external view returns (address)",
	}

	// new contract
	abiContract, err := abi.NewABIFromList(functions)
	if err != nil {
		return nil
	}

	// grpc
	addr := ethgo.HexToAddress("0x879551c63238E0aA828432E579B078FD1EB9C96d") // 001 contract address for miner1
	//addr := ethgo.HexToAddress("0x2B541FAE5323757ef4556312dF9416c1446597D6") // 003 contract address for test1
	client, err := jsonrpc.NewClient(httpUrl)
	if err != nil {
		return nil
	}

	// call contract
	c := contract.NewContract(addr, abiContract, contract.WithJsonRPC(client.Eth()))
	res, err := c.Call("getStakeAddr", ethgo.Latest)
	if err != nil {
		return nil
	}

	addrs := make([]*common.Address, 0)
	for _, v := range res {
		if gethAddr, ok := v.(ethgo.Address); ok {
			addr := common.BytesToAddress(gethAddr.Bytes())
			addrs = append(addrs, &addr)
		}
	}
	log.Info(fmt.Sprintf("stake contract address:%s", addrs))
	return addrs[0]
}
