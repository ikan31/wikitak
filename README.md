# Wikitak

Wikitak returns the first paragraph (summary) of a wikipedia article.

## Installation
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

## Example

```
ikan31@air wikitak % wikitak taekwondo
Fetching summary for Wikipedia page: Taekwondo

Taekwondo (/ˌtaɪkwɒnˈdoʊ, ˌtaɪˈkwɒndoʊ, ˌtɛkwənˈdoʊ/; Korean: 태권도; [t̪ʰɛ.k͈wʌ̹n.d̪o] ⓘ) is a Korean martial art and combat sport involving primarily kicking techniques and punching. "Taekwondo" can be translated as tae ("strike with foot"), kwon ("strike with hand"), and do ("the art or way"). In addition to its five tenets of courtesy, integrity, perseverance, self-control and indomitable spirit, the sport requires three physical skills: poomsae (품새, Form), kyorugi (겨루기, Sparring) and gyeokpa (격파, Breaking Technique).
```

## Donate

Wikipedia is a great tool. Consider donating directly to them to keep them going. Imagine if ads started popping up on their pages. The horror!

[Wikipedia Donation](https://donate.wikimedia.org/w/index.php?title=Special:LandingPage&country=GB&uselang=en&wmf_medium=portal&wmf_source=portalFooter&wmf_campaign=portalFooter)

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
