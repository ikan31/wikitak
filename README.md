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

**NOTE:** If the page has spaces you need to wrap the argument in double quotes (`"`)

For the page name, the tool will replace any spaces with `_` and it will capitalize the first letter. This is to ensure it follows Wikipedia's standard format for page names. 

The following return the same article:

``` sh
wikitak "Artifitial Intelligence" 
wikitak artificial_intelligence
wikitak Artificial_intelligence
```

## Examples
Returns summary of taekwondo:

```
$ wikitak taekwondo

Fetching summary for Wikipedia page: Taekwondo

Taekwondo (/ˌtaɪkwɒnˈdoʊ, ˌtaɪˈkwɒndoʊ, ˌtɛkwənˈdoʊ/; Korean: 태권도; [t̪ʰɛ.k͈wʌ̹n.d̪o] ⓘ) is a Korean martial art and combat sport involving primarily kicking techniques and punching. "Taekwondo" can be translated as tae ("strike with foot"), kwon ("strike with hand"), and do ("the art or way"). In addition to its five tenets of courtesy, integrity, perseverance, self-control and indomitable spirit, the sport requires three physical skills: poomsae (품새, Form), kyorugi (겨루기, Sparring) and gyeokpa (격파, Breaking Technique).

https://en.wikipedia.org/wiki/Taekwondo
```

Article with space:

```
$ wikitak "George Papagheorghe"

Fetching summary for Wikipedia page: George_Papagheorghe

George Papagheorghe (born 22 June 1982, in Constanța) a.k.a. Jorge and at times GEØRGE is a
Romanian singer, dancer and TV host, specially for the talent competition România dansează on
Antena 1 starting 2013.

https://en.wikipedia.org/wiki/George_Papagheorghe
```

"may refer to" page result:
```
$ wikitak tommy                
Fetching summary for Wikipedia page: Tommy

Tommy may refer to:
- Tommy (given name), a list of people and fictional characters
  https://en.wikipedia.org/wiki/Tommy_(given_name)
  wikitak "Tommy_(given_name)"
- Tommy Atkins, or just Tommy, a slang term for a common soldier in the British Army
  https://en.wikipedia.org/wiki/Tommy_Atkins
  wikitak "Tommy_Atkins"
- Tommy Giacomelli (born 1974), Brazilian former footballer also known as simply Tommy
  https://en.wikipedia.org/wiki/Tommy_Giacomelli
  wikitak "Tommy_Giacomelli"
- Tommy (1931 film), a Soviet drama film
  https://en.wikipedia.org/wiki/Tommy_(1931_film)
  wikitak "Tommy_(1931_film)"
- Tomm (disambiguation)
  https://en.wikipedia.org/wiki/Tomm_(disambiguation)
  wikitak "Tomm_(disambiguation)"
- Thomas (disambiguation)
  https://en.wikipedia.org/wiki/Thomas_(disambiguation)
  wikitak "Thomas_(disambiguation)"

https://en.wikipedia.org/wiki/Tommy
```

## Donate

Wikipedia is a great tool. Consider donating directly to them to keep them going. Imagine if ads started popping up on their pages. The horror!

[Wikipedia Donation](https://donate.wikimedia.org/w/index.php?title=Special:LandingPage&country=GB&uselang=en&wmf_medium=portal&wmf_source=portalFooter&wmf_campaign=portalFooter)

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
