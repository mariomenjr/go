package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
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
	h.HandlerFunc("/:id", linkHandler)

	portNumber, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}

	h.Start(portNumber)
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
