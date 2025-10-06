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

	"github.com/rangertaha/gotal"
	"github.com/rangertaha/gotal/internal/trader"
	"github.com/urfave/cli/v2"
)

var Flags = []cli.Flag{
	&cli.TimestampFlag{
		Name:    "start",
		Usage:   "start time `[TIME]`",
		Aliases: []string{"s"},
		Layout:  "2006-01-02",
		Value:   cli.NewTimestamp(time.Now().AddDate(-10, 0, 0)),
	},
	&cli.TimestampFlag{
		Name:    "end",
		Aliases: []string{"e"},
		Layout:  "2006-01-02",
		Usage:   "end time `[TIME]`",
		Value:   cli.NewTimestamp(time.Now()),
	},
}

var BackfillFlags = append(Flags, &cli.StringFlag{
	Name:    "provider",
	Usage:   "data provider to use",
	Aliases: []string{"p"},
	Value:   "",
}, &cli.DurationFlag{
	Name:    "duration",
	Usage:   "duration of data points to download",
	Aliases: []string{"d"},
	Value:   time.Duration(1 * time.Minute),
})

var BackfillCmd = cli.Command{
	Name: "fill",
	// Aliases:                []string{"bf", "fill"},
	Usage:                  "Download historical prices and metadata",
	Description:            "Backfill historical prices and metadata from registered data providers",
	UsageText:              "trader fill [opts..]",
	UseShortOptionHandling: true,
	Flags:                  BackfillFlags,
	Action: func(cCtx *cli.Context) error {
		// cli.ShowSubcommandHelpAndExit(cCtx, 1)
		start := cCtx.Timestamp("start")
		end := cCtx.Timestamp("end")
		duration := cCtx.Duration("duration")
		provider := cCtx.String("provider")

		if err := trader.Backfill(*start, *end, duration, provider); err != nil {
			return err
		}
		return nil
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
	
	EXAMPLE:

    trader fill -p polygon -d 1m -s 2025-01-01 -e 2025-01-02 

    trader fill -p polygon -d 1m -s 2025-01-01 -e 2025-01-02

    trader fill -p polygon -d 1m -s 2025-01-01 -e 2025-01-02

    trader fill -p polygon -d 1m -s 2025-01-01 -e 2025-01-02
 

AUTHOR:
   Rangertaha (rangertaha@gmail.com)
     
     `, cli.SubcommandHelpTemplate),
}

var TrainCmd = cli.Command{
	Name: "train",
	// Aliases:                []string{"trn"},
	Usage:                  "Train a new model",
	Description:            "Train a new model",
	UsageText:              "trader [g opts..] train [opts..] [name]",
	UseShortOptionHandling: true,
	Flags:                  Flags,
	Action: func(cCtx *cli.Context) error {
		start := cCtx.Timestamp("start")
		end := cCtx.Timestamp("end")
		if err := trader.Train(*start, *end); err != nil {
			return err
		}
		return nil
	},
}

var TestCmd = cli.Command{
	Name: "test",
	// Aliases:                []string{"tst"},
	Usage:                  "Test a new model",
	Description:            "Test a new model",
	UsageText:              "trader [g opts..] test [opts..] [name]",
	UseShortOptionHandling: true,
	Flags:                  Flags,
	Action: func(cCtx *cli.Context) error {
		cli.ShowSubcommandHelpAndExit(cCtx, 1)
		return nil
	},
}

var LiveCmd = cli.Command{
	Name: "live",
	// Aliases:                []string{"l"},
	Usage:                  "Live trade a new model",
	Description:            "Live trade a new model",
	UsageText:              "trader [g opts..] live [opts..] [name]",
	UseShortOptionHandling: true,
	Flags:                  Flags,
	Action: func(cCtx *cli.Context) error {
		cli.ShowSubcommandHelpAndExit(cCtx, 1)
		return nil
	},
}

var ExecCmd = cli.Command{
	Name: "exec",
	// Aliases:                []string{"e", "run"},
	Usage:                  "Execute a new model",
	Description:            "Execute a new model",
	UsageText:              "trader [g opts..] exec [opts..] [name]",
	UseShortOptionHandling: true,
	Flags:                  Flags,
	Action: func(cCtx *cli.Context) error {
		cli.ShowSubcommandHelpAndExit(cCtx, 1)
		return nil
	},
}

func main() {
	cli.AppHelpTemplate = fmt.Sprintf(`%s
EXAMPLE:

	trader create myproject

    trader fill -p polygon -d 1m -s 2025-01-01 

    trader train -s 2025-01-01 -e 2025-01-02
    trader test -s 2025-01-01 -e 2025-01-02
    trader live -s 2025-01-01 -e 2025-01-02
    trader exec -s 2025-01-01 -e 2025-01-02

AUTHOR:
   Rangertaha (rangertaha@gmail.com)

     
     `, cli.AppHelpTemplate)

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"V"},
		Usage:   "print the version",
	}

	app := &cli.App{
		Name:        "trader",
		Version:     gotal.VERSION,
		Compiled:    time.Now(),
		Suggest:     true,
		HelpName:    "trader",
		Usage:       "creating, training, testing, and running trading bots",
		Description: "",
		UsageText:   "trader [global opts..] [command] [opts..]",
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
			&CreateCmd,
			&BackfillCmd,
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
