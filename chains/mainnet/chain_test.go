package mainnet_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/triedb"

	"github.com/cerc-io/eth-testing/chains/mainnet"
)

func TestLoadChain(t *testing.T) {
	db := rawdb.NewMemoryDatabase()
	core.DefaultGenesisBlock().MustCommit(db, triedb.NewDatabase(db, nil))
	blocks := mainnet.GetBlocks()
	chain, _ := core.NewBlockChain(db, nil, nil, nil, ethash.NewFaker(), vm.Config{}, nil, nil)
	_, err := chain.InsertChain(blocks[1:])
	if err != nil {
		t.Fatal(err)
	}
}
