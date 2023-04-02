package json

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Raw json.RawMessage

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *Raw) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = Raw(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (j Raw) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}

// Value return json value, implement driver.Valuer interface
func (j *Raw) Set(value interface{}) error {
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}

	raw := Raw(b)
	j = &raw

	return nil
}
