package small

import (
	"github.com/cerc-io/eth-testing/chaindata"
)

var (
	ChainData, err = chaindata.GetFixture("small")
)

func init() {
	if err != nil {
		panic(err)
	}
}
