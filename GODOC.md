# otp
--
    import "github.com/hgfischer/go-otp"

Package go-otp implements one-time-password generators used in 2-factor
authentication systems like RSA-tokens and Google Authenticator. Currently this
supports both HOTP (RFC-4226) and TOTP (RFC-6238).

All tests used in this package, uses reference values from both RFCs to ensure
compatibility with another OTP implementations.

## Usage

```go
const (
	DefaultLength             = 6   // Default length of the generated tokens
	DefaultPeriod             = 30  // Default time period for TOTP tokens, in seconds
	DefaultRandomSecretLength = 100 // Default random secret length
	DefaultWindowBack         = 1   // Default TOTP verification window back steps
	DefaultWindowForward      = 1   // Default TOTP verification window forward steps
	MaxLength                 = 10  // Maximum token length
)
```

#### type HOTP

```go
type HOTP struct {
	Secret  string // The secret used to generate the token
	Length  uint8  // The token size, with a maximum determined by MaxLength
	Counter uint64 // The counter used as moving factor
}
```

HOTP is used to generate tokens based on RFC-4226.

Example:

    hotp := &HOTP{Secret: "your-secret", Counter: 1000, Length: 8}
    token := hotp.Get()

HOTP assumes a set of default values for Secret, Length and Counter. If no
Secret is informed, HOTP will generate a random one that you need to store with
the Counter, for future token verifications. Check this package constants to see
the current default values.

#### func (*HOTP) Get

```go
func (h *HOTP) Get() string
```
Return a token generated with the current HOTP settings

#### type TOTP

```go
type TOTP struct {
	Secret        string    // The secret used to generate a token
	Length        uint8     // The token length
	Time          time.Time // The time used to generate the token
	Period        uint8     // The step size to slice time, in seconds
	WindowBack    uint8     // How many steps HOTP will go backwards to validate a token
	WindowForward uint8     // How many steps HOTP will go forward to validate a token
}
```

TOTP is used to generate tokens based on RFC-6238.

Example:

    totp := &TOTP{Secret: "your-secret"}
    token := totp.Get()

TOTP assumes a set of default values for Secret, Length, Time, Period,
WindowBack and WindowForward.

If no Secret is informed, TOTP will generate a random one that you need to store
with the Counter, for future token verifications.

Check this package constants to see the current default values.

#### func (*TOTP) Get

```go
func (t *TOTP) Get() string
```
Generate a time-based token

#### func (*TOTP) Now

```go
func (t *TOTP) Now() *TOTP
```
Fluent interface to set the TOTP generator's time to the current date/time

#### func (TOTP) Verify

```go
func (t TOTP) Verify(token string) bool
```
Verify a token with the current settings, including the WindowBack and
WindowForward
