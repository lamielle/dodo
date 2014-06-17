package main

import (
	"fmt"
	"sort"

	"github.com/ChimeraCoder/anaconda"
	"github.com/kr/pretty"
)

// A data structure to hold a key/value pair.
type Pair struct {
	Key   string
	Value int
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

// A function to turn a map into a PairList, then sort and return it.
func sortMapByValue(m map[string]int) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	return p
}

func main() {
	anaconda.SetConsumerKey("")
	anaconda.SetConsumerSecret("")
	api := anaconda.NewTwitterApi(
		"",
		"")

	tweets, err := api.GetHomeTimeline(map[string][]string{"count": {"200"}})

	if err != nil {
		panic(err)
	}

	fmt.Println("tweets len:", len(tweets))
	pretty.Println("tweets:", tweets)

	countsForUser := make(map[string]int)
	for _, tweet := range tweets {
		countsForUser[tweet.User.ScreenName]++
	}

	pretty.Println("countsForUser:", countsForUser)

	sorted := sortMapByValue(countsForUser)

	pretty.Println("sorted countsForUser:", sorted)
}
