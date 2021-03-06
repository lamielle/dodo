package main

import (
	"fmt"
	"net/http"
	"os"
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
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("Works!"))
	})
	fmt.Println("Dodo!")
	//err := http.ListenAndServe("", mux)
	//if err != nil {
	//	panic(err)
	//}
	//	return

	dodo_key := os.Getenv("DODO_KEY")
	dodo_secret := os.Getenv("DODO_SECRET")

	if len(dodo_key) == 0 || len(dodo_secret) == 0 {
		panic("Must set DODO_KEY and DODO_SECRET env vars")
	}

	// Set application key and secret
	anaconda.SetConsumerKey(dodo_key)
	anaconda.SetConsumerSecret(dodo_secret)

	// Request an auth token
	url, creds, err := anaconda.AuthorizationURL("http://localhost:8080/testing")
	pretty.Println("url", url)
	pretty.Println("creds", creds)
	pretty.Println("err", err)
	if err != nil {
		panic(err)
	}

	var verifier string
	n, err := fmt.Scanln(&verifier)
	if err != nil {
		panic(err)
	}
	pretty.Println("read n=%s", n)

	creds2, values, err := anaconda.GetCredentials(creds, verifier)
	pretty.Println("creds2", creds2)
	pretty.Println("values", values)
	pretty.Println("err", err)
	if err != nil {
		panic(err)
	}

	api := anaconda.NewTwitterApi(
		creds2.Token,
		creds2.Secret)

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
