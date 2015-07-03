#!fish

## Support functions

function __fish_hyena_using_command
  set cmd (commandline -opc)
  set subcommands $argv
  if [ (count $subcommands) -eq 0 ]
    if [ (count $cmd) -eq 1 ]
      return 0
    else
      return 1
    end
  end
  if [ (count $cmd) = (math (count $subcommands) + 1) ]
    for i in (seq (count $subcommands))
      if not test $subcommands[$i] = $cmd[(math $i + 1)]
        return 1
      end
    end
    return 0
  end
  return 1
end

function __fish_hyena_projects
  hyena ls | tr '\t' '\n'
end

## hyena
complete -f -c hyena -n '__fish_hyena_using_command' -a "init ls add delete save restore"

## hyena add
complete -f -c hyena -n '__fish_hyena_using_command add' -a '(__fish_hyena_projects)'
complete -f -c hyena -n '__fish_hyena_using_command a' -a '(__fish_hyena_projects)'

## hyena delete
complete -f -c hyena -n '__fish_hyena_using_command delete' -a '(__fish_hyena_projects)'
complete -f -c hyena -n '__fish_hyena_using_command d' -a '(__fish_hyena_projects)'


## hyena save
complete -f -c hyena -n '__fish_hyena_using_command save' -a '(__fish_hyena_projects)'
complete -f -c hyena -n '__fish_hyena_using_command s' -a '(__fish_hyena_projects)'

## hyena restore
complete -f -c hyena -n '__fish_hyena_using_command restore' -a '(__fish_hyena_projects)'
complete -f -c hyena -n '__fish_hyena_using_command r' -a '(__fish_hyena_projects)'
