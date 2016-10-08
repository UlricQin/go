package param

import (
	"encoding/json"
	"net/http"

	"github.com/ulricqin/go/errors"
)

func ParseJSON(r *http.Request, v interface{}) {
	if r.ContentLength == 0 {
		errors.Bomb("content is blank")
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(v)
	errors.Dangerous(err)
}
