#!/usr/bin/env osascript -l JavaScript

function chrome_restore_tabs(json_string) {
    app = Application('Google Chrome');
	var data = JSON.parse(json_string);
    var n_windows = Object.keys(data).length;

    for (i = 0; i < n_windows; i++) { // i: window_id
        new_window = app.Window().make();
        var urls = data[i];

        for (j = 0; j < urls.length; j++) {
            var tab = app.Tab({
                url: urls[j]
            });
            new_window.tabs.push(tab);
        }

        new_window.tabs[0].close();
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

	chrome_restore_tabs(data);
}
