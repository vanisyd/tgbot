package bot

import "strings"

func GetCMD(message string) CMD {
	requestedCmd := strings.Split(message, " ")
	for _, command := range Commands {
		if command.Command == requestedCmd[0] {
			return command
		}
	}

	return CMD{}
}

func GetParams(message string) []string {
	return strings.Split(message, " ")[1:]
}
