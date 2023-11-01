curl -X POST http://localhost:8080/payment-subscription \
  -H "Content-Type: application/json" \
  -d '{
    "TransactionStripeID": "pi_3O4N9jAom1IgIvKK1k1LMSgL",
    "UserID": 2,
    "PaymentMethodID": 1,
    "PlanID": 2
  }'
