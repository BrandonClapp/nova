package slug

import (
	"fmt"

	"github.com/brandonclapp/nova/random"
	sl "github.com/gosimple/slug"
)

func MakeSlug(str string) string {
	compiled := sl.Make(str)
	return compiled
}

func MakeUniqueSlug(str string) (string, error) {
	compiled := MakeSlug(str)
	rand, err := random.GetShortRandom()

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s-%s", compiled, rand), nil
}
