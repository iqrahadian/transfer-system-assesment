# paperid-assesment

# command
1. retrieve current data state(transaction,  wallet & user)

curl --location --request GET 'localhost:3000/data'

2. submit disburse request

curl --location --request POST 'localhost:3000/disburse' \
--header 'Content-Type: application/json' \
--data-raw '{
    "sender_id":"user1",
    "source_account_id":"wallet1",
    "destination_number":"456",
    "amount":10000
}' 
