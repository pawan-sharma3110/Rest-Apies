package handlers

import (
	"bank/models"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

var AllPost = []models.Post{
	{
		ID:       1,
		UserName: "Pawan",
		Title:    "About Nature",
		Content:  "Nature is best",
		CreateOn: time.Now(),
	},
	{
		ID:       3,
		UserName: "Rahul",
		Title:    "About me",
		Content:  "Nature",
		CreateOn: time.Now(),
	},
}

func GetAllPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(AllPost)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("only get method recquired")
	}
}
func GetPostByid(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		paramId := r.PathValue("id")
		id, _ := strconv.Atoi(paramId)
		isIdExtis := false
		for _, v := range AllPost {
			if v.ID == id {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				isIdExtis = true
				json.NewEncoder(w).Encode(v)
			}
		}
		if !isIdExtis {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("Provided ID dose not extis.")
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("only get method recquired")
	}
}
func UpdatePostById(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPatch {
		paramId := r.PathValue("id")
		id, _ := strconv.Atoi(paramId)
		isIdExtis := false
		for i, v := range AllPost {
			if v.ID == id {
				var update models.Post
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				isIdExtis = true
				AllPost[i].UserName = update.UserName
				AllPost[i].Title = update.Title
				AllPost[i].Content = update.Content
				AllPost[i].UpdateOn = time.Now()
				json.NewDecoder(r.Body).Decode(&update)
			}
		}
		if !isIdExtis {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("Provided ID dose not extis.")
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("only Patch method recquired")
	}
}
