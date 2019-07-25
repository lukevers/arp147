fs = require "fs"

if #arg < 1
    print "Usage: mkdir [...dirs]"
    return

for argument in *arg
    fs.mkdir argument