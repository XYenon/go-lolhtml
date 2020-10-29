package lolhtml

/*
#include "lol_html.h"
*/
import "C"
import "errors"

var ErrCannotGetErrorMessage = errors.New("cannot get error message from underlying lol-html lib")

// getError is a helper function that gets error message for the last function call.
// You should make sure there is an error when calling this, or the function interprets
// the NULL error message obtained as ErrCannotGetErrorMessage.
func getError() error {
	errC := (*str)(C.lol_html_take_last_error())
	defer errC.Free()
	if errMsg := errC.String(); errMsg != "" {
		return errors.New(errMsg)
	}
	return ErrCannotGetErrorMessage
}
