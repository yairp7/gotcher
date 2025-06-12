# Gotcher

Gotcher is a simple, open-source command-line utility written in Go that watches for filesystem events and executes a user-specified command in response.

## Installation

Install Gotcher using the `go install` command:

`go install github.com/yairp7/gotcher`

## Usage

`gotcher --events=<event1>,<event2>,... --pattern=<pattern> --cmd=<command>`

### Arguments:

* `--events`: Specifies the types of filesystem events to watch for. Valid options: `"write", "remove", "rename", "chmod"`. Multiple events can be specified by separating them with commas.
* `--pattern`: Specifies a regular expression pattern to match files or directories of interest.
* `--cmd`: Specifies the shell command to execute when a matching event occurs.
* `--follow`: When enabled, automatically watches new directories that are created inside the watched path during runtime.
* `--log-level`: Sets the logging level. Valid options: `"debug", "info", "warn", "error"`. Defaults to `"info"`.

### Special Markers

Gotcher supports special markers in the command string that are replaced with event-specific information:

* `#[file]`: Gets replaced with the full path of the file that triggered the event
* `#[op]`: Gets replaced with the operation type that occurred (write, remove, rename, or chmod)

### Example:
`gotcher --events=write,remove --pattern=".txt$" --cmd="echo 'File #[op] event occurred on #[file]!'"`

This command watches for both "write" and "remove" events on files ending with the ".txt" extension. When either event occurs, it will echo a message containing both the operation type and the file path. For example, if `/path/to/document.txt` is modified, it will output: `File write event occurred on /path/to/document.txt!`

To automatically watch new directories that are created:

`gotcher --events=write --pattern=".txt$" --cmd="echo 'Changed: #[file]'" --follow`

This command will watch for write events on .txt files and automatically start watching any new directories that are created within the watched path.

## Contributing
We welcome contributions to Gotcher!

## Versioning
Gotcher follows [Semantic Versioning](https://semver.org/). Version information can be obtained by running:

```bash
gotcher version
```

This will display the version number, git commit hash, and build date.

To create a new release:

```bash
make release TAG=v1.0.0
```

## License
Gotcher is licensed under the MIT License. See the LICENSE file for more details.