package main

// Property holds properties for a Task.
// If it is a branch task, the property will applyied for it and all it's subtasks.
// You can override properties multiple times.
type Property struct {
	WorkerGroup string // If empty, it will inherit the parent's worker group.
	Priority    int    // If 0 or negative, it will inherit the parent's priority.
}
