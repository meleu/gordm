# Gordm - Go Random

This is a project I created to learn the basics of how to create and organize
a project for a CLI written in Go. The focus is more on how to organize
**and test** the project than in the usefulness of creating a random integer.

## ROADMAP

- [x] generate random numbers (mimic `echo $RANDOM`)
- [x] add integration tests
- [x] GitHub Action to run the tests
- [x] specify the min and/or max via arguments
  - [x] e.g.: `gordm -min 0 -max 6`
  - [x] validate that min < max
- [ ] get random integer from random.org
  - [ ] `gordm -web`
- [ ] subcommands: `gordm number` (make it the default when no subcommand is given)
- [ ] `gordm password` to generate a random password
- [ ] `gordm password -size NUM` to define the amount of characters
- [ ] `gordm password -num` to generate a numeric password
- [ ] `gordm password -alpha` to generate a password using the letters of the alphabet
- [ ] `gordm password -symbol` to generate a password allowing symbols
