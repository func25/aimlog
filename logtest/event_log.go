package logtest

func (e *event) Bool(key string, value bool) *event {
	if e.event.Bool(key, value) == nil {
		return nil
	}

	return e
}

func (e *event) Int(key string, value int) *event {
	if e.event.Int(key, value) == nil {
		return nil
	}

	return e
}

func (e *event) Float32(key string, value float32) *event {
	if e.event.Float32(key, value) == nil {
		return nil
	}

	return e
}

func (e *event) Float64(key string, value float64) *event {
	if e.event.Float64(key, value) == nil {
		return nil
	}
	return e
}

func (e *event) Interface(key string, i interface{}) *event {
	if e.event.Interface(key, i) == nil {
		return nil
	}

	return e
}

func (e *event) Str(key string, value string) *event {
	if len(value) > 0 && e.event.Str(key, value) == nil {
		return nil
	}

	return e
}

func (e *event) Err(err error) *event {
	if e.event.Err(err) == nil {
		return nil
	}

	return e.batch("error", err.Error())
}
