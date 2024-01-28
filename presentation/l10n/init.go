package l10n

import (
	"log/slog"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

var SupportedLocale = map[language.Tag]struct{}{
	language.English:    {},
	language.Vietnamese: {},
}

func init() {
	entries := getEntries()

	for _, e := range entries {
		switch msg := e.msg.(type) {
		case string:
			if err := message.SetString(e.tag, e.k, msg); err != nil {
				slog.Error("failed to set language entry", "tag", e.tag.String(), "key", e.k, "msg", msg, "msg_type", "string")
			}
		case catalog.Message:
			if err := message.Set(e.tag, e.k, msg); err != nil {
				slog.Error("failed to set language entry", "tag", e.tag.String(), "key", e.k, "msg", msg, "msg_type", "catalog.Message")
			}
		case []catalog.Message:
			if err := message.Set(e.tag, e.k, msg...); err != nil {
				slog.Error("failed to set language entry", "tag", e.tag.String(), "key", e.k, "msg", msg, "msg_type", "[]catalog.Message")
			}
		}
	}
}
