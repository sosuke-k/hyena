app = Application('Google Chrome');

function chrome_get_tab_info() {
    windows = app.windows();
    var res = {};

    for (i = 0; i < windows.length; i++) {
        w = windows[i];
        res[i] = [];
        for (j = 0; j < w.tabs().length; j++) {
            res[i].push(w.tabs[j].url());
        }
    }
    var s = JSON.stringify(res);
    return s;
}

function chrome_restore_tabs(json_string) {
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

chrome_restore_tabs(chrome_get_tab_info());
