package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	str := "Welcome to the Tech Palace, "
	cus := strings.ToUpper(customer)
	str = str + cus
	return str
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	str := strings.Repeat("*", numStarsPerLine)
	str = str + "\n" + welcomeMsg + "\n"
	str = str + strings.Repeat("*", numStarsPerLine)
	return str
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	str := strings.Trim(oldMsg, "*")
	str = strings.Trim(str, "\n")
	str = strings.ReplaceAll(str, "*", " ")
	str = strings.TrimSpace(str)
	return str
}
