[![CircleCI](https://circleci.com/gh/neko-neko/utmpdump.svg?style=svg)](https://circleci.com/gh/neko-neko/utmpdump)

# utmpdump
utmpdump is output utmp format file to json or tsv or csv.  

## Installation
**[Download latest binary](https://github.com/neko-neko/utmpdump/releases/latest)**

Multi-platform support
- Windows
- Mac OS X
- Linux 32,64bit

Or get the library
```
$ go get github.com/neko-neko/utmpdump
```
Or clone the repository and run
```
$ go install github.com/neko-neko/utmpdump
```

## Usage
```
Usage of utmpdump: utmpdump [options]
-f, --file <file> load a specific file instead of utmp
-t, --until <YYYYMMDDHHMMSS> display the lines until the specified time
-s, --since <YYYYMMDDHHMMSS> display the lines since the specified time
-o, --output <json/tsv/csv> display the lines at specific format. Default format is json.
```

### Example
```
$ utmpdump -f /var/log/wtmp
```

## Contributing
1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D

## Credits
neko-neko

## Contributors
[stapelberg](https://github.com/stapelberg)

## License
MIT
