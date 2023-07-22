package http

import (
	"net/http"

	"github.com/dmytrodemianchuk/boxfit-backend/auth"
	"github.com/dmytrodemianchuk/boxfit-backend/models"
	"github.com/dmytrodemianchuk/boxfit-backend/product"
	"github.com/gin-gonic/gin"
)

type Product struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

type Handler struct {
	useCase product.UseCase
}

func NewHandler(useCase product.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type createInput struct {
	Name     string `json:"url"`
	Category string `json:"title"`
}

func (h *Handler) Create(c *gin.Context) {
	inp := new(createInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.MustGet(auth.CtxUserKey).(*models.User)

	if err := h.useCase.CreateProduct(c.Request.Context(), user, inp.Name, inp.Category); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

type getResponse struct {
	Bookmarks []*Product `json:"bookmarks"`
}

func (h *Handler) Get(c *gin.Context) {
	user := c.MustGet(auth.CtxUserKey).(*models.User)

	bms, err := h.useCase.GetProducts(c.Request.Context(), user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &getResponse{
		Bookmarks: toProducts(bms),
	})
}

type deleteInput struct {
	ID string `json:"id"`
}

func (h *Handler) Delete(c *gin.Context) {
	inp := new(deleteInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.MustGet(auth.CtxUserKey).(*models.User)

	if err := h.useCase.DeleteProduct(c.Request.Context(), user, inp.ID); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func toProducts(bs []*models.Product) []*Product {
	out := make([]*Product, len(bs))

	for i, b := range bs {
		out[i] = toProduct(b)
	}

	return out
}

func toProduct(b *models.Product) *Product {
	return &Product{
		ID:       b.ID,
		Name:     b.Name,
		Category: b.Category,
	}
}
