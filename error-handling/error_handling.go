// Package erratum implements a solution of the exercise titled `Error Handling'.
package erratum

import "errors"

// Use accesses given resource with error handling and recovering from panic.
func Use(o ResourceOpener, input string) (err error) {
	var res Resource
	// Keep trying to open when TransientError
	for res, err = o(); err != nil; res, err = o() {
		if _, ok := err.(TransientError); !ok {
			return err // Give it up
		}

	}
	defer res.Close() // Close when leaving
	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case FrobError:
				// Try to recover from recoverable panic
				res.Defrob(e.defrobTag)
				err = e.inner
			case error:
				err = e
			default:
				err = errors.New("unknown error")
			}
		}
	}()
	res.Frob(input)
	return err
}
