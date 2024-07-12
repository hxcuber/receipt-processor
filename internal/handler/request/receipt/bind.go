package receipt

import "net/http"

func (receipt Receipt) Bind(r *http.Request) error {
	return nil
}
