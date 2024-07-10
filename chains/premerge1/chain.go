// premerge1 is from a Clique-based pre-Merge chain with a single genesis allocation.
package premerge1

import (
	"github.com/cerc-io/eth-testing/chains"
)

var (
	ChainData, err = chains.GetFixture("premerge1")
)

func init() {
	if err != nil {
		panic(err)
	}
}
