# Hygge Media

Hygge Media is a command line application designed to rename and organize media files in a directory to match the format prescribed by Emby. This tool simplifies the organization of your media library, ensuring that your files are named consistently and correctly.

## Features

- Renames media files in a specified directory.
- Organizes media files into season and episode folders.
- Follows the Emby naming conventions for media files.

## Installation

To install Hygge Media, clone the repository and build the application:

```bash
git clone https://github.com/yourusername/hyggemedia.git
cd hyggemedia
make build
```

## Usage

To use Hygge Media, run the following command in your terminal:

```bash
./hyggemedia organize --src-dir <source_directory> --dest-dir <destination_directory> --title "<show_title>"
```

Replace <source_directory> with the path to the directory containing the media files you want to organize, <destination_directory> with the path to the directory where you want the organized files to be placed, and <show_title> with the title of the show.

## Examples
### Organize media files for the show "Friends"
```bash
./hyggemedia organize --src-dir /path/to/source --dest-dir /path/to/destination --title "Friends"
```
### Perform a dry run to see what changes would be made without actually making them
```bash
./hyggemedia organize --src-dir /path/to/source --dest-dir /path/to/destination --title "Friends" --dry-run
```

# Contributing
Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

# License
This project is licensed under the MIT License. See the LICENSE file for details. 