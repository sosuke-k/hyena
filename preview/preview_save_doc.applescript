#!/usr/bin/env osascript -l JavaScript

function preview_get_docs_info(){
    app = Application('Preview');
    var res = {};
    res[0] = [];
    for (i = 0; i < app.documents().length; i++) { // multiple windows
        doc =  app.documents()[i];
        docPath = doc.path();
        res[0].push(docPath);

        // for (j = 0; j < app.windows()[i].length; j++) {
        //     console.log(app.windows()[i].documents()[j].path());
        // }

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
            app.write(
                content, {
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
    var data = preview_get_docs_info();
    console.log(data);
    var exportFileWriter = fileWriter(argv);
    exportFileWriter.write(data);
    exportFileWriter.close();
}
