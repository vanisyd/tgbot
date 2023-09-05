package bot

import "strings"

type CMD struct {
	Command    string
	Parameters []string
	Handler    func([]string) string
}

func (cmd CMD) Signature() (signature string) {
	signature = cmd.Command
	if len(cmd.Parameters) > 0 {
		signature += " {" + strings.Join(cmd.Parameters, "} {") + "}"
	}

	return
}
