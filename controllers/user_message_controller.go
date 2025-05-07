package controllers

import (
	"log"
	"time"

	"github.com/promptlabth/ms-payments/entities"
	"gorm.io/gorm"
)

type UserMessageController struct {
}

func NewUserMessageController() *UserMessageController {
	return &UserMessageController{}
}

// Reset User Message Counts Function. Need Cronjob to run daily
func (t *UserMessageController) ResetMessage(db *gorm.DB) {
	now := time.Now()
	log.Println("Routine Reset Message Running...")

	// 1. Reset message count for free accounts and old active premium accounts on the 1st of the month
	if now.Day() == 1 {
		log.Printf("Resetting messages for free and old active premium accounts...")
		// check free account and old active premium accounts (do not have sub date)
		db.Table("user_balance_messages").
			Where("EXISTS (SELECT 1 FROM users WHERE users.firebase_id = user_balance_messages.firebase_id AND (users.plan_id = 4 OR (users.plan_id != 4 AND users.sub_date IS NULL)))").
			Update("balance_message", 0)
	}

	// 2. Reset message for newly subscribed premium accounts on subscription date
	log.Printf("Resetting messages for newly subscribed premium accounts...")
	// check if the subscription date is the same as today ( Check Day and Month )
	db.Table("user_balance_messages").
		Where("EXISTS (SELECT 1 FROM users WHERE users.firebase_id = user_balance_messages.firebase_id AND (users.plan_id != 4 AND ( EXTRACT(DAY FROM users.end_sub_date) = EXTRACT(DAY FROM ?::timestamp) AND EXTRACT(MONTH FROM users.end_sub_date) = EXTRACT(MONTH FROM ?::timestamp)) AND users.monthly = true))", now, now).
		Update("balance_message", 0)

	// 3. Change expired premium accounts to free, reset message, and set sub_date & end_sub_date to NULL
	log.Printf("Resetting messages and changing expired premium accounts to free...")
	// check if any end_sub_date is expired and reset message
	db.Table("user_balance_messages").
		Where("EXISTS (SELECT 1 FROM users WHERE users.firebase_id = user_balance_messages.firebase_id AND (?::DATE > users.end_sub_date::DATE) AND users.monthly = false)", now).
		Update("balance_message", 0)

	// set their plan_id to free, reset sub_date and end_sub_date to NULL
	db.Model(&entities.User{}).
		Where("(?::DATE > users.end_sub_date::DATE) AND users.monthly = false", now).
		Updates(map[string]interface{}{"plan_id": 4, "sub_date": nil, "end_sub_date": nil})

	log.Println("resetMessage completed.")
}
