package otp

const (
	// Default length of the generated tokens
	DefaultLength = 6
	// Default time period for TOTP tokens
	DefaultPeriod = 30
	// Default random secret length
	DefaultRandomSecretLength = 100
	// Default TOTP verification window back steps
	DefaultWindowBack = 1
	// Default TOTP verification window forward steps
	DefaultWindowForward = 1
	// Maximum token length
	MaxLength = 10
)
