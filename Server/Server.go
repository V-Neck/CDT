package main

import (
	"encoding/json"
	//"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// TODO come up with a way to store this better, maybe a database
var parties []Party


func main() {

	router := mux.NewRouter()

	// routes for modifying parties
	router.HandleFunc("/party", GetParties).Methods("GET")
	router.HandleFunc("/party/{name}", GetParty).Methods("GET")
	router.HandleFunc("/party/{name}", CreateParty).Methods("POST")
	router.HandleFunc("/party/{name}", DeleteParty).Methods("DELETE")

	// routes for modifying songs
	router.HandleFunc("/party/{name}/{songId}", CreatePartySong).Methods("POST")
	router.HandleFunc("/party/{name}/{songId}", DeletePartySong).Methods("DELETE")

	// routes for voting on songs
	router.HandleFunc("/party/{name}/{songId}/upvote", UpvotePartySong).Methods("POST")
	router.HandleFunc("/party/{name}/{songId}/upvote", UndoUpvotePartySong).Methods("DELETE")

	router.HandleFunc("/party/{name}/{songId}/downvote", DownvotePartySong).Methods("POST")
	router.HandleFunc("/party/{name}/{songId}/downvote", UndoDownvotePartySong).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}

// gets all parties
func GetParties(w http.ResponseWriter, r *http.Request)  {

	json.NewEncoder(w).Encode(parties)
}


// gets party specified by name
func GetParty(w http.ResponseWriter, r *http.Request)    {

	params := mux.Vars(r)
	for _, item := range parties {
		if item.Name == params["name"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}


// returns true if the party name is already on the server
func partyExists(partyName string) bool {
	for _, item := range parties {
		if item.Name == partyName {
			return true
		}
	}
	return false
}
// creates a party by name
func CreateParty(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// check to make sure the party does not already exists
	if partyExists(params["name"]) {
		// TODO print error message, there is already a party with this name
	} else {
		// TODO add creation date to party
		parties = append(parties, Party{Name: params["name"]})
		json.NewEncoder(w).Encode(parties)
	}
}


// deletes a party by name
func DeleteParty(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range parties {
		if item.Name == params["name"] {
			parties = append(parties[:index], parties[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(parties)

}



// returns true if a party and song already exists on this server
func songExists(partyName string, songId string) bool {
	for i, item := range parties {
		if item.Name == partyName {
			for _, song := range parties[i].Songs {
				if song.Id == songId {
					return true;
				}
			}
		}
	}
	return false;
}


// creates a party song by party name and song id
func CreatePartySong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range parties {
		if item.Name == params["name"] {

			// check if the song already exists and upvote if it does
			if songExists(params["name"], params["songId"]) {
				UpvotePartySong(w, r)
				return
			} else {
				parties[i].Songs = append(item.Songs, Song{Id: params["songId"], Upvotes: 0, Downvotes: 0})
				json.NewEncoder(w).Encode(parties[i])
			}
		}
	}
}


// deletes a party song by party name and song id
func DeletePartySong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range parties {
		if item.Name == params["name"] {
			for j, song := range parties[i].Songs {
				if song.Id == params["songId"] {
					parties[i].Songs = append(parties[i].Songs[:j], parties[i].Songs[j+1:]...)
					json.NewEncoder(w).Encode(parties[i])
				}
			}
		}
	}
}


// these next 4 methods arent good solutions, too much reused code
// TODO redesign this by passing functions as parameters

// upvotes a song by party name and sond id
func UpvotePartySong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range parties {
		if item.Name == params["name"] {
			for j, song := range parties[i].Songs {
				if song.Id == params["songId"] {
					parties[i].Songs[j].Upvotes += 1
					json.NewEncoder(w).Encode(parties[i])
				}
			}
		}
	}
}
// undo upvotes a song by party name and sond id
func UndoUpvotePartySong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range parties {
		if item.Name == params["name"] {
			for j, song := range parties[i].Songs {
				if song.Id == params["songId"] {
					parties[i].Songs[j].Upvotes -= 1
					json.NewEncoder(w).Encode(parties[i])
				}
			}
		}
	}
}

// downvotes a song by party name and sond id
func DownvotePartySong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range parties {
		if item.Name == params["name"] {
			for j, song := range parties[i].Songs {
				if song.Id == params["songId"] {
					parties[i].Songs[j].Downvotes += 1
					json.NewEncoder(w).Encode(parties[i])
				}
			}
		}
	}
}

// undo downvote to a song by party name and sond id
func UndoDownvotePartySong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range parties {
		if item.Name == params["name"] {
			for j, song := range parties[i].Songs {
				if song.Id == params["songId"] {
					parties[i].Songs[j].Downvotes -= 1
					json.NewEncoder(w).Encode(parties[i])
				}
			}
		}
	}
}


type Party struct {
	Name  string   `json:"name"`
	// TODO add a field for creation date
	Songs []Song `json:"songs"`
}
type Song struct {
	Id string `json:"id"`
	Upvotes int `json:"upvotes"`
	Downvotes int `json:"downvotes"`
}
