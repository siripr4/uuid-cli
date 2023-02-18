# UUID CLI

A simple tool to generate uuid with a cli command.

Generates a random (Version 4) UUID and prints to stdout.

Underlying library: https://pkg.go.dev/github.com/google/uuid

## Installation

**1. Manually build**

1. Clone the repo

```git clone https://github.com/siripr4/uuid-cli.git```

2. cd to the directory

```cd uuid-cli```

3. Build the go file

```go build . -o uuid```

4. Run the binary 

```./uuid```

5. Optionally, add the binary to your PATH to generate uuid from anywhere. Symlink the binary to `usr/local/bin`

```sudo ln -s <absolute-path-to-binary> /usr/local/bin/uuid```

**2. Install via homebrew**

_TODO_

