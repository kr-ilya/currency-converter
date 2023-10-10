# Backend mini-apps 

Mini-application backend (WebApp).

Written in Go.


### Validation of requests from the mini-application

Verification occurs when each request to the API server is executed. If the request fails validation, HTTP Code 401 Unauthorized is returned.

The verification code is located in the file `/internal/app/auth.go`