screen = require "screen"
fs = require "fs"

screen.push!
screen.readonly!

lines = fs.cat "/usr/share/manlib/#{arg[1]}"
for line in *lines
    print line