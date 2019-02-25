fs = require "fs"

dir = fs.listdir(arg[1])
for obj in *dir
    print obj.name