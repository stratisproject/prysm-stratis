package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/stratisproject/prysm-stratis/cmd/prysmctl/checkpointsync"
	"github.com/stratisproject/prysm-stratis/cmd/prysmctl/db"
	"github.com/stratisproject/prysm-stratis/cmd/prysmctl/p2p"
	"github.com/stratisproject/prysm-stratis/cmd/prysmctl/testnet"
	"github.com/stratisproject/prysm-stratis/cmd/prysmctl/validator"
	"github.com/stratisproject/prysm-stratis/cmd/prysmctl/weaksubjectivity"
	"github.com/urfave/cli/v2"
)

var prysmctlCommands []*cli.Command

func main() {
	app := &cli.App{
		Commands: prysmctlCommands,
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	prysmctlCommands = append(prysmctlCommands, checkpointsync.Commands...)
	prysmctlCommands = append(prysmctlCommands, db.Commands...)
	prysmctlCommands = append(prysmctlCommands, p2p.Commands...)
	prysmctlCommands = append(prysmctlCommands, testnet.Commands...)
	prysmctlCommands = append(prysmctlCommands, weaksubjectivity.Commands...)
	prysmctlCommands = append(prysmctlCommands, validator.Commands...)
}
