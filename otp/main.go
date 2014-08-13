package main

import (
	"flag"
	"fmt"

	"github.com/hgfischer/go-otp"
)

var (
	secret  = flag.String("secret", "secret", "Secret")
	length  = flag.Uint("length", otp.DefaultLength, "OTP length")
	period  = flag.Uint("period", otp.DefaultPeriod, "Period in seconds")
	counter = flag.Uint64("counter", 0, "Counter")
)

func main() {
	flag.Parse()

	totp := &otp.TOTP{
		Secret: *secret,
		Length: uint8(*length),
		Period: uint8(*period),
	}
	fmt.Println("TOTP:", totp.Get())

	hotp := &otp.HOTP{
		Secret:  *secret,
		Length:  uint8(*length),
		Counter: *counter,
	}
	fmt.Println("HOTP:", hotp.Get())
}