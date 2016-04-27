package window

import (
	"arp147/computer/line"
	"arp147/computer/window/shell"
	"engo.io/ecs"
	"fmt"
)

type Terminal struct {
	commands map[string]shell.Shell
	world    *ecs.World
	window   *Window
}

func NewTerminal(world *ecs.World, w *Window) *Terminal {
	t := &Terminal{
		commands: make(map[string]shell.Shell),
		world:    world,
		window:   w,
	}

	// Register high level commands
	t.RegisterCommand("help", &shell.Help{})

	// Register commands
	t.RegisterCommand("touch", &shell.Touch{})

	return t
}

func (t *Terminal) RegisterCommand(cmd string, i shell.Shell) {
	t.commands[cmd] = i
}

func (t *Terminal) Dispatch(c *line.Command) error {
	if cmd, ok := t.commands[c.Command]; ok {
		// There are some special cases where we want to define commands at the
		// highest level here, but otherwise we'll
		switch c.Command {
		case "help":
			return t.help(c)
		default:
			return cmd.Run(c)
		}
	} else {
		return fmt.Errorf("%s: command not found", c.Command)
	}
}

// The help func is a high level command defined at the terminal level. It is
// this way because the help command works by using each existing command's
// .Help() func to get the data about how to help.
func (t *Terminal) help(c *line.Command) error {
	// If we send the command "help" with no arguments, we should list all
	// registered commands.
	if len(c.Arguments) < 1 {
		t.window.Write("The commands listed below are are available for use.")
		t.window.Write("Type `help name` to find out more about `name`")
		t.window.NewLine()

		for _, cmd := range t.commands {
			t.window.Write(fmt.Sprintf("%s: %s", cmd.Name(), cmd.Description()))
		}
		return nil
	} else {
		if cmd, ok := t.commands[c.Arguments[0]]; ok {
			t.window.Write(cmd.Help())
			return nil
		} else {
			return fmt.Errorf("%s: command not found", c.Arguments[0])
		}
	}
}
