// Package app board commands
package app

import (
	"fmt"

	"github.com/callerobertsson/wyrm"
)

func (pb *Planban) editBoardCommand() error {

	name, err := wyrm.InputText("Board Name > ", pb.board.Name)
	switch {
	case err == wyrm.ErrAbort:
		return nil
	case err == wyrm.ErrEmpty:
		return fmt.Errorf("board name is mandatory")
	case err != nil:
		return err
	}

	pb.board.Name = name

	info := pb.board.Information

	if pb.board.Config.UseEnvEditor {
		maybeInfo, err := editInEnvEditor("planban-board-info-", pb.board.Information)
		if err != nil {
			return err
		}
		info = maybeInfo
	} else {
		maybeInfo, err := wyrm.InputText("Board Info > ", pb.board.Information)
		switch {
		case err == wyrm.ErrAbort:
			return nil
		case err == wyrm.ErrEmpty:
		case err != nil:
			return err
		}
		info = maybeInfo
	}

	pb.board.Information = info

	pb.RenderBoard()
	return pb.saveBoardFile()
}
