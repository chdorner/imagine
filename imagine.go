package main

import (
	_ "expvar"
	"flag"
	"fmt"
	"github.com/chdorner/imagine/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Set *automatically* at build time
var Version string

var (
	addr    = flag.String("a", ":8080", "address to bind to")
	daddr   = flag.String("debug.a", ":8081", "address to bind to for debug information")
	version = flag.Bool("version", false, "print version and exit")
)

func main() {
	if *version {
		fmt.Println(Version)
		os.Exit(0)
	}
	sigch := make(chan os.Signal)
	go handleSignals(sigch)
	signal.Notify(sigch)

	log.Fatal(server.ListenAndServe(*addr))
}

func handleSignals(sigch chan os.Signal) {
	for sig := range sigch {
		switch sig {
		case syscall.SIGTERM, syscall.SIGINT:
			log.Fatalf("received signal %d, exiting", sig)
		default:
			log.Printf("received signal %d, ignoring", sig)
		}
	}
}

func init() {
	flag.Parse()
}
