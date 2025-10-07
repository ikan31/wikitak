package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

func getWikipediaSummary(page string) (string, error) {
	fmt.Printf("Fetching summary for Wikipedia page: %s\n\n", page)

	url := "https://en.wikipedia.org/wiki/" + page
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	// Set a descriptive User-Agent per Wikipedia's robot policy
	req.Header.Set("User-Agent", "wikitak/1.0 (https://github.com/ikan31/wikitak; ami.ikanovic@gmail.com)")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch page: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return "", fmt.Errorf("page not found")
		}

		return "", fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to parse HTML: %v", err)
	}

	// wikipedia article content is in
	// <div class="mw-content-ltr mw-parser-output" lang="en" dir="ltr">
	content := doc.Find("div.mw-content-ltr.mw-parser-output").First()

	// find the first <p> tag that contains a <b> tag
	// this is usually the first paragraph as the
	// as the first paragraph usually starts with the article title in bold
	var summary string
	content.Find("p").EachWithBreak(func(i int, s *goquery.Selection) bool {
		if s.Find("b").Length() > 0 {
			summary = s.Text()
			return false // break the loop
		}
		return true // continue the loop
	})

	if summary == "" {
		return "", fmt.Errorf("failed to find summary paragraph")
	}

	// Remove bracketed references using regex
	reBracket := regexp.MustCompile(`\[\d+\]`)
	summary = reBracket.ReplaceAllString(summary, "")

	return summary, nil
}

var rootCmd = &cobra.Command{
	Use:   "wikitak [page name]",
	Short: "wikitak is a cli tool for getting wikipedia article summaries.",
	Long:  "wikitak is a cli tool for getting wikipedia article summaries.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		page := args[0]

		// if page contains spaces, replace them with underscores
		// as wikipedia uses underscores in urls for spaces
		// e.g. "New York City" -> "New_York_City"
		page = strings.ReplaceAll(page, " ", "_")

		// if the page first letter is lowercase, capitalize it
		if len(page) > 0 && strings.ToLower(string(page[0])) == string(page[0]) {
			page = strings.ToUpper(string(page[0])) + page[1:]
		}

		summary, err := getWikipediaSummary(page)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(summary)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing wikitak '%s'\n", err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
