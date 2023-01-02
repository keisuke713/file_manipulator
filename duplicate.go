package fm

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"strconv"
)

type DuplicateCmd struct{}

func (d *DuplicateCmd) Name() string {
	return string(DUPLICATE)
}

func (d *DuplicateCmd) Usage() string {
	return "Write same contents of input file multiple times"
}

func (d *DuplicateCmd) Run(args []string) error {
	var (
		in *os.File
		n  int

		err error
	)
	defer in.Close()

	if len(args) < 2 {
		fmt.Println("You have to pass tww argument")
		return nil
	}
	in, err = os.OpenFile(args[0], os.O_RDWR, fs.ModePerm)
	if err != nil {
		return nil
	}
	n, err = strconv.Atoi(args[1])
	if err != nil {
		return err
	}
	if n < 0 {
		fmt.Println("second argument is more than zero")
		return nil
	}
	if err := d.RunWithFile(in, n); err != nil {
		return err
	}
	return nil
}

func (d *DuplicateCmd) RunWithFile(rw io.ReadWriter, n int) error {
	ib, err := io.ReadAll(rw)
	if err != nil {
		return err
	}
	for i := 0; i < n; i++ {
		rw.Write(ib)
	}
	return nil
}
