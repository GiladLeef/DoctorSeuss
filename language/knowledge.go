package language

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/GiladLeef/DoctorSeuss/util"
)

// Subject is the serializer of the knowledge.json file in the res folder
type Subject struct {
	Name     map[string]string `json:"name"`
	Info     string            `json:"info"`
	Opinion  string            `json:"opinion"`
}

var knowledge = SerializeKnowledge()

// SerializeKnowledge returns a list of knowledge, serialized from `res/datasets/knowledge.json`
func SerializeKnowledge() (knowledge []Subject) {
	err := json.Unmarshal(util.ReadFile("res/datasets/knowledge.json"), &knowledge)
	if err != nil {
		fmt.Println(err)
	}

	return knowledge
}

// FindSubject returns the subject found in the sentence and if no subject is found, returns an empty Subject struct
func FindSubject(locale, sentence string) Subject {
	for _, subject := range knowledge {
		name, exists := subject.Name[locale]

		if !exists {
			continue
		}

		// If the actual subject isn't contained in the sentence, continue
		if !strings.Contains(strings.ToLower(sentence), strings.ToLower(name)) {
			continue
		}

		// Returns the right subject
		return subject
	}

	// Returns an empty subject if none has been found
	return Subject{}
}
