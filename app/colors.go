// Package app color wrapper
package app

import "github.com/pterm/pterm"

var white = func(s string) string { return pterm.White(s) }
var gray = func(s string) string { return pterm.Gray(s) }
var yellow = func(s string) string { return pterm.Yellow(s) }
var green = func(s string) string { return pterm.Green(s) }
var blue = func(s string) string { return pterm.Blue(s) }
var red = func(s string) string { return pterm.Red(s) }
