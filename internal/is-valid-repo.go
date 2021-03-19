package internal

import "regexp"

func IsValidRepo(repo string) bool {
	matched, _ := regexp.MatchString(`^[\w\-]+/[\w\-]+$`, repo)

	return matched
}
