fs = require "fs"

if #arg ~= 1
    print "Usage: cat [file]"
    return

lines = fs.cat arg[1]
for line in *lines
    print line