package main

import (
	_ "expvar"
	"flag"
	"fmt"
	"github.com/chdorner/imagine/server"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
)

var (
	addr       string
	daddr      string
	version    bool
	oWhitelist originWhitelist
)

type originWhitelist struct {
	slice []*regexp.Regexp
}

func main() {
	if version {
		fmt.Println(server.Version)
		os.Exit(0)
	}
	sigch := make(chan os.Signal)
	go handleSignals(sigch)
	signal.Notify(sigch)

	log.Fatal(server.ListenAndServe(addr, oWhitelist.slice))
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
	flag.StringVar(&addr, "a", ":8080", "address to bind to")
	flag.StringVar(&daddr, "debug.a", ":8081", "address to bind to for debug information")
	flag.BoolVar(&version, "version", false, "print version and exit")
	flag.Var(&oWhitelist, "o", "origin whitelist regex (may be used multiple times)")
	flag.Parse()
}

func (w *originWhitelist) Set(s string) error {
	r, err := regexp.Compile(s)
	if err != nil {
		return fmt.Errorf(`"%s" is not a valid regular expression`, s)
	}

	w.slice = append(w.slice, r)
	return nil
}

func (w *originWhitelist) String() string {
	ret := make([]string, len(w.slice))
	for i, r := range w.slice {
		ret[i] = fmt.Sprint(r.String())
	}
	str := fmt.Sprintf("[%s]", strings.Join(ret, " "))
	return str
}
