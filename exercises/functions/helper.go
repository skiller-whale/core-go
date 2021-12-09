package main

type Repository struct {
	slug     string
	link     string
	isPublic bool
}

type User struct {
	username     string
	email        string
	repositories []Repository
}

var user = User{
	username: "octocat",
	email:    "octocat@github.com",
	repositories: []Repository{
		{
			slug:     "Hello-World",
			isPublic: true,
		},
		{
			slug:     "octocat.github.io",
			isPublic: true,
		},
		{
			slug:     "git-consortium",
			isPublic: true,
		},
		{
			slug:     "octocat-secrets",
			isPublic: false,
		},
	},
}

var description = map[string]string{
	"https://github.com/octocat/Hello-World":       "My first repository on GitHub!",
	"https://github.com/octocat/octocat.github.io": "My website!",
	"https://github.com/octocat/git-consortium":    "This repo is for demonstration purposes only.",
}

// For the purpose of the exercise, assume that
// fetchDescription fetches description from github
// If the link corresponds to a private repo,
// it terminates the program with a panic
func fetchDescription(link string) string {
	if link == "https://github.com/octocat/octocat-secrets" {
		panic("You don't have permission to access this URL: " + link)
	}
	return description[link]
}
