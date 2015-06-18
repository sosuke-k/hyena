app = Application('Google Chrome');

function chrome_get_tab_info(){
    windows = app.windows();
    var res = {};

    for (i = 0;i < windows.length;i++){
	    w = windows[i];
    	res[i] = [];
    	for (j = 0;j < w.tabs().length;j++){
    		res[i].push(w.tabs[j].url());
    	}
    }
    var s = JSON.stringify(res);
    return s;
}

chrome_get_tab_info();