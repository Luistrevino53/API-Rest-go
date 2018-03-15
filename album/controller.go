package album

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

//Controller ...
type Controller struct {
	Repository Repository
}

//Index GET /
func (c *Controller) Index(w http.ResponseWriter, r *http.Request){
	albums := c.Repository.GetAlbums()
	log.Println(albums)
	data, _ := json.Marshal(albums)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return 
}

//AddAlbum POST /
func (c *Controller) AddAlbum(w http.ResponseWriter, r *http.Request){
	var album Album
	body, err := ioutil.ReadAll(io.LimitedReader(r.Body, 1048576))
	if err != nil {
		log.Fatalln("Error AddAlbum", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddAlbum", err)
	}
	if err := json.Unmarshal(body, &album); err != nil {
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error AddAlbum unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return 
		}
	}
	success := c.Repository.AddAlbum(album)
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}

//UpdateAlbum PUT/
func (c *Controller) UpdateAlbum(w http.ResponseWriter, r *http.Request){
	var album Album
	body, err := ioutil.ReadAll(io.LimitedReader(r.Body, 1048576))
	if err != nil {
		log.Fatalln("Error UpdateAlbum", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error UpdateAlbum", err)	
	}
	if err := json.Unmarshal(body, &album); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error UpdateAlbum unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	success := c.Repository.UpdateAlbum(album)
	if !success {
		e.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOk)
	return
}

//DeleteAlbum DELETE /
func (c *Controller) DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := c.Repository.DeleteAlbum(id); err != "" {
		if strings.Contains(err, "404"){
			w.WriteHeader(http.StatusNotFound)
		} else if strings.Contains(err, "500") {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}