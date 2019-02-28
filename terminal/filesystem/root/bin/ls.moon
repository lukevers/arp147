fs = require "fs"

dir = fs.listdir(arg[1])
if assert dir ~= nil
    for obj in *dir
        print obj.name