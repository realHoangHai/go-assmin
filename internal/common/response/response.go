package response

type Success struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewSuccess(data, paging, filter interface{}) *Success {
	return &Success{
		Data:   data,
		Paging: paging,
		Filter: filter,
	}
}

func SimpleSuccess(data interface{}) *Success {
	return NewSuccess(data, nil, nil)
}
