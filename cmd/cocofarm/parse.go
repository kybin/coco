package main

import (
	"encoding/json"
	"os"

	"github.com/kybin/cocofarm"
)

// parseJobFile parses a job file that is in the form of json.
func parseJobFile(fname string) (*cocofarm.Job, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	dec := json.NewDecoder(f)
	j := &cocofarm.Job{}
	if err := dec.Decode(j); err != nil {
		return nil, err
	}
	for _, root := range j.RootTasks {
		fillRootTask(&root.Task, root)
	}
	return j, nil
}

// fillRootTask fills the task's RootTask.
func fillRootTask(task *cocofarm.Task, root *cocofarm.RootTask) {
	task.Root = root
	for _, t := range task.SubTasks {
		fillRootTask(t, root)
	}
}
