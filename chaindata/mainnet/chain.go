package mainnet

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

const LatestBlock uint = 3

var blocks []*types.Block

// LoadBlocks loads fixture blocks up to and including the passed height or the highest available block
func GetBlocks() []*types.Block {
	if blocks != nil {
		return blocks
	}
	db := rawdb.NewMemoryDatabase()
	genesisBlock := core.DefaultGenesisBlock().MustCommit(db)
	genRLP, err := rlp.EncodeToBytes(genesisBlock)
	if err != nil {
		log.Fatal(err)
	}
	block0RLP, err := loadBlockRLP("./block0_rlp")
	if err != nil {
		panic(err)
	}
	if !bytes.Equal(genRLP, block0RLP) {
		panic("mainnet genesis blocks do not match")
	}

	blocks := make([]*types.Block, LatestBlock+1)
	blocks[0] = new(types.Block)
	rlp.DecodeBytes(block0RLP, blocks[0])

	for i := uint(1); i <= LatestBlock; i++ {
		blockRLP, err := loadBlockRLP(fmt.Sprintf("./block%d_rlp", i))
		if err != nil {
			panic(err)
		}
		blocks[i] = new(types.Block)
		rlp.DecodeBytes(blockRLP, blocks[i])
	}
	return blocks
}

func loadBlockRLP(filename string) ([]byte, error) {
	_, thisPath, _, ok := runtime.Caller(0)
	if !ok {
		panic("could not get function source path")
	}
	path := filepath.Join(filepath.Dir(thisPath), filename)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	blockRLP, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return blockRLP, nil
}
