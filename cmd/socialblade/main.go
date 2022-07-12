package main

import (
	"flag"
	"socialslab/internal/models"
	"socialslab/util"
)

var (
	username *string
	password *string
	profile  *string
)

func init() {
	username = flag.String("user", "please enter", "instagram username")
	password = flag.String("pass", "please enter", "instagram password")
	profile = flag.String("search", "please enter", "instagram profile")
}

func main() {
	log := util.GetLogger()
	log.Info("Starting SocialSlab!")
	flag.Parse()

	if *username == "please enter" || *password == "please enter" || *profile == "please enter" {
		log.Info("Error: Please Enter username, password or profile")
		return
	} else {
		log.Info("username:", *username)
		log.Info("password:", *password)
		log.Info("will search for profile:", *profile)
		models.Run(*username, *password, *profile)
	}
	log.Info("done")
}
