fs = require "fs"

if #arg < 1
    print "Usage: touch [...file]"
    return

for argument in *arg
    fs.touch argument