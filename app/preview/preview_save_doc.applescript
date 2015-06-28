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
    var data = JSON.stringify(res, null, 2);
    return data;
}


function run(argv) {
    fileIO = Library('fileIO');
    var data = preview_get_docs_info();
    fileIO.write(argv, data);
}
