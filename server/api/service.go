package api

var Commands []TCPCommand = []TCPCommand{
	{
		Command:    "attach-bot",
		Parameters: []string{"token", "hash_id"},
		Handler:    AttachBot,
	},
}
