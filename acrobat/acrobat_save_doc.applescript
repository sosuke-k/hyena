#!/usr/bin/env osascript -l JavaScript

function acrobat_get_docs_info() {
  app = Application('Adobe Acrobat');
  var res = {};
  res[0] = [];

  for (i = 0; i < app.documents.length; i++) {
    d = app.documents[i];
    res[0].push(d.fileAlias().toString());
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
    var data = acrobat_get_docs_info();
    var exportFileWriter = fileWriter(argv);
    exportFileWriter.write(data);
    exportFileWriter.close();
}
