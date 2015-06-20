#!/usr/bin/env osascript -l JavaScript

function acrobat_restore_doc(json_string) {
    app = Application('Adobe Acrobat');
    var data = JSON.parse(json_string);
    var docs = data[0];

    for (i = 0; i < docs.length; i++) {
        app.open(Path(docs[i]));
    }
    return;
}


function fileReader(pathAsString) {
    'use strict';

    var app = Application.currentApplication();
    app.includeStandardAdditions = true;
    var path = Path(pathAsString);
    var file = app.openForAccess(path);

    var eof = app.getEof(file);

    return {
        read: function() {
            return app.read(file, {
                to: eof
            });
        },
        close: function() {
            app.closeAccess(file);
        }
    };
}


function run(argv) {
    var reader = fileReader(argv);
    var data = reader.read();
    reader.close();

    acrobat_restore_doc(data);
}
