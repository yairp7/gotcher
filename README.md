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

### Example:
`gotcher --events=write,remove --pattern=".txt$" --cmd="echo 'File event occurred!'"`

This command watches for both "write" and "remove" events on files ending with the ".txt" extension. When either event occurs, the command echo 'File event occurred!' is executed.

## Contributing
We welcome contributions to Gotcher!


## License
Gotcher is licensed under the MIT License. See the LICENSE file for more details.