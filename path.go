package goutils

import (
	"os"
)

// PathExists returns true if the given path exists, otherwise false. If an
// error is encountered, that is returned, otherwise error will be nil.
//
// Since it not existing will cause an os.Stat error, we need to check if the
// error passed back is the IsNotExist error, which would mean the path does
// not exist, instead of an ectual error state.
func PathExists(p string) (bool, error) {
	_, err := os.Stat(p)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}
	
	return false, err
}
