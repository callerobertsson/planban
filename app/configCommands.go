// Package app configuration commands
package app

import (
	"fmt"

	"github.com/callerobertsson/wyrm"
)

func (pb *Planban) configCommands() map[rune]*wyrm.Command {
	sort := 0
	nextSort := func() int {
		sort++
		return sort
	}
	return map[rune]*wyrm.Command{
		'c': {
			Sort:        nextSort(),
			Title:       "color",
			Description: "Toggle color output",
			Function:    pb.toggleNoColorCommand,
		},
		'i': {
			Sort:        nextSort(),
			Title:       "info",
			Description: "Toggle board information text",
			Function:    pb.toggleBoardDescriptionCommand,
		},
		'd': {
			Sort:        nextSort(),
			Title:       "descriptions",
			Description: "Toggle task descriptions",
			Function:    pb.toggleTaskDescriptionsCommand,
		},
		'e': {
			Sort:        nextSort(),
			Title:       "editor",
			Description: "Toggle edit in environment editor",
			Function:    pb.toggleUseEnvEditor,
		},
	}
}

func (pb *Planban) showConfigCommand() error {

	clear()

	fmt.Println(green(fmt.Sprintf("Board %v configuration\n", pb.board.Name)))
	pb.renderConfig()

	fmt.Println("")

	return nil
}

func (pb *Planban) toggleNoColorCommand() error {

	pb.board.Config.NoColors = !pb.board.Config.NoColors

	err := pb.saveBoardFile()
	pb.renderBoard()
	return err
}

func (pb *Planban) toggleBoardDescriptionCommand() error {

	pb.board.Config.HideBoardInformation = !pb.board.Config.HideBoardInformation

	err := pb.saveBoardFile()
	pb.renderBoard()
	return err
}

func (pb *Planban) toggleTaskDescriptionsCommand() error {

	pb.board.Config.HideTaskDescriptions = !pb.board.Config.HideTaskDescriptions

	err := pb.saveBoardFile()
	pb.renderBoard()
	return err
}

func (pb *Planban) toggleUseEnvEditor() error {

	pb.board.Config.UseEnvEditor = !pb.board.Config.UseEnvEditor

	err := pb.saveBoardFile()
	pb.renderBoard()
	return err
}
