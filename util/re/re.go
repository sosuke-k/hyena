package re

import "regexp"

// Split split text by delimeter
func Split(text string, delimeter string) []string {
	reg := regexp.MustCompile(delimeter)
	indexes := reg.FindAllStringIndex(text, -1)
	laststart := 0
	result := make([]string, len(indexes)+1)
	for i, element := range indexes {
		result[i] = text[laststart:element[0]]
		laststart = element[1]
	}
	result[len(indexes)] = text[laststart:len(text)]
	return result
}

// FindString finds most left match string in text string
func FindString(text string, match string) string {
	reg := regexp.MustCompile(match)
	return reg.FindString(text)
}

// FindStringSubmatch finds most left match string in text string
func FindStringSubmatch(text string, match string) []string {
	reg := regexp.MustCompile(match)
	return reg.FindStringSubmatch(text)
}

// FindAllString finds most left match string in text string
func FindAllString(text string, match string) []string {
	reg := regexp.MustCompile(match)
	return reg.FindAllString(text, -1)
}
