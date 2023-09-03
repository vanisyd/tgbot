package bot

var Commands []CMD = []CMD{
	CMD{
		Command:    "/reg",
		Parameters: []string{"password"},
		Handler:    Register,
	},
}
