## hyena

![hyena demo gif](https://raw.githubusercontent.com/sosuke-k/hyena/master/hyena.gif "hyena demo")

## Dependencies

* go
* kobito-cli

## Install

```
go get -u github.com/sosuke-k/hyena
hyena init
```


## Tutorial

1. Create your project.

    `hyena add project_name`

2. Save the current states of applications to your project.

    `hyena save project_name`

3. Restore the states of applications from your project.

    `hyena restore project_name`


### Enable Completion for Bash

1. Install `hyena_completion.bash`. Either:

   1. Place it in your `bash_completion.d` folder, usually something like `/etc/bash_completion.d`,
      `/usr/local/etc/bash_completion.d` or `~/bash_completion.d`.

   2. Or, copy it somewhere (e.g. `~/hyena_completion.bash`) and put the following line in the `.profile` or
      `.bashrc` file in your home directory:
       ```sh
       source ~/hyena_completion.bash
       ```

### Enable Completion for Fish

1. Install `hyena_completion.fish` in your `~/.config/fish/completions` folder.

### Enable Completion for Zsh

1. Install `hyena_completion.zsh` in either ways:

    1. Place it in your `$fpath` directories where zsh-completion functions are stored. It is usually

        - `/usr/share/zsh/site-functions`
        - `/usr/share/zsh/x.y.z/functions` (x.y.z stands for the version)

    2. Or, copy it somewhere (e.g. `~/hyena_completion.zsh`) and put the following lines in the `~/.zshrc` file:
        ```sh
        autoload -U compinit
        compinit
        source /path/to/hyena_completion.zsh
        ```

    3. Done! Enjoy completion :octocat:

    ![](http://i.gyazo.com/3af265e68f994a3c826d364413b85793.gif)

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
* Kobito
* Atom
