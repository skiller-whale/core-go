// Search a web page for a list of words
package main

import "net/http"
import "io"
import "strings"
import "fmt"

// FIXME:
// 1. Embed this interface into WebSearch
// 2. Create two types which implement it based on SearchSimple and SearchIndexed
// 3. Change the use of WebSearch{} in main() to check you can switch between the implementations

// type Searcher interface {
// 	IndexWords(words []string)
// 	Search(word string) bool
// }

type WebSearch struct {
	Url   string
	words []string
}

func (ws *WebSearch) SearchSimple(word string) bool {
	// Just check every word in the list to see if it matches this one.
	for _, w := range ws.words {
		if w == word {
			return true
		}
	}
	return false
}

// func (ws *WebSearch) SearchIndexed(word string) bool {
// 	// First build an index of all the words (slow)
// 	var index map[string]bool = make(map[string]bool)
// 	for _, w := range ws.words {
// 		index[w] = true
// 	}
// 	// Then look up the word in the index (fast)
// 	return index[word]
// }

func (ws *WebSearch) fetch() []string {
	resp, e := http.Get(ws.Url)
	if e != nil {
		panic(e)
	}
	defer resp.Body.Close()
	bytes, e := io.ReadAll(resp.Body)
	if e != nil {
		panic(e)
	}
	return strings.Split(string(bytes), " ")
}

func (l *WebSearch) Index() {
	// Just store the words in memory
	l.words = l.fetch()
}

var wordsToFind string = `Ishmael hark realism Ahab fathomless chasm Victoria brave abyss seaman boat racing submarine whale harpoon`

func main() {
	var ws WebSearch = WebSearch{Url: "http://www.gutenberg.org/files/2701/old/moby10b.txt"} // Moby Dick by Herman Melville
	ws.Index()
	for _, w := range strings.Split(wordsToFind, " ") {
		fmt.Println(w, "→", ws.SearchSimple(w))
	}
}
