# Gordm - Go Random

This is a project created to learn the basics of how to organize a project for
a CLI written in Go. The focus is more on how to organize **and test** the
code than in the usefulness of generating a random integer.

The execution is always following Test-Driven Development.

Topics addressed in such a simple project:

- How to test randomly generated numbers.
- Organization of a typical Go project.
- Integration tests of the CLI.
- Handling command line options with the `flag` package from standard Go lib.
- Make HTTP request to get a number from random.org
- Create a fake HTTP server to be used in tests, without sending requests to the actual server.
- Use lefthook to make sure you don't even commit when tests are failing.
- GitHub actions to:
  - run tests on push
  - create binaries for the release on tag creation

Each commit message has valuable insights about the learnings at each step.

## Commit History
