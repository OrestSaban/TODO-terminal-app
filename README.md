# TODO Terminal App

A simple command-line To-Do list manager written in Go.

## Features

- Add tasks
- Delete tasks
- Update tasks
- List all tasks
- Clear all tasks
- Toggle task completion status
- Help command for usage instructions

## Installation

### macOS / Linux

1. Download the binary for your platform:
   - [macOS](https://github.com/OrestSaban/TODO-terminal-app/releases/download/v1.0.0/todo-darwin-amd64)
   - [Linux](https://github.com/OrestSaban/TODO-terminal-app/releases/download/v1.0.0/todo-linux-amd64)

2. Make the binary executable and move it to your PATH:
   ```bash
   chmod +x todo-darwin-amd64  # or todo-linux-amd64
   sudo mv todo-darwin-amd64 /usr/local/bin/todo  # Adjust for Linux binary
   ```

3. Verify installation:
   ```bash
   todo help
   ```

### Windows

1. Download the binary:
   - [Windows](https://github.com/OrestSaban/TODO-terminal-app/releases/download/v1.0.0/todo-windows-amd64.exe)

2. Move the executable to a directory in your PATH, such as `C:\Program Files\`.

3. Open a command prompt and verify installation:
   ```cmd
   todo help
   ```

## Usage

- List all tasks:
  ```bash
  todo list
  ```

- Add a new task:
  ```bash
  todo add "Buy groceries"
  ```

- Delete a task:
  ```bash
  todo delete 1
  ```

- Update a task:
  ```bash
  todo update 1 "Buy groceries and cook dinner"
  ```

- Clear all tasks:
  ```bash
  todo clear
  ```

- Toggle task completion status:
  ```bash
  todo toggle 1
  ```

- Show help:
  ```bash
  todo help
  ```

## Building from Source

If you prefer to build the tool yourself, ensure you have Go installed and run:
```bash
go install github.com/OrestSaban/TODO-terminal-app@latest
```

## License

MIT License
