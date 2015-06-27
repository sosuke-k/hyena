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
