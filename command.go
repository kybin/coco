package cocofarm

// Cmd is a command to be run.
// First string is a command, and the others are arguments.
// when a Cmd is empty or nil, the Cmd will skipped.
type Cmd []string

// CmdGroup is a bunch of commands that will run in a machine.
// Task itself does not garantee that it is run inside of same machine.
type CmdGroup []Cmd
