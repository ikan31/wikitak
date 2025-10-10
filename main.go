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

var URL_PREFIX = "https://en.wikipedia.org/wiki/"

func getWikipediaSummary(page string) (string, error) {
	fmt.Printf("Fetching summary for Wikipedia page: %s\n\n", page)

	wikiUrl := URL_PREFIX + page

	client := &http.Client{}
	req, err := http.NewRequest("GET", wikiUrl, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	// set a descriptive User-Agent per Wikipedia's robot policy
	req.Header.Set("User-Agent", "wikitak/1.0 (https://github.com/ikan31/wikitak; ami.ikanovic@gmail.com)")

	response, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch page: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		if response.StatusCode == http.StatusNotFound {
			return "", fmt.Errorf("page not found")
		}

		return "", fmt.Errorf("received non-200 response code: %d", response.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to parse HTML: %v", err)
	}

	// wikipedia article content is in
	// <div class="mw-content-ltr mw-parser-output" lang="en" dir="ltr">
	content := doc.Find("div.mw-content-ltr.mw-parser-output").First()

	// find the first <p> tag that contains a <b> tag
	// this is usually the first paragraph (the summary) as the
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

	// there is a case where the page is a "may refer to" page
	// for example if you pass in "george" https://en.wikipedia.org/wiki/George
	// the page will give you a list of articles that contain "george"
	// Names: George (given name), George (surname)
	// People: George (singer), ...
	// Places: George, South Africa, ...

	// In this case we want to parse this list and return it so the user can see
	// the possible articles they might be looking for

	// replace underscores with spaces
	title := strings.ReplaceAll(page, "_", " ")

	if strings.Contains(summary, fmt.Sprintf("%s may refer to:", title)) {
		// parse the list of articles
		var articles []string
		content.Find("li").Each(func(i int, s *goquery.Selection) {
			article := s.Text()

			// lets add the hyperlink if it exists
			// and what the wikitak command would be
			// so the result would be like:
			/*
				- George (given name)
				(indent)  https://en.wikipedia.org/wiki/George_(given_name)
				(indent)  wikitak "George_(given_name)"
			*/
			if link, exists := s.Find("a").Attr("href"); exists {
				if after, ok := strings.CutPrefix(link, "/wiki/"); ok {
					articleTitle := after
					article += fmt.Sprintf("\n  https://en.wikipedia.org%s\n  wikitak \"%s\"", link, articleTitle)
				}
			}

			articles = append(articles, article)
		})

		if len(articles) > 0 {
			summary = fmt.Sprintf("%s may refer to:\n", title)
			for _, article := range articles {
				summary += fmt.Sprintf("- %s\n", article)
			}
		}
	}

	// remove bracketed references using regex
	reBracket := regexp.MustCompile(`\[\d+\]`)
	summary = reBracket.ReplaceAllString(summary, "")

	return summary, nil
}

// wrapText wraps text to the specified width while preserving existing line breaks
func wrapText(text string, width int) string {
	var result string
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		// If the line is a link, command, or already indented, don't wrap it
		if strings.HasPrefix(trimmed, "http") || strings.HasPrefix(trimmed, "wikitak") || (len(line) > 0 && (line[0] == ' ' || line[0] == '\t')) {
			result += line
		} else {
			words := strings.Fields(line)
			var lineLen int
			for _, word := range words {
				space := 0
				if lineLen > 0 {
					space = 1
				}
				if lineLen+len(word)+space > width {
					result += "\n"
					lineLen = 0
					space = 0
				} else if lineLen > 0 {
					result += " "
					lineLen++
				}
				result += word
				lineLen += len(word)
			}
		}
		if i < len(lines)-1 {
			result += "\n"
		}
	}
	return result
}

var rootCmd = &cobra.Command{
	Use:                   "wikitak [page name]",
	Short:                 "wikitak is a cli tool for getting wikipedia article summaries.",
	Long:                  "wikitak is a cli tool for getting wikipedia article summaries.",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
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

		fmt.Println(wrapText(summary, 100))

		fmt.Println()

		fmt.Println("https://en.wikipedia.org/wiki/" + page)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing wikitak: '%s'\n", err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
