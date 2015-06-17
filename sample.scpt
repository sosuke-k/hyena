# display dialog "display dialog 基本"

set the_file to choose file with prompt "Select a text file:"

tell application "Google Chrome"
  # activate
  # set vURL to get URL of active tab of first window
  # copy vURL to stdout

  set out to "[" & return
  set window_list to every window                          # get the windows

  set list_size to count of window_list
  set tmp to "" & list_size & return
  # copy tmp to stdout
  set counter to 0

  repeat with the_window in window_list                    # for every window
    set out to out & "  [" & return
    set tab_list to every tab in the_window                # get the tabs
    repeat with the_tab in tab_list                        # for every tab
      # set stdout to get URL of the_tab
      set vURL to get URL of the_tab
      set out to out & "    '" & vURL & "'," & return
      # set the_title to the title of the_tab                # grab the title
       # set titleString to titleString & the_title & return # concatenate then all
    end repeat
    set out to out & "  ]" & return
    if not (list_size = counter) then
      set out to out & "," & return
    end if
    set counter to counter + 1
  end repeat
  set out to out & "]" & return
  # copy out to stdout
end tell

set tmp to "test"
copy tmp to stdout

# set the_file to (((path to desktop) as string) & "test.txt") as file specification

set record_1 to {record_name:"Record A", record_value:{"This", "is", "a", "list.", 1, 2, 3}}
set record_2 to {record_name:"Record B", record_value:"A string."}
set record_3 to {record_name:"Record C", record_value:1.0 as real}
set the_data to {record_1, record_2, record_3}

try
   open for access the_file with write permission
   set eof of the_file to 0
   write out to the_file starting at eof
   close access the_file
on error
   try
       close access the_file
   end try
end try

# return out
