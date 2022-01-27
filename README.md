# cablemodemcli

[![Build](https://github.com/Tuxdude/cablemodemcli/actions/workflows/build.yml/badge.svg)](https://github.com/Tuxdude/cablemodemcli/actions/workflows/build.yml) [![Tests](https://github.com/Tuxdude/cablemodemcli/actions/workflows/tests.yml/badge.svg)](https://github.com/Tuxdude/cablemodemcli/actions/workflows/tests.yml)

[![Lint](https://github.com/Tuxdude/cablemodemcli/actions/workflows/lint.yml/badge.svg)](https://github.com/Tuxdude/cablemodemcli/actions/workflows/lint.yml) [![CodeQL](https://github.com/Tuxdude/cablemodemcli/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/Tuxdude/cablemodemcli/actions/workflows/codeql-analysis.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/tuxdude/cablemodemcli)](https://goreportcard.com/report/github.com/tuxdude/cablemodemcli)

A CLI (written in go) for querying the status of Cable Modems.

This currently only works (and has been tested with) an Arris S33 Cable
Modem. The CLI uses the [cablemodemutil go library](https://github.com/Tuxdude/cablemodemutil)
for interacting with the cable modem.

If you would like to add support for other cable modems, please
file an Issue or submit a pull request with details for further discussion.

# Usage

Just build and run the binary with the necessary flags depending on your
cable modem configuration.

## Relevant `cablemodemcli` Flags

These are the common set of flags that will be required for typically
querying your cable modem. In many cases the default values of these
flags might already have the right value and you can skip specifying
such flag altogether.

See the [example](#example-usage) for a typical usage.

```
  -host string
        Hostname or IP of your Arris S33 Cable modem (default "192.168.100.1")
  -protocol string
        HTTP or HTTPS protocol to use (default "https")
  -skipverifycert
        Skip SSL cert verification (because of self-signed certs on the cable modem) (default true)
  -username string
        Admin username (default "admin")
  -password string
        Admin password (default "password")
```

## Advanced Debugging and Miscellaneous `cablemodemcli` Flags

These are more advanced flags meant for either debugging the CLI
and/or the `cablemodemutil` library along with debugging the
control (eg. run the program in a loop and validate that the
library/CLI behaves correctly).

You could potentially use that to run a daemon although logging
so much information to STDOUT might not be relevant for most
scenarios.

```
  -debug
        Log additional debug information except for requests and responses
  -debugReq
        Log additional debug information for requests
  -debugResp
        Log additional debug information for responses
  -delay int
        Number of seconds delay between successive query attempts in a loop, setting a negative value here will result in applying a random delay each time (default 300)
  -loop int
        Number of times to query in a loop, -1 to loop forever until an error (default 1)
  -print
        Whether to display the output of each status query (default true)
  -status_file string
        Instead of querying the cable modem input, read the specified status file (in JSON format) just to verify parsing
```

# Example Usage

```
$ cablemodemcli -host 192.168.100.1 -username user1 -password pass1

{
  "DeviceInfo": {
    "Model": "S33",
    "SerialNumber": "123456789012345",
    "MACAddress": "12:34:56:78:90:AB"
  },
  "DeviceSettings": {
    "FrontPanelLightsOn": true,
    "EnergyEfficientEthernetOn": false,
    "AskMeLater": false,
    "NeverAsk": true
  },
  "AuthSettings": {
    "CurrentLogin": "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
    "CurrentNameAdmin": "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
    "CurrentNameUser": "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
    "CurrentPasswordAdmin": "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
    "CurrentPasswordUser": "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"
  },
  "SoftwareStatus": {
    "FirmwareVersion": "123_456_789_012_345.67890",
    "CertificateInstalled": true,
    "CustomerVersion": "Foo_Bar",
    "HDVersion": "V1.0",
    "DOCSISSpecVersion": "DOCSIS 3.1"
  },
  "StartupStatus": {
    "Boot": {
      "Status": "OK",
      "Comment": "Operational"
    },
    "ConfigurationFile": {
      "Status": "OK",
      "Comment": ""
    },
    "Connectivity": {
      "Status": "OK",
      "Comment": "Operational"
    },
    "DownstreamConnection": {
      "FrequencyHZ": 507000000,
      "Comment": "Locked"
    },
    "Security": {
      "Status": "Enabled",
      "Comment": "BPI+"
    }
  },
  "ConnectionStatus": {
    "ConnectionEstablishedTime": "2022-01-20T00:15:21-08:00",
    "SystemTime": "2022-01-26T18:37:39-08:00",
    "UpTime": 584538000000000,
    "DOCSISNetworkAccess": "Allowed",
    "InternetConnectionStatus": "Connected",
    "DownstreamPlan": "NorthAmerica",
    "DownstreamFrequencyHZ": 507000000,
    "DownstreamSignalPowerDBMV": -3,
    "DownstreamSignalSNRDB": 39,
    "UpstreamChannelID": 2,
    "DownstreamChannels": [
      {
        "LockStatus": "Locked",
        "Modulation": "QAM256",
        "ChannelID": 20,
        "FrequencyHZ": 507000000,
        "SignalPowerDBMV": -4,
        "SignalSNRMERDB": 39,
        "CorrectedErrors": 332,
        "UncorrectedErrors": 1229
      },
      {
        "LockStatus": "Locked",
        "Modulation": "QAM256",
        "ChannelID": 17,
        "FrequencyHZ": 483000000,
        "SignalPowerDBMV": -5,
        "SignalSNRMERDB": 39,
        "CorrectedErrors": 446,
        "UncorrectedErrors": 1200
      },
      {
        "LockStatus": "Locked",
        "Modulation": "QAM256",
        "ChannelID": 18,
        "FrequencyHZ": 489000000,
        "SignalPowerDBMV": -3,
        "SignalSNRMERDB": 40,
        "CorrectedErrors": 379,
        "UncorrectedErrors": 1272
      },
      {
        "LockStatus": "Locked",
        "Modulation": "QAM256",
        "ChannelID": 19,
        "FrequencyHZ": 495000000,
        "SignalPowerDBMV": -3,
        "SignalSNRMERDB": 40,
        "CorrectedErrors": 358,
        "UncorrectedErrors": 1190
      },
      {
        "LockStatus": "Locked",
        "Modulation": "OFDM PLC",
        "ChannelID": 48,
        "FrequencyHZ": 850000000,
        "SignalPowerDBMV": -8,
        "SignalSNRMERDB": 37,
        "CorrectedErrors": 470417943,
        "UncorrectedErrors": 60
      }
    ],
    "UpstreamChannels": [
      {
        "LockStatus": "Locked",
        "Modulation": "SC-QAM",
        "ChannelID": 1,
        "WidthHZ": 3200000,
        "FrequencyHZ": 10400000,
        "SignalPowerDBMV": 40
      },
      {
        "LockStatus": "Locked",
        "Modulation": "SC-QAM",
        "ChannelID": 2,
        "WidthHZ": 6400000,
        "FrequencyHZ": 16400000,
        "SignalPowerDBMV": 41.3
      }
    ]
  },
  "Logs": [
    {
      "Timestamp": "2022-01-26T13:11:09-08:00",
      "Log": "RNG-RSP CCAP Commanded Power Exceeds Value Corresponding to the Top of the DRW;CM-MAC=<12:34:56:78:90:ab>;CMTS-MAC=<cd:ef:12:34:56:78>;CM-QOS=1.1;CM-VER=3.1;"
    },
    {
      "Timestamp": "2022-01-26T13:11:09-08:00",
      "Log": "Dynamic Range Window violation"
    },
    {
      "Timestamp": "2022-01-26T13:46:21-08:00",
      "Log": "Successful LAN WebGUI login from 192.168.1.1 on 22/01/26 at 1:46 PM."
    }
  ]
}
```
