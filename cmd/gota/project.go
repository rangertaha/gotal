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

	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/trader"

	"github.com/urfave/cli/v2"
)

var InitCmd = cli.Command{
	Name:                   "init",
	Usage:                  "Initialize user configs",
	Description:            "Initialize and create user config directory",
	UsageText:              "gota [g opts..] init",
	UseShortOptionHandling: true,
	Action: func(cCtx *cli.Context) error {
		trader.Init(cCtx.String("path"))
		return nil
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLE:
   %s init 
   %s init project/path
 
AUTHOR:
   Rangertaha (rangertaha@gmail.com)

`, cli.SubcommandHelpTemplate, internal.CLI, internal.CLI),
}

var NewCmd = cli.Command{
	Name:                   "new",
	Usage:                  "Create a new project directory",
	Description:            "Create a new project directory",
	UsageText:              "gota [g opts..] new [opts..] [dirname]",
	UseShortOptionHandling: true,
	Action: func(cCtx *cli.Context) error {
		cli.ShowSubcommandHelpAndExit(cCtx, 1)
		return nil
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLE:
   %s new myproject
 
AUTHOR:
   Rangertaha (rangertaha@gmail.com)
     
`, cli.SubcommandHelpTemplate, internal.CLI),
}
