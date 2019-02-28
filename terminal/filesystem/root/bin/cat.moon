fs = require "fs"

lines = fs.cat arg[1]
for line in *lines
    print line