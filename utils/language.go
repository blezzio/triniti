package utils

import (
	"fmt"

	"golang.org/x/text/language"
)

func GetLanguage(acceptedLang string, supportedLocales map[language.Tag]struct{}) (language.Tag, error) {
	if tags, _, err := language.ParseAcceptLanguage(acceptedLang); err != nil {
		return language.English, err
	} else {
		for _, tag := range tags {
			if _, ok := supportedLocales[tag]; ok {
				return tag, nil
			}
		}
	}
	return language.English, fmt.Errorf("all language are unsupported")
}
