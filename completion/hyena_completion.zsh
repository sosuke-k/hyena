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
        'delete[Delete a project]' \
        'save[Save current state to the project]' \
        'restore[Restore project]' \
        'browser[Open hyena page with browser]'
}

__hyena_projects(){
    _values \
        'projects' \
        $(hyena ls | tr "\t" " ")
}
