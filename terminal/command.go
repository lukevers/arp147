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
	args := strings.Split(str, " ")
	if len(args) < 1 {
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
		// TODO: throw error in console
		return
	} else {
		defer (*f).Close()
	}

	state := newState(args)
	bytes, err := ioutil.ReadAll(*f)
	if err != nil {
		log.Println(err)
	}

	switch ftype {
	case "moon":
		// TODO: better import, not every time...
		bridge, err := ioutil.ReadFile("terminal/moonscript-bridge.lua")
		if err != nil {
			log.Println(err)
			// TODO ...
			return
		}

		if err := state.DoString(
			fmt.Sprintf(
				string(bridge),
				string(bytes),
			),
		); err != nil {
			log.Println(err)
		}
	case "lua":
		if err := state.DoString(
			string(bytes),
		); err != nil {
			log.Println(err)
		}
	}
}

func newState(args []string) *lua.LState {
	state := lua.NewState()
	state.PreloadModule("moonc", gmoonscript.Loader)

	// Re-define `print` to print to the screen
	state.SetGlobal("print", state.NewFunction(func(L *lua.LState) int {
		log.Println(L.ToString(1))
		// TOOD: print to screen
		return 0
	}))

	// Re-define `arg` to add arguments
	argTable := state.NewTable()
	for i, arg := range args {
		argTable.Insert(i, lua.LString(arg))
	}

	state.SetGlobal("arg", argTable)

	return state
}
