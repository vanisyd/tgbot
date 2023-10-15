package api

type TCPCommand struct {
	Command    string
	Parameters []string
	Handler    func(args []string)
}
