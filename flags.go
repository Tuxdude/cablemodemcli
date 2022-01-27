package main

import "flag"

var (
	host           = flag.String("host", "192.168.100.1", "Hostname or IP of your Arris S33 Cable modem")
	protocol       = flag.String("protocol", "https", "HTTP or HTTPS protocol to use")
	skipVerifyCert = flag.Bool("skipverifycert", true, "Skip SSL cert verification (because of self-signed certs on the cable modem)")
	username       = flag.String("username", "admin", "Admin username")
	password       = flag.String("password", "password", "Admin password")
)
