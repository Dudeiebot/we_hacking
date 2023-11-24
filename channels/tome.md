DONT STRESS IT, JUST ME TALKING

So i create a function which take in a channel string `func sendData(ch chan string)`
I assign a some string to the var message, and then send the message to ch.

I create my func main
Then i make my unbuffered channel `myChannel := make(chan string)`
start a goroutine to that executes my sendData func `go sendData(myChannel)`
Then i receive data through myChannel `receivedData := <-myChannel`
PrintOut
close myChannel
