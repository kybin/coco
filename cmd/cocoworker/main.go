package main

// Worker waits and accepts a task from server.
//
// Log from the task saves as a file.
//
// You can run multiple workers in a machine,
// if then the port must specified, except a default one.
//
// If the machine's cpu has more than 8 threads, it will automatically drop
// 1 core for other job (is usually for monitoring).
