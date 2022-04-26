package main

import (
	"net/http"
	"strings"

	"github.com/Kamva/mgm"
	"github.com/mariomenjr/handlr"
	"github.com/mariomenjr/links/models"
	"github.com/mariomenjr/links/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func Init() {
	utils.LoadEnv()
	models.Init()
}

func main() {
	Init()

	h := handlr.New()

	h.Handler("/:id", linkHandler)
	h.Start(1993)
}

func linkHandler(w http.ResponseWriter, r *http.Request) {
	link := &models.Link{}
	coll := mgm.Coll(link)

	err := coll.First(bson.M{"id": strings.Trim(r.URL.Path, "/")}, link)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	http.Redirect(w, r, link.Url, http.StatusSeeOther)
}
