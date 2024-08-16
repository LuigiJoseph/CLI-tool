package main

import (
	"time"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

const outOfRange = 99999
const daysInLastSixMonths = 183
const weeksInLastSixMonths = 26

func stats(email string) {
	commits := processRepositories(email)
	printCommitsStats(commits)
}

// given the user emai, return the commits he made in the past 6 months
func processRepositories(email string) map[int]int {
	filepath := getDotFilePath()
	repos := parseFileLinesToSlice(filepath)
	daysInMap := daysInLastSixMonths

	commits := make(map[int]int, daysInMap)
	for i := daysInMap; i > 0; i-- {
		commits[i] = 0
	}

	for _, path := range repos {
		commits = fileCommits(email, path, commits)
	}

	return commits
}

// given a repo found in "path", gets the commits
// and puts them in the "commits" map, returning it when completed.
func fileCommits(email string, path string, commits map[int]int) map[int]int {
	//instantiate a git repo object from path
	repo, err := git.PlainOpen(path)
	if err != nil {
		panic(err)
	}
	//get the HEAD reference
	ref, err := repo.Head()
	if err != nil {
		panic(err)
	}
	//get the commits history from HEAD
	iterator, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		panic(err)
	}
	//iterate the commits
	offset := calcOffSet()
	err = iterator.ForEach(func(c *object.Commit) error {
		daysAgo := countDaysSinceDate(c.Author.When) + offset

		if c.Author.Email != email {
			return nil
		}

		if daysAgo != outOfRange {
			commits[daysAgo]++
		}

		return nil
	})

	if err != nil {
		return (err)
	}

	return commits
}

func getBeginningOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	startOfDay := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	return startOfDay
}

func countDaysSinceDate(date time.Time) int {
	days := 0
	now := getBeginningOfDay(time.Now())
	for date.Before(now) {
		date = date.Add(time.Hour * 24)
		days++
		if days > daysInLastSixMonths {
			return outOfRange
		}
	}
	return days
}

func printCommitsStats(commits []string) {

}
