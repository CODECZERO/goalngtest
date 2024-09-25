package main

import (
	"net/http"
	"context"
	"io"
	"log"
	"sync"
	"time"
	"github.com/CODECZERO/goalngtest/internal/db"
)

func startScraper(db *db.,councrry int,timeBetweenRequest time.Duration ){

	log.Printf("collecting feeds every %s on %v goroutiness",timeBetweenRequest,councrry)
	ticker:=time.NewTicker(timeBetweenRequest)
	for ; ;<-ticker.C{
		feeds,err:=db.GetNextFeedsToFetch(context.Backgorund(),int32(councrry))
		if err!=nil {
			log.Println("couldn't get next feeds to fecth",err)
			continue
		}

		log.Printf("Found %v feeds to fetch!",len(feeds))
		wg:=&sync.WaitGroup()

		for _,feed:=range feeds {
			wg.Add(1)
			go scrapeFeed(db,wg,feed)
		}

		wg.Wait()
	}
}

func scrapeFeed(db *db.	,wg *sync.WaitGroup,feed db.feeds){
	defer wg.Done()
	_,err:=db.MarkFeedFetched(context.Backgorund(),feed.ID)
	if err!=nil {
		log.Printf("couldn't mark feed %s fetched %v",feed.Name,err)
		return
	}

	feedData,err:=fetchFeed(feed.Url)
	if err !=nil{
		log.Printf("couldn't collect feed %s : %v",feed.Name,err)
		return
	}

	for _,iteam :=range feedData.Channel.iteam{
		log.Println("Found Post",iteam.Title)
	}

	log.Printf("feed %s collected , %v posts found",feed.Name,len(feedData.Channel.iteam))
}

type RSSFeed struct{
	Channel struct{
		Title string `xml:"title"`
		Link string `xml:"link"`
		Description string `xml:"description"`
		Language string `xml:"language"`
		iteam []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct{
	Title string `xml:"title"`
	Link string `xml:"link"`
	Description string `xml:"description"`
	PubDate string `xml:"pubDate"`
}

func fecthFeed(){
	
}
