// Package app wyrm root command defintion
package app

import (
	"github.com/callerobertsson/wyrm"
)

// getRootCommand returns the wyrm command structure.
func (pb *Planban) getRootCommand() *wyrm.Command {
	sort := 0
	nextSort := func() int {
		sort++
		return sort
	}

	return &wyrm.Command{
		Title:       "planban",
		Description: "Planban Kanban",
		Commands: map[rune]*wyrm.Command{
			// Task commands
			'a': {
				Sort:        nextSort(),
				Title:       "add task",
				Description: "Add task to selected stack",
				Function:    pb.addTaskCommand,
			},
			'e': {
				Sort:        nextSort(),
				Title:       "edit task",
				Description: "Edit selected task",
				Function:    pb.editTaskCommand,
			},
			'D': {
				Sort:        nextSort(),
				Title:       "delete task",
				Description: "Delete selected task",
				Function:    pb.deleteTaskCommand,
			},
			// Stack commands
			's': {
				Sort:        nextSort(),
				Title:       "stack",
				Description: "Stack commands",
				Commands:    pb.stackCommands(),
			},
			// Board commands
			'b': {
				Sort:        nextSort(),
				Title:       "board",
				Description: "Board commands",
				Function:    pb.editBoardCommand,
			},
			// Config commands
			'c': {
				Sort:        nextSort(),
				Title:       "config",
				Description: "Config commands",
				Function:    pb.showConfigCommand,
				Commands:    pb.configCommands(),
			},
			// Navigation
			'h': {
				Sort:        nextSort(),
				Title:       "prev stack",
				Description: "Select previous stack",
				Function:    pb.selectPrevStackCommand,
			},
			'l': {
				Sort:        nextSort(),
				Title:       "next stack",
				Description: "Select next stack",
				Function:    pb.selectNextStackCommand,
			},
			'j': {
				Sort:        nextSort(),
				Title:       "task below",
				Description: "Select task below",
				Function:    pb.selectTaskBelowCommand,
			},
			'k': {
				Sort:        nextSort(),
				Title:       "task above",
				Description: "Select task above",
				Function:    pb.selectTaskAboveCommand,
			},
			// Moving Tasks
			'H': {
				Sort:        nextSort(),
				Title:       "move left",
				Description: "Move task left",
				Function:    pb.moveTaskLeftCommand,
			},
			'L': {
				Sort:        nextSort(),
				Title:       "move right",
				Description: "Move task right",
				Function:    pb.moveTaskRightCommand,
			},
			'K': {
				Sort:        nextSort(),
				Title:       "task up",
				Description: "Move task up",
				Function:    pb.moveTaskUpCommand,
			},
			'J': {
				Sort:        nextSort(),
				Title:       "task down",
				Description: "Move task down",
				Function:    pb.moveTaskDownCommand,
			},
			// Misc
			'.': {
				Sort:        nextSort(),
				Title:       "refresh",
				Description: "Redraw the board",
				Function:    func() error { pb.RenderBoard(); return nil },
			},
		},
	}
}
