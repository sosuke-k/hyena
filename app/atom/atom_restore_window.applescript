#!/usr/bin/env osascript -l JavaScript

function open_atom(dir) {
  var app = Application.currentApplication();
  app.includeStandardAdditions = true;
  cmd = "cd " + dir + ";atom";
  app.doShellScript(cmd);
}

function restore_atom_windows(json_string) {
  var data = JSON.parse(json_string);

  for (i = 0; i < data[0].length; i++) {
      var dir = data[0][i];
      open_atom(dir);
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
    try {
      var app = Application('com.github.atom');
    } catch (e) {
      console.log(e);
      return false;
    }

    var reader = fileReader(argv);
    var data = null;
    try {
      data = reader.read();
    } catch (e) {
      console.log(e);
      return false;
    } finally {
      reader.close();
    }

    restore_atom_windows(data);
}
