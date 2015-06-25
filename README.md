## hyena

![hyena demo gif](https://raw.githubusercontent.com/sosuke-k/hyena/master/hyena.gif "hyena demo")

## Dependencies

* go

## Install

```
go get github.com/sosuke-k/hyena
```

### Enable Completion for Bash

1. Install `hyena_completion.bash`. Either:

   1. Place it in your `bash_completion.d` folder, usually something like `/etc/bash_completion.d`,
      `/usr/local/etc/bash_completion.d` or `~/bash_completion.d`.

   2. Or, copy it somewhere (e.g. `~/hyena_completion.bash`) and put the following line in the `.profile` or
      `.bashrc` file in your home directory:

           source ~/hyena_completion.bash


### Enable Completion for Fish

1. Install `hyena_completion.fish` in your `~/.config/fish/completions` folder.


## Usage

<usage>

```

Usage:
  hyena [command]

Available Commands:
  ls                                            List all projects
  add [project name]                            Add project to list
  save [project name]                           Save project
  restore [project name]                        Restore project
  help [command]                                Help about any command


Use "hyena help [command]" for more information about that command.
```

</usage>

## Support Application

* Google Chrome
* Adobe Acrobat
