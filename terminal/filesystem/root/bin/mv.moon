fs = require "fs"

if #arg ~= 2
    print "Usage: mv [from] [to]"
    return

err = fs.mv arg[1], arg[2]
if err
    print err