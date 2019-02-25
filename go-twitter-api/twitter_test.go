package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"testing"
)

func Test_Search_Tweets(t *testing.T) {

	client := createClient()

	// Search Tweets
	search, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: "gopher",
	})
	if err != nil {
		t.Errorf("error %v", err)
	}

	t.Logf("%v", resp)

	t.Logf("%v", search)
}

func Test_Filter_Stream(t *testing.T) {

	client := createClient()

	params := &twitter.StreamFilterParams{
		Track: []string{"iphone"},
		StallWarnings: twitter.Bool(true),
		Language:[]string{"ja"},
	}
	stream, err := client.Streams.Filter(params)
	if err != nil {
		t.Errorf("error %v", err)
	}

	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		fmt.Println("--------")
		fmt.Println(tweet.Text)
		fmt.Println("------------------")
	}
	demux.DM = func(dm *twitter.DirectMessage) {
		fmt.Println("~~~~~~~~")
		fmt.Println(dm.SenderID)
		fmt.Println("~~~~~~~~~~~~~~~~~~")
	}
	demux.HandleChan(stream.Messages)

	//for message := range stream.Messages {
	//	fmt.Println(message)
	//}
}
