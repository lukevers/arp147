fs = require "fs"

dir = if arg[1] != nil
    arg[1]
else
    "/home"

err = fs.cd dir
print err if err != nil