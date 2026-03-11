package ncentralclient

import (
	"encoding/json"
	"time"
)

type GenericResult[T any] struct {
	Data  T `json:"data"`
	Links struct {
	} `json:"_links"`
	Warning any `json:"_warning"`
}

type GenericPage[T any] struct {
	Data       []T `json:"data"`
	PageNumber int `json:"pageNumber"`
	PageSize   int `json:"pageSize"`
	ItemCount  int `json:"itemCount"`
	TotalItems int `json:"totalItems"`
	TotalPages int `json:"totalPages"`
	Links      struct {
		FirstPage    *string `json:"firstPage"`
		PreviousPage *string `json:"previousPage"`
		NextPage     *string `json:"nextPage"`
		LastPage     *string `json:"lastPage"`
	} `json:"_links"`
	Warning any `json:"_warning"`
}

const DateTimeFormat = "2006-01-02T15:04:05.999"

type DateTime time.Time

func (d *DateTime) UnmarshalJSON(b []byte) error {
	var stringValue string
	if err := json.Unmarshal(b, &stringValue); err != nil {
		return err
	}
	timeValue, err := time.Parse(DateTimeFormat, stringValue)
	if err != nil {
		return err
	}
	*d = DateTime(timeValue)
	return nil
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d).Format(DateTimeFormat))
}
