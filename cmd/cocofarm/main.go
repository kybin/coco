package main

import (
	"flag"
	"net/http"
)

// Farm communicates both clients and workers.
// It will keep watching jobs and workers, and match those.
//
// When it got a job from a client, it will split the job to tasks.
// it will look available workers and pick one randomly. (or just next one?)
// It sends a task to the worker via rpc, and waiting until the task is done by the worker.
// If the worker succeed it will mark the task done.
// If the worker fails to do the task, it will retry with another worker
// until reached to certern limit. Then it will mark the task failed.
//
// Farm also shows current status to it's web page.
// The page should show all current jobs and workers and their status.

func main() {
	var (
		addr string
		cert string
		key  string
	)
	flag.StringVar(&addr, "addr", ":3187", "https bind address")
	flag.StringVar(&cert, "cert", "cert/cert.pem", "https certificate file")
	flag.StringVar(&key, "key", "cert/key.pem", "https key file")
	flag.Parse()

	http.HandleFunc("/", rootHandler)
	http.ListenAndServeTLS(addr, cert, key, nil)
}
