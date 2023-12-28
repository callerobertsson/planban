// Package app task commands
package app

import (
	"fmt"

	"github.com/callerobertsson/wyrm"
)

func (pb *Planban) addTaskCommand() error {

	t, err := pb.editTask(Task{})
	if err != nil {
		return err
	}

	pb.board.Stacks[pb.stackIndex].Tasks = append(pb.board.Stacks[pb.stackIndex].Tasks, Task{t.Name, t.Description})

	pb.RenderBoard()
	return pb.saveBoardFile()
}

func (pb *Planban) editTaskCommand() error {

	t, err := pb.editTask(pb.board.Stacks[pb.stackIndex].Tasks[pb.taskIndex])
	if err != nil {
		return err
	}

	pb.board.Stacks[pb.stackIndex].Tasks[pb.taskIndex] = t

	pb.RenderBoard()
	return pb.saveBoardFile()
}

func (pb *Planban) deleteTaskCommand() error {

	name := pb.board.Stacks[pb.stackIndex].Tasks[pb.taskIndex].Name
	prompt := fmt.Sprintf("WARNING: Do you want to delete the %q task? [yes] > ", name)

	input, err := wyrm.InputText(prompt, "")
	if err != nil || input != "yes" {
		return nil
	}

	err = pb.deleteTask(pb.stackIndex, pb.taskIndex)
	if err != nil {
		return err
	}

	pb.adjustTaskIndex()
	pb.RenderBoard()
	return pb.saveBoardFile()
}

func (pb Planban) editTask(t Task) (Task, error) {

	name, err := wyrm.InputText("Task Name > ", t.Name)
	switch {
	case err == wyrm.ErrEmpty:
		return t, fmt.Errorf("task name is mandatory")
	case err != nil:
		return t, err
	}

	if pb.board.Config.UseEnvEditor {
		desc, err := editInEnvEditor("planban-task-description-", t.Description)
		if err != nil {
			return t, err
		}
		return Task{name, desc}, nil
	}

	desc, err := wyrm.InputText("Task Description > ", t.Description)
	switch {
	case err == wyrm.ErrEmpty: // Description is not mandatory
	case err != nil:
		return t, err
	}

	return Task{name, desc}, nil
}

func (pb *Planban) addTask(si, ti int, t Task) error {
	if si < 0 || si > len(pb.board.Stacks) {
		return fmt.Errorf("stack out of range")
	}

	if ti < 0 || ti > len(pb.board.Stacks[si].Tasks) {
		return fmt.Errorf("task out of range")
	}

	if ti == len(pb.board.Stacks[si].Tasks) {
		pb.board.Stacks[si].Tasks = append(pb.board.Stacks[si].Tasks, t)
		return nil
	}

	pb.board.Stacks[si].Tasks = append(pb.board.Stacks[si].Tasks[:ti+1], pb.board.Stacks[si].Tasks[ti:]...)
	pb.board.Stacks[si].Tasks[ti] = t

	return nil
}

func (pb *Planban) deleteTask(si, ti int) error {
	if si < 0 || si > len(pb.board.Stacks) {
		return fmt.Errorf("stack out of range")
	}

	if ti < 0 || ti > len(pb.board.Stacks[si].Tasks) {
		return fmt.Errorf("task out of range")
	}

	pb.board.Stacks[si].Tasks = append(pb.board.Stacks[si].Tasks[:ti], pb.board.Stacks[si].Tasks[ti+1:]...)

	return nil
}
