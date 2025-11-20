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
	"time"

	"github.com/rangertaha/gotal/internal"
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

var FillFlags = append(Flags, &cli.StringFlag{
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

var FillCmd = cli.Command{
	Name:                   "fill",
	Category:               "trading",
	Usage:                  "Download historical prices and metadata",
	Description:            "Backfill historical prices and metadata from registered data providers",
	UsageText:              fmt.Sprintf(`%s [g opts..] fill [opts..]`, internal.CLI),
	UseShortOptionHandling: true,
	Flags:                  FillFlags,
	Action: func(cCtx *cli.Context) error {
		start := cCtx.Timestamp("start")
		end := cCtx.Timestamp("end")
		duration := cCtx.Duration("duration")
		provider := cCtx.String("provider")

		if err := trader.Fill(*start, *end, duration, provider); err != nil {
			return err
		}
		return nil
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLE:
   %s fill -p polygon -d 1m -s 2025-01-01 -e 2025-01-02 

AUTHOR:
   Rangertaha (rangertaha@gmail.com)

`, cli.SubcommandHelpTemplate, internal.CLI),
}

var TrainCmd = cli.Command{
	Name:                   "train",
	Category:               "trading",
	Usage:                  "Train the strategy with historical prices",
	Description:            "Train the strategy with historical prices",
	UsageText:              fmt.Sprintf(`%s [g opts..] train [opts..]`, internal.CLI),
	UseShortOptionHandling: true,
	Flags:                  Flags,
	Action: func(cCtx *cli.Context) error {
		start := cCtx.Timestamp("start")
		end := cCtx.Timestamp("end")

		// Train a new model
		if err := trader.Train(*start, *end); err != nil {
			return err
		}
		return nil
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLE:
   %s train -s 2025-01-01 -e 2025-01-02 

AUTHOR:
   Rangertaha (rangertaha@gmail.com)

`, cli.SubcommandHelpTemplate, internal.CLI),
}

var TestCmd = cli.Command{
	Name:                   "test",
	Category:               "trading",
	Usage:                  "Test the strategy with historical prices",
	Description:            "Test the strategy with historical prices",
	UsageText:              fmt.Sprintf(`%s [g opts..] test [opts..]`, internal.CLI),
	UseShortOptionHandling: true,
	Flags:                  Flags,
	Action: func(cCtx *cli.Context) error {
		start := cCtx.Timestamp("start")
		end := cCtx.Timestamp("end")

		// Test trading
		if err := trader.Test(*start, *end); err != nil {
			return err
		}
		return nil
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLE:
   %s test -s 2025-01-01 -e 2025-01-02 

AUTHOR:
   Rangertaha (rangertaha@gmail.com)
     
`, cli.SubcommandHelpTemplate, internal.CLI),
}

var LiveCmd = cli.Command{
	Name:                   "live",
	Category:               "trading",
	Usage:                  "Live trading with fake money",
	Description:            "Live trading with fake money",
	UsageText:              fmt.Sprintf(`%s [g opts..] live [opts..]`, internal.CLI),
	UseShortOptionHandling: true,
	Flags:                  Flags,
	Action: func(cCtx *cli.Context) error {
		start := cCtx.Timestamp("start")
		end := cCtx.Timestamp("end")

		// Live trading
		if err := trader.Live(*start, *end); err != nil {
			return err
		}
		return nil
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLE:
   %s live -s 2025-01-01 -e 2025-01-02 

AUTHOR:
   Rangertaha (rangertaha@gmail.com)
     
`, cli.SubcommandHelpTemplate, internal.CLI),
}

var ExecCmd = cli.Command{
	Name:                   "exec",
	Category:               "trading",
	Usage:                  "Execute live trading with real money",
	Description:            "Execute live trading with real money",
	UsageText:              fmt.Sprintf(`%s [g opts..] exec [opts..]`, internal.CLI),
	UseShortOptionHandling: true,
	Flags:                  Flags,
	Action: func(cCtx *cli.Context) error {
		start := cCtx.Timestamp("start")
		end := cCtx.Timestamp("end")

		// Execute a new model
		if err := trader.Exec(*start, *end); err != nil {
			return err
		}
		return nil
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLE:
   %s exec -s 2025-01-01 -e 2025-01-02 
 
AUTHOR:
   Rangertaha (rangertaha@gmail.com)
     
`, cli.SubcommandHelpTemplate, internal.CLI),
}
