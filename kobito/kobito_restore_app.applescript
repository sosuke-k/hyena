#!/usr/bin/env osascript -l JavaScript

function get_n_by_title(title) {
  var app = Application.currentApplication();
  app.includeStandardAdditions = true;
  cmd = "kobito ls | cat -n | grep '" + title + "' | tr -s ' ' ',' | tr -s '\t' ',' | cut -d ',' -f 2";
  var out = app.doShellScript(cmd);
  return parseInt(out) - 1;
}

// Kobito must be activated
function open_n_th_item(n) {
  var se = Application('System Events')
  se.includeStandardAdditions = true;
  for (var i = 0; i < n - 1; i++) {
    se.keyCode(125); // Press Key Down
  }
  se.keystroke("o");
}

function activate_kobito() {
  var kobito = Application("Kobito");
  kobito.quit();
  var app = Application.currentApplication();
  app.includeStandardAdditions = true;
  app.doShellScript("open -a Kobito");
  kobito.activate();
  delay(1);
}

fucntion get_one_title(json_string) {
  var data = JSON.parse(json_string);
  var l = Object.keys(data).length;
  return l > 0 ? data[0] : "";
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

  var title = get_one_title(data);
  if (title != "") {
    activate_kobito();
    var n = get_n_by_title(title);
    open_n_th_item(n);
  }
}
