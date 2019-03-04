print ""
print "   ###    ########  ########    ##   ##        ######## "
print "  ## ##   ##     ## ##     ## ####   ##    ##  ##    ## "
print " ##   ##  ##     ## ##     ##   ##   ##    ##      ##   "
print "##     ## ########  ########    ##   ##    ##     ##    "
print "######### ##   ##   ##          ##   #########   ##     "
print "##     ## ##    ##  ##          ##         ##    ##     "
print "##     ## ##     ## ##        ######       ##    ##     "
print ""

screen = require "screen"
fs = require "fs"

if login!
    print "Login ... SUCCESS!"
    print "Directory " .. fs.cwd()
    screen.writable!
else
    print "Login ... FAILURE!"
    print "Could not load saved game."
    print "Initiating..."
    -- TODO

print ""