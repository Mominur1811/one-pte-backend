package utils

import (
	"github.com/gin-gonic/gin"
)

type Page struct {
	Items        interface{} `json:"items"`
	ItemsPerPage int         `json:"itemsPerPage"`
	PageNumber   int         `json:"pageNumber"`
	TotalItems   int         `json:"totalItems"`
	TotalPages   int         `json:"totalPages"`
}

func SendPage(c *gin.Context, page Page) {
	SendData(c, page)
}
