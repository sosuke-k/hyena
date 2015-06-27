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


function run(argv) {
    fileIO = Library('fileIO');
    var data = chrome_get_tab_info();
    var exportFileWriter = fileIO.fileWriter(argv);
    exportFileWriter.write(data);
    exportFileWriter.close();
}
