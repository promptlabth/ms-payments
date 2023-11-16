curl -X POST http://localhost:8080/payment-subscription \
  -H "Content-Type: application/json" \
  -d '{
    "PaymentIntentId": "pi_3NzipoAom1IgIvKK1rvmu0rK",
    "UserID": 2,
    "PaymentMethodID": 1,
    "PlanID": 2
  }'
