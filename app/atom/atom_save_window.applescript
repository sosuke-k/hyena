#!/usr/bin/env osascript -l JavaScript

function atom_get_window_info() {
    var systemEvent = Application('System Events');
    systemEvent.includeStandardAdditions = true;

    var atomProcess = systemEvent.processes.byName('Atom');
    var windowMenu = atomProcess.menuBars[0].menuBarItems.byName('Window');
    var windowItems = windowMenu.menus[0].menuItems;
    var res = {};
    res[0] = [];

    var i = windowItems.length - 1;
    while (i > 0) {
        if (windowItems[i].name() == null) {
            break;
        }
        var info = windowItems[i].name().split(" - ");
        //e.g. sample.txt - /Users/name/Desktop - Atom
        //e.g. /Users/name/Desktop - Atom
        var dir = info[info.length-2];
        res[0].push(dir);
        i--;
    }
    var data = JSON.stringify(res);
    return data;
}

function fileWriter(pathAsString) {
    'use strict';

    var app = Application.currentApplication();
    app.includeStandardAdditions = true;
    var path = Path(pathAsString);
    var file = app.openForAccess(path, {
        writePermission: true
    });

    /* reset file length */
    app.setEof(file, {
        to: 0
    });

    return {
        write: function(content) {
            app.write(content, {
                to: file,
                as: 'text'
            });
        },
        close: function() {
            app.closeAccess(file);
        }
    };
}

function run(argv) {
    var activeWindows = atom_get_window_info();
    var exportFileWriter = fileWriter(argv);
    exportFileWriter.write(activeWindows);
    exportFileWriter.close();
}
