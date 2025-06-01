package cart

import (
	"context"
	"github.com/labstack/echo/v4"
	"inventory/config"
	"inventory/internal/util"
	"net/http"
)

type ICartHandler interface {
	Add(e echo.Context) error
	Get(e echo.Context) error
	List(e echo.Context) error
	Buy(e echo.Context) error
	Verify(e echo.Context) error
	BuyCount(e echo.Context) error
}
type SCartHandler struct {
	appConfig *config.AppConfig
	service   ICartService
}

func NewCartHandler(client *config.AppConfig) ICartHandler {
	return &SCartHandler{
		appConfig: client,
		service:   NewCartService(client),
	}
}

// Add godoc
// @Summary      Add cart
// @Description  Returns a user based on the authenticated user ID from the context
// @Tags         cart
// @Param user body RQCart true "cart data"
// @Produce      json
// @Success      200  {object} response.BaseRS(data=([]cart.RSCart),error={})
// @Failure 400 {object} response.BaseRS(data={},error=(errorhandler.BaseErr))  // Return error in case of failure (adjust this to match your actual error handler)
// @Failure 404 {object} response.BaseRS(data={},error=(errorhandler.BaseErr))  // Return error if the user to delete is not found
// @Security     BearerAuth
// @Router       /api/v1/cart [post]
func (h *SCartHandler) Add(e echo.Context) error {
	var rq RQCart
	err := util.BindAndValidate(&rq, e)
	if err != nil {
		return err
	}
	rq.UserID = e.Get("userId").(string)

	ctx := context.Background()
	rs, err := h.service.Add(ctx, rq)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusCreated, util.ConvertToResponse(rs))
}

// Get godoc
// @Summary      Get cart by ID
// @Description  Returns a cart based on the provided ID from the path
// @Tags         cart
// @Produce      json
// @Param        id   path      string  true  "Cart ID"
// @Success      200  {object}  response.BaseRS{data=cart.RSCart,error=interface{}}
// @Failure      400  {object}  response.BaseRS{data=interface{},error=errorhandler.BaseErr}
// @Failure      404  {object}  response.BaseRS{data=interface{},error=errorhandler.BaseErr}
// @Security     BearerAuth
// @Router       /api/v1/cart/{id} [get]
func (h *SCartHandler) Get(e echo.Context) error {
	id := e.Param("id")

	ctx := context.Background()
	rs, err := h.service.Get(ctx, id)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusCreated, util.ConvertToResponse(rs))
}

// List godoc
// @Summary      List cart
// @Description  Returns a carts
// @Tags         cart
// @Produce      json
// @Success      200  {object}  response.BaseRS{data=[]cart.RSCart,error=interface{}}
// @Failure      400  {object}  response.BaseRS{data=interface{},error=errorhandler.BaseErr}
// @Failure      404  {object}  response.BaseRS{data=interface{},error=errorhandler.BaseErr}
// @Security     BearerAuth
// @Router       /api/v1/cart [get]
func (h *SCartHandler) List(e echo.Context) error {
	userId := e.Get("userId").(string)
	status := e.QueryParam("status")

	ctx := context.Background()
	rs, err := h.service.List(ctx, userId, status)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusCreated, util.ConvertToListResponse(rs))
}

// Buy godoc
// @Summary      Submit cart for purchase
// @Description  Submits the authenticated user's cart for checkout
// @Tags         cart
// @Accept       json
// @Produce      json
// @Param        request  body  cart.RQCart  true  "Cart purchase request"
// @Success      201  {object}  response.BaseRS{data=cart.RSCart,error=interface{}}
// @Failure      400  {object}  response.BaseRS{data=interface{},error=errorhandler.BaseErr}
// @Failure      500  {object}  response.BaseRS{data=interface{},error=errorhandler.BaseErr}
// @Security     BearerAuth
// @Router       /api/v1/cart/buy [post]
func (h *SCartHandler) Buy(e echo.Context) error {
	var rq RQCart
	err := util.BindAndValidate(&rq, e)
	if err != nil {
		return err
	}
	userId := e.Get("userId").(string)
	rq.UserID = userId

	ctx := context.Background()
	rs, err := h.service.Buy(ctx, rq)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusCreated, util.ConvertToResponse(rs))
}

// Verify godoc
// @Summary      Verify cart payment
// @Description  Verifies payment for a submitted cart
// @Tags         cart
// @Accept       json
// @Produce      json
// @Param request body cart.RQPayment true "Payment verification request"
// @Success      200  {object}  response.BaseRS{data=map[string]string,error=interface{}}
// @Failure      400  {object}  response.BaseRS{data=interface{},error=errorhandler.BaseErr}
// @Failure      500  {object}  response.BaseRS{data=interface{},error=errorhandler.BaseErr}
// @Security     BearerAuth
// @Router       /api/v1/cart/verify [post]
func (h *SCartHandler) Verify(e echo.Context) error {
	var rq RQPayment
	err := util.BindAndValidate(&rq, e)
	if err != nil {
		return err
	}
	//
	ctx := context.Background()

	err = h.service.Verify(ctx, rq)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, util.ConvertToResponse(map[string]string{"message": "payment verified successfully"}))

}

// BuyCount godoc
// @Summary      Get user's purchase count
// @Description  Returns the number of purchases made by the authenticated user
// @Tags         cart
// @Produce      json
// @Success      200  {object}  response.BaseRS{data=int,error=interface{}}
// @Failure      400  {object}  response.BaseRS{data=interface{},error=errorhandler.BaseErr}
// @Failure      500  {object}  response.BaseRS{data=interface{},error=errorhandler.BaseErr}
// @Security     BearerAuth
// @Router       /api/v1/cart/buy-count [get]
func (h *SCartHandler) BuyCount(e echo.Context) error {
	userId := e.Get("userId").(string)

	ctx := context.Background()
	rs, err := h.service.BuyCount(ctx, userId)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, util.ConvertToResponse(rs))
}
