An implementation of a command line text editor written in Golang as part of a course project for CS455 (2023-24 I Sem.) in a team of four (Abhishek ["abhishekshree"](https://github.com/abhishekshree) Shree, Jahnvi ["Janhvi-Rochwani"](https://github.com/Janhvi-Rochwani) Rochwani, Parinay ["parinayc20"](https://github.com/parinayc20) Chauhan, and me, Aditya ["cliche-niche"](https://github.com/cliche-niche) Tanwar).

## Running
To compile the project, please ensure you have `1.21x` (or a higher) version of Go installed. Thereafter, you can simply clone the repository and run the following commands:
```bash
go build main.go
./main -location="."
```
Alternatively, after <i>building</i> the project, you can also run `./main -location=filename.txt` or `./main`. The differences between the three usages are:
+ `./main -location="path/to/directory/`: Opens the editor in the directory given in the path with an interactive view of the directory available on the left.
+ `./main -location="."`: Opens the editor in the same directory as the executable is in.
+ `./main -location="filename.ext"`: Opens just the give file (and not a directory).

## Features
Upon opening a file, the editor supports features like:
+ Basic editing (cut, copy, paste, undo, redo, etc.)
+ Navigation using keys as well as key combinations
+ Navigation using scrolling
+ Autosave (at regular intervals of 30s)
+ Reminder to save a file before closing
+ Keyboard shortcuts:
    + `Ctrl-C`: Close the application
    + `Ctrl-Q`: Copy text if there is any selection. If there is no selection, copy the line the cursor is in
    + `Ctrl-X`: Cut text if there is any selection. If there is no selection, cut the line the cursor is in
    + `Ctrl-V`: Paste previously copied/cut text
    + `Ctrl-S`: Save all changes to the file
    + `Ctrl-O`: Toggle autosave on/off (initially off)
    + `F1`: Display help containing shortcuts and more
    + Some of these shortcuts are available on the front page as well when the application is run