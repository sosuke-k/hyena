#!/usr/bin/env osascript -l JavaScript

function chrome_get_tab_info() {
    app = Application('Google Chrome');
    windows = app.windows();
    var res = {};

    for (i = 0; i < windows.length; i++) {
        w = windows[i];
        res[i] = [];
        for (j = 0; j < w.tabs().length; j++) {
            res[i].push(w.tabs[j].url());
        }
    }
    var s = JSON.stringify(res);
    return s;
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
    var data = chrome_get_tab_info();
    var exportFileWriter = fileWriter(argv);
    exportFileWriter.write(data);
    exportFileWriter.close();
}
