package main

import (
	"fmt"
	"github.com/WANNA959/sdsctl/pkg/cmds"
	"github.com/WANNA959/sdsctl/pkg/cmds/disk"
	"github.com/WANNA959/sdsctl/pkg/cmds/pool"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := cmds.NewApp()
	app.Commands = []*cli.Command{
		// pool commands
		pool.NewShowPoolCommand(),
		pool.NewCreatePoolCommand(),
		pool.NewStopPoolCommand(),
		pool.NewAutoStartPoolCommand(),
		pool.NewStartPoolCommand(),
		pool.NewStopPoolCommand(),

		// disk commands
		disk.NewShowDiskCommand(),
		disk.NewCreateDiskCommand(),
		disk.NewDeleteDiskCommand(),
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("error options: %s\n", err.Error())
		os.Exit(-1)
	}
}
