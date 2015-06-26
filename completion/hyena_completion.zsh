#compdef _hyena hyena
#autoload

_hyena(){
    _arguments \
        '1: :__hyena_modes' \
        '2: :__hyena_projects'
}

__hyena_modes(){
    _values \
        'mode' \
        'init[Initialize hyena (Should be executed only when you installed hyena)]' \
        'ls[List all your projects]' \
        'add[Add a new project]' \
        'save[Save current state to the project]' \
        'restore[Restore project]'
}

__hyena_projects(){
    _values \
        'projects' \
        $(hyena ls | tr "\t" " ")
}
