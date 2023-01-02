package fm

import (
	"fmt"
	"io"
	"os"
)

type CopyCmd struct{}

func (c *CopyCmd) Name() string {
	return string(COPY)
}

func (c *CopyCmd) Usage() string {
	return "Write same contents of input file to output one"
}

func (c *CopyCmd) Run(args []string) error {
	var (
		in  *os.File
		out *os.File

		err error
	)
	defer func() {
		in.Close()
		out.Close()
	}()
	// args[0]が存在すること
	if len(args) < 1 {
		fmt.Println("You have to pass one file name at least")
		return nil
	}
	in, err = os.Open(args[0])
	if err != nil {
		return err
	}
	// args[1]がなかったらデフォルトのファイルを作る
	out, err = orCreate(args)
	if err != nil {
		return nil
	}

	if err := c.RunWithFile(in, out); err != nil {
		return err
	}
	return nil
}

func (c *CopyCmd) RunWithFile(in io.Reader, out io.Writer) error {
	ib, err := io.ReadAll(in)
	if err != nil {
		return err
	}
	out.Write(ib)
	return nil
}
