# Planban

A simple command line interface Kanban program.

## Synopsis

    planban [<path_to_planban_json_file>]

If `planban` is invoked without arguments it will create or open a
default board using the `planban.json` file for storage.

It can also be started with a JSON file path as first argument. If the
file doesn't exist it will be created and filled with the default board.

## Installation

Clone this repo and build using Golang. I have only tested on Linux at
the moment.

## Usage

Planban is based on (Wyrm)[https://github.com/callerobertsson/wyrm], a
single key command executor, so it's very fast to use, no long and
tedious command names to enter, only one key press most of the time.

Navigating between tasks is done using `hjkl` (vim-like) and moving
tasks around by `HJKL`.

The `?` key shows all commands availiable, including global commands.

Complete list of currently implemented commands:
```
Available command keys (recursive):
    [a] "add task" - Add task to selected stack
    [e] "edit task" - Edit selected task
    [D] "delete task" - Delete selected task
    [s] "stack" - Stack commands
        [a] "add" - Add a new stack
        [e] "edit" - Edit stack
        [D] "delete" - Delete stack
        [L] "move right" - Move stack right
        [H] "move left" - Move stack left
    [b] "board" - Board commands
    [c] "config" - Config commands
        [c] "color" - Toggle color output
        [i] "info" - Toggle board information text
        [d] "descriptions" - Toggle task descriptions
    [h] "prev stack" - Select previous stack
    [l] "next stack" - Select next stack
    [j] "task below" - Select task below
    [k] "task above" - Select task above
    [H] "move left" - Move task left
    [L] "move right" - Move task right
    [K] "task up" - Move task up
    [J] "task down" - Move task down
    [.] "refresh" - Redraw the board
Global command keys:
         [?] - display detailed help
    [ctrl-l] - clear screen
    [escape] - abort input
     [space] - print a new prompt
         [!] - execute shell command
         [q] - quit program nicely
   [newline] - print help on the current commands
```

The file `devboard.json` contains a Kanban with development tasks and
shows the current status of Planban development.

## Example usage

### Add task

It's as simple as navigating to the stack (colum) you want to add the task to
using `h` and `l`, then pressing `a` for "add task", fill out the name
and optional description. The new task should appear at the bottom of
the stack.

### Change config settings

Press `c` for config and then choose what setting to toggle. For example `c` for
color output and `d` to show or hide task descriptions.

The config is stored in the JSON-file so they will be set when you open
the board the next time.

Enjoy!
/Calle Robertsson


