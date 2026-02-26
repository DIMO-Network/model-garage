package ruptela_test

import (
	"math/big"

	"github.com/DIMO-Network/cloudevent"
	"github.com/ethereum/go-ethereum/common"
)

var (
	vehicleContract     = common.HexToAddress("0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF")
	aftermarketContract = common.HexToAddress("0x06012c8cf97BEaD5deAe237070F9587f8E7A266d")

	vehicleSubject1      = cloudevent.ERC721DID{ChainID: 1, ContractAddress: vehicleContract, TokenID: big.NewInt(1)}.String()
	vehicleSubject33     = cloudevent.ERC721DID{ChainID: 1, ContractAddress: vehicleContract, TokenID: big.NewInt(33)}.String()
	vehicleSubject162682 = cloudevent.ERC721DID{ChainID: 137, ContractAddress: vehicleContract, TokenID: big.NewInt(162682)}.String()

	aftermarketSubject2  = cloudevent.ERC721DID{ChainID: 1, ContractAddress: aftermarketContract, TokenID: big.NewInt(2)}.String()
	aftermarketSubject33 = cloudevent.ERC721DID{ChainID: 1, ContractAddress: aftermarketContract, TokenID: big.NewInt(33)}.String()
)
