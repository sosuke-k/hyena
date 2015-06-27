function write(pathAsString, text) {
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

    app.write(text, {
        to: file,
        as: 'text'
    });

    app.closeAccess(file);
}


function read(pathAsString) {
    'use strict';

    var app = Application.currentApplication();
    app.includeStandardAdditions = true;
    var path = Path(pathAsString);
    var file = app.openForAccess(path);

    var eof = app.getEof(file);

    var data = null;
    try {
        data = app.read(file, {
            to: eof
        });
    } catch (e) {
        return false;
    } finally {
        app.closeAccess(file);
    }
    return data;
}
