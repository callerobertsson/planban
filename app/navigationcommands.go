// Package app board navigation commands.
package app

func (pb *Planban) selectTaskAboveCommand() error {

	if pb.taskIndex > 0 {
		pb.taskIndex--
	}

	pb.RenderBoard()

	return nil
}

func (pb *Planban) selectTaskBelowCommand() error {

	max := len(pb.board.Stacks[pb.stackIndex].Tasks) - 1

	if pb.taskIndex < max {
		pb.taskIndex++
	}

	pb.RenderBoard()

	return nil
}

func (pb *Planban) selectNextStackCommand() error {

	i := pb.stackIndex + 1
	max := len(pb.board.Stacks)

	if i < max {
		pb.stackIndex = i
	}

	pb.adjustTaskIndex()

	pb.RenderBoard()

	return nil
}

func (pb *Planban) selectPrevStackCommand() error {

	if pb.stackIndex > 0 {
		pb.stackIndex--
	}

	pb.adjustTaskIndex()

	pb.RenderBoard()

	return nil
}

func (pb *Planban) adjustTaskIndex() {

	if pb.taskIndex >= len(pb.board.Stacks[pb.stackIndex].Tasks) {
		pb.taskIndex = len(pb.board.Stacks[pb.stackIndex].Tasks) - 1
	}

	if pb.taskIndex < 0 {
		pb.taskIndex = 0
	}
}
