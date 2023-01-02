package fm

type (
	CmdName string

	Cmd interface {
		Name() string
		Usage() string
		Run([]string) error
	}
)

const (
	BINARY_NAME = "fm"
)

const (
	REVERSE   CmdName = "reverse"
	COPY      CmdName = "copy"
	DUPLICATE CmdName = "dup"
	REPLACE   CmdName = "replace"
	HELP      CmdName = "help"
)

var CmdMap = map[CmdName]Cmd{
	HELP:      &HelpCmd{},
	REVERSE:   &ReverseCmd{},
	COPY:      &CopyCmd{},
	DUPLICATE: &DuplicateCmd{},
	REPLACE:   &ReplaceCmd{},
}
