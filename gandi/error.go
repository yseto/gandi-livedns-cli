package gandi

import (
	"encoding/json"
	"fmt"
	"io"
)

func ShowError(r io.Reader) error {
	var data Response
	err := json.NewDecoder(r).Decode(&data)
	if err != nil {
		return err
	}
	return fmt.Errorf("API error : %s", data.Message)
}
