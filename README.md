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
### Very first version, mimic the $RANDOM behavior

[link to commit](https://github.com/meleu/gordm/commit/933d2e0630cac2b5817ed374e7f3648d528e7260)

Creating a gordm package with a single function just to return a random
integer.

The gordm package needs to expose the Seed so we can test it in a
deterministic way.

In this commit I'm also adding code for the CLI offering a way for the
users to interact with the functionalities of the gordm package.

### Add integration tests for the CLI

[link to commit](https://github.com/meleu/gordm/commit/f7d1454f7b425a401842d2daa60bafeae9b11434)

Using the techniques I learned from the "Powerful Command-Line
Applications in Go" (Chapter 2).

It's a useful way to make sure the CLI is also working as expected.

### Add GitHub Action workflow to run the tests

[link to commit](https://github.com/meleu/gordm/commit/47a8c03e9a85f771d9da5b2cfa8712fd0d8041d4)

Nothing fancy, just running tests on push.

### Add lefthook config

[link to commit](https://github.com/meleu/gordm/commit/a3af0b4fe9fd9480cd9482fe63887202bbebd086)

Run tests before committing changes to the codebase.

It requires lefthook installed (which is easy, via Homebrew).

More info in https://lefthook.dev/

### Add --min and --max command line options

[link to commit](https://github.com/meleu/gordm/commit/3fdd250231b485d4c0ab1426e9a3088f9de9c8c3)

This is an example of how to use the 'flag' package, from standard
library, to parse command line arguments.

Also adding (integration) tests to validate behavior.

### Add function to get random number from random.org

[link to commit](https://github.com/meleu/gordm/commit/4609067d2b509b101ce340b047d3fbf08fd68c05)

This commit has a simple, but useful, example of how to make an HTTP
request and how to create a mocked server to be used in tests (without
sending requests to the actual server).

It also has some examples of conversion tricks between:
byte <--> string <--> int

### Add --web option

[link to commit](https://github.com/meleu/gordm/commit/169471e0178da8b438b9fb8628298821fa66cce2)

Here I'm not adding another test for the CLI to actually call the
program with '-web', as that would be like a repetition of what was made
in `gordm_test.go`.

What I'm doing is adding a test to check if the '--help' message shows
information about the '-web' flag.

By the way, in this commit we also have an example of how to use a
RegExp with a pattern able to match a string with multiple lines.

### Add GoReleaser workflow

[link to commit](https://github.com/meleu/gordm/commit/38402ad5d8fc5a41e1511528ba7d010fafacdd21)

See https://goreleaser.com/ci/actions/

### Tell the project's story from the commit messages

[link to commit](https://github.com/meleu/gordm/commit/fa433f0f0ffdc9b9745eb4d21f5f4797a35f705b)


