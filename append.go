package multierror

// Append is a helper function that will append more errors
// onto an Error in order to create a larger multi-error.
//
// If err is not a multierror.Error, then it will be turned into
// one. If any of the errs are multierr.Error, they will be flattened
// one level into err.
func Append(err error, errs ...error) *Error {
	me,ok := err.(*Error)
	
	if me == nil {
		me = &Error{
			Errors: make([]error, 0, len(errs)),
		}
	}
	
	if !ok {
		me = &Error{
			Errors: make([]error, 1, len(errs)+1),
		}
		me.Errors[0] = err
	}	
	
	if me.Errors == nil {
		me.Errors = make([]error, 0, len(errs))
	}
	
	for _,e := range errs {
		if e != nil {
			me.Errors = append(me.Errors, e)
		}
	}
	
	return me
}
