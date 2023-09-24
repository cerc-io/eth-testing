package small2

import (
	"github.com/cerc-io/eth-testing/chaindata"
)

var (
	ChainData, err = chaindata.GetFixture("small2")
)

func init() {
	if err != nil {
		panic(err)
	}
}
