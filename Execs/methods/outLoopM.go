package methods

import "math/rand"

func StepGiving(code string) int {
	// Step giving detemines whether func have to checked long or short

	var from int = 25
	var to int = 95 // in reality it goes to "to" + "from" `400`

	var time int = rand.Intn(to) + from
	return time
}
