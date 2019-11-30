package util

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

// NewUUID 生成新的uuid
func NewUUID() string {
	u2 := uuid.NewV4()
	no := strings.Replace(u2.String(), "-", "", -1)
	return no
}
