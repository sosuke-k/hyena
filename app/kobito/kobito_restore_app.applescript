#!/usr/bin/env osascript -l JavaScript

function get_n_by_title(title) {
  var app = Application.currentApplication();
  app.includeStandardAdditions = true;
  cmd = "kobito ls | cat -n | grep '" + title + "' | cut -f 1";
  var out = app.doShellScript(cmd);
  out = parseInt(out) - 1;
  return typeof(out) == "number" ? out : 0;
}

// Kobito must be activated
function open_n_th_item(n) {
  var se = Application('System Events')
  se.includeStandardAdditions = true;
  for (var i = 0; i < n - 1; i++) {
    delay(0.1);
    se.keyCode(125); // Press Key Down
  }
  delay(0.1);
  se.keystroke("e", {
    using: "command down"
  });
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

function get_one_title(json_string) {
  var data = JSON.parse(json_string);
  var array = data[0];
  var l = array.length;
  var title = l > 0 ? array[0] : "";
  return (title == "Kobito" && l > 1) ? array[1] : title;
}


function run(argv) {
    try {
      var app = Application('Kobito');
    } catch (e) {
      console.log(e);
      return false;
    }

  fileIO = Library('fileIO');
  var data = fileIO.read(argv);
  if (!data) {
    return false;
  }

  var title = get_one_title(data);
  if (title != "") {
    activate_kobito();
    var n = get_n_by_title(title);
    open_n_th_item(n);
  }
}
