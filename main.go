package main

const (
	DISCORD_WEBHOOK string = "INSERT WEBHOOK URL HERE"
	USER_AGENT      string = "github.com/un4gi/feedthebaby"
)

func main() {
	// SetupCloseHandler()

	// Could make this a for loop with a sleep timer to send messages every x seconds
	CheckTargetItems()

}

// SetupCloseHandler will allow the user to exit the program gracefully
// func SetupCloseHandler() {
// 	c := make(chan os.Signal)
// 	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
// 	go func() {
// 		<-c
// 		log.Println("\r- Ctrl+C pressed in Terminal")
// 		os.Exit(0)
// 	}()
// }
