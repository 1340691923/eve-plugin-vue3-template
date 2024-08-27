package vo

import "encoding/json"

type HelloWorld struct {
	Text json.RawMessage `json:"text"`
}
