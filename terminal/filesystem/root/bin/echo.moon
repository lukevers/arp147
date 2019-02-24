class Echo
    phrase: ""
    say: => print @phrase

phrase = ""
for argument in *arg
    phrase ..= argument .. " "

with Echo!
    .phrase = phrase
    \say!