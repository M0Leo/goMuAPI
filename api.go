package main

import (
	"encoding/json"
	"fmt"
	"goMuAPI/main/db"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type APIServer struct {	
	listenAddr  string
	SongService db.SongService
}

func NewAPIServer(listenAddr string, songService db.SongService) *APIServer {
	return &APIServer{
		listenAddr:  listenAddr,
		SongService: songService,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.Handle("/song", HttpHandleFunc(s.handleMuxSong))
	router.Handle("/song/{id}", HttpHandleFunc(s.handleGetSongByID))

	log.Println("Server running on port: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleMuxSong(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		songs, err := s.SongService.GetSongs()
		if err != nil {
			return err
		}

		return JSON(w, http.StatusOK, songs)
	case "POST":
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

		return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetSongByID(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		id, err := getID(r)
		if err != nil {
			return err
		}

		song, err := s.SongService.GetSongByID(id)
		if err != nil {
			return err
		}

		return JSON(w, http.StatusOK, song)
	case "DELETE":
		return JSON(w, http.StatusCreated, "")
	}

	return fmt.Errorf("method not allowed %s", r.Method)
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

func getID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid id given %s", idStr)
	}
	return id, nil
}