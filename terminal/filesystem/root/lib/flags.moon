M = {}

append = table.insert
usage_ = nil

M.quit = (msg) ->
    if msg then io.stderr\write msg, '\n'
    if usage_
        print '\n', usage_
    os.exit 1

glob = (args) ->
    if #args == 1 and args[1]\match '[%*%?]'
        -- no shell globbing, it's Windows :(
        wildcard = args[1]
        table.remove args, 1
        f = io.popen 'dir /b '..wildcard
        path = wildcard\match [[(.-\)[^\]+$]] or ''
        for line in f\lines!
            append args, path..line
        f\close!

M.parse = (usage) ->
    usage_ = usage
    takes_value, known_flags = {},{}
    for line in usage\gmatch '[^\n]+'
        flag,rest = line\match '^%s+%-(%S+)%s+(.*)'
        if flag
            known_flags[flag] = true
            takes_value[flag] = rest\match '^<'

    quit = (flag,msg) -> M.quit '-'..flag..' '..msg
    args,i = {},1
    while i <= #arg
        a,val = arg[i]
        flag = a\match '^%-(.+)'
        if flag
            if not known_flags[flag] then quit flag,'unknown flag'
            if takes_value[flag]
                i += 1
                if i > #arg or arg[i]\match '^%-[^%-]'
                    quit flag,'needs a value'
                val = arg[i]
            args[flag] = val or true
        else
            append args, a
        i += 1
    glob args
    return args

return M