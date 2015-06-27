#!/usr/bin/env osascript -l JavaScript

function acrobat_restore_doc(json_string) {
  var app = Application('Adobe Acrobat');
  var data = JSON.parse(json_string);
  var docs = data[0];

  for (i = 0; i < docs.length; i++) {
    app.open(Path(docs[i]));
  }
  return;
}


function run(argv) {
  try {
    var app = Application('Adobe Acrobat');
  } catch (e) {
    console.log(e);
    return false;
  }
  fileIO = Library('fileIO');
  var reader = fileIO.fileReader(argv);
  try {
    var data = reader.read();
  } catch (e) {
    console.log(e);
    return false;
  } finally {
    reader.close();
  }

  acrobat_restore_doc(data);
}
