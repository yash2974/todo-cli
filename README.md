# todo-cli

A command-line internal tool that stores tasks in an Excel (`.xlsx`) file.

<img width="599" height="450" alt="todo-cli list output" src="https://github.com/user-attachments/assets/ef65c77a-9737-4d76-a712-8a5803f43a4b" />

## Requirements

- [Go](https://go.dev/dl/) 1.25 or later

## Install

Clone the repo and build the binary:

```sh
git clone https://github.com/yash2974/todo-cli.git
cd todo-cli
go build -o todocli .
```

This produces a `todocli` (or `todocli.exe` on Windows) executable in the project folder. Move it somewhere on your `PATH` if you want to run `todocli` from anywhere.

## Setup

todo-cli reads two environment variables to know where to store your data:

| Variable | Purpose |
|---|---|
| `TODOCLI_FILEPATH` | Path to the Excel file where tasks are stored |
| `TODOCLI_BACKUPFILEPATH` | Path to the backup Excel file |

Both are optional — if unset, todo-cli defaults to `~/.todocli/Book1.xlsx` and `~/.todocli/Backup.xlsx`.

### macOS / Linux

```sh
echo 'export TODOCLI_FILEPATH="$HOME/Documents/todocli/Book1.xlsx"' >> ~/.zshrc
echo 'export TODOCLI_BACKUPFILEPATH="$HOME/Documents/todocli/Backup.xlsx"' >> ~/.zshrc
source ~/.zshrc
```

(Use `~/.bashrc` instead of `~/.zshrc` if you're on bash.)

Verify it's set:

```sh
env | grep TODOCLI
```

### Windows

On Command Prompt:

```bat
setx TODOCLI_FILEPATH "C:\Users\you\Documents\todocli\Book1.xlsx"
setx TODOCLI_BACKUPFILEPATH "C:\Users\you\Documents\todocli\Backup.xlsx"
```

`setx` only applies to new terminal windows, so open a fresh Command Prompt after running it.

Verify it's set:

```bat
echo %TODOCLI_FILEPATH%
```

## Usage

```sh
# Add a task
todocli add -T "Buy groceries" -t errands -p high

# List tasks (defaults to 5 most recent, not done)
todocli list
todocli list -c 20 -d true -t errands

# Update a task
todocli update -I 3 -D true              # mark task 3 as done
todocli update -I 3 -E "New title"       # edit title
todocli update -I 3 -p urgent            # change priority
```

### `add`

| Flag | Shorthand | Description | Default |
|---|---|---|---|
| `--title` | `-T` | Task title (required) | `empty task` |
| `--tag` | `-t` | Tag | `""` |
| `--priority` | `-p` | `urgent`, `high`, `normal`, or `low` | `normal` |

### `list`

| Flag | Shorthand | Description | Default |
|---|---|---|---|
| `--count` | `-c` | Number of tasks to show | `5` |
| `--done` | `-d` | Filter by done status (`true`/`false`) | `false` |
| `--tag` | `-t` | Filter by tag | `""` |

### `update`

| Flag | Shorthand | Description |
|---|---|---|
| `--id` | `-I` | Task ID (required) |
| `--done` | `-D` | Mark task done/not done (`true`/`false`) |
| `--edit` | `-E` | New title |
| `--tag` | `-t` | New tag |
| `--priority` | `-p` | New priority (`urgent`, `high`, `normal`, `low`) |
