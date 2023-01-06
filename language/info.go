package modules

import (
	"fmt"

	"github.com/GiladLeef/DoctorSeuss/language"
	"github.com/GiladLeef/DoctorSeuss/util"
)

var (
	// InfoTag is the intent tag for its module
	InfoTag = "info"
	// ArticleKnowledge is the map of functions to find the article in front of a subject
	// in different languages
	ArticleKnowledge = map[string]func(string) string{}
)

// InfoReplacer replaces the pattern contained inside the response by the info of the subject
// specified in the message.
// See modules/modules.go#Module.Replacer() for more details.
func InfoReplacer(locale, entry, response, _ string) (string, string) {
	subject := language.FindSubject(locale, entry)

	// If there isn't a subject respond with a message from res/datasets/messages.json
	if subject.Info == "" {
		responseTag := "no subject"
		return responseTag, util.GetMessage(locale, responseTag)
	}

	articleFunction, exists := ArticleKnowledge[locale]
	subjectName := subject.Name[locale]
	if exists {
		subjectName = articleFunction(subjectName)
	}

	return InfoTag, fmt.Sprintf(response, subjectName, subject.Info)
}
