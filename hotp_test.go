package otp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHOTP(t *testing.T) {
	secret := `12345678901234567890`
	table := map[uint64]string{
		0: `755224`,
		1: `287082`,
		2: `359152`,
		3: `969429`,
		4: `338314`,
		5: `254676`,
		6: `287922`,
		7: `162583`,
		8: `399871`,
		9: `520489`,
	}

	for cnt, expected := range table {
		hotp := &HOTP{Secret: secret, Counter: cnt, Length: 6}
		result := hotp.Get()
		assert.Equal(t, expected, result)
	}
}

func TestHOTPShouldBeCroppedToMaxLength(t *testing.T) {
	hotp := &HOTP{Length: 20}
	result := hotp.Get()
	assert.Equal(t, MaxLength, len(result))
}

func TestHOTPShouldUseDefaultValues(t *testing.T) {
	hotp := &HOTP{}
	result := hotp.Get()
	assert.Equal(t, DefaultLength, hotp.Length)
	assert.NotEmpty(t, hotp.Secret)
	assert.Equal(t, hotp.Length, len(result))
}
