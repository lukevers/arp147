package terminal

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/blang/vfs"
	"github.com/rucuriousyet/gmoonscript"
	lua "github.com/yuin/gopher-lua"
)

func (ts *TerminalSystem) loginScript(result *chan bool) {
	file, err := ts.vfs.FS.OpenFile("/home/login.moon", os.O_RDONLY, 0777)

	if err != nil {
		// No login script. Nothing to do
		return
	}

	state := newState([]string{}, ts)
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		ts.WriteLine("Could not read script:")
		ts.WriteError(err)
		return
	}

	eval("moon", preprocessBytes(bytes, ts), state, ts)
	*result <- true
}

func (ts *TerminalSystem) command(str string) {
	args := strings.Split(strings.Trim(str, " "), " ")
	if len(args) < 1 || args[0] == "" {
		return
	}

	exts := []string{"moon", "lua"}
	var f *vfs.File
	var ftype string
	for _, ext := range exts {
		path := fmt.Sprintf("/bin/%s.%s", args[0], ext)
		file, err := ts.vfs.FS.OpenFile(path, os.O_RDONLY, 0777)

		if err != nil {
			continue
		} else {
			f = &file
			ftype = ext
			break
		}
	}

	if f == nil {
		ts.WriteLine(fmt.Sprintf("%s: command not found", args[0]))
		return
	} else {
		defer (*f).Close()
	}

	state := newState(args, ts)
	bytes, err := ioutil.ReadAll(*f)
	if err != nil {
		ts.WriteLine("Could not read script:")
		ts.WriteError(err)
		return
	}

	eval(ftype, preprocessBytes(bytes, ts), state, ts)
}

func preprocessBytes(bytes []byte, ts *TerminalSystem) (source string) {
	lines := strings.Split(string(bytes), "\n")

	for _, line := range lines {
		if !strings.Contains(line, "use ") {
			source += fmt.Sprintf(
				"%s\n",
				line,
			)

			continue
		}

		parts := strings.Split(line, "=")
		parts[0] = strings.TrimSpace(parts[0])
		declarations := strings.Split(parts[1], "use")
		file := strings.Trim(strings.TrimSpace(declarations[1]), "\"")

		f, err := ts.vfs.FS.OpenFile(file, os.O_RDONLY, 0777)
		bytes, err := ioutil.ReadAll(f)
		if err != nil {
			ts.WriteError(err)
			continue
		}

		source += fmt.Sprintf(
			"%s = ->\n",
			parts[0],
		)

		for _, fline := range strings.Split(string(bytes), "\n") {
			source += fmt.Sprintf(
				"\t%s\n",
				fline,
			)
		}

		// Self invoke function for result
		source += fmt.Sprintf(
			"%s = %s()\n",
			parts[0],
			parts[0],
		)
	}

	return
}

func eval(ftype, source string, state *lua.LState, ts *TerminalSystem) {
	switch ftype {
	case "moon":
		// TODO: better import, not every time...
		bridge, err := ioutil.ReadFile("terminal/moonscript-bridge.lua")
		if err != nil {
			ts.WriteLine("Could not load lua-moon bridge:")
			ts.WriteError(err)
			return
		}

		if err := state.DoString(
			fmt.Sprintf(
				string(bridge),
				string(source),
			),
		); err != nil {
			ts.WriteLine("Could not run moon script:")
			ts.WriteError(err)
			return
		}
	case "lua":
		if err := state.DoString(
			string(source),
		); err != nil {
			ts.WriteLine("Could not run lua script:")
			ts.WriteError(err)
			return
		}
	}
}

func newState(args []string, ts *TerminalSystem) *lua.LState {
	state := lua.NewState(lua.Options{
		CallStackSize: 40960,
		RegistrySize:  40960 * 20,
	})

	state.PreloadModule("moonc", gmoonscript.Loader)
	state.PreloadModule("fs", ts.vfs.ScriptLoader)

	state.PreloadModule("screen", func(state *lua.LState) int {
		mod := state.SetFuncs(state.NewTable(), map[string]lua.LGFunction{
			"push": func(L *lua.LState) int {
				ts.pages[ts.page].hide()

				ts.page++
				ts.pages[ts.page] = &page{
					lines:     make(map[int]*line),
					line:      0,
					escapable: true,
				}

				return 0
			},
			"pop": func(L *lua.LState) int {
				ts.pages[ts.page].hide()
				delete(ts.pages, ts.page)
				ts.page--
				ts.pages[ts.page].show()
				ts.WriteLine("")

				return 0
			},
			"readonly": func(L *lua.LState) int {
				ts.pages[ts.page].readonly = true
				return 0
			},
			"writable": func(L *lua.LState) int {
				ts.pages[ts.page].readonly = false
				return 0
			},
			"count": func(L *lua.LState) int {
				L.Push(lua.LNumber(len(ts.pages)))
				return 1
			},
		})

		state.Push(mod)
		return 1
	})

	// Re-define `print` to print to the screen
	state.SetGlobal("print", state.NewFunction(func(L *lua.LState) int {
		str := L.ToString(1)
		ts.WriteLine(str)
		return 0
	}))

	state.PreloadModule("user", func(state *lua.LState) int {
		mod := state.SetFuncs(state.NewTable(), map[string]lua.LGFunction{
			"login": func(L *lua.LState) int {
				login := true

				// TODO: load saved game
				if false {
					login = true
				}

				time.Sleep(2 * time.Second)

				L.Push(lua.LBool(login))
				return 1
			},
			"new": func(L *lua.LState) int {
				// TODO
				return 0
			},
		})

		state.Push(mod)
		return 1
	})

	state.PreloadModule("shields", func(state *lua.LState) int {
		mod := state.SetFuncs(state.NewTable(), map[string]lua.LGFunction{
			"up": func(L *lua.LState) int {
				ts.ship.Shield.Increase(ts.ship)
				return 0
			},
			"down": func(L *lua.LState) int {
				ts.ship.Shield.Decrease(ts.ship)
				return 0
			},
			"max": func(L *lua.LState) int {
				ts.ship.Shield.Max(ts.ship)
				return 0
			},
			"min": func(L *lua.LState) int {
				ts.ship.Shield.Min(ts.ship)
				return 0
			},
		})

		state.Push(mod)
		return 1
	})

	state.SetGlobal("include", state.NewFunction(func(L *lua.LState) int {
		str := L.ToString(1)

		file, err := ts.vfs.FS.OpenFile(str, os.O_RDONLY, 0777)
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			ts.WriteLine("Could not read script:")
			ts.WriteError(err)
			return 0
		}

		ftype := strings.SplitAfterN(file.Name(), ".", 2)[1]
		eval(ftype, string(bytes), state, ts)
		state.Push(state.Get(1))
		return 1
	}))

	// Re-define `arg` to add arguments
	argTable := state.NewTable()
	for i, arg := range args {
		argTable.Insert(i, lua.LString(arg))
	}

	state.SetGlobal("arg", argTable)

	return state
}
