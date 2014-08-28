package main

import (
	"encoding/base32"
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
	key, _ := base32.StdEncoding.DecodeString(*secret)

	totp := &otp.TOTP{
		Secret: string(key),
		Length: uint8(*length),
		Period: uint8(*period),
	}
	fmt.Println("TOTP:", totp.Get())

	hotp := &otp.HOTP{
		Secret:  string(key),
		Length:  uint8(*length),
		Counter: *counter,
	}
	fmt.Println("HOTP:", hotp.Get())
}