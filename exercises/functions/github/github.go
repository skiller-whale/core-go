// package github exists as a dummy data source for exercises.
//
// For the purpose of these exercises, you should treat github as if it is an
// external dependency that you cannot change.
package github

import "fmt"

type repo struct {
	description string
	public      bool
}

var repos map[string]repo = map[string]repo{
	"https://github.com/octocat/Hello-World": {
		description: "My first repository on GitHub!",
		public:      true,
	},
	"https://github.com/octocat/octocat.github.io": {
		description: "My website!",
		public:      true,
	},
	"https://github.com/octocat/git-consortium": {
		description: "This repo is for demonstration purposes only.",
		public:      true,
	},
	"https://github.com/octocat/octocat-secrets": {
		description: "None of your business",
		public:      false,
	},
}

// GetDescription returns the description string for a github repo
// It will panic if used to retrieve the description for a private repo
//
// Note: In real life this function would probably return an error code,
// it would be very unusual to raise a panic for this sort of situtation.
func GetDescription(url string) string {
	repo, ok := repos[url]
	if !ok {
		fmt.Println("No repo found at: \"" + url + "\"")
		return ""
	}

	if !repo.public {
		panic("Repo is private. Cannot access " + url)
	}

	return repo.description
}

// IsPublic returns a boolean declaring whether or not a repo is public
func IsPublic(url string) bool {
	repo, ok := repos[url]
	if !ok {
		fmt.Println("No repo found at", url)
	}
	return repo.public
}
