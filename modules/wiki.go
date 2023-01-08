package modules

import (
	"fmt"
	"strings"
	"github.com/trietmn/go-wiki"
)

var WikiTag = "wiki"

func WikiReplacer(locale, entry, response, _ string) (string, string) {

	res, err := gowiki.Summary(entry, 2, -1, false, true)
    if err != nil {
		responseTag := "no wiki"
		return WikiTag, responseTag
    }
    res = strings.Replace(res, "(listen)", "", -1)
	return WikiTag, fmt.Sprintf(response, res)
}
