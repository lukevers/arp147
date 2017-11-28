package terminal

import (
	"engo.io/ecs"
	"engo.io/engo"
	"errors"
	"github.com/lukevers/arp147/input"
	"log"
	"strings"
)

// TODO
type line struct {
	text []string

	// TODO
	Locked bool
}

type Shell struct {
	world   *ecs.World
	Program Program

	lines map[int]*line
	line  int
}

func (s *Shell) HandleKey(key engo.Key, mods *input.Modifiers) {
	// If there is a program currently running that's not the shell, send the
	// key event to the program to handle instead of handling it here.
	if s.Program != nil {
		s.Program.HandleKey(key, mods)
		return
	}

	if s.lines[s.line] == nil {
		s.lines[s.line] = &line{}
	}

	// TODO:
	//   - signals
	//     - ^C
	//   - file system
	//     - users
	//     - permissions
	//     - files
	//     - executable files `./some-file` that run in shell
	//   - Pipe output:
	//     - >
	//     - >>
	//     - |
	//   - SDK (for program development)
	//     - flag/argument parsing
	//     - os package for file system, ...
	//     - ui toolsets

	length := len(s.lines[s.line].text)

	switch key {
	case engo.Backspace:
		if length > 0 {
			s.lines[s.line].text = s.lines[s.line].text[0 : length-1]
		}
	case engo.Enter:
		s.lines[s.line].Locked = true
		s.line++

		cmd, args := s.Parse(s.lines[s.line-1].text)
		program, err := s.GetProgram(cmd)
		if err != nil {
			// TODO: handle error
			log.Println("ERROR! Program not found:", cmd)
			break
		}

		// Handle UI based programs differently than headless programs.
		if !program.Headless() {
			// TODO: handle error
			log.Println("ERROR! UI based programs not currently supported.")
			break
		}

		if code := program.Run(args); code > 0 {
			// TODO: handle error
			log.Println("ERROR:", program.Error())
			break
		}

		log.Println("output:", program.Output())
	default:
		var symbol string

		// If the key is [a-z] apply shift rules.
		if key >= engo.A && key <= engo.Z {
			if mods.Shift {
				symbol = string(key)
			} else {
				symbol = string(key + 32)
			}
		} else {
			symbol = string(key)
		}

		s.lines[s.line].text = append(s.lines[s.line].text, symbol)
	}
}

func (s *Shell) Parse(text []string) (string, []string) {
	long := strings.Join(text, "")
	parts := strings.Split(long, " ")
	return parts[0], parts[1:]
}

func (s *Shell) GetProgram(name string) (Program, error) {
	for _, program := range registeredPrograms {
		if program.Name() == name {
			return program, nil
		}
	}

	return nil, errors.New("Program not found")
}
