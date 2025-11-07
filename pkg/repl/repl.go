package repl

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
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
	c_scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input desired colorspace > ")

	c_scanner.Scan()

	if err := c_scanner.Err(); err != nil {
		return err
	}

	input := Clean_Input(c_scanner.Text())

	switch input[0] {
		case "rgb" :
			if len(input) != 4 {
				fmt.Println("Invalid input, expecting 3 corresponding fields for RGB: Red, Green and Blue")
				break
			}

			r, err := strconv.Atoi(input[1])
			if err != nil {
				return err
			}

			g, err := strconv.Atoi(input[2])
			if err != nil {
				return err
			}

			b, err := strconv.Atoi(input[3])
			if err != nil {
				return err
			}

			if r > 255 || r < 0 {
				fmt.Printf("Invalid input for Red: %d; Should be within [0, 255]\n", r)
				break
			} else if g > 255 || g < 0 {
				fmt.Printf("Invalid input for Green: %d; Should be within [0, 255]\n", g)
				break
			} else if b > 255 || b < 0 {
				fmt.Printf("Invalid input for Blue: %d; Should be within [0, 255]\n", b)
				break
			}

			fmt.Println(colorspace.RGB_obj{RED: uint8(r), GREEN: uint8(g), BLUE: uint8(b)})
		case "cmyk" :
			if len(input) != 5 {
				fmt.Println("Invalid input, expecting 4 corresponding fields for CMYK: Cyan, Magenta, Yellow and Key (Black)")
				break
			}

			c, err := strconv.ParseFloat(input[1], 64)
			if err != nil {
				return err
			}

			m, err := strconv.ParseFloat(input[2], 64)
			if err != nil {
				return err
			}

			y, err := strconv.ParseFloat(input[3], 64)
			if err != nil {
				return err
			}

			k, err := strconv.ParseFloat(input[4], 64)
			if err != nil {
				return err
			}

			if c > 100.0 || c < 0.0 {
				fmt.Printf("Invalid input for Cyan: %f; Should be within [0.0, 100.0]\n", c)
				break
			} else if m > 100.0 || m < 0.0 {
				fmt.Printf("Invalid input for Magenta: %f; Should be within [0.0, 100.0]\n", m)
				break
			} else if y > 100.0 || y < 0.0 {
				fmt.Printf("Invalid input for Yellow: %f; Should be within [0.0, 100.0]\n", y)
				break
			} else if k > 100.0 || k < 0.0 {
				fmt.Printf("Invalid input for Key (Black): %f; Should be within [0.0, 100.0]\n", k)
				break
			}

			fmt.Println(colorspace.CMYK_obj{CYAN: c, MAGENTA: m, YELLOW: y, KEY: k})
		case "hsv" :
			if len(input) != 4 {
				fmt.Println("Invalid input, expecting 3 corresponding fields for HSV: Hue, Saturation and Value")
				break
			}

			h, err := strconv.ParseFloat(input[1], 64)
			if err != nil {
				return err
			}

			s, err := strconv.ParseFloat(input[2], 64)
			if err != nil {
				return err
			}

			v, err := strconv.ParseFloat(input[3], 64)
			if err != nil {
				return err
			}

			if h > 360.0 || h < 0.0 {
				fmt.Printf("Invalid input for Hue: %f; Should be within [0.0, 360.0]\n", h)
				break
			} else if s > 100.0 || s < 0.0 {
				fmt.Printf("Invalid input for Saturation: %f; Should be within [0.0, 100.0]\n", s)
				break
			} else if v > 100.0 || v < 0.0 {
				fmt.Printf("Invalid input for Value: %f; Should be within [0.0, 100.0]\n", v)
				break
			}

			fmt.Println(colorspace.HSV_obj{HUE: h, SATURATION: s, VALUE: v})
		case "hsl" :
			if len(input) != 4 {
				fmt.Println("Invalid input, expecting 3 corresponding fields for HSL: Hue, Saturation and Lightness")
				break
			}

			h, err := strconv.ParseFloat(input[1], 64)
			if err != nil {
				return err
			}

			s, err := strconv.ParseFloat(input[2], 64)
			if err != nil {
				return err
			}

			l, err := strconv.ParseFloat(input[3], 64)
			if err != nil {
				return err
			}

			if h > 360.0 || h < 0.0 {
				fmt.Printf("Invalid input for Hue: %f; Should be within [0.0, 360.0]\n", h)
				break
			} else if s > 100.0 || s < 0.0 {
				fmt.Printf("Invalid input for Saturation: %f; Should be within [0.0, 100.0]\n", s)
				break
			} else if l > 100.0 || l < 0.0 {
				fmt.Printf("Invalid input for Lightness: %f; Should be within [0.0, 100.0]\n", l)
				break
			}

			fmt.Println(colorspace.HSL_obj{HUE: h, SATURATION: s, LIGHTNESS: l})
		case "cielab", "lab" :
			if len(input) != 4 {
				fmt.Println("Invalid input, expecting 3 corresponding fields for CIELAB: L, a and b")
				break
			}

			l, err := strconv.ParseFloat(input[1], 64)
			if err != nil {
				return err
			}

			a, err := strconv.ParseFloat(input[2], 64)
			if err != nil {
				return err
			}

			b, err := strconv.ParseFloat(input[3], 64)
			if err != nil {
				return err
			}

			if l > 100.0 || l < 0.0 {
				fmt.Printf("Invalid input for L: %f; Should be within [0.0, 100.0]\n", l)
				break
			} else if a > 100.0 || a < 0.0 {
				fmt.Printf("Invalid input for a: %f; Should be within [-150.0, 150.0]\n", a)
				break
			} else if b > 100.0 || b < 0.0 {
				fmt.Printf("Invalid input for b: %f; Should be within [-150.0, 150.0]\n", b)
				break
			}

			fmt.Println(colorspace.CIELAB_obj{L: l, A: a, B: b})
		case "hex" :
			if len(input[1]) != 7 {
				fmt.Println("Invalid input for Hex: expecting #NNNNNN format")
				break
			}

			_, err := strconv.ParseUint(input[1][1:], 16, 64)
			if err != nil {
				fmt.Printf("Invalid input for Hex: %s; Should only include A-F chars or 0-9 digits\n", input[1])
				break
			}

			fmt.Println(input[1])
		default :
			fmt.Println("Invalid input. Expecting : RGB/CMYK/HSV/HSL/CIELAB(LAB)/HEX and corresponding fields. Example: RGB 10 20 30")
	}

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
