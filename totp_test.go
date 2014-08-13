package otp

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTOTP(t *testing.T) {
	secret := `12345678901234567890`
	table := map[time.Time]string{
		time.Date(1970, 1, 1, 0, 0, 59, 0, time.UTC):     `94287082`,
		time.Date(2005, 3, 18, 1, 58, 29, 0, time.UTC):   `07081804`,
		time.Date(2005, 3, 18, 1, 58, 31, 0, time.UTC):   `14050471`,
		time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC):  `89005924`,
		time.Date(2033, 5, 18, 3, 33, 20, 0, time.UTC):   `69279037`,
		time.Date(2603, 10, 11, 11, 33, 20, 0, time.UTC): `65353130`,
	}

	for tm, expected := range table {
		totp := &TOTP{Secret: secret, Length: 8, Time: tm, Period: 30}
		result := totp.Get()
		assert.Equal(t, expected, result, tm.String())
		assert.True(t, totp.Verify(result))
	}
}
