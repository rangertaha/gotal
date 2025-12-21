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

	"github.com/urfave/cli/v2"
)

var CreateCmd = cli.Command{
	Name:                   "create",
	Usage:                  "Create a new project",
	Description:            "Create a new project",
	UsageText:              "gota [g opts..] create [opts..] [projectname]",
	UseShortOptionHandling: true,
	Action: func(cCtx *cli.Context) error {
		fmt.Println("Creating project...", cCtx.String("name"))
		return nil
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLE:
   %s create myproject

 
AUTHOR:
   Rangertaha (rangertaha@gmail.com)

`, cli.SubcommandHelpTemplate, internal.CLI, internal.CLI),
}

var UpdateCmd = cli.Command{
	Name:                   "update",
	Usage:                  "Update a project",
	Description:            "Update a project",
	UsageText:              "gota [g opts..] update [opts..] [projectname]",
	UseShortOptionHandling: true,
	Action: func(cCtx *cli.Context) error {
		fmt.Println("Updating project...", cCtx.String("name"))
		return nil
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLE:
   %s update              # update current project
   %s update myproject    # update specific project directory
 
AUTHOR:
   Rangertaha (rangertaha@gmail.com)
     
`, cli.SubcommandHelpTemplate, internal.CLI),
}
