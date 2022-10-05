package main

import "log"

func Run() error {
	// should be responsible for initializing the gRPC server
	return nil
}

func main() {
	if err := Run(); err != null {
		log.Fatal(err)
	}
}
