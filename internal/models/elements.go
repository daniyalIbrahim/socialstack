package models

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

func InstagramLoginElementsChain(username string, password string) ElementsChained {
	var elements = ElementsChained{}
	elements.Add(WebPageElement{ElementName: "accept cookies", Xpath: "/html/body/div[4]/div/div/button[1]", Action: "Click", ActionArg: ""})
	elements.Add(WebPageElement{ElementName: "Phone/username/Email", Xpath: "/html/body/div[1]/section/main/article/div[2]/div[1]/div[2]/form/div/div[1]/div/label/input", Action: "SetValue", ActionArg: username + "\n"})
	elements.Add(WebPageElement{ElementName: "Password", Xpath: "/html/body/div[1]/section/main/article/div[2]/div[1]/div[2]/form/div/div[2]/div/label/input", Action: "SetValue", ActionArg: password + "\n"})
	elements.Add(WebPageElement{ElementName: "Login button", Xpath: "/html/body/div[1]/section/main/article/div[2]/div[1]/div[2]/form/div/div[3]/button", Action: "Click", ActionArg: ""})
	elements.Add(WebPageElement{ElementName: "Don't Save login Details", Xpath: "/html/body/div[1]/div/div/section/main/div/div/div/div/button", Action: "Click", ActionArg: ""})
	elements.Add(WebPageElement{ElementName: "Don't Turn On Notifications", Xpath: "/html/body/div[1]/div/div[1]/div/div[2]/div/div/div[1]/div/div[2]/div/div/div/div/div/div/div/div[3]/button[2]", Action: "Click", ActionArg: ""})
	return elements
}

func InstagramFindProfileChain(profileName string) ElementsChained {
	var elements = ElementsChained{}
	elements.Add(WebPageElement{ElementName: "Searching For The Profile", Xpath: "/html/body/div[1]/div/div[1]/div/div[1]/div/div/div/div[1]/div[1]/section/nav/div[2]/div/div/div[2]/input", Action: "SetValue", ActionArg: profileName + "\n"})
	elements.Add(WebPageElement{ElementName: "Clicking On The Profile", Xpath: "/html/body/div[1]/div/div[1]/div/div[1]/div/div/div/div[1]/div[1]/section/nav/div[2]/div/div/div[2]/div[3]/div/div[2]/div/div[1]/a/div", Action: "Click", ActionArg: ""})
	return elements
}
