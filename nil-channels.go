

// https://www.godesignpatterns.com/2014/05/nil-channels-always-block.html


// Safely disable channels after they are closed.
for ch1 != nil || ch2 != nil {
	select {
	case _, ok := <-ch1:
		if !ok {
			ch1 = nil
		}
	case _, ok := <-ch2:
		if !ok {
			ch2 = nil
		}
	}
}