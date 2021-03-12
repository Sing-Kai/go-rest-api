package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Sing-Kai/go-rest-api/internal/comment"
	"github.com/gorilla/mux"
)

// Handler - stores pointer to our comment service
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

//Response - object to store responses
type Response struct {
	Message string
	Error   string
}

// NewHandler - returns a pointer to a Handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

//SetupRoutes - set up all the routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("setting up routes")
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods("DELETE")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Response{Message: "I am healthy"}); err != nil {
			panic(err)
		}
	})
}

//GetComment - retrieve a comment by ID
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		sendErrorResponse(w, "unable to parse UINT from id", err)
	}

	comment, err := h.Service.GetComment(uint(i))

	if err != nil {
		sendErrorResponse(w, "Error Retrieving Comment By ID", err)
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

//GetAllComments - retrieves all comments from comment service
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	comments, err := h.Service.GetAllComments()

	if err != nil {
		sendErrorResponse(w, "Failed to retrieve all comments", err)
	}

	if err := json.NewEncoder(w).Encode(comments); err != nil {
		panic(err)
	}
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		sendErrorResponse(w, "Failed to post new comment", err)
	}

	comment, err := h.Service.PostComment(comment)

	if err != nil {
		sendErrorResponse(w, "Failed to post comment", err)
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

//UpdateComment - updates comment
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		sendErrorResponse(w, "Failed to update new comment", err)
	}

	vars := mux.Vars(r)
	id := vars["id"]
	commentID, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		sendErrorResponse(w, "Failed to parse uint from id", err)
	}

	comment, err = h.Service.UpdateComment(uint(commentID), comment)

	if err != nil {
		sendErrorResponse(w, "Failed to update comment", err)
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

//DeleteComment - deletes comment by id
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]
	commentID, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		sendErrorResponse(w, "Failed to parse uint from id", err)
	}

	err = h.Service.DeleteComment(uint(commentID))

	if err != nil {
		sendErrorResponse(w, "Failed to delete comment by commentID", err)
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Comment successfully deleted"}); err != nil {
		panic(err)
	}
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}
