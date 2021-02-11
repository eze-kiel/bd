# bd

Manage a database of birthday dates from the terminal !

<img align="center" src="./img/21.png" width="250" height="250" />

## Demo

[![bd in action](https://asciinema.org/a/wrCm3ZEcBpDuLpIxwDSMo27jh.svg)](https://asciinema.org/a/wrCm3ZEcBpDuLpIxwDSMo27jh)

## Gettin' started

### Get the binary

You can get the program on your machine...

* ...using Golang:

```
$ git clone github.com/eze-kiel/bd.git
$ cd bd/
$ go build
```

* ...from the [releases](https://github.com/eze-kiel/bd/releases)

I highly recommend you to move the freshly installed binary into your `$PATH` scope.

### Setup the files

Once `bd` is installed, you have one command to initialize everything:

```
$ bd init
```

This will create `$HOME/.bd/dates.json`.

## Usage

From the help menu:

```
Usage:
  bd [command]

Available Commands:
  coming      Display coming birthdays
  completion  Generate completion script
  delete      Delete an entry from the birthday database
  help        Help about any command
  init        Initialize bd components
  insert      Insert a birthday date into the base
  list        List all the saved birthdays
  search      Search a specific birthday

Flags:
  -h, --help   help for bd
```

## Completion

### Bash

You can add the courcing command to your `.bashrc`:

```
$ echo "source <(yourprogram completion bash)" > .bashrc
```

Or exectute the following:

```
$ bd completion bash > /etc/bash_completion.d/bd
```

### Zsh

If shell completion is not already enabled in your environment, you will need to enable it.  You can execute the following once:

```
$ echo "autoload -U compinit; compinit" >> ~/.zshrc
```

Then to load completion for each session, execute once:

```
$ bd completion zsh > "${fpath[1]}/_bd"
```

### Fish

To load completions for each session, execute once:

```
$ yourprogram completion fish > ~/.config/fish/completions/yourprogram.fish
```

## License

MIT

## Credits

Gopher from [Maria Letta](https://github.com/MariaLetta/free-gophers-pack)