package otp

// Default settings for all generators
const (
	DefaultHashAlgo			  = SHA1 // Default hash algorithm to SHA1
	DefaultLength             = 6    // Default length of the generated tokens
	DefaultPeriod             = 30   // Default time period for TOTP tokens, in seconds
	DefaultRandomSecretLength = 100  // Default random secret length
	DefaultWindowBack         = 1    // Default TOTP verification window back steps
	DefaultWindowForward      = 1    // Default TOTP verification window forward steps
)

// Maximum values for all generators
const (
	MaxLength = 10 // Maximum token length
)

// Valid hash algorithm
type Hash int
const (
	SHA1 	Hash = iota
	SHA256
	SHA512
)
