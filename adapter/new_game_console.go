package adapter

type NewGameConsole struct{}

func (c *NewGameConsole) OutputAV() string {
	return "AV signal from New Game Console"
}
