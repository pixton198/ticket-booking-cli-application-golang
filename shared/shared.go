package shared

import "strings"

var Myvar = "somevalue"
func Vallidateuserinput (FirstName string, LastName string, Email string, UserTickets uint, remainingtickets uint ) (bool, bool, bool) {
	isvalidname := len(FirstName) >= 2 && len(LastName) >= 2
	isvalidemail := strings.Contains(Email, "@")
	isvalidticket := UserTickets > 0 && UserTickets <= remainingtickets
	return isvalidname, isvalidemail, isvalidticket
}