package http

import (
	"strings"
	"url-shortener/internal/service"


	"encoding/json"
	"net/http"
)

type URLHandler struct {
	service *service.URLService
	baseURL string
}

type shortenRequest struct {
	URL string `json:"url"`
}
type shortenResponse struct {
	ShortCode string `json:"short_code"`
	ShortURL  string `json:"short_url"`
}

func NewURLHandler(service *service.URLService, baseURL string) *URLHandler {
	return &URLHandler{
		service: service,
		baseURL: baseURL,
	}
}

func (h *URLHandler) Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req shortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	result, err := h.service.Create(r.Context(), req.URL)
	if err != nil {
		if err == service.ErrInvalidURL {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	resp := shortenResponse{
		ShortCode: result.ShortCode,
		ShortURL:  h.baseURL + "/" + result.ShortCode,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(resp)
}

func (h *URLHandler) Redirect(w http.ResponseWriter , r *http.Request){

	if r.Method != http.MethodGet{
		http.Error (w , "method not allowed" , http.StatusMethodNotAllowed)
		return 
	}


	code := strings.TrimPrefix(r.URL.Path  ,"/")
		if code == ""{
			http.NotFound(w , r)
			return 
		}

	result ,err := h.service.GetByCode(r.Context() , code )
	if err != nil {
		http.NotFound(w , r)
		return 
	}

	http.Redirect( w , r , result.LongURL , http.StatusFound)

}