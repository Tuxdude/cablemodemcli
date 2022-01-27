package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/tuxdude/cablemodemutil"
)

// Returns the JSON formatted string representation of the specified object.
func prettyPrintJSON(x interface{}) string {
	p, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		return fmt.Sprintf("%#v", x)
	}
	return string(p)
}

func run() int {
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
	status, err := cm.Status()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
		return 1
	}
	fmt.Println(prettyPrintJSON(status))
	return 0
}

func main() {
	flag.Parse()
	os.Exit(run())
}
