package cloudevent

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

var errInvalidDID = errors.New("invalid DID")

// NFTDID is a Decentralized Identifier for NFTs.
type NFTDID struct {
	ChainID         uint64         `json:"chainId"`
	ContractAddress common.Address `json:"contract"`
	TokenID         uint32         `json:"tokenId"`
}

// DecodeNFTDID decodes a DID string into a DID struct.
func DecodeNFTDID(did string) (NFTDID, error) {
	// sample did "did:nft:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF_1"
	parts := strings.Split(did, ":")
	if len(parts) != 4 {
		return NFTDID{}, errInvalidDID
	}
	if parts[0] != "did" {
		return NFTDID{}, fmt.Errorf("%w, incorrect DID prefix %s", errInvalidDID, parts[0])
	}
	if parts[1] != "nft" {
		return NFTDID{}, fmt.Errorf("%w, incorrect DID method %s", errInvalidDID, parts[1])
	}
	nftParts := strings.Split(parts[3], "_")
	if len(nftParts) != 2 {
		return NFTDID{}, fmt.Errorf("%w, incorrect NFT format %s", errInvalidDID, parts[3])
	}
	tokenID, err := strconv.ParseUint(nftParts[1], 10, 32)
	if err != nil {
		return NFTDID{}, fmt.Errorf("%w, invalid token ID %s", errInvalidDID, nftParts[1])
	}
	addrBytes := nftParts[0]
	if !common.IsHexAddress(addrBytes) {
		return NFTDID{}, fmt.Errorf("%w, invalid contract address %s", errInvalidDID, addrBytes)
	}
	chainID, err := strconv.ParseUint(parts[2], 10, 64)
	if err != nil {
		return NFTDID{}, fmt.Errorf("%w, invalid chain ID %s", errInvalidDID, parts[2])
	}

	return NFTDID{
		ChainID:         chainID,
		ContractAddress: common.HexToAddress(addrBytes),
		TokenID:         uint32(tokenID),
	}, nil
}

// String returns the string representation of the NFTDID.
func (n NFTDID) String() string {
	return fmt.Sprintf("did:nft:%d:%s_%d", n.ChainID, n.ContractAddress.Hex(), n.TokenID)
}
