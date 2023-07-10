![Ducks](https://github.com/mitjafelicijan/gddg/assets/296714/b979dc03-b84c-450a-ab1b-728ea2fc3484)

# Go duck duck go

This is a small CLI utility that searches with the help of
[DuckDuckGo](https://duckduckgo.com/). By default it outputs results to
console. But it also supports other formats like JSON and XML.

This by itself can be useful, but there is also `--cgi` flag available and you
can use this as a cgi-bin utility.

## Install

```console
git clone git@github.com:mitjafelicijan/gddg.git
cd gddg
go install .
```

## Usage

```console
$ gddg -h
Usage: gddg [--format FORMAT] [--cgi] [--region REGION] [--list-regions] [TERM]

Positional arguments:
  TERM

Options:
  --format FORMAT, -f FORMAT
                         output format of results (text, json, xml) [default: text]
  --cgi                  outputs in a form that CGI understands
  --region REGION, -r REGION
                         region of the search results [default: en-us]
  --list-regions, -l     print all available regions and exist
  --help, -h             display this help and exit
```

**Region list includes**: ar-es, au-en, at-de, be-fr, be-nl, br-pt, bg-bg, ca-en,
ca-fr, ct-ca, cl-es, cn-zh, co-es, hr-hr, cz-cs, dk-da, ee-et, fi-fi, fr-fr,
de-de, gr-el, hk-tzh, hu-hu, is-is, in-en, id-en, ie-en, il-en, it-it, jp-jp,
kr-kr, lv-lv, lt-lt, my-en, mx-es, nl-nl, nz-en, no-no, pk-en, pe-es, ph-en,
pl-pl, pt-pt, ro-ro, ru-ru, xa-ar, sg-en, sk-sk, sl-sl, za-en, es-ca, es-es,
se-sv, ch-de, ch-fr, tw-tzh, th-en, tr-tr, us-en, us-es, ua-uk, uk-en, vn-en

## Available formats

### Text (default format)

You can get this output with `gddg 'niels bohr'`.

```text
1. Niels Bohr - Wikipedia
Niels Henrik David Bohr ( Danish: [ˈne̝ls ˈpoɐ̯ˀ]; 7 October 1885 - 18 November 1962) was a Danish physicist who made foundational contributions to understanding atomic structure and quantum theory, for which he received the Nobel Prize in Physics in 1922. Bohr was also a philosopher and a promoter of scientific research.
https://en.wikipedia.org/wiki/Niels_Bohr

2. Niels Bohr | Biography, Education, Accomplishments, & Facts
Niels Bohr, in full Niels Henrik David Bohr, (born October 7, 1885, Copenhagen, Denmark—died November 18, 1962, Copenhagen), Danish physicist who is generally regarded as one of the foremost physicists of the 20th century.
https://www.britannica.com/biography/Niels-Bohr
```

### JSON

You can get this output with `gddg 'niels bohr' -f json`.

```json
[
  {
    "title": "Niels Bohr - Wikipedia",
	"link": "https://en.wikipedia.org/wiki/Niels_Bohr",
	"snippet": "Niels Henrik David Bohr ( Danish: [ˈne̝ls ˈpoɐ̯ˀ]; 7 October 1885 - 18 November 1962) was a Danish physicist who made foundational contributions to understanding atomic structure and quantum theory, for which he received the Nobel Prize in Physics in 1922. Bohr was also a philosopher and a promoter of scientific research."
  }, {
	"title": "Niels Bohr | Biography, Education, Accomplishments, & Facts",
	"link": "https://www.britannica.com/biography/Niels-Bohr",
	"snippet": "November 18, 1962 (aged 77) Copenhagen Denmark Awards And Honors: Copley Medal (1938) Nobel Prize (1922) Notable Family Members: son Aage N. Bohr brother Harald August Bohr Subjects Of Study: atomic model atomic theory complementarity principle correspondence principle liquid-drop model ... (Show more) See all related content → Top Questions"
  }
]
```

### XML

You can get this output with `gddg 'niels bohr' -f xml`.

```xml
<results>
  <item>
    <title>Niels Bohr - Wikipedia</title>
    <link>https://en.wikipedia.org/wiki/Niels_Bohr</link>
    <snippet>Niels Henrik David Bohr ( Danish: [ˈne̝ls ˈpoɐ̯ˀ]; 7 October 1885 - 18 November 1962) was a Danish physicist who made foundational contributions to understanding atomic structure and quantum theory, for which he received the Nobel Prize in Physics in 1922. Bohr was also a philosopher and a promoter of scientific research.</snippet>
  </item>
  <item>
	<title>Niels Bohr | Biography, Education, Accomplishments, &amp; Facts</title>
    <link>https://www.britannica.com/biography/Niels-Bohr</link>
    <snippet>November 18, 1962 (aged 77) Copenhagen Denmark Awards And Honors: Copley Medal (1938) Nobel Prize (1922) Notable Family Members: son Aage N. Bohr brother Harald August Bohr Subjects Of Study: atomic model atomic theory complementarity principle correspondence principle liquid-drop model ... (Show more) See all related content → Top Questions</snippet>
  </item>
</results>
```

## cgi-bin

This can also be used as cgi bin script and it respects [Common Gateway
Interface](https://en.wikipedia.org/wiki/Common_Gateway_Interface).

If you pass `QUERY_STRING` as environmental variable as per specification, to
the utility it will use that as a search term.

## License

[gddg](https://github.com/mitjafelicijan/gddg) was written by [Mitja
Felicijan](https://mitjafelicijan.com) and is released under the BSD two-clause
license, see the LICENSE file for more information.
