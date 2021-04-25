package try

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

func readFile(path string) error {
	rf, err := os.Open(path)
	defer rf.Close()
	if err != nil {
		er := errors.WithMessage(err, "Failed to open")
		return errors.Wrap(er, "Failed to open11111")
	}
	return  nil
}

func a() error {
	return errors.WithMessage(readFile(filepath.Join(".", ".gitignore1")), "aaaaaaaaaa")
}

func Err1Start() {
	err := a()
	if err != nil {
		fmt.Printf("original error: %T -- aaa: %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack: \n%+v\n", err)
		os.Exit(0)
	}
}
