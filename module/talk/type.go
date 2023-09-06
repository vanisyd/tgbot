package talk

type Meaning struct {
	Keywords []string
	Handler  func([]string) string
}
