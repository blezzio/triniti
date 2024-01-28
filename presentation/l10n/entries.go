package l10n

import (
	"golang.org/x/text/language"
)

func getEntries() []entry {
	return []entry{
		{language.English, HeaderLang, "en"},
		{language.Vietnamese, HeaderLang, "vi"},

		{language.English, HeaderDesc, "An anonymous and free URL shortener"},
		{language.Vietnamese, HeaderDesc, "Website rút gọn link, đường dẫn, URL miễn phí"},

		{language.English, HeaderKeywords, "triniti, url shortener, tiny url, short url, small url, link shortener, tiny link, short link, small link, free, anonymous, online"},
		{language.Vietnamese, HeaderKeywords, "triniti, rút gọn đường dẫn, rút gọn link, rút gọn url, link ngắn, url ngắn, đường dẫn ngắn, miễn phí, online, ẩn danh"},

		{language.English, IndexGreeting, "Hi you there!"},
		{language.Vietnamese, IndexGreeting, "Chào đằng ấy!"},
	}
}
