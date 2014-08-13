package otp

import "time"

type TOTP struct {
	Secret string
	Length uint8
	Time   time.Time
	Period uint8
	Window struct {
		Back    uint8
		Forward uint8
	}
}

func (t *TOTP) setDefaults() {
	if len(t.Secret) == 0 {
		t.Secret = randomString(100)
	}
	if t.Length == 0 {
		t.Length = 6
	}
	if t.Time.IsZero() {
		t.Time = time.Now()
	}
	if t.Period == 0 {
		t.Period = 30
	}
	if t.Window.Back == 0 {
		t.Window.Back = 1
	}
	if t.Window.Forward == 0 {
		t.Window.Forward = 1
	}
}

func (t *TOTP) normalize() {
	if t.Length > 10 {
		t.Length = 10
	}
}

func (t *TOTP) Get() string {
	t.setDefaults()
	t.normalize()
	ts := uint64(t.Time.Unix() / int64(t.Period))
	hotp := &HOTP{Secret: t.Secret, Counter: ts, Length: t.Length}
	return hotp.Get()
}

func (t *TOTP) GetNow() string {
	t.Time = time.Now()
	return t.Get()
}

func (t TOTP) Verify(token string) bool {
	t.setDefaults()
	t.normalize()
	for i := int(t.Window.Back) * -1; i <= int(t.Window.Forward); i++ {
		t.Time = t.Time.Add(time.Second * time.Duration(int(t.Period)*i))
		if t.Get() == token {
			return true
		}
	}
	return false
}