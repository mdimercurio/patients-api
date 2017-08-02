package main

import (
	"fmt"
)

type Pagination struct {
	Self string `json:"self"`
	Next string `json:"next"`
}

func newPagination(page, l int) Pagination {
	var next string
	if l >= PatientsPerPage {
		next = pageURL(page + 1)
	}
	return Pagination{
		Self: pageURL(page),
		Next: next,
	}
}

var paginationPrefix = "http://localhost:8080/v1/patients?page="

func pageURL(page int) string {
	return fmt.Sprintf("%s%d", paginationPrefix, page)
}
