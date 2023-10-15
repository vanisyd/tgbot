package bot

import (
	"errors"
	database "github.com/vanisyd/tgbot-db"
	"github.com/vanisyd/tgbot/tgapi"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

var CurrentMSG tgapi.Message
var CurrentCMD CMD
var CurrentDBUser database.User

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

func GetCurrentDBUser() database.User {
	if CurrentDBUser == (database.User{}) {
		user := database.FindUser(CurrentMSG.From.ID)
		if _, ok := user.(database.User); !ok {
			uid := database.AddUser(database.User{TgID: CurrentMSG.From.ID})
			if userObjId, ok := uid.(primitive.ObjectID); ok {
				user = database.GetUser(userObjId)
			}
		}

		if userObj, ok := user.(database.User); ok {
			CurrentDBUser = userObj
		}
	}

	return CurrentDBUser
}
