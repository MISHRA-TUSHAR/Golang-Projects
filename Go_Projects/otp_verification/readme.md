# OTP Verification using Twilio in Go

This project demonstrates how to implement OTP (One-Time Password) verification using Twilio in Go.

Create a Twilio account and get your account SID and auth token

Create a Twilio verify service and get the service SID

Create a .env file in the root of the project and add the in the above credentials from Twilio

```
TWILIO_ACCOUNT_SID=<ACCOUNT SID>
TWILIO_AUTHTOKEN=<AUTH TOKEN>
TWILIO_SERVICES_ID=<SERVICE ID>
```

Install dependencies
```
go mod download
```

Run the server
```
go run cmd/main.go
```

Send OTP
Send a POST request to the /otp endpoint with the following body to send an OTP to a user's phone number

POST /otp

Request Body
```
{
  "phoneNumber": "<phone-number-with-country-code>"
}
```

```
curl -H "Content-Type: application/json" -X POST -d '{"phoneNumber": "+91xxxxxxxxxx"}' http://localhost:8000/otp
```

Response
```
{
  "status": 202,
  "message": "success",
  "data": "OTP sent successfully"
}
```

Verify OTP
Verify a user's OTP by sending a POST request to the /verify endpoint with the following body that contains the phone number and the OTP code received by the user

POST /verifyOTP

Request Body
```
{
  "user": {
    "phoneNumber": "<phone-number-with-country-code>"
  },
  "code": "<code here>"
}

```

Response
```
{
  "status": 202,
  "message": "success",
  "data": "OTP verified successfully"
}
```
