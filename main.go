package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

// number of NUM_PARALLEL
const NUM_PARALLEL = 5

// A Response struct to map the Entire Response
type Response struct {
	Category string    `json:"category"`
	Articles []Article `json:"articles"`
}

// An Article struct to map news entries
type Article struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}
type result struct {
	News Response
	err  error
}

func makeUrl(category string) string {
	// This function return an URL for topnewsheadlines request.
	godotenv.Load(".env")
	API_KEY := os.Getenv("API_KEY")
	if category != "" {
		return fmt.Sprintf("https://newsapi.org/v2/top-headlines?country=us&category=%s&apiKey=%s", category, API_KEY)
	}
	return fmt.Sprintf("https://newsapi.org/v2/top-headlines?country=us&apiKey=%s", API_KEY)
}

func fetchNewsHeadline(category string) (Response, error) {
	url := makeUrl(category)
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	responseObject.Category = category
	return responseObject, err
}

// Stream news to input channel
func streamNews(done <-chan struct{}, inputs []string) <-chan string {
	inputCh := make(chan string)
	go func() {
		defer close(inputCh)
		for _, input := range inputs {
			select {
			case inputCh <- input:
			case <-done:
				// in case done is closed prematurely (because error midway),
				// finish the loop (closing input channel)
				break
			}
		}
	}()
	return inputCh
}

func AsyncHTTP(categories []string) ([]Response, error) {
	// This is a function that spawn 5 worker goroutines, to fetch news
	done := make(chan struct{})
	defer close(done)

	inputCh := streamNews(done, categories)

	var wg sync.WaitGroup
	// bulk add goroutine counter at the start
	wg.Add(NUM_PARALLEL)

	resultCh := make(chan result)

	for i := 0; i < NUM_PARALLEL; i++ {
		// spawn N worker goroutines, each is consuming a shared input channel.
		go func() {
			for input := range inputCh {
				response, err := fetchNewsHeadline(input)
				resultCh <- result{response, err}
			}
			wg.Done()
		}()
	}

	// Wait all worker goroutines to finish. Happens if there's no error (no early return)
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	results := []Response{}
	for result := range resultCh {
		if result.err != nil {
			// return early. done channel is closed, thus input channel is also closed.
			// all worker goroutines stop working (because input channel is closed)
			return nil, result.err
		}
		results = append(results, result.News)
	}

	return results, nil
}

func WriteToJson(jsonObject []byte) {
	// writing json to file
	_ = ioutil.WriteFile("topnewsheadlines.json", jsonObject, 0644)

	// to append to a file
	// create the file if it doesn't exists with O_CREATE, Set the file up for read write, add the append flag and set the permission
	f, err := os.OpenFile("./debug-web.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		log.Fatal(err)
	}
	// write to file, f.Write()
	f.Write(jsonObject)
}

func main() {
	categories := []string{"business", "general", "science", "technology", "health"}
	start := time.Now()

	results, err := AsyncHTTP(categories)
	if err != nil {
		fmt.Println(err)
		return
	}
	resultJson, _ := json.MarshalIndent(results, "", " ")
	WriteToJson(resultJson)

	fmt.Println("finished in ", time.Since(start))
}
