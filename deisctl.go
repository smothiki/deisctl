package main

import (
	"fmt"
	"os"
  "path/filepath"
  "github.com/deis/deisctl/systemd"
  "github.com/deis/deis-without-fleet/tests/utils"
	docopt "github.com/docopt/docopt-go"
)

const (
  downloadDir     = "/home/core/deis/systemd/"
  )


func cmdInstall( ){
  fmt.Println("starting systemd units")
	files, _ := utils.ListFiles(downloadDir + "*.service")
  conn, _ :=systemd.NewSystemdUnitManager()
	conn.Enable(files)
	for _, file := range files {
		_, file = filepath.Split(file)
		conn.Start(file)
	}
}

func cmdUninstall(){
  fmt.Println("starting systemd units")
  files, _ := utils.ListFiles(downloadDir + "*.service")
  conn, _ :=systemd.NewSystemdUnitManager()
  for _, file := range files {
    _, file = filepath.Split(file)
    conn.Stop(file)
  }
  conn.Disable(files)
}


func main() {
	usage := `Deis Control Utility

Usage:
  deisctl <command> [<target>...] [options]

Example Commands:

  deisctl install
  deisctl uninstall
  deisctl list
  deisctl start router
  deisctl stop router builder
  deisctl status controller

Options:
  --debug                     print debug information to stderr
  --endpoint=<url>            etcd endpoint for fleet [default: http://127.0.0.1:4001]
  --etcd-key-prefix=<path>    keyspace for fleet data in etcd [default: /_coreos.com/fleet/]
  --known-hosts-file=<path>   file used to store remote machine fingerprints [default: ~/.fleetctl/known_hosts]
  --strict-host-key-checking  verify SSH host keys [default: true]
  --tunnel=<host>             establish an SSH tunnel for communication with fleet and etcd
  --verbosity=<level>         log at a specified level of verbosity to stderr [default: 0]
`
	// parse command-line arguments
	args, err := docopt.Parse(usage, nil, true, "", true)
	if err != nil {
		os.Exit(2)
	}
	command := args["<command>"]
	//targets := args["<target>"].([]string)
	//setGlobalFlags(args)
	// construct a client

	// dispatch the command
	switch command {
	case "list":
		fmt.Println("In progress")
	case "scale":
		fmt.Println("In progress")
	case "start":
		fmt.Println("In progress")
	case "stop":
		fmt.Println("In progress")
	case "status":
		fmt.Println("In progress")
	case "install":
		cmdInstall()
	case "uninstall":
		cmdUninstall()
	default:
		fmt.Printf(usage)
		os.Exit(2)
	}
	if err != nil {
		os.Exit(2)
	}
}
