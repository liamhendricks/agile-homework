package swapi

import "errors"

var ErrNoData = errors.New("no data")

type SwapiNotFoundError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e SwapiNotFoundError) Error() string {
	return e.Msg
}
