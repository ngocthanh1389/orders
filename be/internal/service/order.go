package service

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	dataaccess "github.com/ngocthanh1389/orders/be/internal/data_access"
	"github.com/ngocthanh1389/orders/be/pkg/server"
	"github.com/ngocthanh1389/orders/be/pkg/util"
	"go.uber.org/zap"
)

type orderRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type order struct {
	orderDAL dataaccess.OrderDAL
	userDAL  dataaccess.UserDAL
	l        *zap.SugaredLogger
}

func AddOrderService(s *server.Server, db *sqlx.DB, l *zap.SugaredLogger) {
	o := order{
		l:        l,
		orderDAL: *dataaccess.NewOrderDAL(db),
		userDAL:  *dataaccess.NewUserDAL(db),
	}
	o.register(s)
}

func (o *order) register(s *server.Server) {
	s.Register(http.MethodPost, "order", o.createOrder)
}

func (o *order) createOrder(c *gin.Context) {
	var req orderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		o.l.Errorw("invalid request", "err", err)
		c.JSON(http.StatusBadRequest, util.ConstructErrResponse(err))
		return
	}
	o.l.Infow("receive create order request", "req", req)

	ok, err := o.userDAL.ValidateUser(req.Name, req.Password)
	if err != nil {
		o.l.Errorw("error when validate user", "err", err)
		c.JSON(http.StatusInternalServerError, util.ConstructErrResponse(err))
		return
	}
	if !ok {
		o.l.Info("user pass isn't valid")
		c.JSON(http.StatusNotAcceptable, util.ConstructErrResponse(errors.New("user pass isn't valid")))
		return
	}
}
