screen = require "screen"
fs = require "fs"

if #arg ~= 1
    print "Usage: manlib [lib]"
    return

screen.push!
screen.readonly!

lines = fs.cat "/usr/share/manlib/#{arg[1]}"
for line in *lines
    print line