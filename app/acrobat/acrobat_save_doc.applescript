#!/usr/bin/env osascript -l JavaScript

function acrobat_get_docs_info() {
    app = Application('com.adobe.Acrobat.Pro');
    var res = {};
    res[0] = [];

    for (i = 0; i < app.documents.length; i++) {
        d = app.documents[i];
        res[0].push(d.fileAlias().toString());
    }
    var s = JSON.stringify(res, null, 2);
    return s;
}


function run(argv) {
    try {
        var app = Application('com.adobe.Acrobat.Pro');
    } catch (e) {
        console.log(e);
        return false;
    }
    fileIO = Library('fileIO');
    var data = acrobat_get_docs_info();
    fileIO.write(argv, data);
}
