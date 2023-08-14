package repository

import (
	"promptlabth/ms-payments/entities"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestStore(t *testing.T) {
	// Create a new mock SQL connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := PaymentRepository{DB: db}

	// Define test payment data
	payment := entities.Payment{
		// ... set properties here ...
	}

	// Set up your expected SQL query
	mock.ExpectExec(`INSERT INTO payments`).WithArgs(
		payment.UserID, payment.PaymentMethodID, payment.Coin, payment.TransactionStripeID, sqlmock.AnyArg(), payment.FeatureID,
	).WillReturnResult(sqlmock.NewResult(1, 1)) // Simulate a successful insert

	// Call the store method
	err = repo.Store(payment)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Ensure the mock expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
