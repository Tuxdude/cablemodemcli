// Command cablemodemcli is a tool to query the cable modem.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"github.com/tuxdude/cablemodemutil"
)

const (
	minRandDelaySeconds = 5 * 60
	maxRandDelaySeconds = 30 * 60
)

// Returns the JSON formatted string representation of the specified object.
func prettyPrintJSON(x interface{}) string {
	p, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		return fmt.Sprintf("%#v", x)
	}
	return string(p)
}

func randDelay() time.Duration {
	max := maxRandDelaySeconds
	min := minRandDelaySeconds
	d := rand.Intn(max-min) + min
	return time.Duration(d) * time.Second
}

func handleErr(err error) int {
	fmt.Fprintf(os.Stderr, "Error: %s", err)
	return 1
}

func runInFileMode() int {
	f, err := ioutil.ReadFile(*readFromFile)
	if err != nil {
		return handleErr(err)
	}
	var raw cablemodemutil.CableModemRawStatus
	err = json.Unmarshal(f, &raw)
	if err != nil {
		return handleErr(err)
	}
	status, err := cablemodemutil.ParseRawStatus(raw)
	if err != nil {
		return handleErr(err)
	}
	if *showOutput {
		fmt.Println(prettyPrintJSON(status))
	}
	return 0
}

func run() int {
	if *readFromFile != "" {
		return runInFileMode()
	}

	input := cablemodemutil.RetrieverInput{
		Host:           *host,
		Protocol:       *protocol,
		SkipVerifyCert: *skipVerifyCert,
		Username:       *username,
		ClearPassword:  *password,
	}
	input.Debug.Debug = *debug
	input.Debug.DebugReq = *debugReq
	input.Debug.DebugResp = *debugResp
	cm := cablemodemutil.NewStatusRetriever(&input)
	beginTime := time.Now()
	nextRequestTime := beginTime

	if *loop == 0 {
		fmt.Fprintln(os.Stderr, "Warning: -loop flag set to zero, hence not querying status from the cable modem.")
	}

	forever := *loop < 0
	pending := *loop
	for forever || pending > 0 {
		pending--
		currTime := time.Now()
		if *delay < 0 {
			nextRequestTime = nextRequestTime.Add(randDelay())
		} else {
			nextRequestTime = nextRequestTime.Add(time.Duration(*delay) * time.Second)
		}
		if *debug {
			fmt.Printf("\n\n************************************************\n")
			fmt.Printf("Begin time:%s\n", beginTime)
			fmt.Printf("Status Query time: %s\n", currTime)
			fmt.Printf("************************************************\n\n")
		}

		status, err := cm.Status()
		if err != nil {
			return handleErr(err)
		}
		if *showOutput {
			fmt.Println(prettyPrintJSON(status))
		}

		if *debug {
			if forever || pending > 0 {
				fmt.Printf("\n\n************************************************\n")
				fmt.Printf("Begin time:%s\n", beginTime)
				fmt.Printf("Last Query time: %s\n", currTime)
				fmt.Printf("Next Query time: %s\n", nextRequestTime)
				fmt.Printf("************************************************\n\n")
				time.Sleep(time.Until(nextRequestTime))
			} else {
				fmt.Printf("\n\n************************************************\n")
				fmt.Printf("Begin time:%s\n", beginTime)
				fmt.Printf("Last Query time: %s\n", currTime)
				fmt.Printf("************************************************\n\n")
			}
		}
	}

	return 0
}

func main() {
	flag.Parse()
	os.Exit(run())
}
