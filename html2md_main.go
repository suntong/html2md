////////////////////////////////////////////////////////////////////////////
// Program: html2md
// Purpose: HTML to Markdown
// Authors: Tong Sun (c) 2020-2024, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

//go:generate sh -v html2md_cliGen.sh

import (
	"fmt"
	"os"

	"github.com/mkideal/cli"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	progname = "html2md"
	version  = "1.6.0"
	date     = "2025-02-28"

	rootArgv *rootT
)

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
func main() {
	cli.SetUsageStyle(cli.DenseNormalStyle)
	if err := cli.Root(root).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("")
}
