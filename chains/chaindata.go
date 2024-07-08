package chains

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

type Paths struct {
	ChainData, Ancient string
}

// List of names of chaindata fixtures accessible via ChainDataPaths
var FixtureChains = []string{
	"premerge1",
	"premerge2",
	"postmerge1",
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

// GetFixture returns the paths to the fixture chaindata for the given name.  This copies the
// fixture data and returns paths to temp directories, due to the fact that Go modules are installed
// read-only and Leveldb/Pebble will attempt to create lock files when opening a DB.
func GetFixture(chain string) (*Paths, error) {
	if !IsFixture(chain) {
		return nil, errors.New("no fixture named " + chain)
	}

	_, thisPath, _, ok := runtime.Caller(0)
	if !ok {
		return nil, errors.New("could not get function source path")
	}

	chaindataPath := filepath.Join(filepath.Dir(thisPath), "data", chain, "geth", "chaindata")
	if _, err := os.Stat(chaindataPath); err != nil {
		return nil, errors.New("cannot access chaindata at " + chaindataPath)
	}

	// Copy chaindata directory to a temporary directory
	// Note that we assume the ancient path is a subdirectory
	copyTo, err := os.MkdirTemp("", chain+"-chaindata-*")
	if err != nil {
		return nil, err
	}
	if err := copyDir(chaindataPath, copyTo); err != nil {
		return nil, err
	}

	return &Paths{
		copyTo,
		filepath.Join(copyTo, "ancient"),
	}, nil
}

func copyDir(src, dest string) error {
	err := os.MkdirAll(dest, 0755)
	if err != nil {
		return err
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		if entry.IsDir() {
			if err := copyDir(srcPath, destPath); err != nil {
				return err
			}
		} else {
			if err := copyFile(srcPath, destPath); err != nil {
				return err
			}
		}
	}
	return nil
}

func copyFile(src, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	return err
}
