package openapi

import "net/http"

// Controller TODO
type Controller struct{}

// GetAllInstruments TODO
func (controller *Controller) GetAllInstruments(
	interact getallinstruments.Interactor,
	r *http.Request,
) {
	interact()
}
