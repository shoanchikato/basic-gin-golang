package route

import (
	"errors"
	"strconv"
)

func getIDParam(stringID string) (uint, error) {
	idInt, err := strconv.Atoi(stringID)
	if err != nil {
		return 0, errors.New("bad id value, can't be parsed to an interger")
	}

	id := uint(idInt)

	return id, nil
}
