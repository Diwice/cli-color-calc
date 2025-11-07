package repl

import (
	"os"
	"fmt"
	//"bufio"
	"strings"
	"pkg/colorspace"
)

type cli_command struct {
	name        string
	description string
	Callback    func() error
}

func Clean_Input(input string) []string {
	return strings.Fields(strings.Trim(strings.ToLower(input), " "))
}

func help_cmd() error {
	fmt.Println("Usage :\n")

	for _, v := range Get_cmds() {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}

	return nil
}

func exit_cmd() error {
	fmt.Println("Closing the color calculator...")
	os.Exit(0)
	return fmt.Errorf("Failed to exit, try again")
}

func calc_cmd() error {
	new_rgb := colorspace.RGB_obj{RED: 10, GREEN: 10, BLUE: 10}
	fmt.Println(new_rgb)
	fmt.Println(new_rgb.To_cmyk())
	return nil
}

func Get_cmds() map[string]cli_command {
	return map[string]cli_command{
		"help":{
			name: "help",
			description: "Display a help message",
			Callback: help_cmd,
		},
		"exit":{
			name: "exit",
			description: "Exit the app",
			Callback: exit_cmd,
		},
		"calculate":{
			name: "calculate",
			description: "Input and calculate the colors in different colorspaces",
			Callback: calc_cmd,
		},
	}
}
