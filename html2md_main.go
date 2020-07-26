////////////////////////////////////////////////////////////////////////////
// Program: html2md
// Purpose: HTML to Markdown
// Authors: Tong Sun (c) 2020, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

//go:generate sh -v html2md_cliGen.sh

import (
	"fmt"
	"os"

	"github.com/mkideal/cli"
	clix "github.com/mkideal/cli/ext"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The OptsT type defines all the configurable options from cli.
type OptsT struct {
	Filei                       *clix.Reader
	Sel                         string
	OptHeadingStyle             string
	OptHorizontalRule           string
	OptBulletListMarker         string
	OptCodeBlockStyle           string
	OptFence                    string
	OptEmDelimiter              string
	OptStrongDelimiter          string
	OptLinkStyle                string
	OptLinkReferenceStyle       string
	PluginConfluenceAttachments bool
	PluginConfluenceCodeBlock   bool
	PluginFrontMatter           bool
	PluginGitHubFlavored        bool
	PluginStrikethrough         bool
	PluginTable                 bool
	PluginTaskListItems         bool
	PluginVimeoEmbed            bool
	PluginYoutubeEmbed          bool
	Verbose                     int
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	progname = "html2md"
	version  = "0.1.0"
	date     = "2020-07-25"

	rootArgv *rootT
	// Opts store all the configurable options
	Opts OptsT
)

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
func main() {
	cli.SetUsageStyle(cli.DenseNormalStyle) // left-right, for up-down, use ManualStyle
	//NOTE: You can set any writer implements io.Writer
	// default writer is os.Stdout
	if err := cli.Root(root).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("")
}

//==========================================================================
// Dumb root handler

func html2md(ctx *cli.Context) error {
	ctx.JSON(ctx.RootArgv())
	ctx.JSON(ctx.Argv())
	fmt.Println()

	return nil
}
