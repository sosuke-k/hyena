#!/usr/bin/env osascript -l JavaScript

function preview_restore_doc(json_string) {
    app = Application.currentApplication();
    app.includeStandardAdditions = true;
    var data = JSON.parse(json_string);

    for (i = 0; i < data[0].length; i++) {
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


function run(argv) {
    try {
        var app = Application('com.apple.Preview');
    } catch (e) {
        console.log(e);
        return false;
    }

    fileIO = Library('fileIO');
    var data = fileIO.read(argv);
    if (!data) {
        return false;
    }

    preview_restore_doc(data);
}
