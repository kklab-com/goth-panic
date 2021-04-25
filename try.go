package kkpanic

func Try(try func()) SafeCatch {
	safe := &Safe{}
	Catch(try, func(r Caught) {
		safe.caught = r
	})

	return safe
}

type SafeCatch interface {
	Catch(catch func(caught Caught)) SafeFinally
}

type SafeFinally interface {
	Finally(finally func())
}

type Safe struct {
	caught Caught
}

func (s *Safe) Catch(catch func(caught Caught)) SafeFinally {
	if s.caught != nil {
		catch(s.caught)
	}

	return s
}

func (s *Safe) Finally(finally func()) {
	finally()
}
