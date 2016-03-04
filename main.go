package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-signals
		fmt.Printf("received signal %s, but waiting 8 seconds before exiting\n", sig)
		time.Sleep(8 * time.Second)
		os.Exit(0)
	}()

	http.HandleFunc("/", hello)
	fmt.Println("running...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	rel := os.Getenv("RELEASE")
	t := time.Now()

	fmt.Fprintf(res, "Hello. it is %s and I'm on release %s\n", t.Format(time.UnixDate), rel)
}
