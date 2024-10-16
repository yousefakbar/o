# Obsidian CLI - 'o'

- [Introduction](#introduction)
- [Installation](#installation)
  - [Go Installation](#go-installation)
  - [Build from Source](#build-from-source)
- [Usage](#usage)
  - [Today](#today)
  - [Search](#search)
  - [New](#new)
- [Configuration](#configuration)
- [Contributing](#contributing)

## Introduction

'o' is a command-line tool written in Go that allows you to interact seamlessly with your Obsidian vault from the terminal. It was created to help streamline note-taking and reduce the need to context switch between the terminal and the Obsidian GUI. With 'o', you can create new notes, search through existing markdown files, and quickly access daily notes, all while staying in your favorite command-line environment.

Key features include:
- **Today Command**: Open today's daily note for quick editing.
- **Search Command**: List and search markdown files using fuzzy finder (fzf).
- **New Command**: Create new notes with ease.

## Installation

> [!IMPORTANT]
> **Prerequisites**:
> - [fzf](https://github.com/junegunn/fzf)
> - (optional) [fd](https://github.com/sharkdp/fd)

### Go Installation

To install 'o' using Go, ensure you have Go installed and set up properly. Then run:

```sh
go install github.com/yousefakbar/o/cmd/o@latest
```

This will install the binary in your `$GOPATH/bin` or `$GOBIN` directory. Make sure that directory is in your system's `PATH` to use 'o' directly from the terminal.

### Build from Source

If you want to build 'o' from the source code, follow these steps:

1. Clone the repository:

   ```sh
   git clone https://github.com/yousefakbar/o.git
   cd o
   ```

2. Build the binary:

   ```sh
   go build -o o cmd/o/main.go
   ```

3. Optionally, move the binary to a directory in your `PATH`:

   ```sh
   mv o /usr/local/bin/
   ```

## Usage

'o' provides several commands to help you manage your Obsidian vault efficiently.

> [!NOTE]
> Each command has a shorthand alias (e.g. `t` for `today`). You can learn more in the help screen `o help`

### Today

The `today` command opens today's daily note for quick editing in your preferred text editor:

```sh
o today
```

If you haven't set a custom location for daily notes, the command will create and open the note in the default directory.

### Search

The `search` command allows you to search for existing markdown files using `fzf` for fuzzy finding:

```sh
o search
```

This command will list all `.md` files in your Obsidian vault, and you can use `fzf` to quickly select the one you want to open.

### New

The `new` command helps you create a new note in your vault:

```sh
o new "My New Note"
```

If you don't provide a note name as an argument, the command will prompt you to enter one interactively.

## Configuration

'o' uses environment variables to configure some settings:

- **`OBSIDIAN_VAULT_PATH`**: Set this variable to the path of your Obsidian vault:

  ```sh
  export OBSIDIAN_VAULT_PATH="/path/to/your/vault"
  ```

- **`EDITOR`**: Set this variable to specify your preferred editor for opening notes. If not set, 'nano' will be used by default.

  ```sh
  export EDITOR="vim"
  ```

## Contributing

Contributions are welcome! If you'd like to contribute, feel free to fork the repository and create a pull request. For any questions or suggestions, open an issue on GitHub.
