package main

import "html"
import "fmt"
func main() {
	const s = `"Fran & Freedle's Dinner" <tasty@example.com>`
	// It escapes only five such characters: <, >, &, ' and ".
	fmt.Println(html.EscapeString(s))

	// It unescapes a larger range of entities than EscapeString escapes.
	//"&aacute;" unescapes to "á"
	fmt.Println(html.UnescapeString("&aacute;&#225;"))
}