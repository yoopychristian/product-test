package services

import (
	"net/http"
	cfg "product-test/config"
	tables "product-test/database"
	h "product-test/helpers"
	"product-test/shared"

	"github.com/gin-gonic/gin"
)

func ProductList(ctx cfg.RepositoryContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		process := "|services|product-list|"
		p := tables.Product{}
		sort := c.Param("sort")

		if sort == "new" {
			list, err := p.ProductListTime(ctx.DB)
			if err != nil {
				h.BadResponse(h.RespParams{
					Log:      ctx.Log,
					Context:  c,
					Severity: h.DEBUG,
					Section:  process + "bind",
					Reason:   "missing input",
				})
				return
			}
			data := []shared.Product{}
			for _, row := range list {
				data = append(data, shared.Product{
					IDProduct:   row.IDProduct,
					ProductName: row.ProductName,
					Price:       row.Price,
					Description: row.Description,
					Quantity:    row.Quantity,
				})
			}
			c.JSON(http.StatusOK, gin.H{
				"status": true,
				"data":   data,
			})
		} else if sort == "high" {
			list, err := p.ProductPriceHigh(ctx.DB)
			if err != nil {
				h.BadResponse(h.RespParams{
					Log:      ctx.Log,
					Context:  c,
					Severity: h.DEBUG,
					Section:  process + "bind",
					Reason:   "missing input",
				})
				return
			}
			data := []shared.Product{}
			for _, row := range list {
				data = append(data, shared.Product{
					IDProduct:   row.IDProduct,
					ProductName: row.ProductName,
					Price:       row.Price,
					Description: row.Description,
					Quantity:    row.Quantity,
				})
			}
			c.JSON(http.StatusOK, gin.H{
				"status": true,
				"data":   data,
			})
		} else if sort == "low" {
			list, err := p.ProductPriceLow(ctx.DB)
			if err != nil {
				h.BadResponse(h.RespParams{
					Log:      ctx.Log,
					Context:  c,
					Severity: h.DEBUG,
					Section:  process + "bind",
					Reason:   "missing input",
				})
				return
			}
			data := []shared.Product{}
			for _, row := range list {
				data = append(data, shared.Product{
					IDProduct:   row.IDProduct,
					ProductName: row.ProductName,
					Price:       row.Price,
					Description: row.Description,
					Quantity:    row.Quantity,
				})
			}
			c.JSON(http.StatusOK, gin.H{
				"status": true,
				"data":   data,
			})
		} else if sort == "a-z" {
			list, err := p.ProductNameAZ(ctx.DB)
			if err != nil {
				h.BadResponse(h.RespParams{
					Log:      ctx.Log,
					Context:  c,
					Severity: h.DEBUG,
					Section:  process + "bind",
					Reason:   "missing input",
				})
				return
			}
			data := []shared.Product{}
			for _, row := range list {
				data = append(data, shared.Product{
					IDProduct:   row.IDProduct,
					ProductName: row.ProductName,
					Price:       row.Price,
					Description: row.Description,
					Quantity:    row.Quantity,
				})
			}
			c.JSON(http.StatusOK, gin.H{
				"status": true,
				"data":   data,
			})
		} else if sort == "z-a" {
			list, err := p.ProductNameZA(ctx.DB)
			if err != nil {
				h.BadResponse(h.RespParams{
					Log:      ctx.Log,
					Context:  c,
					Severity: h.DEBUG,
					Section:  process + "bind",
					Reason:   "missing input",
				})
				return
			}
			data := []shared.Product{}
			for _, row := range list {
				data = append(data, shared.Product{
					IDProduct:   row.IDProduct,
					ProductName: row.ProductName,
					Price:       row.Price,
					Description: row.Description,
					Quantity:    row.Quantity,
				})
			}
			c.JSON(http.StatusOK, gin.H{
				"status": true,
				"data":   data,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": false,
				"data":   nil,
			})
		}
	}
}
