package cocofarm

// Job consists of root tasks.
type Job struct {
	Name      string      `json:"name"`
	RootTasks []*RootTask `json:"root_tasks"`
}
