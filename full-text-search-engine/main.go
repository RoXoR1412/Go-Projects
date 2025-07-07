package main

import (
	"flag"
	"log"
	"time"

	utils "github.com/RoXoR1412/full-text-search-engine/utils"
)

func main() {
	var dumpPath, query string
	flag.StringVar(&dumpPath, "p", "enwiki-20250401-pages-articles-multistream1.xml-p1p41242", "wiki abstract dump path")
	flag.StringVar(&query, "q", "small wild cat", "search query")
	flag.Parse()
	log.Println("Full text search is in progress")
	start := time.Now()
	docs, err := utils.LoadDocuments(dumpPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d documents in %s", len(docs), time.Since(start))
	start = time.Now()
	idx := make(utils.Index)
	idx.Add(docs)
	log.Printf("Indexed %d documents in %s time", len(idx), time.Since(start))
	start = time.Now()
	matchedIDs := idx.Search(query)
	log.Printf("Search found %d documents in %v", len(matchedIDs), time.Since(start))
	for _, id := range matchedIDs {
		doc := docs[id]
		log.Printf("ID: %d, Title: %s\n", id, doc.Text)
	}
}
