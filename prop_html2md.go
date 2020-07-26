////////////////////////////////////////////////////////////////////////////
// Program: html2md
// Purpose: HTML to Markdown
// Authors: Tong Sun (c) 2020, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"

	"github.com/mkideal/cli"
)

//==========================================================================
// root handler

func html2md(ctx *cli.Context) error {
	ctx.JSON(ctx.RootArgv())
	ctx.JSON(ctx.Argv())
	fmt.Println()

	return nil
}
