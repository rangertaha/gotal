// Copyright (C) 2026 Rangertaha <rangertaha@gmail.com>
// SPDX-License-Identifier: GPL-3.0-or-later
package internal

import "fmt"

const (
	NAME     = "gotal"
	CLI      = "gota"
	VERSION  = "0.0.0"
	COMPILED = "2025-06-27"
	COMMIT   = "35a7441"
	BANNER   = `   
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
)

var Banner = fmt.Sprintf(BANNER, COMMIT, VERSION, COMPILED)

func PrintBanner() {
	if Banner != "" {
		fmt.Print(Banner)
	}
}
