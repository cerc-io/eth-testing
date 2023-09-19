package util

import (
	"os"
	"path/filepath"
	"runtime"
)

// ChainDataPath returns the absolute paths to fixture chaindata for the given name.
// Currently the only chains are:
// - small: 32-block small-size chain
// - medium: a short medium-sized chain
func ChainDataPaths(chain string) (string, string) {
	_, thisPath, _, ok := runtime.Caller(0)
	if !ok {
		panic("could not get function source path")
	}

	chainPath := filepath.Join(filepath.Dir(thisPath), "..", "data", chain)

	chaindataPath, err := filepath.Abs(chainPath)
	if err != nil {
		panic("cannot resolve path " + chainPath)
	}
	ancientdataPath := filepath.Join(chaindataPath, "ancient")

	if _, err := os.Stat(chaindataPath); err != nil {
		panic("must populate chaindata at " + chaindataPath)
	}

	return chaindataPath, ancientdataPath
}
