// Search a web page for a list of words
package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Embed this interface into WebSearch, then refactor the indexing & searching out into a separate type.
//
// type Searcher interface {
// 	IndexWords(words []string)
// 	Search(word string) bool
// }

type WebSearch struct {
	Url   string
	words []string // An implementation detail to be refactored out
}

func (ws *WebSearch) Search(word string) bool {
	// Just check every word in the list to see if it matches this one.
	for _, w := range ws.words {
		if w == word {
			return true
		}
	}
	return false
}

// (optional) An alternative implementation to attach to a new `IndexedSearcher` type:
//
// func (ws *WebSearch) Search(word string) bool {
// 	// First build an index of all the words (slow)
// 	var index map[string]bool = make(map[string]bool)
// 	for _, w := range ws.words {
// 		index[w] = true
// 	}
// 	// Then look up the word in the index (fast)
// 	return index[word]
// }

func (l *WebSearch) IndexWords(words []string) {
	// just store them to search later, nothing clever
	l.words = words
}

// fetch and FetchAndIndex methods don't need changing! 👇🏻

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

func (l *WebSearch) FetchAndIndex() {
	// Just store the words in memory
	l.IndexWords(l.fetch())
}

func main() {
	var wordsToFind string = `Ishmael hark realism Ahab fathomless chasm Victoria brave abyss seaman boat racing submarine whale harpoon`
	var ws WebSearch = WebSearch{Url: "http://www.gutenberg.org/files/2701/old/moby10b.txt"} // Moby Dick by Herman Melville

	ws.FetchAndIndex()

	for _, w := range strings.Split(wordsToFind, " ") {
		fmt.Println(w, "→", ws.Search(w))
	}
}
