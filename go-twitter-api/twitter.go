package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"os"
)

func createClient() *twitter.Client {

	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_TOKEN")
	tokenSecret := os.Getenv("TWITTER_TOKEN_SECRET")

	fmt.Printf("%s %s %s %s", consumerKey, consumerSecret, accessToken, tokenSecret)

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, tokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	return client

	//// Send a Tweet
	//tweet, resp, err := client.Statuses.Update("just setting up my twttr", nil)
	//
	//// Status Show
	//tweet, resp, err := client.Statuses.Show(585613041028431872, nil)
	//
	//// Search Tweets
	//search, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
	//	Query: "gopher",
	//})
	//
	//// User Show
	//user, resp, err := client.Users.Show(&twitter.UserShowParams{
	//	ScreenName: "dghubble",
	//})
	//
	//// Followers
	//followers, resp, err := client.Followers.List(&twitter.FollowerListParams{})

}
