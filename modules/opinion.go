package modules

import (
	"fmt"

	"github.com/GiladLeef/DoctorSeuss/language"
	"github.com/GiladLeef/DoctorSeuss/util"
)

var (
	// OpinionTag is the intent tag for its module
	OpinionTag = "opinion"
	ArticleKnowledge = map[string]func(string) string{}
)

// OpinionReplacer replaces the pattern contained inside the response by the opinion of the subject
// specified in the message.
// See modules/modules.go#Module.Replacer() for more details.
func OpinionReplacer(locale, entry, response, _ string) (string, string) {
	subject := language.FindSubject(locale, entry)

	// If there isn't a subject respond with a message from res/datasets/messages.json
	if subject.Opinion == "" {
		responseTag := "no opinion"
		return responseTag, util.GetMessage(locale, responseTag)
	}

	articleFunction, exists := ArticleKnowledge[locale]
	subjectName := subject.Name[locale]
	if exists {
		subjectName = articleFunction(subjectName)
	}

	return OpinionTag, fmt.Sprintf(response, subjectName, subject.Opinion)
}
