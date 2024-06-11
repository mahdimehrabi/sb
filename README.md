
# Create a user with its addresses
```
curl -X POST \
http://localhost:8080/api/articles \
-H 'Content-Type: application/json' \
-d '{
"name": "John",
"lastname": "Doe",
"addresses": {
"street": "123 Main St",
"city": "Anytown",
"state": "Anystate",
"zipCode": "12345",
"country": "USA"
}
}' 
```