package main

import (
	"flag"
	"socialslab/internal/models"
	"socialslab/util"
)

var (
	username      *string
	password      *string
	searchKeyword *string
)

func init() {
	username = flag.String("user", "", "instagram username")
	password = flag.String("pass", "", "instagram password")
	searchKeyword = flag.String("search", "", "instagram search keyword")
}

func main() {
	log := util.GetLogger()
	log.Info("Starting SocialClout!")
	flag.Parse()
	if *username == "" || *password == "" || *searchKeyword == "" {
		log.Info("Error:  username, password or profile")
		return
	} else {
		log.Info("username:", *username)
		log.Info("password:", *password)
		log.Infof("will search for %s in instagram", *searchKeyword)
		models.RunInstagramScraper(*username, *password, *searchKeyword)
	}

	log.Info("done")
}
