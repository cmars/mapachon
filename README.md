# mapachón

mapachón 🦝 likes to dig through dumps. There's no telling what kind of things this feral creature might find.

## Installation

    go install github.com/cmars/mapachon@main

## Usage

### Run an Apache Tika server

    podman-compose up

### Run mapachón

    mapachon /path/to/dump

The thing about los mapachónes is they like to wash their food. So the result
is clean text, rather than questionable document files.

## Security

This isn't some [Record Player Omega](https://geecologist.org/2019/04/godel-escher-bach/), it's a silly
project named after a tricksy wild animal.

Know your threat model and [don't eat stuff off the sidewalk](https://www.youtube.com/watch?v=nFYv7ly-7EM).

* Apache Tika might itself be exploited by malicious input. Consider the
  container runtime surface area.
* Apache Tika might be exploited to produce content that exploits sqlite3 C
  libraries. Consider where mapachon is executed. But already, we've reduced
  this down to text inputs to prepared statements.
* sqlite3 clients might be exploited by malicious databases. Consider whether
  the db is _really_ clean (would _you_ eat food washed by a raccoon?). But again.
  is this very likely through string inputs alone?

That's already several layers to get through, a much-reduced surface area, a
more robust record player (if not a perfect one), and you might add
[more](https://www.qubes-os.org/) [layers](https://webassembly.org/) to your
analysis pipeline.
