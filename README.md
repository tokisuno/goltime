# Goltime (read as "goal time")

## What this is?
Goltime is a CLI tool (functionally a script) written in Go!
* used to output how many minutes are left in the given day
* script can be used with any other interface that accepts a command output as a valid form of input

This is a simple script and the first CLI tool I have build for myself, so error-handling and testing isn't verbose and robust by any means. Creating the base script ``root.go`` is something I do whenever I am using a new language. I've been using Go recently and have been really loving it, so creating this small project was a test of mine to see if I could actually solve a problem I have in my day-to-day.

## Example
```bash
~/code/golang/goltime % ./bin/goltime | figlet
  __   _ _  _
 / /_ / | || |
| '_ \| | || |_
| (_) | |__   _|
 \___/|_|  |_|

```

## Installation
1. Download & extract the .zip contents into a folder
2. Installing
    * Run the ``install.sh`` file inside of the project folder

    OR

    * Execute ``go build -o bin/goltime main.go`` inside the project folder
3. Enjoy!

## Usage
* ``goltime``: Returns minutes left until midnight in your local time
    - Does not print \n to be compatible with my Qtile bar/other widgets
* ``goltime endt {TIME}``: Returns time left until user-defined endpoint
    - Time must be written in HHMM format (ex. 1430)
    - Time must be set in the future
        * Rejects time that's in the past (for now, read TODO)

# TODO
- [ ] Finish endt
    - [ ] Don't crash when the user's input has reached past the "goal time"
- [ ] Add option to display string with text
    - [ ] Default: ("{minutes} minutes left")
    - [ ] Roll over until new start time (default: 00:00)
- [ ] Add option to output with newline
