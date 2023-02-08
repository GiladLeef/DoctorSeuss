package modules

import (
  "fmt"
  "strings"
  "regexp"
  "github.com/trietmn/go-wiki"
)

var WikiTag = "wiki"

func WikiReplacer(locale, entry, response, _ string) (string, string) {

  res, err := gowiki.Summary(entry, 2, -1, false, true)
    if err != nil {
    responseTag := "no wiki"
    return WikiTag, responseTag
    }
  
  re := regexp.MustCompile(\[[^\]]*\])
  output := re.ReplaceAllString(res, "")

  // Remove text inside parentheses
  re = regexp.MustCompile(\([^\)]*\))
  output = re.ReplaceAllString(output, "")
  
  // Remove (born XX Month YYYY)
  re = regexp.MustCompile(\(born [0-9]+ [A-Z][a-z]+ [0-9]+\))
  output = re.ReplaceAllString(output, "")

    // Remove (XX Month YYYY – XX Month YYYY)
  re = regexp.MustCompile(\([0-9]+ [A-Z][a-z]+ [0-9]+ – [0-9]+ [A-Z][a-z]+ [0-9]+\))
    output = re.ReplaceAllString(output, "")

  // Remove text inside slashes
  re = regexp.MustCompile(/[^/]*/)
  output = re.ReplaceAllString(output, "")

  // Remove semi-colons
  output = strings.Replace(output, ";", "", -1)
    res = strings.Replace(output, "(listen)", "", -1)

  // Remove all non-English characters
  re = regexp.MustCompile([^\x00-\x7F]+)
  output = re.ReplaceAllString(output, "")
  
  return WikiTag, fmt.Sprintf(response, res)
}
