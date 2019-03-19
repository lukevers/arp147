ship = require "ship"

if #arg ~= 2
    print "Usage: goto [x] [y]"
    return

ship.go(arg[1], arg[2])