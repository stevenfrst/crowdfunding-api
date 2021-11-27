package transaction

import (
	"errors"
	"fmt"
	payment "github.com/stevenfrst/crowdfunding-api/drivers/midtrans"
	converter "github.com/stevenfrst/crowdfunding-api/helper/accounting"
	"github.com/stevenfrst/crowdfunding-api/usecase/campaign"
	"github.com/stevenfrst/crowdfunding-api/usecase/reward"
	"github.com/stevenfrst/crowdfunding-api/usecase/users"
	"gopkg.in/gomail.v2"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type TransactionUseCase struct {
	repoTransaction TransactionRepoInterface
	repoCampaign campaign.CampaignRepoInterface
	repoReward reward.RewardRepoInterface
	repoUser users.UserRepoInterface
	payment payment.MidtransInterface
	dialer gomail.Dialer
}

// NewUsecase function to create a new UsecaseInterface
func NewUsecase(transactionRepo TransactionRepoInterface,campaignRepository campaign.CampaignRepoInterface,payment payment.MidtransInterface,email gomail.Dialer, rewardRepo reward.RewardRepoInterface, repoUser users.UserRepoInterface) TransactionUsecaseInterface {
	return TransactionUseCase{
		transactionRepo,
		campaignRepository,
		rewardRepo,
		repoUser,
		payment,
		email,
	}
}

// CreateTransaction method to create new transaction
func (t TransactionUseCase) CreateTransaction(campaignID,userID,Nominal int) (Domain,error) {
	var transaction Domain
	rand.Seed(time.Now().UTC().UnixNano())
	id := rand.Intn(1000)
	log.Println(id)
	resp := t.payment.GetLinkResponse(id,Nominal)
	transaction.CampaignID = uint(campaignID)
	transaction.UserID = uint(userID)
	transaction.PaymentLink = resp.RedirectURL
	transaction.Nominal = Nominal
	transaction.ID = uint(id)
	transactionReturned,err := t.repoTransaction.CreateTransaction(&transaction)
	if err != nil {
		return Domain{},errors.New("Gagal Membuat Transaksi/Internal Error")
	}
	return transactionReturned,nil
}

// GetStatusByID method return domain transaction via id
func (t TransactionUseCase) GetStatusByID(ID int) (Domain,error) {
	transaction,err := t.repoTransaction.GetByID(ID)
	//log.Println(transaction)
	if err != nil  {
		return Domain{},errors.New("Internal Error")
	} else if transaction.ID == 0 {
		return Domain{},errors.New("Not Found")
	}

	return transaction,nil
}

// GetNotificationPayment method to get payment notification
func (t TransactionUseCase) GetNotificationPayment(input DomainNotification) (Domain,error) {
	transactionID, err := strconv.Atoi(input.OrderID)
	if err != nil {
		return Domain{},err
	}
	transaction,err := t.repoTransaction.GetByID(transactionID)
	if err != nil {
		return Domain{},err
	}
	if input.PaymentType == "bank_transfer" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" || input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
		transaction.PaymentType = input.PaymentType
		transaction.TransactionStatus = input.TransactionStatus
		transaction.FraudStatus = input.FraudStatus
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	updatedTransaction,_ := t.repoTransaction.UpdateTransaction(&transaction)

	campaign,err := t.repoCampaign.FindByID(int(updatedTransaction.CampaignID))
	if err != nil {
		return Domain{},err
	}


	if  updatedTransaction.Status == "paid" {
		campaign.Supporters = campaign.Supporters + 1
		campaign.AmountNow = campaign.AmountNow + updatedTransaction.Nominal
		rewardId,rewards,_ := t.GetRewardByAmount(updatedTransaction.Nominal)
		//if err != nil {
		//	return Domain{},err
		//}
		//
		userEmail,_ := t.repoUser.GetEmailByID(int(updatedTransaction.UserID))
		//if err != nil {
		//	return Domain{},err
		//}
		//log.Println(rewards,userEmail)
		var newMail = Email{
			Sender: "oppaidaisuki363@gmail.com",
			ToEmail:userEmail,
			Subject: "Notifikasi Pembayaran",
			Reward  :rewards,
			Nominal : converter.GoalAmountFormatIDR(updatedTransaction.Nominal),
		}

		_ = t.dialer.DialAndSend(SendEmailNotification(newMail))
		_, err = t.repoCampaign.UpdateCampaign(campaign)
		if err != nil {
			return Domain{},err
		}
		var domainHistory reward.DomainHistory

		domainHistory.UserID = updatedTransaction.UserID
		domainHistory.TransactionID = updatedTransaction.ID
		domainHistory.RewardID = uint(rewardId)

		_ = t.repoReward.SaveRewardHistory(domainHistory)
	}


	return *updatedTransaction,err
}

func (t TransactionUseCase) GetRewardByAmount(amount int) (int,string,error) {
	id,reward, err := t.repoReward.GetRewardByAmount(amount)
	if err != nil {
		return 0,"Failed to get reward",err
	}
	return id,reward,nil
}

// SendEmailNotification function to create gomail Message
func SendEmailNotification(sender Email) *gomail.Message  {
	var bodyEmail string
	if sender.Reward != "" {
		bodyEmail = fmt.Sprintf("Transaksi Sukses, Anda Menyumbang <b>%v</b> anda mendapatkan hadiah berupa <b>%v</b>",sender.Nominal,sender.Reward)
	} else {
		bodyEmail = fmt.Sprintf("Transaksi Sukses, Anda Menyumbang <b>%v</b>, dimengerti semoga hari anda cerah",sender.Nominal)
	}
	mailer := gomail.NewMessage()
	mailer.SetHeader("From",sender.Sender)
	mailer.SetHeader("To",sender.ToEmail)
	mailer.SetHeader("Subject",sender.Subject)
	mailer.SetBody("text/html",bodyEmail)

	//err := dialer.DialAndSend(mailer)
	//log.Println(sender)
	return mailer

}