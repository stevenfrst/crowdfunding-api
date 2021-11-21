package transaction

import (
	"errors"
	payment "github.com/stevenfrst/crowdfunding-api/drivers/midtrans"
	"github.com/stevenfrst/crowdfunding-api/usecase/campaign"
	"strconv"
)

type TransactionUseCase struct {
	repoTransaction TransactionRepoInterface
	repoCampaign campaign.CampaignRepoInterface
	payment payment.ConfigMidtrans
}

func NewUsecase(transactionRepo TransactionRepoInterface,campaignRepository campaign.CampaignRepoInterface,payment payment.ConfigMidtrans) TransactionUsecaseInterface {
	return TransactionUseCase{
		transactionRepo,
		campaignRepository,
		payment,
	}
}

func (t TransactionUseCase) CreateTransaction(campaignID,userID,Nominal int) (Domain,error) {
	var transaction Domain
	id, err := t.repoTransaction.GetLastTransactionID()
	if err != nil {
		return Domain{},err
	}
	resp := t.payment.GetLinkResponse(id,Nominal)
	transaction.CampaignID = uint(campaignID)
	transaction.UserID = uint(userID)
	transaction.PaymentLink = resp.RedirectURL
	transaction.Nominal = Nominal
	transactionReturned,err := t.repoTransaction.CreateTransaction(&transaction)
	if err != nil {
		return Domain{},err
	}

	return transactionReturned,nil
}

func (t TransactionUseCase) GetStatusByID(ID int) (Domain,error) {
	transaction,err := t.repoTransaction.GetByID(ID)
	//log.Println(transaction)
	if err != nil  {
		return Domain{},err
	} else if transaction.ID == 0 {
		return Domain{},errors.New("Not Found")
	}

	return transaction,nil
}

func (t TransactionUseCase) GetNotificationPayment(input DomainNotification) (Domain,error) {
	transactionID, err := strconv.Atoi(input.OrderID)
	if err != nil {
		return Domain{},err
	}
	transaction,err := t.repoTransaction.GetByID(transactionID)
	if err != nil {
		return Domain{},err
	}
	if input.PaymentType == "bank_transfer" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
		transaction.PaymentType = input.PaymentType
		transaction.TransactionStatus = input.TransactionStatus
		transaction.FraudStatus = input.FraudStatus
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
		transaction.PaymentType = input.PaymentType
		transaction.TransactionStatus = input.TransactionStatus
		transaction.FraudStatus = input.FraudStatus
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}
	//log.Println("RANGGENAH",transaction)
	// error ng kene ranggenah wkwk
	updatedTransaction,err := t.repoTransaction.UpdateTransaction(&transaction)
	if err != nil {
		return Domain{},err
	}
	campaign,err := t.repoCampaign.FindByID(int(updatedTransaction.CampaignID))
	if err != nil {
		return Domain{},err
	}

	if  updatedTransaction.Status == "paid" {
		campaign.Supporters = campaign.Supporters + 1
		campaign.AmountNow = campaign.AmountNow + updatedTransaction.Nominal

		_, err := t.repoCampaign.UpdateCampaign(campaign)
		if err != nil {
			return Domain{},err
		}
	}

	return *updatedTransaction,err
}