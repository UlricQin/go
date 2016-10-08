package param

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ulricqin/go/errors"
)

func ParseJSON(r *http.Request, v interface{}) {
	if r.ContentLength == 0 {
		errors.Bomb("content is blank")
	}

	if r.Body == nil {
		errors.Bomb("body is nil")
	}

	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errors.Bomb("cannot read body")
	}

	err = json.Unmarshal(bs, v)
	if err != nil {
		errors.Bomb("cannot decode body: %s", err.Error())
	}
}
