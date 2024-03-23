package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

type IRequester interface {
	Request(idx int, id int)
	GetIds() []int
	GetResponses() []Response
}

type Requester struct {
	ids       []int
	responses []Response
}

func (r *Requester) Request(idx int, id int) {
	r.responses[idx] = request(id)
}

func (r Requester) GetIds() []int {
	return r.ids
}

func (r Requester) GetResponses() []Response {
	return r.responses
}

func main() {
	ids := []int{1, 2, 12, 31}

	requester := Requester{
		ids:       ids,
		responses: make([]Response, len(ids)),
	}
	RunParallel(&requester)
}

func RunParallel(requester IRequester) {
	var wg sync.WaitGroup

	for idx, id := range requester.GetIds() {
		wg.Add(1)
		idx := idx
		id := id
		go func() {
			defer wg.Done()
			requester.Request(idx, id)
		}()
	}

	wg.Wait()

	resp, err := json.MarshalIndent(requester.GetResponses(), "", "  ")
	if err != nil {
		return
	}
	fmt.Printf("Responses: %s\n", resp)
}

type Response struct {
	index int
	Name  string `json:"title,omitempty"`
	ID    int    `json:"id,omitempty"`
	Err   error  `json:"error,omitempty"`
}

func request(id int) Response {
	endpoint := "https://jsonplaceholder.typicode.com/todos/" + strconv.Itoa(id)
	parsedUrl, err := url.Parse(endpoint)
	if err != nil {
		return Response{
			Err: err,
		}
	}
	endpoint = parsedUrl.String()
	fmt.Println("URL: ", endpoint)
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("Error while fetching url \"%s\"\nError: %+v\n", endpoint, err)
		return Response{
			Err: err,
		}
	}

	body, err := io.ReadAll(res.Body)
	err = res.Body.Close()
	if err != nil {
		return Response{
			Err: err,
		}
	}

	result := Response{
		index: id,
	}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Error in json unmarshal: ", err)
		return Response{
			Err: err,
		}
	}

	fmt.Printf("Response: %+v\n", result)
	return result
}

// Develop a concurrent program in Golang that fetches
// data from multiple APIs simultaneously.
// Each API call should be executed concurrently, and
// the program should aggregate the results efficiently.
// Ensure proper synchronization and handling of errors.

func Fact(num int) int {
	if num <= 1 {
		return 1
	}

	return num * Fact(num-1)
}
