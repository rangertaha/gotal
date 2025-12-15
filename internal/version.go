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
	CLI = "gota"
	VERSION = "0.0.0"
	COMPILED = "2025-06-27"
	COMMIT = "35a7441"
)

const CLI_BANNER = `   
  ____  ___ _____  _    
 / ___|/ _ \_   _|/ \   
| |  _| | | || | / _ \  
| |_| | |_| || |/ ___ \ 
 \____|\___/ |_/_/   \_\
                
 Go Technical Analysis
______________________________________________
COMMIT:  %s
AUTHOR:  Rangertaha
VERSION: v%s
DATE:    %s
`

const LIB_BANNER = `   
  ____  ___ _____  _    _     
 / ___|/ _ \_   _|/ \  | |    
| |  _| | | || | / _ \ | |    
| |_| | |_| || |/ ___ \| |___ 
 \____|\___/ |_/_/   \_\_____|
              
 Go Technical Analysis Library
______________________________________________
COMMIT:  %s
AUTHOR:  Rangertaha
VERSION: v%s
DATE:    %s

`

func CliBanner() string {
	return fmt.Sprintf(CLI_BANNER, COMMIT, VERSION, COMPILED)
}

func LibBanner() string {
	return fmt.Sprintf(LIB_BANNER, COMMIT, VERSION, COMPILED)
}

func PrintBanner() {
	fmt.Print(CliBanner())
}


// Init initializes the gotal library
func Init() {
	fmt.Println("Initializing gotal library")
}
