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

	"github.com/urfave/cli/v2"
)

var InitCmd = cli.Command{
	Name:                   "init",
	// Aliases:                []string{"p"},
	Usage:                  "Initialize user config directory",
	Description:            "Initialize and create user config directory",
	UsageText:              "trader [g opts..] init",
	UseShortOptionHandling: true,
	// Flags:                  Flags,
	Action: func(cCtx *cli.Context) error {
		cli.ShowSubcommandHelpAndExit(cCtx, 1)
		return nil
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLE:

    trader create myproject
 

AUTHOR:
   Rangertaha (rangertaha@gmail.com)
     
     `, cli.SubcommandHelpTemplate),
}


var CreateCmd = cli.Command{
	Name:                   "create",
	// Aliases:                []string{"p"},
	Usage:                  "Create a new project directory",
	Description:            "Create a new project directory",
	UsageText:              "trader [g opts..] create [opts..] [name]",
	UseShortOptionHandling: true,
	// Flags:                  Flags,
	Action: func(cCtx *cli.Context) error {
		cli.ShowSubcommandHelpAndExit(cCtx, 1)
		return nil
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLE:

    trader create myproject
 

AUTHOR:
   Rangertaha (rangertaha@gmail.com)
     
     `, cli.SubcommandHelpTemplate),
}
