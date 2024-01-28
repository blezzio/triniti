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

		{language.English, HeaderTitle, "Triniti | Just Another URL Shortener"},
		{language.Vietnamese, HeaderTitle, "Triniti | Rút gọn đường dẫn của bạn"},

		{language.English, IndexGreeting, "Hi you there!"},
		{language.Vietnamese, IndexGreeting, "Chào đằng ấy!"},

		{language.English, IndexSummary, "is a free to use and open source URL shortener service!"},
		{language.Vietnamese, IndexSummary, "là dịch vụ rút gọn đường dẫn hoàn toàn miễn phí, không có ràng buộc và mã nguồn mở!"},

		{language.English, IndexURLTrans, "URL"},
		{language.Vietnamese, IndexURLTrans, "Đường dẫn"},

		{language.English, IndexInputPlaceholder, "Place your URL here!"},
		{language.Vietnamese, IndexInputPlaceholder, "Nhập đường dẫn ở đây!"},

		{language.English, IndexCreate, "Shorten"},
		{language.Vietnamese, IndexCreate, "Rút gọn"},

		{language.English, IndexImgAlt, "Three iris flowers surrounded by bubblesThree iris flowers surrounded by bubbles on the night background, painted in Van Gogh style."},
		{language.Vietnamese, IndexImgAlt, "Ba bông hoa diên vĩ được bao quanh bởi bong bóng trên nền đêm, được vẽ theo phong cách Van Gogh."},
	}
}
