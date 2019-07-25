package filesystem

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	lua "github.com/yuin/gopher-lua"
)

// ScriptLoader creates a set of filesystem level functions to be used in
// scripts.
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

				if !strings.HasPrefix(dir, "/") {
					dir = fmt.Sprintf("%s/%s", fs.cwd, dir)
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

			fs.cwd = fulldir
			return 0
		},
		"touch": func(L *lua.LState) int {
			file := L.ToString(1)
			fullfile := file
			if !strings.HasPrefix(file, "/") {
				fullfile = fmt.Sprintf("%s/%s", fs.cwd, file)
			}

			_, err := fs.FS.OpenFile(fullfile, os.O_CREATE, 0777)
			if err != nil {
				L.Push(lua.LString(fmt.Sprintf("%s: could not touch file", file)))
				return 1
			}

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
				fs.WriteError(err)
				return 0
			}

			i, err := fs.FS.Stat(fulldir)
			if err != nil {
				fs.WriteError(err)
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
			_, err := fs.FS.Stat("/")
			if err != nil {
				fs.WriteError(err)
				return 0
			}

			directory := L.NewTable()
			dir := L.NewTable()
			dir.RawSetString("dir", lua.LBool(true))
			dir.RawSetString("name", lua.LString("/"))
			dir.RawSetString("size", lua.LString(strconv.Itoa(int(fs.DirSize("/")))))
			dir.RawSetString("max", lua.LString("128000"))
			directory.Append(dir)

			L.Push(directory)
			return 1
		},
		"cat": func(L *lua.LState) int {
			dir := L.ToString(1)
			fulldir := dir

			if !strings.HasPrefix(dir, "/") {
				fulldir = fmt.Sprintf("%s/%s", fs.cwd, dir)
			}

			_, err := fs.FS.Stat(fulldir)
			if err != nil {
				fs.WriteError(err)
				return 0
			}

			file, err := fs.FS.OpenFile(fulldir, os.O_RDONLY, 0777)
			if err != nil {
				fs.WriteError(err)
				return 0
			}

			bytes, err := ioutil.ReadAll(file)
			if err != nil {
				fs.WriteError(err)
				return 0
			}

			lines := L.NewTable()
			for _, str := range strings.Split(string(bytes), "\n") {
				lines.Append(lua.LString(str))
			}

			L.Push(lines)
			return 1
		},
	})

	state.Push(mod)
	return 1
}
