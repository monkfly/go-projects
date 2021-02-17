package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"strings"
)

type discover struct {
	file     string
	keywords []string
}

type keyword struct {
	name  string
	count int
}

func (k keyword) String() string {
	return fmt.Sprintf("Keyword: %s\nCount: %d\n-----", k.name, k.count)
}

type byCount []keyword

func (a byCount) Len() int           { return len(a) }
func (a byCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byCount) Less(i, j int) bool { return a[i].count > a[j].count }

// FindKeywords ...
func FindKeywords() {
	fmt.Println("Start find keywords!")

	discover := discover{
		file:     getFile("file.txt"),
		keywords: getKeywords("keywords.txt"),
	}

	var keywords []keyword

	for _, v := range discover.keywords {
		if strings.Contains(discover.file, v) {
			keywords = append(keywords, keyword{name: v, count: strings.Count(discover.file, v)})
		}
	}

	fmt.Println("Top 3 most common:\n-----")

	sort.Sort(byCount(keywords))

	for i := range keywords {
		if i < 3 {
			fmt.Println(keywords[i])
		}
	}
}

func getFile(fileName string) string {
	file, err := ioutil.ReadFile(filepath.Join("/Users/u17628152/go/src/keywords/config", fileName))
	if err != nil {
		log.Fatal("Error while reading file!", err)
	}

	return string(file)
}

func getKeywords(file string) []string {
	return strings.Split(getFile(file), ", ")
}
