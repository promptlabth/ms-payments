package repository

import (
	"errors"
	"promptlabth/ms-payments/entities"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestStoreSubscription(t *testing.T) {
	// Sample plan data
	plan := entities.Plan{
		PlanType: "premium",
		Datetime: "2023-01-01T00:00:00Z",
	}

	// Set up the test cases
	tests := []struct {
		name    string
		payment entities.PaymentSubscription
		setup   func(mock sqlmock.Sqlmock)
		wantErr bool
	}{
		{
			name: "success",
			payment: entities.PaymentSubscription{
				UserID:          uintPtr(1),
				PaymentMethodID: uintPtr(2),
				Plan:            plan,
				PaymentIntentId: "tx_123",
			},
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`INSERT INTO payments`).
					WithArgs(1, 2, 1, "tx_123", sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), "active").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name: "database error",
			payment: entities.PaymentSubscription{
				UserID:          uintPtr(1),
				PaymentMethodID: uintPtr(2),
				Plan:            plan,
				PaymentIntentId: "tx_123",
			},
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`INSERT INTO payments`).
					WithArgs(1, 2, 1, "tx_123", sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), "active").
					WillReturnError(errors.New("database error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up the mock database and repository
			mockDb, mock, err := sqlmock.New()
			dialector := postgres.New(postgres.Config{
				Conn:       mockDb,
				DriverName: "postgres",
			})
			assert.NoError(t, err)
			db, err := gorm.Open(dialector, &gorm.Config{})
			repo := &PaymentSubscriptionsRepository{conn: db}

			// Set up the expected database interactions
			if tt.setup != nil {
				tt.setup(mock)
			}

			// Call the method under test
			err = repo.Store(tt.payment)

			// Check the results
			assert.Equal(t, tt.wantErr, err != nil)

			// Ensure all expectations are met
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func uintPtr(i uint) *uint {
	return &i
}
