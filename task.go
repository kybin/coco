package cocofarm

// RootTask is a task that is hold property for it's sub tasks.
type RootTask struct {
	Task
	Property Property `json:"property"`
}

// Property defines how, and who will run the task.
type Property struct {
	WorkerGroup string `json:"worker_group"`
	Priority    int    `json:"priority"`
}

// Task holds commands to make a job done.
//
// A Task could also have it's subtasks.
// If so, it's running order is as follows.
//
// PreCmds -> (Subtasks's Commands) -> Cmds
//
// If one of the commands fails, the task will marked as fail,
// and the following commands will not run.
//
// Note that this does not support clean-up command
// for making the process simple.
// If you need that, please make your own batch script.
type Task struct {
	Root *RootTask

	Name string `json:"name"`

	PreCmds CmdGroup `json:"pre_cmds"`
	Cmds    CmdGroup `json:"cmds"`

	SubTasks       []*Task `json:"sub_tasks"`
	SerialSubTasks bool    `json:"serial_sub_tasks"` // If true, it will run next sub task when prior sub task is done.
}

// Property get it's root task's property.
func (t *Task) Property() Property {
	return t.Root.Property
}
