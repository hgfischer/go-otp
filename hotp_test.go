package otp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hmacsha1(t *testing.T) {
	secret := `12345678901234567890`
	table := map[int64]string{
		0: `cc93cf18508d94934c64b65d8ba7667fb7cde4b0`,
		1: `75a48a19d4cbe100644e8ac1397eea747a2d33ab`,
		2: `0bacb7fa082fef30782211938bc1c5e70416ff44`,
		3: `66c28227d03a2d5529262ff016a1e6ef76557ece`,
		4: `a904c900a64b35909874b33e61c5938a8e15ed1c`,
		5: `a37e783d7b7233c083d4f62926c7a25f238d0316`,
		6: `bc9cd28561042c83f219324d3c607256c03272ae`,
		7: `a4fb960c0bc06e1eabb804e5b397cdc4b45596fa`,
		8: `1b3c89f65e6c9e883012052823443f048b4332db`,
		9: `1637409809a679dc698207310c8c7fc07290d9e5`,
	}

	for cnt, expected := range table {
		result := hmacSHA1([]byte(secret), movingFactorToCounter(cnt))
		assert.Equal(t, expected, fmt.Sprintf("%x", result))
	}
}

func TestGenerateHOTP(t *testing.T) {
	secret := `12345678901234567890`
	table := map[int64]string{
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
		result := GenerateHOTP(secret, cnt, 6)
		assert.Equal(t, expected, result)
	}
}

func TestGeneratedHOTPShouldBeCroppedTo10CodeDigits(t *testing.T) {
	secret := `12345678901234567890`
	result := GenerateHOTP(secret, 0, 20)
	assert.Equal(t, 10, len(result))
}
