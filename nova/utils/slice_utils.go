package utils

import "errors"

// TODO: Convert ALL of this to generics as soon as 1.18 is released.

func FindSingle[T any](list []*T) (*T, error) {
	if len(list) != 1 {
		return nil, errors.New("length of list must equal exactly one")
	}

	single := list[0]

	if single == nil {
		return nil, errors.New("list is empty")
	}

	return single, nil
}

func Contains[T comparable](list []T, e T) bool {
	for _, a := range list {
		if a == e {
			return true
		}
	}
	return false
}

func Prepend[T any](list []T, e T) []T {
	list = append([]T{e}, list...)
	return list
}

func PopEnd[T any](list *[]T) T {
	f := len(*list)
	rv := (*list)[f-1]
	*list = (*list)[:f-1]
	return rv
}

func PopStart[T any](list *[]T) T {
	f := len(*list)
	rv := (*list)[0]
	*list = (*list)[1:f]
	return rv
}
