# Top News Headline Fetcher

[![Go Report Card](https://goreportcard.com/badge/github.com/yinnyC/MakeUtility)](https://goreportcard.com/report/github.com/yinnyC/MakeUtility)
[![code size](https://img.shields.io/github/languages/code-size/yinnyC/MakeUtility)](https://img.shields.io/github/languages/code-size/yinnyC/MakeUtility)
[![last commit](https://img.shields.io/github/last-commit/yinnyC/MakeUtility)](https://img.shields.io/github/last-commit/yinnyC/MakeUtility)

This is a program written in Golang to fetch the top 10 news headlines from [NewsAPI](https://newsapi.org/). The program was built with **Goroutines** libery so that it can fetch news from the 5 categories(business, general, science, technology, health)  independently and simultaneously.

## Instructions

- Clone the repo

  ```zsh
  git clone https://github.com/yinnyC/MakeUtility.git
  ```

- Configure the environmnet file, and put in your own API Key.  

  ```zsh
  # change file name
  mv .env.example .env
  ```

- Run the program

  ```zsh
  go run main.go
  ```

## Sample Outcome

Please refer to  directories /SampleOutput and there will be two files:

- topnewsheadlines.json

```json
[
 {
  "category": "science",
  "articles": [
   {
    "title": "Among the Stars chronicles daring space mission to repair physics experiment - Ars Technica",
    "url": "https://arstechnica.com/science/2021/10/among-the-stars-chronicles-daring-space-mission-to-repair-physics-experiment/"
   },
   {
    "title": "NASA reassigns 2 astronauts from Boeing's Starliner to SpaceX's Crew Dragon - Space.com",
    "url": "https://www.space.com/nasa-reassigns-astronauts-boeing-spacex"
   },
   {
    "title": "Burial that included a racy love goddess inscription held multiple people - Livescience.com",
    "url": "https://www.livescience.com/tomb-nestors-cup-multiple-individuals"
   },
   {
    "title": "Creepy New Drone That Walks and Flies Is a Robopocalypse Nightmare Come True - Gizmodo",
    "url": "https://gizmodo.com/creepy-new-drone-that-walks-and-flies-is-a-robopocalyps-1847809268"
   },
   ...
  ]
```

- debug-web.log
