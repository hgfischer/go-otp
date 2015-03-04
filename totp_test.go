package otp

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTOTP(t *testing.T) {
	secret := `JBSWY3DPEHPK3PXP`
	table := map[time.Time]string{
		time.Date(1970, 1, 1, 0, 0, 59, 0, time.UTC):     `41996554`,
		time.Date(2005, 3, 18, 1, 58, 29, 0, time.UTC):   `33071271`,
		time.Date(2005, 3, 18, 1, 58, 31, 0, time.UTC):   `28358462`,
		time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC):  `94742275`,
		time.Date(2033, 5, 18, 3, 33, 20, 0, time.UTC):   `28890699`,
		time.Date(2603, 10, 11, 11, 33, 20, 0, time.UTC): `94752434`,
	}

	for tm, expected := range table {
		totp := &TOTP{Secret: secret, Length: 8, Time: tm, Period: 30}
		result := totp.Get()
		assert.Equal(t, expected, result, tm.String())
		assert.True(t, totp.Verify(result))
	}
}

func TestTOTPShouldBeCroppedToMaxLength(t *testing.T) {
	totp := &TOTP{Length: 20}
	result := totp.Get()
	assert.Equal(t, MaxLength, len(result))
}

func TestTOTPShouldUseDefaultValues(t *testing.T) {
	totp := &TOTP{}
	result := totp.Get()
	assert.NotEmpty(t, totp.Secret)
	assert.Equal(t, uint8(DefaultLength), totp.Length)
	assert.False(t, totp.Time.IsZero())
	assert.Equal(t, totp.Length, uint8(len(result)))
}

func TestTOTPShouldUseCurrentTimeWithFluentInterface(t *testing.T) {
	past := time.Date(1979, 3, 26, 19, 30, 0, 0, time.Local)
	now := time.Now().Format(time.Kitchen)
	totp := &TOTP{Time: past}
	totp.Now().Get()
	assert.Equal(t, now, totp.Time.Format(time.Kitchen))
}

func TestTOTPVerifyShouldFail(t *testing.T) {
	past := time.Now().Add(time.Second * DefaultPeriod * -2)
	totp := &TOTP{Time: past}
	token := totp.Get()
	assert.False(t, totp.Now().Verify(token))
}
