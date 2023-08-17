package game

type Player struct {
	Name string
}

type State struct {
	Players []Player
}

type Context struct {
	State    State
	Id       string
	Commands []Command
}

func NewContext() *Context{
	context := Context{
		Commands: make([]Command, 0),
	}
	return &context 
}

func (context *Context) Execute(command Command) {
	context.Commands = append(context.Commands, command)
	command.Execute(context)
}

type Command interface {
	Execute(*Context)
}
