package models

import (
	"fmt"
	"github.com/go-rod/rod/lib/utils"
	"log"
	"strconv"
)

type InstagramScraper struct {
	UserName          string
	Password          string
	KeywordSearch     string
	KeywordAction     string
	ProfilePostsCount int
}

func (I *InstagramScraper) New(username, password, keyword string) {
	I = &InstagramScraper{
		UserName:      username,
		Password:      password,
		KeywordSearch: keyword,
		KeywordAction: "Instagram",
	}
}
func (I *InstagramScraper) InstagramLoginElementsChain() ElementsChained {
	var elements = ElementsChained{}
	elements.Add(WebPageElement{ElementName: "accept cookies", Xpath: "/html/body/div[4]/div/div/button[1]", Action: "Click", ActionArg: ""})
	elements.Add(WebPageElement{ElementName: "Phone/username/Email", Xpath: "/html/body/div[1]/section/main/article/div[2]/div[1]/div[2]/form/div/div[1]/div/label/input", Action: "SetValue", ActionArg: I.UserName + "\n"})
	elements.Add(WebPageElement{ElementName: "Password", Xpath: "/html/body/div[1]/section/main/article/div[2]/div[1]/div[2]/form/div/div[2]/div/label/input", Action: "SetValue", ActionArg: I.Password + "\n"})
	elements.Add(WebPageElement{ElementName: "Login button", Xpath: "/html/body/div[1]/section/main/article/div[2]/div[1]/div[2]/form/div/div[3]/button", Action: "Click", ActionArg: ""})
	//elements.Add(WebPageElement{ElementName: "Don't Save login Details", Xpath: "/html/body/div[1]/div/div/section/main/div/div/div/div/button", Action: "Click", ActionArg: ""})
	elements.Add(WebPageElement{ElementName: "Don't Turn On Notifications", Xpath: "/html/body/div[1]/div/div[1]/div/div[2]/div/div/div[1]/div/div[2]/div/div/div/div/div/div/div/div[3]/button[2]", Action: "Click", ActionArg: ""})
	return elements
}

func (I *InstagramScraper) InstagramSearch() ElementsChained {
	var elements = ElementsChained{}
	elements.Add(WebPageElement{ElementName: "Searching For The Profile", Xpath: "/html/body/div[1]/div/div[1]/div/div[1]/div/div/div/div[1]/div[1]/section/nav/div[2]/div/div/div[2]/input", Action: "SetValue", ActionArg: I.KeywordSearch + "\n"})
	elements.Add(WebPageElement{ElementName: "Clicking On The Profile", Xpath: "/html/body/div[1]/div/div[1]/div/div[1]/div/div/div/div[1]/div[1]/section/nav/div[2]/div/div/div[2]/div[3]/div/div[2]/div/div[1]/a/div", Action: "Click", ActionArg: ""})
	elements.Add(WebPageElement{ElementName: "Getting Posts Count", Xpath: "/html/body/div[1]/div/div[1]/div/div[1]/div/div/div/div[1]/div[1]/section/main/div/header/section/ul/li[1]/div/span", Action: "GetText", ActionArg: ""})
	return elements
}

func (I *InstagramScraper) GenerateProfilePostsXPathsList() ElementsChained {
	var elements = ElementsChained{}
	name := "Image at Row %d Column %d"
	//xpath := "/html/body/div[1]/div/div[1]/div/div[1]/div/div/div/div[1]/div[1]/section/main/div[1]/div[3]/div[2]/div/div/div[%d]/div[%d]/a/div/div[1]/img"
	xpath := "/html/body/div[1]/div/div[1]/div/div[1]/div/div/div/div[1]/div[1]/section/main/div/div[3]/article/div[1]/div/div[%d]/div[%d]"
	logger.Info("Generating Profile Images List Max Images:", I.ProfilePostsCount)
	for i := 1; i < I.ProfilePostsCount; i++ {
		for j := 1; j < 4; j++ {
			elements.Add(WebPageElement{ElementName: fmt.Sprintf(name, i, j), Xpath: fmt.Sprintf(xpath, i, j), Action: "Click", ActionArg: ""})
			elements.Add(WebPageElement{ElementName: fmt.Sprintf("Find Image R:%v C:%v", i, j), Xpath: "/html/body/div[1]/div/div[1]/div/div[2]/div/div/div[1]/div/div[3]/div/div/div/div/div[2]/div/article/div/div[1]/div/div/div/div[1]/img", Action: "ImageSrc", ActionArg: ""})
		}
	}
	return elements
}

//
func (I *InstagramScraper) GenerateMultiPostsXPathsList() ElementsChained {
	var elements = ElementsChained{}
	name := "Image at Row %d Column %d"
	//xpath := "/html/body/div[1]/div/div[1]/div/div[1]/div/div/div/div[1]/div[1]/section/main/div[1]/div[3]/div[2]/div/div/div[%d]/div[%d]/a/div/div[1]/img"
	xpath := "/html/body/div[1]/div/div[1]/div/div[1]/div/div/div/div[1]/div[1]/section/main/div/div[3]/article/div[1]/div/div[%d]/div[%d]"
	logger.Info("Generating Profile Images List Max Images:", I.ProfilePostsCount)
	for i := 1; i < I.ProfilePostsCount; i++ {
		for j := 1; j < 4; j++ {
			elements.Add(WebPageElement{ElementName: fmt.Sprintf(name, i, j), Xpath: fmt.Sprintf(xpath, i, j), Action: "Click", ActionArg: ""})
			elements.Add(WebPageElement{ElementName: fmt.Sprintf("Find Image R:%v C:%v", i, j), Xpath: "/html/body/div[1]/div/div[1]/div/div[2]/div/div/div[1]/div/div[3]/div/div/div/div/div[2]/div/article/div/div[2]/div/div[1]/div[2]/div/div/div/ul/li[2]/div/div/div/div/div[1]/img", Action: "MultiImageSrc", ActionArg: ""})
		}
	}
	return elements
}

func RunInstagramScraper(username, password, keyword string) {
	var instagram = &InstagramScraper{
		UserName:      username,
		Password:      password,
		KeywordSearch: keyword,
	}
	var scraper = Scraper{}
	scraper.NewBrowser()
	scraper.NewPage("https://www.instagram.com/")
	// scraper.Page.MustNavigate("https://www.instagram.com/")

	processLogin := instagram.InstagramLoginElementsChain()
	_, err := scraper.ProcessElementsChain(processLogin)
	if err != nil {
		logger.Info("Error: %v", err)
	}
	logger.Info("Logged in Successfully!")
	// Get the profile posts count
	processSearch := instagram.InstagramSearch()
	val, err := scraper.ProcessElementsChain(processSearch)
	if err != nil {
		logger.Info("Error: %v", err)
	}
	instagram.ProfilePostsCount, err = strconv.Atoi(val[2].(string))
	logger.Infof("Result from Search %v has type %T", val[2], val[2])
	log.Printf("Profile Posts Count: %d", instagram.ProfilePostsCount)
	// Get the number of posts
	processPosts := instagram.GenerateProfilePostsXPathsList()
	val, err = scraper.ProcessElementsChain(processPosts)
	if err != nil {
		logger.Info("Error: %v", err)
	}
	logger.Info("Profile Posts Count: %d", len(val))
	processMultiPosts := instagram.GenerateProfilePostsXPathsList()
	val, err = scraper.ProcessElementsChain(processMultiPosts)
	if err != nil {
		logger.Info("Error: %v", err)
	}
	logger.Info("Profile Multi Posts Count: %d", len(val))
	utils.Pause()
	defer scraper.Browser.MustClose()
}
