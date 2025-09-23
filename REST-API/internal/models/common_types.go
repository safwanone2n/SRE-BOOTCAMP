package models

type FilterParams struct {
	OffSet int `json:"offset"`
	Limit  int `json:"limit"`
}

type IdRequest struct {
	ID int64 `json:"id"`
}

type IdResponse struct {
	Id int64 `json:"id"`
}

type ListRequest struct {
	FilterParams FilterParams `json:"filter_params"`
}
