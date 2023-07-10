package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/alexflint/go-arg"
	"github.com/anaskhan96/soup"
)

type ResultItem struct {
	XMLName xml.Name `json:"-" xml:"item"`
	Title   string   `json:"title" xml:"title"`
	Link    string   `json:"link" xml:"link"`
	Snippet string   `json:"snippet" xml:"snippet"`
}

func search(query string, region string) []ResultItem {
	params := url.Values{
		"q":  {query},
		"kl": {query},
	}
	encodedURL := fmt.Sprintf("https://html.duckduckgo.com/html/?%s", params.Encode())

	client := &http.Client{}
	req, err := http.NewRequest("POST", encodedURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		os.Exit(1)
	}
	defer response.Body.Close()

	htmlBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		os.Exit(1)
	}

	doc := soup.HTMLParse(string(htmlBytes))
	results := []ResultItem{}

	links := doc.FindAll("div", "class", "result__body")
	for _, item := range links {
		var title string = item.Find("a", "class", "result__a").FullText()
		var link string = item.Find("a", "class", "result__a").Attrs()["href"]
		var snippet string = item.Find("a", "class", "result__snippet").FullText()

		results = append(results, ResultItem{
			Title:   title,
			Link:    link,
			Snippet: snippet,
		})
	}

	return results
}

func main() {
	var args struct {
		Format      string `arg:"-f,--format" default:"text" help:"output format of results (text, json, xml)"`
		CGI         bool   `arg:"--cgi" help:"outputs in a form that CGI understands"`
		Term        string `arg:"positional,env:QUERY_STRING"`
		Region      string `arg:"-r,--region" default:"en-us" help:"region of the search results"`
		ListRegions bool   `arg:"-l,--list-regions" help:"print all available regions and exist"`
	}

	arg.MustParse(&args)

	// Check if format is correct one.
	if args.Format != "text" && args.Format != "json" && args.Format != "xml" {
		fmt.Printf("Provided format `%s` is not supported. Supported formats are (text, json, xml).\n", args.Format)
		os.Exit(1)
	}

	// Print all available regions and exist
	if args.ListRegions {
		fmt.Println("List of available regions:")
		fmt.Println("ar-es, au-en, at-de, be-fr, be-nl, br-pt, bg-bg, ca-en, ca-fr, ct-ca, cl-es, cn-zh, co-es, hr-hr, cz-cs, dk-da, ee-et, fi-fi, fr-fr, de-de, gr-el, hk-tzh, hu-hu, is-is, in-en, id-en, ie-en, il-en, it-it, jp-jp, kr-kr, lv-lv, lt-lt, my-en, mx-es, nl-nl, nz-en, no-no, pk-en, pe-es, ph-en, pl-pl, pt-pt, ro-ro, ru-ru, xa-ar, sg-en, sk-sk, sl-sl, za-en, es-ca, es-es, se-sv, ch-de, ch-fr, tw-tzh, th-en, tr-tr, us-en, us-es, ua-uk, uk-en, vn-en")
		os.Exit(0)
	}

	// If no term provided exit.
	if len(args.Term) == 0 {
		fmt.Println("Search term must be provided. Try with `ddg 'niels bohr'`")
		os.Exit(1)
	}

	// If CGI is enable print obligatory header first.
	if args.CGI {
		fmt.Println("Content-type: text/plain\n")
	}

	// Fetch results.
	results := search(args.Term, args.Region)

	// Handle results display.
	switch args.Format {
	case "text":
		for idx, item := range results {
			fmt.Printf("%d. %s\n", idx+1, item.Title)
			fmt.Printf("%s\n", item.Snippet)
			fmt.Printf("%s\n\n", item.Link)
		}
	case "json":
		jsonData, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			fmt.Println("Error marshaling to JSON:", err)
			os.Exit(1)
		}
		fmt.Println(string(jsonData))
	case "xml":
		xmlData, err := xml.MarshalIndent(results, "", "  ")
		if err != nil {
			fmt.Println("Error marshaling to XML:", err)
			os.Exit(1)
		}
		fmt.Printf("<results>%s</results>", string(xmlData))
	}
}
