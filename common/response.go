package common

type successRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

type PagingResponse struct {
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	Total      int    `json:"total"`
	Cursor     string `json:"cursor"`
	NextCursor string `json:"next_cursor"`
}

func NewSuccessResponse(data, paging, filter interface{}) *successRes {
	return &successRes{
		Data:   data,
		Paging: paging,
		Filter: filter,
	}
}

func NewPaging(page int, limit int, total int) PagingResponse {
	return PagingResponse{
		Page:  page,
		Limit: limit,
		Total: total,
	}
}

func SimpleSuccessResponse(data interface{}) *successRes {
	return NewSuccessResponse(data, nil, nil)
}
