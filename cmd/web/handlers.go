package main

import "net/http"

func (app *application) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	err := app.renderTemplate(w, r, "virtual-terminal", nil)
	if err != nil {
		app.errorLog.Println(err)
	}

}
