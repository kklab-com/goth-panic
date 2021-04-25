package kkpanic

func Try(try func()) SafeCatch {
	safe := &Safe{}
	Catch(try, func(r Caught) {
		safe.caught = r
	})

	return safe
}

type SafeCatch interface {
	Catch(err interface{}, catch func(caught Caught)) SafeCatch
	CatchAll(catch func(caught Caught)) SafeCatch
	Finally(finally func())
}

type Safe struct {
	caught Caught
}

func (s *Safe) Catch(err interface{}, catch func(caught Caught)) SafeCatch {
	if s.caught != nil {
		if err == s.caught.Data() {
			catch(s.caught)
		}
	}

	return s
}

func (s *Safe) CatchAll(catch func(caught Caught)) SafeCatch {
	if s.caught != nil {
		catch(s.caught)
	}

	return s
}

func (s *Safe) Finally(finally func()) {
	finally()
}
