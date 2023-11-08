curl -X POST http://localhost:8080/payment-subscription \
  -H "Content-Type: application/json" \
  -d '{
    "PaymentIntentId": "pi_3O817lAom1IgIvKK0jNizE1P",
    "UserID": 2,
    "PaymentMethodID": 1,
    "PlanID": 2
  }'
