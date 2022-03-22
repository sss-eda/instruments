package openapi

import "net/http"

// GetAllInstruments TODO
func GetAllInstruments(
	interact getallinstruments.Interactor,
	controller getallinstruments.Controller,
	presenter getallisntruments.Presenter,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		interact(
			controller.OpenAPI(r),
			presenter.OpenAPI(w),
		)
	}
}
