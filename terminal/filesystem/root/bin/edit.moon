screen = require "screen"
fs = require "fs"

if #arg ~= 1
    print "Usage: edit [file]"
    return

fs.touch arg[1]

screen.push!
screen.editable!

lines = fs.cat arg[1]
for line in *lines
    print line

screen.bind('escape', ->
    screen.unbind('escape')
    screen.readonly!
    fs.write(arg[1], screen.get_page_contents!)
    screen.pop!
)