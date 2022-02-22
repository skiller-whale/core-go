package main

import (
	"fmt"
	// "github.com/skiller-whale/core-go/exercises/functions/github"
)

type Repository struct {
	path string
	url string
}

type User struct {
	username     string
	email        string
	repositories []Repository
}

func main() {
	user := User{
		username: "octocat",
		email:    "octocat@github.com",
		repositories: []Repository{
			{path: "Hello-World"},
			{path: "octocat.github.io"},
			{path: "git-consortium"},
			{path: "octocat-secrets"},
		},
	}
	baseURL := "https://github.com/octocat"

	// TODO: call buildURL with baseURL and "Hello-World" and print the results

	// Uncomment this when you write the processRepos function
	// processedUser := processRepos(user, baseURL)
	// for _, repo := range processedUser.repositories {
	//	fmt.Println(getDescription(repo))
	// }
	}
}

// TODO: Write a function called buildURL here

// func getDescription(repo) {
// 	return github.GetDescription(repo.url)
// }
