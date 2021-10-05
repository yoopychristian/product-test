package services

import (
	cfg "product-test/config"
	tables "product-test/database"
	h "product-test/helpers"
	shared "product-test/shared"
	"strconv"
	"time"

	"product-test/tools"

	"github.com/gin-gonic/gin"
)

func AddProduct(ctx cfg.RepositoryContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		process := "|services|add-product|"
		now := time.Now()
		randNum := tools.RandomNumber(100000, 999999)
		strRandNum := strconv.Itoa(randNum)
		input := shared.ParamProduct{}
		if err := c.Bind(&input); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "bind",
				Reason:   "missing input",
			})
			return
		}

		//name-product
		if err := h.MustNotEmpty(input.ProductName, "product-name"); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "idproduct-mustnotempty",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		//price
		if err := h.NotZero(input.Price, "price"); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "namaBarang-mustnotempty",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		//description
		if err := h.MustNotEmpty(input.Description, "description"); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "description-mustnotempty",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		//quantity
		if err := h.NotZero(input.Quantity, "quantity"); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "harga-mustnotempty",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		product := tables.Product{}
		if err := product.Create(ctx.DB, strRandNum, input.ProductName, input.Description, input.Price, input.Quantity, now, true); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.ERROR,
				Section:  process + "result",
				Error:    err,
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		//Account Information
		h.GoodResponse(c, nil)
	}
}
