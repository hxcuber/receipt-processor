package errordesc

import "net/http"

func (e ErrorDesc) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
