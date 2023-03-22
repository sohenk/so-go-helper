package password

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "sjm1006"
	encrptPassword, err := HashPassword(password)
	if err != nil {
		panic(err)
	}
	t.Log("hashpassword", encrptPassword)
	t.Log("passwordcheck", CheckPasswordHash("sjm1006", encrptPassword))
}
