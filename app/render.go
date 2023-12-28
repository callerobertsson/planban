// Package app render functions
package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/pterm/pterm"
	"golang.org/x/term"
)

func clear() {
	fmt.Printf(string("\x1b[2J\x1b[H"))
}

func enableColor(b bool) {
	pterm.DisableColor()
	if b {
		pterm.EnableColor()
	}
}

func header(f string, a ...any) {
	pterm.Printf(pterm.Green(fmt.Sprintf(f, a...)))
}

func subheader(f string, a ...any) {
	pterm.Printf(pterm.Gray(fmt.Sprintf(f, a...)))
}

func text(f string, a ...any) {
	pterm.Printf(pterm.White(fmt.Sprintf(f, a...)))
}

func (pb *Planban) renderConfig() {

	var yesno = func(b bool) string {
		if b {
			return "yes"
		}
		return "no"
	}

	fmt.Printf("  No [C]olors                  %s\n", yesno(pb.board.Config.NoColors))
	fmt.Printf("  Hide Board [I]nformation     %s\n", yesno(pb.board.Config.HideBoardInformation))
	fmt.Printf("  Hide Task [D]escriptions     %s\n", yesno(pb.board.Config.HideTaskDescriptions))
	fmt.Printf("  [E]dit in environment editor %s\n", yesno(pb.board.Config.UseEnvEditor))
}

func (pb *Planban) renderBoard() {

	enableColor(!pb.board.Config.NoColors)

	clear()

	header("Board: %s\n", pb.board.Name)
	if !pb.board.Config.HideBoardInformation && pb.board.Information != "" {
		subheader("%s\n\n", pb.board.Information)
	}

	width, _, err := term.GetSize(0)
	if err != nil {
		return
	}

	colWidth := width / len(pb.board.Stacks)

	tbl := tablewriter.NewWriter(os.Stdout)
	applyDefaultTableSettings(tbl)
	tbl.SetColWidth(colWidth)

	// Header: stack names
	header := []string{}
	for i, s := range pb.board.Stacks {
		name := blue(strings.ToUpper(s.Name))
		if i == pb.stackIndex {
			name = red("> " + strings.ToUpper(s.Name) + " <")
		}
		header = append(header, name)
	}
	tbl.SetHeader(header)

	// Content: stack tasks
	stacks := []string{}
	for i, s := range pb.board.Stacks {
		stacks = append(stacks, pb.stackTasks(i, s.Tasks))
	}
	tbl.Append(stacks)

	tbl.Render()

	fmt.Println("")
}

func (pb *Planban) stackTasks(index int, ts []Task) string {
	tblStr := &strings.Builder{}
	tbl := tablewriter.NewWriter(tblStr)
	applyDefaultTableSettings(tbl)

	for i, t := range ts {
		name := green(t.Name)
		if index == pb.stackIndex && i == pb.taskIndex {
			name = yellow("> " + t.Name + " <")
		}
		tbl.Append([]string{name})
		if !pb.board.Config.HideTaskDescriptions && t.Description != "" {
			tbl.Append([]string{gray(t.Description)})
		}
	}

	tbl.Render()

	return tblStr.String()
}

func applyDefaultTableSettings(tbl *tablewriter.Table) {
	tbl.SetAlignment(tablewriter.ALIGN_LEFT)
	tbl.SetBorder(false)
	tbl.SetAutoWrapText(false)
	tbl.SetAutoFormatHeaders(false)
	tbl.SetTablePadding("")
	tbl.SetColumnSeparator("|")
	tbl.SetRowSeparator("-")
	tbl.SetCenterSeparator("+")
}
