package main

import (
	"encoding/json"
	"goMuAPI/main/db"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	SongService db.SongService      
}

func NewAPIServer(listenAddr string, songService db.SongService) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		SongService:      songService,
	}
}


func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.Handle("/songs", HttpHandleFunc(s.handleGetSong))
	router.Handle("/song", HttpHandleFunc(s.handleCreateSong))

	log.Println("Server running on port: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleGetSong(w http.ResponseWriter, r *http.Request) error {
	songs, err := s.SongService.GetSongs()
	if err != nil {
		return err
	}

	return JSON(w, http.StatusOK, songs)
}

func (s *APIServer) handleCreateSong(w http.ResponseWriter, r *http.Request) error {
	req := new(CreateSong)

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	song, err := NewSong(req.Title, req.Artist, req.Title, req.Year)
	if err != nil {
		return err
	}

	if err := s.SongService.CreateSong(song); err != nil {
		return err
	}

	return JSON(w, http.StatusCreated, song)
}

type ApiError struct {
	Error string `json:"error"`
}

func JSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func HttpHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			JSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}