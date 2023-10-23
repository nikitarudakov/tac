package groupio

type ExprGroup struct {
	GroupName string
	Repl      func(string) string
}

type ExprGroupMapper map[string]ExprGroup
