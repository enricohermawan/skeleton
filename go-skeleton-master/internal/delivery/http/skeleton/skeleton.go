package skeleton

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"go-skeleton/pkg/response"
)

// ISkeletonSvc is an interface to Skeleton Service
// Masukkan function dari service ke dalam interface ini
type ISkeletonSvc interface {
}

type (
	// Handler ...
	Handler struct {
		skeletonSvc ISkeletonSvc
	}
)

// New for bridging product handler initialization
func New(is ISkeletonSvc) *Handler {
	return &Handler{
		skeletonSvc: is,
	}
}

// SkeletonHandler will receive request and return response
func (h *Handler) SkeletonHandler(w http.ResponseWriter, r *http.Request) {
	var (
		resp     *response.Response
		result   interface{}
		metadata interface{}
		err      error
		errRes   response.Error
	)
	resp = &response.Response{}
	defer resp.RenderJSON(w, r)

	switch r.Method {
	// Check if request method is GET
	case http.MethodGet:

	// Check if request method is POST
	case http.MethodPost:

	// Check if request method is PUT
	case http.MethodPut:

	// Check if request method is DELETE
	case http.MethodDelete:

	default:
		err = errors.New("404")
	}

	// If anything from service or data return an error
	if err != nil {
		// Error response handling
		errRes = response.Error{
			Code:   101,
			Msg:    "Data Not Found",
			Status: true,
		}
		// If service returns an error
		if strings.Contains(err.Error(), "service") {
			// Replace error with server error
			errRes = response.Error{
				Code:   201,
				Msg:    "Failed to process request due to server error",
				Status: true,
			}
		}

		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		resp.Error = errRes
		return
	}

	resp.Data = result
	resp.Metadata = metadata
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
}
