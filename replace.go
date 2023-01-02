package fm

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"
)

type ReplaceCmd struct{}

func (r *ReplaceCmd) Name() string {
	return string(REPLACE)
}

func (r *ReplaceCmd) Usage() string {
	return "replace word you specified"
}

func (r *ReplaceCmd) Run(args []string) error {
	if len(args) != 3 {
		fmt.Println("You have to pass three arguments")
		return nil
	}

	if err := r.RunWithFile(args[0], args[1], args[2]); err != nil {
		return err
	}
	return nil
}

func (r *ReplaceCmd) RunWithFile(fn, old, new string) error {
	in, err := os.Open(fn)
	if err != nil {
		return err
	}
	ib, err := io.ReadAll(in)
	if err != nil {
		return err
	}
	in.Close()
	is := string(ib)
	rs := strings.ReplaceAll(is, old, new)

	// ファイルは再読み込み出来ないので開き直す
	in, err = os.OpenFile(fn, os.O_RDWR, fs.ModePerm)
	if err != nil {
		return nil
	}
	in.Write([]byte(rs))

	diff := len(is) - len(rs)
	if diff > 0 {
		for i := 0; i < diff; i++ {
			in.Write([]byte(" "))
		}
	}
	return nil
}
