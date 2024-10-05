package main

import (
	datastore "github.com/MiyamotoAkira/gotweet/datastore"
	routes "github.com/MiyamotoAkira/gotweet/routes"
)

func main() {
	repo := datastore.Setup("tweet.db")
	r := routes.SetupRouter(repo)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
