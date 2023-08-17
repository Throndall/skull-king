package game

type CommandAddPlayer struct {
	Player Player
}

func (c CommandAddPlayer) Execute(context *Context) {
	context.State.Players = append(context.State.Players, c.Player)
}
