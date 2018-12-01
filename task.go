package coco

// Task is a task handled by a worker.
// A task is either a parent task or a child task.
//
// A parent task could define how child tasks are excuted. (parrellel or series)
// A parent task could hold 'pre-command', 'post-command', 'pre-task-command', 'post-task-command'.
// Note that the environments from the commands will not inherited. Use 'env' for that.
//
// A Child holds series of (usually one) commands.
