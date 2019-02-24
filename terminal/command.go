package terminal

import (
	"log"
	"strings"

	"github.com/rucuriousyet/gmoonscript"
	lua "github.com/yuin/gopher-lua"
)

func command(str string) {
	args := strings.Split(str, " ")
	if len(args) < 1 {
		return
	}

	log.Println(args)
	// state := newState(args)

	// 	err := state.DoString(`
	// local moonscript_code = [[
	// class Echo
	//     phrase: ""
	//     say: => print @phrase

	// phrase = ""
	// for argument in *arg
	//     phrase ..= argument .. " "

	// with Echo!
	//     .phrase = phrase
	//     \say!
	// ]]

	// local moonc = require("moonc")

	// lua_code, err = moonc.compile(moonscript_code)
	// if err ~= nil then
	// 	print(err)
	// else
	// 	loadstring(lua_code)()
	// end
	// 		`)
	// 	if err != nil {
	// 		panic(err)
	// 	}
}

// 	err := state.DoString(`
// local moonscript_code = [[
// class Thing
//   name: "unknown"

// class Person extends Thing
//   say_name: => print "Hello, I am #{@name}!"

// with Person!
//   .name = "MoonScript"
//   \say_name!
// ]]

// local moonc = require("moonc")

// lua_code, err = moonc.compile(moonscript_code)
// if err ~= nil then
// 	print(err)
// else
// 	loadstring(lua_code)()
// end
// 	`)
// 	if err != nil {
// 		panic(err)
// 	}
// }

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
