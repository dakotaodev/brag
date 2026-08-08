# brag

`brag` is a small command-line journal for recording accomplishments, useful feedback, and work you want to remember later.

## Install

Build from a local checkout:

```sh
go build -o brag .
```

Or install the command with Go:

```sh
go install github.com/dakotaodev/brag@latest
```

Confirm the installation:

```sh
brag --help
```

## Usage

Add an entry. Quote values containing spaces:

```sh
brag add "Led the project retrospective"
```

List entries. The leading value on each line is the entry ID:

```sh
brag list
```

Example output:

```text
===========================
e4a1b2 - Aug 7, 2026 at 9:30 AM - Led the project retrospective
===========================
```

Update an entry using its ID:

```sh
brag update e4a1b2 "Led a successful project retrospective"
```

Delete an entry using its ID:

```sh
brag delete e4a1b2
```

For command-specific help, run:

```sh
brag <command> --help
```

## Data storage

Entries are stored as JSON in your operating system's per-user configuration directory:

| Platform | Default location |
| --- | --- |
| macOS | `~/Library/Application Support/brag/brag.json` |
| Linux | `$XDG_CONFIG_HOME/brag/brag.json`, or `~/.config/brag/brag.json` when `XDG_CONFIG_HOME` is unset |
| Windows | `%AppData%\\brag\\brag.json` |

The directory is created automatically on first use.