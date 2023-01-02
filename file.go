package fm

import (
	"io/fs"
	"os"
)

const (
	dstName = "dst.txt"
)

func orCreate(args []string) (*os.File, error) {
	if len(args) < 2 {
		return os.Create(dstName)
	}
	return os.OpenFile(args[1], os.O_RDWR|os.O_CREATE, fs.ModePerm)
}
