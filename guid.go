package neutron

import (
	"crypto/rand"
	"encoding/hex"
)

func GUID() string {
	buf := make([]byte, 8)
	n, _ := rand.Read(buf)
	return hex.EncodeToString(buf[:n])
}
