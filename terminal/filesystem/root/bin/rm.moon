fs = require "fs"

if #arg < 1
    print "Usage: rm [...file]"
    return

for argument in *arg
    err = fs.rm argument
    if err
        print err