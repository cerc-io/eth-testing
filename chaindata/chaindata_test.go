package chaindata_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"

	"github.com/cerc-io/eth-testing/chaindata"
)

func testReadChainData(t *testing.T, data *chaindata.Paths) {
	kvdb, ldberr := rawdb.NewLevelDBDatabase(data.ChainData, 1024, 256, t.Name(), true)
	if ldberr != nil {
		t.Fatal(ldberr)
	}
	edb, err := rawdb.NewDatabaseWithFreezer(kvdb, data.Ancient, t.Name(), true)
	if err != nil {
		t.Fatal(err)
	}
	defer edb.Close()

	// Check that we can open and traverse a trie at the head block
	hash := rawdb.ReadHeadHeaderHash(edb)
	height := rawdb.ReadHeaderNumber(edb, hash)
	if height == nil {
		t.Fatalf("unable to read header height for header hash %s", hash)
	}
	header := rawdb.ReadHeader(edb, hash, *height)
	if header == nil {
		t.Fatalf("unable to read canonical header at height %d", height)
	}
	sdb := state.NewDatabase(edb)
	tree, err := sdb.OpenTrie(header.Root)
	if err != nil {
		t.Fatal(err)
	}
	it := tree.NodeIterator(nil)
	for it.Next(true) {
	}
	if err := it.Error(); err != nil {
		t.Fatal(err)
	}
}

func TestReadChainData(t *testing.T) {
	for _, name := range []string{"small", "small2"} {
		t.Run(name, func(t *testing.T) {
			data, err := chaindata.GetFixture(name)
			if err != nil {
				t.Fatal(err)
			}
			testReadChainData(t, data)
		})
	}
}
