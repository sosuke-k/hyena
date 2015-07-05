#!bash

_hyena()
{
  local cur=${COMP_WORDS[COMP_CWORD]}
  case "$COMP_CWORD" in
  1)
    COMPREPLY=( $(compgen -W "init ls add delete save restore browser" -- $cur) );;
  2)
    COMPREPLY=( $(compgen -W "$(hyena ls)" -- $cur) );;
  esac
}
complete -F _hyena hyena
