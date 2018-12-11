package coco

// Task is either a leaf task or a branch task.
// If a task holds children, it is a branch task, or it is a leaf task.

// Cmd is a command to be run.
// First string is a command, and the others are arguments.
// when a Cmd is empty or nil, the Cmd will skipped.
type Cmd []string

// CmdGroup is a bunch of commands that will run in a machine.
// Task itself does not garantee that it is run inside of same machine.
type CmdGroup []Cmd

// Task's running order is as follows.
//
// PreCmds -> (Subtasks's Commands) -> Cmds
//
// If one of the commands fails, task will also marked as fail,
// and the following commands will not run.
// Note that this does not support clean-up command
// for making the process simple.
// If you need this, please make your own batch script.
type Task struct {
	Property Property

	PreCmds CmdGroup
	Cmds    CmdGroup

	Subtasks       []Task
	SerialSubtasks bool // If true, it will run next sub task when prior sub task is done.
}
