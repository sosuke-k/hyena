JsOsaDAS1.001.00bplist00�Vscript_�slack = Application('Slack');
slack.activate();

delay(0.1);

systemEvent = Application('System Events');
systemEvent.includeStandardAdditions = true;

emoji = ':full_moon_with_face:';
systemEvent.setTheClipboardTo(emoji);

process = systemEvent.processes.byName('Slack');
editMenu = process.menuBars[0].menuBarItems.byName('Edit');
paste = editMenu.menus[0].menuItems.byName('Paste');

paste.click();
systemEvent.keyCode(36);                              �jscr  ��ޭ