local moonscript_code = [[
    %s
]]

local moonc = require("moonc")

lua_code, err = moonc.compile(moonscript_code)
if err ~= nil then
	print(err)
else
	loadstring(lua_code)()
end