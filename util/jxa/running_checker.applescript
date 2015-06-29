#!/usr/bin/env osascript -l JavaScript

function run(argv) {
    console.log("check if " + argv[0] + " is running");
    var isRunning = check(argv[0]);
    return isRunning;
}

function check(identifier) {

    // Look for iTunes
    ObjC.import('stdlib')
    ObjC.import('AppKit')
    var isRunning = false
    var apps = $.NSWorkspace.sharedWorkspace.runningApplications // Note these never take () unless they have arguments
    apps = ObjC.unwrap(apps) // Unwrap the NSArray instance to a normal JS array
    var app, itunes
    for (var i = 0, j = apps.length; i < j; i++) {
        app = apps[i]

        // Another option for comparison is to unwrap app.bundleIdentifier
        // ObjC.unwrap(app.bundleIdentifier) === 'org.whatever.Name'

        // Some applications do not have a bundleIdentifier as an NSString
        if (typeof app.bundleIdentifier.isEqualToString === 'undefined') {
            continue;
        }

        if (app.bundleIdentifier.isEqualToString(identifier)) {
            isRunning = true;
            break;
        }
    }

    //if (!isRunning) {
    //  $.exit(1)
    //}

    return isRunning;

}