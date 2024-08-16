package main

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
		daysAgo := countryDaysSinceDate(c.Author.When) + offset

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

func printCommitsStats(commits []string) {

}
