screen = require "screen"
fs = require "fs"

screen.push!
screen.readonly!

lines = fs.cat "/usr/share/man/#{arg[1]}"
for line in *lines
    print line