package chaindata

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

// GetFixture returns the paths to the fixture chaindata for the given name.  This copies the
// fixture data and returns paths to temp directories, due to the fact that Go modules are installed
// read-only and Leveldb will attempt to create lock files when opening a DB.
func GetFixture(chain string) (*Paths, error) {
	if !IsFixture(chain) {
		return nil, errors.New("no fixture named " + chain)
	}

	_, thisPath, _, ok := runtime.Caller(0)
	if !ok {
		return nil, errors.New("could not get function source path")
	}

	chaindataPath := filepath.Join(filepath.Dir(thisPath), "_data", chain)
	if _, err := os.Stat(chaindataPath); err != nil {
		return nil, errors.New("cannot access chaindata at " + chaindataPath)
	}

	// Copy chaindata directory to a temporary directory
	// Note that we assume the ancient path is a subdirectory
	copyTo := filepath.Join(os.TempDir(), chain)
	if err := copyDir(chaindataPath, copyTo); err != nil {
		return nil, err
	}
	ancientCopy := filepath.Join(copyTo, "ancient")

	return &Paths{copyTo, ancientCopy}, nil
}

func copyDir(src, dest string) error {
	err := os.MkdirAll(dest, 0755)
	if err != nil {
		return err
	}

	srcDir, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcDir.Close()

	fileInfos, err := srcDir.Readdir(-1)
	if err != nil {
		return err
	}

	for _, fileInfo := range fileInfos {
		srcPath := filepath.Join(src, fileInfo.Name())
		destPath := filepath.Join(dest, fileInfo.Name())

		if fileInfo.IsDir() {
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
	if err != nil {
		return err
	}
	return nil
}
