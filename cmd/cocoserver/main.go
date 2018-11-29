package main

// Server communicates both clients and workers.
// It will keep watching jobs and workers, and match those.
//
// When it got a job from a client, it will split the job to tasks.
// it will look available workers and pick one randomly. (or just next one?)
// It sends a task to the worker, and waiting until the task is done by the worker.
// If the worker succeed it will mark the task done.
// If the worker fails to do the task, it will retry with another worker
// until reached to certern limit. Then it will mark the task failed.
//
// Server also shows current status to it's web page.
// The page should show all current jobs and workers and their status.
