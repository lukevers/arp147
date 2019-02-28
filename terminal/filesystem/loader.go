package filesystem

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	lua "github.com/yuin/gopher-lua"
)

func (fs *VirtualFS) ScriptLoader(state *lua.LState) int {
	mod := state.SetFuncs(state.NewTable(), map[string]lua.LGFunction{
		"cwd": func(L *lua.LState) int {
			state.Push(lua.LString(fs.cwd))
			return 1
		},
		"mkdir": func(L *lua.LState) int {
			for i := 1; i > 0; i++ {
				dir := L.ToString(i)
				if dir == "" {
					break
				}

				fs.FS.Mkdir(dir, 0777)
			}

			return 0
		},
		"cd": func(L *lua.LState) int {
			dir := L.ToString(1)
			fulldir := dir

			if !strings.HasPrefix(dir, "/") {
				fulldir = fmt.Sprintf("%s/%s", fs.cwd, dir)
			}

			info, err := fs.FS.Stat(fulldir)
			if err != nil {
				L.Push(lua.LString(fmt.Sprintf("%s: no such directory", dir)))
				return 1
			}

			if !info.IsDir() {
				L.Push(lua.LString(fmt.Sprintf("%s: not a directory", dir)))
				return 1
			}

			fs.cwd = info.Name()
			return 0
		},
		"touch": func(L *lua.LState) int {
			// dir := L.ToString(1)
			// fs.FS.Mkdir(dir, 0777)

			return 0
		},
		"listdir": func(L *lua.LState) int {
			dir := L.ToString(1)
			fulldir := dir

			if !strings.HasPrefix(dir, "/") {
				fulldir = fmt.Sprintf("%s/%s", fs.cwd, dir)
			}

			info, err := fs.FS.ReadDir(fulldir)
			if err != nil {
				log.Println(err)
				return 0
			}

			i, err := fs.FS.Stat(fulldir)
			if err != nil {
				log.Println(err)
				return 0
			}

			directory := L.NewTable()
			currentdir := L.NewTable()
			currentdir.RawSetString("dir", lua.LBool(true))
			currentdir.RawSetString("name", lua.LString("."))
			directory.Append(currentdir)

			if i.Name() != "/" {
				prevdir := L.NewTable()
				prevdir.RawSetString("dir", lua.LBool(true))
				prevdir.RawSetString("name", lua.LString(".."))
				directory.Append(prevdir)
			}

			for _, i := range info {
				obj := L.NewTable()
				obj.RawSetString("dir", lua.LBool(i.IsDir()))
				obj.RawSetString("name", lua.LString(i.Name()))
				obj.RawSetString("size", lua.LNumber(i.Size()))

				directory.Append(obj)
			}

			L.Push(directory)
			return 1
		},
		"volumes": func(L *lua.LState) int {
			// TODO: support /mnt and not hard-coded
			info, err := fs.FS.Stat("/")
			if err != nil {
				log.Println(err)
				return 0
			}

			log.Println(info.Size())

			directory := L.NewTable()
			dir := L.NewTable()
			dir.RawSetString("dir", lua.LBool(true))
			dir.RawSetString("name", lua.LString("/"))
			dir.RawSetString("size", lua.LString(strconv.Itoa(int(fs.DirSize("/")))))
			dir.RawSetString("max", lua.LString("16384")) // 16kb
			directory.Append(dir)

			L.Push(directory)
			return 1
		},

		// "compile": func(L *lua.LState) int {
		// 	code := L.CheckString(1)

		// 	luaCode, err := Compile(L, code)
		// 	if err != nil {
		// 		state.Push(lua.LNil)
		// 		state.Push(lua.LString(err.Error()))

		// 		return 2
		// 	}

		// 	L.Push(lua.LString(luaCode))
		// 	return 1
		// },
	})

	state.Push(mod)
	return 1
}
