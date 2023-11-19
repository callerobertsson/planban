// Package app movement commands.
package app

import "fmt"

func (pb *Planban) moveStackRightCommand() error {

	if pb.stackIndex > len(pb.board.Stacks)-2 {
		pb.renderBoard()
		return nil
	}

	s := pb.board.Stacks[pb.stackIndex]

	err := pb.deleteStack(pb.stackIndex)
	if err != nil {
		return err
	}

	pb.stackIndex++

	err = pb.addStack(pb.stackIndex, s)
	if err != nil {
		return err
	}

	pb.renderBoard()
	return pb.saveBoardFile()
}

func (pb *Planban) moveStackLeftCommand() error {

	if pb.stackIndex < 1 {
		pb.renderBoard()
		return nil
	}

	s := pb.board.Stacks[pb.stackIndex]

	err := pb.deleteStack(pb.stackIndex)
	if err != nil {
		return err
	}

	pb.stackIndex--
	fmt.Println("stackindex", pb.stackIndex)

	err = pb.addStack(pb.stackIndex, s)
	if err != nil {
		return err
	}

	pb.renderBoard()
	return pb.saveBoardFile()
}

func (pb *Planban) moveTaskUpCommand() error {

	if pb.taskIndex < 1 {
		pb.renderBoard()
		return nil
	}

	t := pb.board.Stacks[pb.stackIndex].Tasks[pb.taskIndex]

	err := pb.deleteTask(pb.stackIndex, pb.taskIndex)
	if err != nil {
		return err
	}

	pb.taskIndex--

	err = pb.addTask(pb.stackIndex, pb.taskIndex, t)
	if err != nil {
		return err
	}

	pb.renderBoard()
	return pb.saveBoardFile()
}

func (pb *Planban) moveTaskDownCommand() error {

	if pb.taskIndex > len(pb.board.Stacks[pb.stackIndex].Tasks)-2 {
		pb.renderBoard()
		return nil
	}

	t := pb.board.Stacks[pb.stackIndex].Tasks[pb.taskIndex]

	err := pb.deleteTask(pb.stackIndex, pb.taskIndex)
	if err != nil {
		return err
	}

	pb.taskIndex++

	err = pb.addTask(pb.stackIndex, pb.taskIndex, t)
	if err != nil {
		return err
	}

	pb.renderBoard()
	return pb.saveBoardFile()
}

func (pb *Planban) moveTaskLeftCommand() error {

	if pb.stackIndex < 1 {
		pb.renderBoard()
		return nil
	}

	t := pb.board.Stacks[pb.stackIndex].Tasks[pb.taskIndex]

	err := pb.deleteTask(pb.stackIndex, pb.taskIndex)
	if err != nil {
		return err
	}

	pb.taskIndex = 0
	pb.stackIndex--

	err = pb.addTask(pb.stackIndex, pb.taskIndex, t)
	if err != nil {
		return err
	}

	pb.renderBoard()
	return pb.saveBoardFile()
}

func (pb *Planban) moveTaskRightCommand() error {

	if pb.stackIndex > len(pb.board.Stacks)-2 {
		pb.renderBoard()
		return nil
	}

	t := pb.board.Stacks[pb.stackIndex].Tasks[pb.taskIndex]

	err := pb.deleteTask(pb.stackIndex, pb.taskIndex)
	if err != nil {
		return err
	}

	pb.taskIndex = 0
	pb.stackIndex++

	err = pb.addTask(pb.stackIndex, pb.taskIndex, t)
	if err != nil {
		return err
	}

	pb.renderBoard()
	return pb.saveBoardFile()
}
