package main

// Client creates a job and send it to server.
// It will only send a job and exit, do not wait until the job is done.
//
// It upload a job as json file to cocoserver, and got the job id.
// Client does not know what will be the job id before submitted,
// so it will be empty.
