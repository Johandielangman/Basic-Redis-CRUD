#!/bin/bash

curl -X POST -H "Content-Type: application/json" -d '{
  "customer_id": "'$(uuidgen)'",
  "line_items": [
    {
      "item_id": "'$(uuidgen)'",
      "quantity": 5,
      "price": 1999
    }
  ]
}' http://localhost:3000/orders
