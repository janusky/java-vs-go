#!/bin/bash
count=10
server=localhost:8080
while [ $count -gt 0 ]
do
  echo "create new client"
  curl -skw "\ntime_total: %{time_total}s\n\n" -X POST http://$server/client/new/100

  echo "create new transaction"
  curl -skw "\ntime_total: %{time_total}s\n\n" --header "Content-Type: application/json" \
    --request POST \
    --data '{"from_client_id":1,"to_client_id":2, "amount":1}' \
    http://$server/transaction

  echo "get balance"
  curl -skw "\ntime_total: %{time_total}s\n\n" -X GET http://$server/client/1/balance

  count=$(( $count - 1 ))
  echo "-------------------------------------------"
done
echo "EMD curl-api.sh"