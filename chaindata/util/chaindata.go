package util

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

type ChainData struct {
	Path, AncientPath string
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

// GetChainData returns the absolute paths to fixture chaindata for the given name.
func GetChainData(chain string) (*ChainData, error) {
	// fail if chain not in FixtureChains
	if !IsFixture(chain) {
		return nil, errors.New("no fixture named " + chain)
	}

	_, thisPath, _, ok := runtime.Caller(0)
	if !ok {
		return nil, errors.New("could not get function source path")
	}

	chainPath := filepath.Join(filepath.Dir(thisPath), "..", "data", chain)

	chaindataPath, err := filepath.Abs(chainPath)
	if err != nil {
		return nil, errors.New("cannot resolve path " + chainPath)
	}
	ancientdataPath := filepath.Join(chaindataPath, "ancient")

	if _, err := os.Stat(chaindataPath); err != nil {
		return nil, errors.New("must populate chaindata at " + chaindataPath)
	}

	return &ChainData{chaindataPath, ancientdataPath}, nil
}
