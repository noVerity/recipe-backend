package rest

type URIElement struct {
	Id string `uri:"id" binding:"required"`
}

type QueryPagination struct {
	Offset int `form:"offset"`
	Limit  int `form:"limit"`
}

type PagingInformation struct {
	Count  int `json:"count"`
	Offset int `json:"offset"`
}

type PagedResult struct {
	Pagination PagingInformation `json:"pagination"`
}
