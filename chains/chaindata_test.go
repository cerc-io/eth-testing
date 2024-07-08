package chains_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"

	"github.com/cerc-io/eth-testing/chains"
)

func testReadChainData(t *testing.T, data *chains.Paths) {
	edb, err := rawdb.Open(rawdb.OpenOptions{
		Directory:         data.ChainData,
		AncientsDirectory: data.Ancient,
		Namespace:         t.Name(),
		Cache:             1024,
		Handles:           256,
		ReadOnly:          true,
	})
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
		t.Fatalf("unable to read canonical header at height %d", *height)
	}
	sdb := state.NewDatabase(edb)
	tree, err := sdb.OpenTrie(header.Root)
	if err != nil {
		t.Fatal(err)
	}
	it, err := tree.NodeIterator(nil)
	if err != nil {
		t.Fatal(err)
	}
	for it.Next(true) {
	}
	if err := it.Error(); err != nil {
		t.Fatal(err)
	}
}

func TestReadChainData(t *testing.T) {
	for _, name := range chains.FixtureChains {
		t.Run(name, func(t *testing.T) {
			data, err := chains.GetFixture(name)
			if err != nil {
				t.Fatal(err)
			}
			testReadChainData(t, data)
		})
	}
}
