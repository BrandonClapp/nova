package random

import (
	"strings"

	uuid "github.com/nu7hatch/gouuid"
)

func GetShortRandom() (string, error) {
	slugHash, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	hash := strings.Split(slugHash.String(), "-")[0]

	return hash, nil
}
