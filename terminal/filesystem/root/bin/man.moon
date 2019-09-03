screen = require "screen"
fs = require "fs"

if #arg ~= 1
    print "Usage: man [command]"
    return

screen.push!
screen.readonly!

lines = fs.cat "/usr/share/man/#{arg[1]}"
for line in *lines
    print line

screen.bind('escape', ->
    screen.unbind('escape')
    screen.pop!
)