package groupio

type ExprGroup struct {
	GroupName string
	Repl      func(string) string
}
