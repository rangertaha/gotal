// Copyright 2024 Rangertaha. All Rights Reserved.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/rangertaha/gotal/internal"
	"github.com/urfave/cli/v2"
)

func main() {
	cli.AppHelpTemplate = fmt.Sprintf(`%s
EXAMPLE:

   %s init
   %s new myproject

   %s fill  -p polygon -d 1m -s 2025-01-01 
   %s train -s 2025-01-01 -e 2025-01-02
   %s test  -s 2025-01-01 -e 2025-01-02
   %s live  -s 2025-01-01 -e 2025-01-02
   %s exec  -s 2025-01-01 -e 2025-01-02

AUTHOR:
   Rangertaha (rangertaha@gmail.com)
   
`, cli.AppHelpTemplate, internal.CLI, internal.CLI, internal.CLI, internal.CLI, internal.CLI, internal.CLI, internal.CLI)

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"V"},
		Usage:   "print the version",
	}

	app := &cli.App{
		Name:        "gota",
		Version:     internal.VERSION,
		Compiled:    time.Now(),
		Suggest:     true,
		HelpName:    "gota",
		Usage:       "Used to create, traipathn, test, and run financial trading bots",
		Description: "A framework for creating, training, testing, and running financial trading bots based on gotal (Go Technical Analysis Library).",
		UsageText:   fmt.Sprintf(`%s [global opts..] [command] [opts..]`, internal.CLI),
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Value: false,
				Usage: "Log debug messags for development",
				Action: func(ctx *cli.Context, v bool) error {
					return nil
				},
			},
		},
		Action: func(ctx *cli.Context) error {
			cli.ShowAppHelpAndExit(ctx, 0)
			return nil
		},
		Commands: []*cli.Command{
			&InitCmd,
			&NewCmd,
			&FillCmd,
			&TrainCmd,
			&TestCmd,
			&LiveCmd,
			&ExecCmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
