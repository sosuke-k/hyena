#!/usr/bin/env osascript -l JavaScript

function open_atom(dir) {
    var app = Application.currentApplication();
    app.includeStandardAdditions = true;
    cmd = "cd " + dir + ";atom";
    app.doShellScript(cmd);
}

function restore_atom_windows(json_string) {
    var data = JSON.parse(json_string);

    for (i = 0; i < data[0].length; i++) {
        var dir = data[0][i];
        open_atom(dir);
    }
    return;
}


function run(argv) {
    try {
        var app = Application('com.github.atom');
    } catch (e) {
        console.log(e);
        return false;
    }

    fileIO = Library('fileIO');
    var data = fileIO.read(argv);
    if (!data) {
        return false;
    }

    restore_atom_windows(data);
}