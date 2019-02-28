fs = require "fs"
Table = use "lib/table.moon"

Table\addHeader("USED")
Table\addHeader("AVAILABLE")
Table\addHeader("MAX")
Table\addHeader("CAPACITY")
Table\addHeader("MOUNT LOCATION")

volumes = fs.volumes!
for vol in *volumes
    capacity = math.floor((vol.size / vol.max) * 100)
    available = vol.max - vol.size
    Table\addRow({vol.size, "#{available}", vol.max, "#{capacity}%", vol.name})

Table\render()