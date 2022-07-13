package models

import "socialslab/util"

var (
	logger = util.GetLogger()
)

type WebPageElement struct {
	ElementName string
	Xpath       string
	Action      string
	ActionArg   string
}

type ElementsChained []WebPageElement

func (e *ElementsChained) Add(element WebPageElement) {
	*e = append(*e, element)
}

func (w *WebPageElement) New(name string, xPath string, action string, args string) *WebPageElement {
	return &WebPageElement{
		ElementName: name,
		Xpath:       xPath,
		Action:      action,
		ActionArg:   args,
	}
}

type ProcessScraper interface {
	ProcessElementsChain(chained ElementsChained) error
	ProcessElement(element WebPageElement) error
	ProcessError(xpath string, name string) error
}
