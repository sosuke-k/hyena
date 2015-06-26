#!/usr/bin/env osascript -l JavaScript

function preview_restore_doc(json_string){
    app = Application.currentApplication();
    app.includeStandardAdditions = true;
    var data = JSON.parse(json_string);

    for (i = 0; i < data[0].length; i++){
        app.doShellScript("open -a Preview.app " + data[0][i]);

        // var docInfo = app.document({
        //     path: data[0][i]
        // });
        // d = app.open(docInfo);
        // console.log(d.path);
        // app.documents().push(d);
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

    preview_restore_doc(data);
}
