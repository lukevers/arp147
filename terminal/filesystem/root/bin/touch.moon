fs = require "fs"

if #arg < 1
    print "Usage: touch [...file]"
    return

for argument in *arg
    err = fs.touch argument
    if err
        print err