package crawlerService

import (
    "strings"
	"fmt"

    "golang.org/x/net/html"
)

// ExtractDataDays extracts the data from div elements with the specified class and data-day attribute.
func ExtractDataDays(body string) []string {
    var dataDays []string
    tokenizer := html.NewTokenizer(strings.NewReader(body))

    for {
        nextToken := tokenizer.Next()
        switch nextToken {
        case html.ErrorToken:
            // End of the document
            return dataDays
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			if token.Data == "div" {
				var hasClass, hasDataDay bool
				var dataDayValue string
				for _, attr := range token.Attr {
					if attr.Key == "class" && strings.Contains(attr.Val, "col-12 col-sm-4 col-md-3 col-lg-3 mx-0 day-wrappper") {
						hasClass = true
					}
					if attr.Key == "data-day" {
						hasDataDay = true
						dataDayValue = attr.Val
					}
				}
				if hasClass && hasDataDay {
					fmt.Printf("Found data-day: %s\n", dataDayValue)
					dataDays = append(dataDays, dataDayValue)

					//more logic come here
				}
			}
		}
    }
}