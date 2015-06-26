#!/usr/bin/env osascript -l JavaScript

function kobito_get_window_info() {
    var systemEvent = Application('System Events');
    systemEvent.includeStandardAdditions = true;

    var kobitoProcess = systemEvent.processes.byName('Kobito');
    var windowMenu = kobitoProcess.menuBars[0].menuBarItems.byName('Window');
    try {
      console.log(windowMenu.menus[0].menuItems.length);
    } catch (e) {
      console.log(e);
      windowMenu = kobitoProcess.menuBars[0].menuBarItems.byName('ウィンドウ');
    }
    var windowItems = windowMenu.menus[0].menuItems;
    var res = {};
    res[0] = [];

    var i = windowItems.length - 1;
    while (i > 0) {
        if (windowItems[i].name() == null) {
            break;
        }
        res[0].push(windowItems[i].name());
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


function run(argv){
    var activeWindows = kobito_get_window_info();
    console.log(activeWindows);
    var exportFileWriter = fileWriter(argv);
    exportFileWriter.write(activeWindows);
    exportFileWriter.close();
}
