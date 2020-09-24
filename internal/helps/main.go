package helps

// StartupMessage : startup message loader for thermide
func StartupMessage() {
	var logo string
	logo += "\n%s"
	logo += " ┌───────────────────────────────────────────────────┐\n"
	logo += " │ %s │\n"
	logo += " │ %s │\n"
	logo += " │                                                   │\n"
	logo += " │ Handlers %s  Threads %s │\n"
	logo += " │ Prefork .%s  PID ....%s │\n"
	logo += " └───────────────────────────────────────────────────┘"
	logo += "%s\n\n"

}
