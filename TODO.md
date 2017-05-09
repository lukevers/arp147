This "TODO" document doesn't mean every thing on here is an actual TODO note, but that I had thought about it and want to get some ideas down and might do them in the future. I still have no actual gameplay, but want to think about how things functionally would work prior to figuring out gameplay. My guess is that gameplay will rely heavily on the computer system that I want to control the entire ship.

# Overview of ideas

This is just a shortlist of things that I want to do, in no particular order:

- [ ] Create a health system for ships
- [ ] Create some sort of a cargo bay limit for supplies/weight
- [ ] Create the computer system
- [ ] Figure out a "room" system for the ship (for example, there should be an engine room, a bridge, a server room for the computers).
- [ ] Building off of the health system and room system, individual rooms can be damaged.
- [ ] Humans? Aliens? Some sort of crew.
- [x] Some sort of time/clock system.
- [ ] Save/load games.
- [ ] Pause the game with `esc` and have a paused menu screen show up.

# More in depth ideas

## Room system

- Ships have individual rooms. TBD if you can look into the ships. Since I am horrible at designing graphics, I imagine I won't allow users to look into ships.
- Each individual room has a purpose.
- Each ship can have a certain number of rooms depending on the type of ship it is.
- Either each ship can decide the type of rooms it wants to have (by adding/removing them at any/certain times), or each ship is locked in to the type of rooms it should have. I can see both options being the right way to handle this, as forcing a ship to have certain rooms could help with game balance, but picking each room type per ship would be nice from a customization point.
- Room ideas:
    - Needed on every ship:
        - Bridge
        - Engines
        - Oxygen
        - Server
        - Living quarters
        - Cargo (size varying)
    - Optional:
        - Shields

## Health system

- Ships have an overall health rating.
- Each individual room has an overall health rating.
- Each member of the crew has an overall health rating. The crew probably should sleep and eat, but let's come back to that later.

## Computer system

- Everything is done on the ship via the computer and programs.
- There is console access with a shell similar to bash.
- There is a text editor in the shell so users can write programs and keep a journal.
- If the server room in the ship is damaged, things won't work properly.
- The server room should be able to overheat if:
    - "CPU intense" programs are running and the server room is not cool enough
    - Fans break (assuming there are fans..TBD)
- SDK to access every system on the ship.
    - https://github.com/avelino/awesome-go#embeddable-scripting-languages
        - I've used both of these before and liked them both:
            - https://github.com/yuin/gopher-lua
            - https://github.com/robertkrimen/otto
    - Maybe multiple SDKs? (through code generation--later though, will stick with one at first.)
- Basic commands for every system on the ship also.
    - If the system does not exist on the ship, the command will still be there but it will either error out or just say that it's not connected.
    - Examples:
        - `shields [status|set|on|off]`
        - `engines [status|set|on|off]`
- Cronjobs for the shell (need to figure out time first)
- Filesystem
