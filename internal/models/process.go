package models

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/devices"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
	"time"
)

var (
	ERRORS = 0
)

type Scraper struct {
	Browser *rod.Browser
	Page    *rod.Page
}

func (S *Scraper) NewBrowser() *rod.Browser {
	logger.Debug("Initializing scraper")
	url := launcher.New().Headless(false).Delete("use-mock-keychain").MustLaunch()
	S.Browser = rod.New().ControlURL(url).MustConnect()
	return S.Browser
}
func (S *Scraper) NewPage(URL string) *rod.Page {
	logger.Debug("Browsing To website: " + URL)
	S.Page = S.Browser.MustPage(URL)
	S.Page.MustEmulate(devices.Clear)
	return S.Page
}
func (S *Scraper) GetHijacker() *rod.HijackRouter {
	logger.Info("Initializing hijacker")
	router := S.Browser.HijackRequests()
	defer router.MustStop()
	return router
}

func (S *Scraper) ProcessElementsChain(chained ElementsChained) ([]interface{}, error) {
	var result []interface{}
	logger.Info("Processing elements chain")
	for _, element := range chained {
		res, err := S.ProcessElement(element)
		if err != nil {
			logger.Errorf("Error: %v", err)
			return nil, err
		}
		logger.Debugf("Result: %v", res)
		result = append(result, res)
	}
	return result, nil
}

func (S *Scraper) ProcessElement(element WebPageElement) (interface{}, error) {
	var result interface{}
	logger.Infof("Processing element: %v", element.ElementName)
	switch element.Action {
	case "Click":
		elem := S.Page.Timeout(15 * time.Second).MustElementX(element.Xpath)
		elem.ScrollIntoView()
		elem.MustClick()
		logger.Debugf("Clicked: %v", element.ElementName)
	case "SetValue":
		elem, err := S.Page.ElementX(element.Xpath)
		if err != nil {
			logger.Errorf("Error: %v", err)
			return nil, err
		}
		elem.MustInput(element.ActionArg)
		logger.Debugf("Set value: %v", element.ActionArg)
	case "GetText":
		elem, err := S.Page.ElementX(element.Xpath)
		if err != nil {
			logger.Errorf("Error: %v", err)
			return nil, err
		}
		text, err := elem.Text()
		logger.Debugf("Text value: %v", text)
		return text, err
	case "ImageSrc":
		el, err := S.Page.Timeout(3 * time.Second).ElementX(element.Xpath)
		if err != nil {
			logger.Errorf("Error: %v", err)
			// recover from panic
			defer func() {
				if r := recover(); r != nil {
					logger.Errorf("Recover from Error: %v", r)
					S.Page.KeyActions().Press(input.Escape).Release(input.Escape).MustDo()
				}
			}()
			panic(err)
		}
		el.ScrollIntoView()
		value, err := el.Attribute("src")
		logger.Debugf("image: %s \n", *value)
		S.Page.KeyActions().Press(input.Escape).Release(input.Escape).MustDo()
		logger.Debugf("Closed Image: %v", element.ElementName)
		return value, nil
	case "MultiImageSrc":
		el, err := S.Page.Timeout(3 * time.Second).ElementX(element.Xpath)
		if err != nil {
			logger.Errorf("Error: %v", err)
			// recover from panic
			defer func() {
				if r := recover(); r != nil {
					logger.Errorf("Recover from Error: %v", r)
					S.Page.KeyActions().Press(input.Escape).Release(input.Escape).MustDo()
				}
			}()
			panic(err)
		}
		el.ScrollIntoView()
		value, err := el.Attribute("src")
		logger.Debugf("image: %s \n", *value)
		S.Page.KeyActions().Press(input.Escape).Release(input.Escape).MustDo()
		logger.Debugf("Closed Image: %v", element.ElementName)
		return value, nil
	}
	return result, nil
}

func (S *Scraper) ProcessError(xpath string, name string) error {
	logger.Infof("Processing error: %v", name)
	el, err := S.Page.ElementX(xpath)
	if err != nil {
		logger.Errorf("Error: %v", err)
		return err
	}
	text, err := el.Text()
	if err != nil {
		logger.Errorf("Error: %v", err)
		return err
	}
	logger.Debugf("Text: %v", text)
	return nil
}
