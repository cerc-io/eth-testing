package premerge2

import (
	"github.com/cerc-io/eth-testing/chains"
	"github.com/cerc-io/eth-testing/utils"
)

var (
	ChainData, err = chains.GetFixture("premerge2")

	// keybytes-encoded leaf keys of all state nodes
	Block1_StateNodeLeafKeys [][]byte
)

func init() {
	if err != nil {
		panic(err)
	}

	for _, path := range Block1_StateNodePaths {
		if len(path) > 0 && path[len(path)-1] == 16 {
			hash := utils.HexToKeyBytes(path)
			Block1_StateNodeLeafKeys = append(Block1_StateNodeLeafKeys, hash)
		}
	}
}
