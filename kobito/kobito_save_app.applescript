function kobito_get_window_info(){
	systemEvent = Application('System Events');
	systemEvent.includeStandardAdditions = true;

	kobitoProcess = systemEvent.processes.byName('Kobito');

	windowMenu = kobitoProcess.menuBars[0].menuBarItems.byName('Window');

	// TODO:
	// For loop to get windows' list
	// JSON.stringify
	// return JSON string
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


function run(argv){
	var activeWindows = kobito_get_window_info();
	
	var exportFileWriter = fileWriter(argv);
	exportFileWriter.write(activeWindows);
	
	exportFileWriter.close();
}
