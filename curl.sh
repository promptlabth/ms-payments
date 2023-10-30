curl -X POST http://localhost:8080/payment-subscription \
  -H "Content-Type: application/json" \
  -d '{
    "TransactionStripeID": "your_transaction_stripe_id",
    "Datetime": "your_datetime",
    "StartDatetime": "your_start_datetime",
    "EndDatetime": "your_end_datetime",
    "SubscriptionStatus": "your_subscription_status",
    "UserID": 1,
    "PaymentMethodID": 1,
    "PlanID": 1
  }'
