package terminal

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/blang/vfs"
	"github.com/rucuriousyet/gmoonscript"
	lua "github.com/yuin/gopher-lua"
)

func (ts *TerminalSystem) command(str string) {
	args := strings.Split(strings.Trim(str, " "), " ")
	if len(args) < 1 || args[0] == "" {
		return
	}

	log.Println(args)

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
		log.Println(err)
		ts.WriteLine("Could not read script")
		return
	}

	eval(ftype, bytes, state, ts)
}

func eval(ftype string, source []byte, state *lua.LState, ts *TerminalSystem) {
	switch ftype {
	case "moon":
		// TODO: better import, not every time...
		bridge, err := ioutil.ReadFile("terminal/moonscript-bridge.lua")
		if err != nil {
			log.Println(err)
			ts.WriteLine("Could not load lua-moon bridge")
			return
		}

		if err := state.DoString(
			fmt.Sprintf(
				string(bridge),
				string(source),
			),
		); err != nil {
			log.Println(err)
			ts.WriteLine("Could not run moon script")
		}
	case "lua":
		if err := state.DoString(
			string(source),
		); err != nil {
			log.Println(err)
			ts.WriteLine("Could not run lua script")
		}
	}
}

func newState(args []string, ts *TerminalSystem) *lua.LState {
	state := lua.NewState()
	state.PreloadModule("moonc", gmoonscript.Loader)
	state.PreloadModule("fs", ts.vfs.ScriptLoader)

	// Re-define `print` to print to the screen
	state.SetGlobal("print", state.NewFunction(func(L *lua.LState) int {
		str := L.ToString(1)

		log.Println(str)
		ts.WriteLine(str)
		return 0
	}))

	state.SetGlobal("include", state.NewFunction(func(L *lua.LState) int {
		str := L.ToString(1)

		file, err := ts.vfs.FS.OpenFile(str, os.O_RDONLY, 0777)
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Println(err)
			ts.WriteLine("Could not read script")
			return 0
		}

		ftype := strings.SplitAfterN(file.Name(), ".", 2)[1]
		eval(ftype, bytes, state, ts)
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
