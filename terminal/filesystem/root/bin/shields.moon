shields = require "shields"

switch arg[1]
    when "up"
        shields.up!
    when "down"
        shields.down!
    when "max"
        shields.max!
    when "min"
        shields.min!
    else
        print "Usage: shields [up|down|max|min]"