package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func main() {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("active.es.toml")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lang := r.FormValue("lang")
		accept := r.Header.Get("Accept-Language")
		localizer := i18n.NewLocalizer(bundle, lang, accept)

		name := r.FormValue("name")
		if name == "" {
			name = "Developer"
		}

		greeting := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "Greeting",
				Other: "Hello, my name is {{.Name}}",
			},
			TemplateData: map[string]string{
				"Name": name,
			},
		})

		resp := map[string]string{
			"payload": greeting,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resp)
	})

	log.Println("Listening on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
