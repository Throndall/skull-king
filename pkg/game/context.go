package game

import (
	"bytes"
	"encoding/gob"
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
	// The problem is that our data contains slices [] that are stored
	// by reference. If we simply return the State, it will copy
	// Plain field but slices will be pointing to the same underlying data
	// This can be handled manually by coping the slices, but it requires
	// a lot of typing and is error prone.
	// SO because we don't care a about performance, the following trick works
	// by serializing and deserializing the data
	b := &bytes.Buffer{}
	gob.NewEncoder(b).Encode(context.state)

	state := State{}
	gob.NewDecoder(b).Decode(&state)
	return state
}
