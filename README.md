# Wikitak

Wikitak returns the first paragraph of a wikipedia article.

## Installation

You can install Wikitak using Go (requires Go 1.18+):

```sh
go install github.com/ikan31/wikitak@latest
```

This will install the `wikitak` binary in your `$GOPATH/bin` or `$HOME/go/bin` directory.

## Usage

After installation, run:

```sh
wikitak [page name]
```

For the page name, the tool will replace any spaces with `_` and it will capitalize the first letter. This is to ensure it follows Wikipedia's standard format for page names. 

The following return the same article:

``` sh
wikitak Artifitial Intelligence 
wikitak artificial_intelligence
wikitak Artificial_intelligence
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
