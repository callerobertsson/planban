// Package app implements the Planban application
package app

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/callerobertsson/wyrm"
)

// Planban defines the application type
type Planban struct {
	file       string
	board      *Board
	wyrm       *wyrm.Wyrm
	stackIndex int
	taskIndex  int
}

var initialBoard = Board{
	Name:        "Planban Board",
	Information: "Board information text",
	Stacks: []Stack{
		{
			Name: "Backlog",
			Tasks: []Task{
				Task{
					Name:        "Task Example",
					Description: "A description",
				},
			},
		},
		{
			Name:  "Doing",
			Tasks: []Task{},
		},
		{
			Name:  "Done",
			Tasks: []Task{},
		},
	},
}

// New creates a new Planban instance
func New(f string) (*Planban, error) {
	pb := Planban{file: f}

	if _, err := os.Stat(f); os.IsNotExist(err) {
		err = writeBoardToFile(f, &initialBoard)
		if err != nil {
			return nil, err
		}
	}

	pb.wyrm = wyrm.New(pb.getRootCommand())

	err := pb.loadBoardFile()

	return &pb, err
}

// Run starts the interactive Planban session
func (pb *Planban) Run() error {

	pb.RenderBoard()

	pb.wyrm.Run()

	return nil
}

func (pb *Planban) saveBoardFile() error {

	return writeBoardToFile(pb.file, pb.board)
}

func (pb *Planban) loadBoardFile() error {

	data, err := ioutil.ReadFile(pb.file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &pb.board)

	return err
}

func writeBoardToFile(f string, b *Board) error {

	json, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(f, []byte(json), 0664)
}
