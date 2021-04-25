package kkpanic

func Try(try func()) Safe {
	safe := &SafeImpl{}
	Catch(try, func(r Caught) {
		safe.caught = r
	})

	return safe
}

type Safe interface {
	Catch(err interface{}, catch func(caught Caught)) Safe
	CatchAll(catch func(caught Caught)) Safe
	Finally(finally ...func())
}

type SafeImpl struct {
	caught Caught
}

func (s *SafeImpl) Catch(err interface{}, catch func(caught Caught)) Safe {
	if s.caught != nil {
		if err == s.caught.Data() {
			catch(s.caught)
		}
	}

	return s
}

func (s *SafeImpl) CatchAll(catch func(caught Caught)) Safe {
	if s.caught != nil {
		catch(s.caught)
	}

	return s
}

func (s *SafeImpl) Finally(finally ...func()) {
	for i := len(finally) - 1; i >= 0; i-- {
		defer finally[i]()
	}
}
