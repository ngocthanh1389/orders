package service

import (
	"net/http"
	"time"

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
	s.Register(http.MethodGet, "orders", o.GetTodayOrder)
}

func (o *order) createOrder(c *gin.Context) {
	var req orderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		o.l.Errorw("invalid request", "err", err)
		c.JSON(http.StatusBadRequest, util.ConstructErrResponse("yêu cầu không hợp lệ"))
		return
	}
	o.l.Infow("receive create order request", "req", req)

	ok, err := o.userDAL.ValidateUser(req.Name, req.Password)
	if err != nil {
		o.l.Errorw("error when validate user", "err", err)
		c.JSON(http.StatusInternalServerError, util.ConstructErrResponse("hệ thống lỗi"))
		return
	}
	if !ok {
		o.l.Info("user pass isn't valid")
		c.JSON(http.StatusNotAcceptable, util.ConstructErrResponse("tên hoặc mật khẩu không đúng"))
		return
	}
	ok, err = o.orderDAL.CheckHasOrderToday(req.Name)
	if err != nil {
		o.l.Errorw("error when check order today", "err", err)
		c.JSON(http.StatusInternalServerError, util.ConstructErrResponse("hệ thống lỗi"))
		return
	}
	if ok {
		o.l.Info("user pass made order")
		c.JSON(http.StatusNotAcceptable, util.ConstructErrResponse("đã đặt cơm rồi"))
		return
	}
	if err := o.orderDAL.Insert(req.Name, time.Now()); err != nil {
		o.l.Errorw("error when insert order", "err", err)
		c.JSON(http.StatusInternalServerError, util.ConstructErrResponse("hệ thống lỗi"))
		return
	}
	c.JSON(http.StatusAccepted, util.ConstructSuccessResponse("đặt cơm thành công"))
}

func (o *order) GetTodayOrder(c *gin.Context) {
	o.l.Infow("receive get orders request")
	res, err := o.orderDAL.GetTodayOrder()
	if err != nil {
		o.l.Errorw("error when get orders", "error", err)
		c.JSON(http.StatusInternalServerError, util.ConstructErrResponse("hệ thống lỗi"))
		return
	}
	c.JSON(http.StatusAccepted, util.ConstructSuccessResponse(res))
}
