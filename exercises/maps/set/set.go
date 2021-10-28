package main

import (
	"fmt"
	"strings"
)

func main() {
	// TODO: Declare the set variable here


	// TODO: Populate set, so it represents the set of words in the slice
	// returned by the getWords() function.

	// search(set)
}

func search(words map[string]bool) {
	for {
		fmt.Print("\nEnter your search term: ")
		var word string
		fmt.Scanln(&word)

		// TODO: Print "Found" if word is in words, and "Not Found" otherwise

	}
}

/* -------------------------------------------------- */
/* You don't need to worry about any code below here. */
/* -------------------------------------------------- */

const text = `Be thou blest, Bertram, and succeed thy father
In manners, as in shape! thy blood and virtue
Contend for empire in thee, and thy goodness
Share with thy birthright! Love all, trust a few,
Do wrong to none: be able for thine enemy
Rather in power than use, and keep thy friend
Under thy own life's key: be cheque'd for silence,
But never tax'd for speech. What heaven more will,
That thee may furnish and my prayers pluck down,
Fall on thy head! Farewell, my lord;
'Tis an unseason'd courtier; good my lord,
Advise him.`

// Return all the words from `text` in a slice
// This transforms all words to lowercase, but does not remove punctuation
func getWords() []string {
	return strings.Split(
		strings.ReplaceAll(strings.ToLower(text), "\n", " "),
		" ",
	)
}
