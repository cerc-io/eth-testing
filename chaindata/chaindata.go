package chaindata

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

type Paths struct {
	ChainData, Ancient string
}

// List of names of chaindata fixtures accessible via ChainDataPaths
var FixtureChains = []string{
	"small", "small2",
}

func IsFixture(chain string) bool {
	has := false
	for _, fixture := range FixtureChains {
		if chain == fixture {
			has = true
			break
		}
	}
	return has
}

// GetFixture returns the absolute paths to fixture chaindata for the given name.
func GetFixture(chain string) (*Paths, error) {
	if !IsFixture(chain) {
		return nil, errors.New("no fixture named " + chain)
	}

	_, thisPath, _, ok := runtime.Caller(0)
	if !ok {
		return nil, errors.New("could not get function source path")
	}

	chaindataPath := filepath.Join(filepath.Dir(thisPath), "_data", chain)
	ancientdataPath := filepath.Join(chaindataPath, "ancient")

	if _, err := os.Stat(chaindataPath); err != nil {
		return nil, errors.New("cannot access chaindata at " + chaindataPath)
	}

	return &Paths{chaindataPath, ancientdataPath}, nil
}
