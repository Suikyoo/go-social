package authutils

import (
	"encoding/json"
	"log"
)
type Password []byte

func (p *Password) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*p = Password(s)
	return nil
}

func (p *Password) MarshalJSON() ([]byte, error) {
  s := string(*p)
  b, err := json.Marshal(s)
  if (err != nil) {
    return nil, err
  }

  return b, nil

}

func (p *Password) Clear() {
  for i := range len(*p) {
    (*p)[i] = '0'
  }
}

func (p *Password) Compare(q *Password) bool {
  log.Println(len(*p), len(*q))
  if len(*p) != len(*q) {
    return false
  }
  for i := range len(*p) {
    log.Println((*p)[i] , (*q)[i])
    if (*p)[i] != (*q)[i] {
      return false
    }
  }
  return true
}
