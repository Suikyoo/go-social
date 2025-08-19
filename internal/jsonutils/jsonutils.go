// collection of json utility functions
package jsonutils

import (
	"encoding/json"
	"net/http"
)

func Read(w http.ResponseWriter, r *http.Request, data any) error{
  //maxBytes := 1_048_578
  //r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

  decoder := json.NewDecoder(r.Body)
  decoder.DisallowUnknownFields()

  return decoder.Decode(data)
}

func Write(w http.ResponseWriter, data any) error {
  w.Header().Set("Content-Type", "application/json")
  encoder := json.NewEncoder(w)
  return encoder.Encode(data)

}
