// Package erratum implements a solution of the exercise titled `Error Handling'.
package erratum

import "errors"

// Use accesses given resource with error handling and recovering from panic.
func Use(o ResourceOpener, input string) (err error) {
	// Keep trying to open when TransientError
	resource, err := o()
	for err != nil {
		if _, ok := err.(TransientError); !ok {
			return err // Give it up
		}
		resource, err = o()

	}
	defer resource.Close() // Close when leaving
	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case FrobError:
				// Try to recover from recoverable panic
				resource.Defrob(e.defrobTag)
				err = e.inner
			case error:
				err = e
			default:
				err = errors.New("unknown error")
			}
		}
	}()
	resource.Frob(input)
	return err
}
