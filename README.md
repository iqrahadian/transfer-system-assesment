# paperid-assesment

# simple design & flow diagram

https://docs.google.com/document/d/1Rqa4FevkjujxPm5KzzC16yJtXGOyfW7OKxZyIynSzXM/edit?usp=sharing

# command
1. to run serivce simply do `go run main.go`

2. retrieve current data state(transaction,  wallet & user)

curl --location --request GET 'localhost:3000/data'

3. submit disburse request

curl --location --request POST 'localhost:3000/disburse' \
--header 'Content-Type: application/json' \
--data-raw '{
    "sender_id":"user1",
    "source_account_id":"wallet1",
    "destination_number":"456",
    "amount":10000
}' 
