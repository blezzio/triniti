package views

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/blezzio/triniti/apis/types"
	"github.com/blezzio/triniti/presentation/l10n"
	"github.com/blezzio/triniti/utils"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Index struct {
	printers   map[language.Tag]*message.Printer
	defPrinter *message.Printer
	templ      *template.Template
}

func NewIndex(fs embed.FS) *Index {
	templ := template.Must(template.New(indexTemplName).ParseFS(fs, indexTemplFN...))
	prts := make(map[language.Tag]*message.Printer)
	var defprt *message.Printer
	first := true
	for tag := range l10n.SupportedLocale {
		prts[tag] = message.NewPrinter(tag)
		if first {
			defprt = prts[tag]
			first = false
		}
	}

	return &Index{
		printers:   prts,
		defPrinter: defprt,
		templ:      templ,
	}
}

func (t *Index) Exec(wr http.ResponseWriter, data any) error {
	validData, ok := data.(*types.HTMLIndexView)
	if !ok {
		return utils.Trace(fmt.Errorf("expected type %T got %T", &types.HTMLIndexView{}, data), "failed to parse data type")
	}

	lang := language.English
	if accept, err := utils.GetLanguage(validData.AcceptLanguage, l10n.SupportedLocale); err == nil {
		lang = accept
	}

	printer := t.defPrinter
	if p, ok := t.printers[lang]; ok {
		printer = p
	}

	param := struct {
		HeaderData
		Greet, Summary, URLTrans, Placeholder, Create, ImgAlt, TrinitiURL string
	}{
		HeaderData:  NewHeaderData(printer),
		Greet:       printer.Sprintf(l10n.IndexGreeting),
		Summary:     printer.Sprintf(l10n.IndexSummary),
		URLTrans:    printer.Sprintf(l10n.IndexURLTrans),
		Placeholder: printer.Sprintf(l10n.IndexInputPlaceholder),
		Create:      printer.Sprintf(l10n.IndexCreate),
		ImgAlt:      printer.Sprintf(l10n.IndexImgAlt),
		TrinitiURL:  os.Getenv("TRINITI_URL"),
	}

	t.addHeaders(wr)
	if err := t.templ.Execute(wr, param); err != nil {
		return utils.Trace(err, "failed to execute template")
	}
	return nil
}

func (t *Index) addHeaders(wr http.ResponseWriter) {
	wr.Header().Add("Content-Type", "text/html")
}
