# Hygge Media

![Go](https://github.com/simonnordberg/hyggemedia/actions/workflows/go.yml/badge.svg)

Hygge Media is a command line application designed to rename and organize media files in a directory to match the format prescribed by [Emby](https://emby.media/). This tool simplifies the organization of your media library, ensuring that your files are named consistently and correctly.

## Features

- Renames media files in a specified directory.
- Organizes media files into season and episode folders.
- Follows the [Emby naming naming conventions](https://emby.media/support/articles/TV-Naming.html) for media files.

## Installation

To install Hygge Media, clone the repository and build the application:

```bash
git clone https://github.com/simonnordberg/hyggemedia.git
cd hyggemedia
make build
```

## Usage

To use Hygge Media, run the following command in your terminal:

```bash
./hyggemedia tv \ 
    --source-dir <source directory> \
    --target-dir <target directory> \
    --title "<title>"
```

Replace <source directory> with the path to the directory containing the media files you want to organize, <target directory> with the path to the directory where you want the organized files to be placed, and <title> with the title of the show.

**Note:** Files will be copied by default (rather than moved) as a non-destructive operation.

## Command Line Options

### Global Options
- `--title, -t` (required): Title of the show.
- `--exec, -n`: Perform a dry run without making changes.
- `--move, -m`: Move files instead of copying.
- `--source-dir, -s` (required): Source directory to scan for files.
- `--target-dir, -d` (required): Target directory to organize files into.

## Examples
### Organize media files for the show "Friends"
```bash
./hyggemedia tv \ 
    --source-dir /path/to/source \
    --target-dir /path/to/target \
    --title "Friends"

./hyggemedia organize \
    --src-dir /path/to/source \
    --dest-dir /path/to/destination \
    --title "Friends"
```
### Execute the changes, i.e. actually copy/move (-m) the files
```bash
./hyggemedia organize \
    --src-dir /path/to/source \
    --dest-dir /path/to/destination \
    --title "Friends" \
    --exec
```
### Move files instead of copying them
```bash
./hyggemedia organize \
    --src-dir /path/to/source \
    --dest-dir /path/to/destination \
    --title "Friends" \
    --move \
    --exec
```

# Contributing
Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

# License
This project is licensed under the MIT License. See the LICENSE file for details.