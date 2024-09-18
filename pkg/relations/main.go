package relations

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/exp/slices"
)

var relations []Relation
var visited []string

type Relation struct {
	Source string
	Target string
}

func GetRelations(domain string) []Relation {
	relations = []Relation{}
	visited = []string{}

	var wg sync.WaitGroup
	var mx sync.Mutex
	wg.Add(1)
	go addRelations(&wg, &mx, domain)
	wg.Wait()

	return relations
}

func addRelations(wg *sync.WaitGroup, mx *sync.Mutex, domain string) {
	defer wg.Done()
	mx.Lock()
	if slices.Contains(visited, domain) {
		mx.Unlock()
		return
	}
	visited = append(visited, domain)
	mx.Unlock()

	friends := getFriends(domain)
	for _, friend := range friends {
		relations = append(relations, Relation{Source: domain, Target: friend})
		wg.Add(1)
		go addRelations(wg, mx, friend)
	}
}

func getFriends(domain string) []string {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://%s/friends.txt", domain), nil)
	if err != nil {
		log.Fatalf("Error while creating request: %v", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return []string{}
	}

	if isInvalid(res) {
		return []string{}
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading body of %s: %v", domain, err)
	}

	friends := filter(strings.Split(string(body), "\n"), isNotEmptyString)
	return friends
}

func isInvalid(res *http.Response) bool {
	if res.StatusCode != 200 {
		return true
	}
	if !(strings.Contains(res.Header.Get("Content-Type"), "text/plain")) {
		return true
	}
	return false
}

func filter[T any](items []T, fn func(item T) bool) []T {
	filteredItems := []T{}
	for _, value := range items {
		if fn(value) {
			filteredItems = append(filteredItems, value)
		}
	}
	return filteredItems
}

func isNotEmptyString(friend string) bool {
	return friend != ""
}
