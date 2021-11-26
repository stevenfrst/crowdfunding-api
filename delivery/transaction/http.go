package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/stevenfrst/crowdfunding-api/delivery"
	"github.com/stevenfrst/crowdfunding-api/delivery/transaction/request"
	"github.com/stevenfrst/crowdfunding-api/delivery/transaction/response"
	"github.com/stevenfrst/crowdfunding-api/usecase/transaction"
	"net/http"
	"strconv"
)

type TransactionDelivery struct {
	usecase transaction.TransactionUsecaseInterface
}

// NewTransactionDelivery function used to assign all handlers to usecase
func NewTransactionDelivery(uc transaction.TransactionUsecaseInterface) *TransactionDelivery {
	return &TransactionDelivery{
		uc,
	}
}

// CreateTransaction used to make transactions via formvalue and processed to usecase
func (t *TransactionDelivery) CreateTransaction(c echo.Context) error {
	campaignid,_ := strconv.Atoi(c.FormValue("campaign_id"))
	userid,_ := strconv.Atoi(c.FormValue("user_id"))
	nominal,_ := strconv.Atoi(c.FormValue("nominal"))

	transactionResp,err := t.usecase.CreateTransaction(campaignid,userid,nominal)
	if err != nil {
		return delivery.ErrorResponse(c,http.StatusInternalServerError,"gagal membuat transaksi",err)
	}
	return delivery.SuccessResponse(c,response.FromDomain(transactionResp))

}

// GetStatusByID used to view transaction status via ID
func (t TransactionDelivery) GetStatusByID(c echo.Context) error {
	id,err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return delivery.ErrorResponse(c,http.StatusBadRequest,"Failed",err)
	}
	resp,err := t.usecase.GetStatusByID(id)
	if err != nil {
		return delivery.ErrorResponse(c,http.StatusInternalServerError,"Failed",err)
	}
	return delivery.SuccessResponse(c,response.FromDomain(resp))

}
// GetNotificationPayment used by midtrans to send payment notifications
func (t TransactionDelivery) GetNotificationPayment(c echo.Context) error {
	var input request.TransactionNotification

	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,map[string]string{
			"messages": err.Error(),
		})
	}
	//	usecase

	notification,err := t.usecase.GetNotificationPayment(input.ToDomainNotification())

	if err != nil {
		return delivery.ErrorResponse(c,http.StatusInternalServerError,"Wrong Input",err)
	}

	return delivery.SuccessResponse(c,response.FromDomain(notification))
}
