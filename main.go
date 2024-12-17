package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/games", getGames)

	log.Fatal(http.ListenAndServe("10.0.0.29:9547", router))
}

type GameStruct struct {
	ImageUrl string `json:"imageUrl"`
	Name     string `json:"name"`
}

type GamesStruct struct {
	Items []GameStruct `json:"items"`
}

func getGames(w http.ResponseWriter, r *http.Request) {
	exampleGame1 := GameStruct{
		Name:     "meow",
		ImageUrl: "https://picsum.photos/536/354",
	}
	exampleGame2 := GameStruct{
		Name:     "mrrp",
		ImageUrl: "https://picsum.photos/536/332",
	}
	exampleGame3 := GameStruct{
		Name:     "mraow",
		ImageUrl: "https://picsum.photos/536/234",
	}
	games := []GameStruct{exampleGame1, exampleGame2, exampleGame3}

	if err := json.NewEncoder(w).Encode(games); err != nil {
		log.Fatal(err)
	}

}
