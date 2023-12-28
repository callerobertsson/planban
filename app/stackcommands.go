// Package app Stack commands
package app

import (
	"fmt"
	"math"
	"strconv"

	"github.com/callerobertsson/wyrm"
)

func (pb *Planban) stackCommands() map[rune]*wyrm.Command {
	sort := 0
	nextSort := func() int {
		sort++
		return sort
	}
	return map[rune]*wyrm.Command{
		'a': {
			Sort:        nextSort(),
			Title:       "add",
			Description: "Add a new stack",
			Function:    pb.addStackCommand,
		},
		'e': {
			Sort:        nextSort(),
			Title:       "edit",
			Description: "Edit stack",
			Function:    pb.editStackCommand,
		},
		'D': {
			Sort:        nextSort(),
			Title:       "delete",
			Description: "Delete stack",
			Function:    pb.deleteStackCommand,
		},
		'L': {
			Sort:        nextSort(),
			Title:       "move right",
			Description: "Move stack right",
			Function:    pb.moveStackRightCommand,
		},
		'H': {
			Sort:        nextSort(),
			Title:       "move left",
			Description: "Move stack left",
			Function:    pb.moveStackLeftCommand,
		},
	}
}

func (pb *Planban) addStackCommand() error {

	name, err := wyrm.InputText("New Stack Name > ", "")
	switch {
	case err == wyrm.ErrEmpty:
		return nil
	case err != nil:
		return err
	}

	s := Stack{Name: name}

	r, err := wyrm.InputRune("Insert Stack Before or After [ba] > ")
	switch {
	case err == wyrm.ErrAbort:
		return nil
	case err != nil:
		return err
	}

	offs := 0
	if r == 'a' {
		offs = 1
	}

	err = pb.addStack(pb.stackIndex+offs, s)
	if err != nil {
		return err
	}

	pb.RenderBoard()
	return pb.saveBoardFile()
}

func (pb *Planban) deleteStackCommand() error {

	name := pb.board.Stacks[pb.stackIndex].Name
	prompt := fmt.Sprintf("WARNING: Do you want to delete the %q stack? [yes] > ", name)

	input, err := wyrm.InputText(prompt, "")
	if err != nil || input != "yes" {
		return nil
	}

	err = pb.deleteStack(pb.stackIndex)
	if err != nil {
		return err
	}

	if pb.stackIndex >= len(pb.board.Stacks) {
		pb.stackIndex = len(pb.board.Stacks) - 1
	}

	pb.RenderBoard()
	return pb.saveBoardFile()
}

func (pb *Planban) editStackCommand() error {

	prompt := fmt.Sprintf("Rename %q stack > ", pb.board.Stacks[pb.stackIndex].Name)
	name, err := wyrm.InputText(prompt, pb.board.Stacks[pb.stackIndex].Name)
	switch {
	case err == wyrm.ErrAbort:
		return nil
	case err == wyrm.ErrEmpty:
		return fmt.Errorf("stack must have a name")
	case err != nil:
		return err
	}

	pb.board.Stacks[pb.stackIndex].Name = name

	showMax, err := wyrm.InputInt("Max tasks to show > ", strconv.Itoa(pb.board.Stacks[pb.stackIndex].ShowMax), math.MaxInt)
	if err != nil {
		return err
	}

	pb.board.Stacks[pb.stackIndex].ShowMax = showMax

	pb.RenderBoard()
	return pb.saveBoardFile()
}

func (pb *Planban) addStack(i int, s Stack) error {
	if i < 0 || i > len(pb.board.Stacks) {
		return fmt.Errorf("out of range")
	}

	if i == len(pb.board.Stacks) {
		pb.board.Stacks = append(pb.board.Stacks, s)
		return nil
	}

	pb.board.Stacks = append(pb.board.Stacks[:i+1], pb.board.Stacks[i:]...)
	pb.board.Stacks[i] = s

	return nil
}

func (pb *Planban) deleteStack(i int) error {
	if i < 0 || i > len(pb.board.Stacks) {
		return fmt.Errorf("out of range")
	}

	pb.board.Stacks = append(pb.board.Stacks[:i], pb.board.Stacks[i+1:]...)

	return nil
}
