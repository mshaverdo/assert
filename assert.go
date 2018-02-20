package assert

// If true, failed assertions will panics. False means 'no assertions'
var Enabled = true

func True(t bool, message string) {
	if Enabled && t != true {
		panic(message)
	}
}

func False(f bool, message string) {
	if Enabled && f != false {
		panic(message)
	}
}