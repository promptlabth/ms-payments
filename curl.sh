curl -X POST http://localhost:8080/payment-subscription \
  -H "Content-Type: application/json" \
  -d '{
    "TransactionStripeID": "pi_3NtR3HAom1IgIvKK1rqZaC0G",
    "SubscriptionStatus": "active",
    "UserID": 1,
    "PaymentMethodID": 1,
    "PlanID": 1
  }'
