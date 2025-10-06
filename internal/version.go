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
package internal

import "fmt"


const (
	NAME = "gotal"
	VERSION = "0.0.0"
	COMPILED = "2025-06-27"
	COMMIT = "35a7441"
)

const BANNER = `   
  ____  ___ _____  _     
 / ___|/ _ \_   _|/ \ 
| |  _| | | || | / _ \
| |_| | |_| || |/ ___ \
 \____|\___/ |_/_/   \_\
                  
 Go Trading Agent 
______________________________________________
COMMIT:  %s
AUTHOR:  Rangertaha
VERSION: v%s
DATE:    %s

`

func Banner() string {
	return fmt.Sprintf(BANNER, COMMIT, VERSION, COMPILED)
}

func PrintBanner() {
	fmt.Print(Banner())
}


// Init initializes the gotal library
func Init() {
	fmt.Println("Initializing gotal library")
}
