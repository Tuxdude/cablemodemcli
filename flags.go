package main

import "flag"

var (
	debug          = flag.Bool("debug", false, "Log additional debug information except for requests and responses")
	debugReq       = flag.Bool("debugReq", false, "Log additional debug information for requests")
	debugResp      = flag.Bool("debugResp", false, "Log additional debug information for responses")
	host           = flag.String("host", "192.168.100.1", "Hostname or IP of your Arris S33 Cable modem")
	protocol       = flag.String("protocol", "https", "HTTP or HTTPS protocol to use")
	skipVerifyCert = flag.Bool("skipverifycert", true, "Skip SSL cert verification (because of self-signed certs on the cable modem)")
	username       = flag.String("username", "admin", "Admin username")
	password       = flag.String("password", "password", "Admin password")
	loop           = flag.Int("loop", 1, "Number of times to query in a loop, -1 to loop forever until an error")
	delay          = flag.Int("delay", 300, "Number of seconds delay between successive query attempts in a loop, setting a negative value here will result in applying a random delay each time")
	showOutput     = flag.Bool("print", true, "Whether to display the output of each status query")
	readFromFile   = flag.String("status_file", "", "Instead of querying the cable modem input, read the specified status file (in JSON format) just to verify parsing")
)
