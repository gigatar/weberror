package weberror

import "encoding/json"

// Error data type.
type Error struct {
	Code    int    `json:"code,omitempty" example:"35"`
	Message string `json:"message,omitempty" example:"Invalid input"`
}

// ToJSON returns the JSON string of the error struct.
// Note: Empty Code or Message will be ommited.
func (w Error) ToJSON() []byte {
	result, _ := json.Marshal(w)

	return result
}
