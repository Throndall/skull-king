package game

import (
	"fmt"
)

// Command is anything that modifies the game state
type Command interface {
	Execute(*State) error
}

// Context keeps record of all the commands executed.
type Context struct {
	state    State
	commands []Command
}

func NewContext() *Context {
	context := Context{
		commands: make([]Command, 0),
	}
	return &context
}

// Push attempts to execute a command and if successfully applied,
// it will be added to the command queu
// Commands will fail if the game state doesn't allow the given command
func (context *Context) Push(command Command) error {
	err := command.Execute(&context.state)
	if err != nil {
		return fmt.Errorf("Unable to execute command, %w", err)
	}
	context.commands = append(context.commands, command)
	return nil
}

// Pop removes the last executed command and restores the state.
// It is like undo
func (context *Context) Pop() Command {
	panic("Not Implemented Yet")
}

// State Creates a deep copy of the game State.
// This prevents external users from modifying it
func (context *Context) State() State {
	state := State{}
	for _, c := range context.commands {
		err := c.Execute(&state)
		if err != nil {
			panic(fmt.Errorf("Invalid command found in context, %#v", c))
		}
	}
	return state
}
