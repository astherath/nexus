# Nexus

Nexus is a Go-driven command line interface used for fetching, parsing, and displaying match data from major professional League of Legends regions

## Installation

The latest version of Go can be downloaded from the [Golang Website](https://golang.org/dl/)

Once installed, use ```go get``` to install the package

```bash
go get -u github.com/astherath/nexus
```

To verify a successful installation use
```bash
nexus version
```

## Usage
To initially fetch the match data for a region:
```bash
nexus fetch [region abbreviation]
```
Currently the supported regions are: LCK, LPL, LEC, LCS.

To display the upcoming matches for the region:
```bash
nexus upcoming -a
```

At any point use ```nexus [command] --help``` for detailed use.

## Troubleshooting
The most common error is:
```bash
package github.com/astherath/nexus: cannot find package "github.com/astherath/nexus" in any of:
	/usr/local/go/src/github.com/astherath/nexus (from $GOROOT)
	/root/work/src/github.com/astherath/nexus (from $GOPATH)
```

To fix, re-instate the ```$PATH``` and ```$GOPATH``` variables (assuming default installation):
[Setting GOPATH](https://golang.org/doc/code.html#GOPATH)

For any other problems, please open an issue.


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[Apache](http://www.apache.org/licenses/LICENSE-2.0)
