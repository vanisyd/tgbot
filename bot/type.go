package bot

type CMD struct {
	Command    string
	Parameters []string
	Handler    func([]string) string
}
