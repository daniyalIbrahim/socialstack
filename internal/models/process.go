package models

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"net/http"
)

type ProcessScraper interface {
	ProcessElementsChain(chained ElementsChained) error
	ProcessElement(element WebPageElement) error
	ProcessError(xpath string, name string) error
}

type Scraper struct {
	browser *rod.Browser
	page    *rod.Page
}

func InitScraper() *rod.Browser {
	logger.Debug("Initializing scraper")
	url := launcher.New().Headless(false).Delete("use-mock-keychain").MustLaunch()
	browser := rod.New().ControlURL(url).MustConnect()
	return browser
}

func (S *Scraper) GetHijacker() *rod.HijackRouter {
	logger.Info("Initializing hijacker")
	router := S.browser.HijackRequests()
	defer router.MustStop()
	router.MustAdd("*.js", func(ctx *rod.Hijack) {
		// Here we update the request's header. Rod gives functionality to
		// change or update all parts of the request. Refer to the documentation
		// for more information.
		ctx.Request.Req().Header.Set("My-Header", "test")
		// LoadResponse runs the default request to the destination of the request.
		// Not calling this will require you to mock the entire response.
		// This can be done with the SetXxx (Status, Header, Body) functions on the
		// ctx.Response struct.
		_ = ctx.LoadResponse(http.DefaultClient, true)
		// Here we append some code to every js file.
		// The code will update the document title to "hi"
		ctx.Response.SetBody(ctx.Response.Body() + "\n document.title = 'hi' ")
	})
	return router
}

func (S *Scraper) ProcessElementsChain(chained ElementsChained) error {
	logger.Info("Processing elements chain")
	for _, element := range chained {
		err := S.ProcessElement(element)
		if err != nil {
			logger.Errorf("Error: %v", err)
			return err
		}
	}
	return nil
}

func (S *Scraper) ProcessElement(element WebPageElement) error {
	logger.Infof("Processing element: %v", element.ElementName)
	switch element.Action {
	case "Click":
		elem, err := S.page.ElementX(element.Xpath)
		if err != nil {
			logger.Errorf("Error: %v", err)
			return err
		}
		elem.MustClick()
		logger.Debugf("Clicked: %v", element.ElementName)
	case "SetValue":
		elem, err := S.page.ElementX(element.Xpath)
		if err != nil {
			logger.Errorf("Error: %v", err)
			return err
		}
		elem.MustInput(element.ActionArg)
		logger.Debugf("Set value: %v", element.ActionArg)
	}
	return nil
}

func (S *Scraper) ProcessError(xpath string, name string) error {
	logger.Infof("Processing error: %v", name)
	el, err := S.page.ElementX(xpath)
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
