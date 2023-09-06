package bot

import (
	"errors"
	"github.com/vanisyd/tgbot/tgapi"
	"strings"
)

var CurrentMSG tgapi.Message
var CurrentCMD CMD

func GetCMD(message string) (CMD, error) {
	requestedCmd := strings.Split(message, " ")
	for _, command := range Commands {
		if command.Command == requestedCmd[0] {
			CurrentCMD = command
			return command, nil
		}
	}

	return CMD{}, errors.New("command not found")
}

func GetParams(message string) []string {
	return strings.Split(message, " ")[1:]
}
