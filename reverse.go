package fm

import (
	"fmt"
	"io"
	"os"
)

type ReverseCmd struct{}

func (r *ReverseCmd) Name() string {
	return string(REVERSE)
}

func (r *ReverseCmd) Usage() string {
	return "Write reversed contents of input file to output one"
}

func (r *ReverseCmd) Run(args []string) error {
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
		fmt.Println("Here is an example")
		r.Example()
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

	if err := r.RunWithFile(in, out); err != nil {
		return err
	}
	return nil
}

func (r *ReverseCmd) RunWithFile(in io.Reader, out io.Writer) error {
	ib, err := io.ReadAll(in)
	if err != nil {
		return err
	}
	is := string(ib)
	rns := []rune(is)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	out.Write([]byte(string(rns)))
	return nil
}

func (r *ReverseCmd) Example() string {
	return fmt.Sprintf("`%s %s %s %s`", BINARY_NAME, r.Name(), "input.txt", "output.txt")
}
