package nableclient

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
