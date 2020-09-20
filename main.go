package main

import "watt/server"

func main() {
	// token,  err := utils.Token(1234)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(token)
	// return
	server.Run()

	// str := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyaWQiOjEyMzQsImV4cCI6MTYwMDU3OTkyMSwiaXNzIjoid2F0dF9qd3RfaXNzdWVyIn0.W9wy-3SOZ6r121CmNr2B75RezM6apG2_96nTpxB3hZY"
	// uid, err := utils.Parse(str)

	// if err != nil {
	// jwt.ValidationErrorExpired
	// fmt.Println(err.(*jwt.ValidationError).Errors)
	// }

	// fmt.Println(uid)

}
