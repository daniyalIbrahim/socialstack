package models

import (
	"github.com/go-rod/rod/lib/utils"
	"socialslab/util"
)

var (
	logger = util.GetLogger()
)

func Run(username, password, profile string) {
	var scraper Scraper
	scraper.browser = InitScraper()

	scraper.page = scraper.browser.MustPage("https://instagram.com")

	err := scraper.ProcessElementsChain(InstagramLoginElementsChain(username, password))
	if err != nil {
		logger.Info("Error: %v", err)
	}
	logger.Info("Logged in Successfully!")
	err = scraper.ProcessElementsChain(InstagramFindProfileChain(profile))
	if err != nil {
		logger.Info("Error: %v", err)
	}
	_, err = scraper.page.Eval("()=>scroll(0, 1000);")
	if err != nil {
		logger.Info("Error: %v", err)
	}
	logger.Info("Scrolled Successfully!")
	utils.Pause()
}
