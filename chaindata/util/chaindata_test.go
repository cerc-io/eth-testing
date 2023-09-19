package util_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/core/rawdb"

	"github.com/cerc-io/eth-testing/chaindata/util"
)

func testReadChainData(t *testing.T, name string) {
	ChainDataPath, AncientDataPath := util.ChainDataPaths(name)

	kvdb, ldberr := rawdb.NewLevelDBDatabase(ChainDataPath, 1024, 256, "vdb-geth", false)
	if ldberr != nil {
		t.Fatal(ldberr)
	}
	edb, err := rawdb.NewDatabaseWithFreezer(kvdb, AncientDataPath, "vdb-geth", false)
	if err != nil {
		t.Fatal(err)
	}
	if err != nil {
		t.Fatal(err)
	}
	defer edb.Close()

	hash := rawdb.ReadHeadHeaderHash(edb)
	height := rawdb.ReadHeaderNumber(edb, hash)
	if height == nil {
		t.Fatalf("unable to read header height for header hash %s", hash)
	}
	header := rawdb.ReadHeader(edb, hash, *height)
	if header == nil {
		t.Fatalf("unable to read canonical header at height %d", height)
	}
}

func TestReadChainData(t *testing.T) {
	for _, name := range []string{"small", "medium"} {
		t.Run(name, func(t *testing.T) { testReadChainData(t, name) })
	}
}
