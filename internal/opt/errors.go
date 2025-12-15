package opt

func (o *Option) AddError(err error) {
	o.Params["error"] = err
}

func (o *Option) HasErrors() bool {
	if _, ok := o.Params["error"]; ok {
		return true
	}
	return false
}

func (o *Option) DelErrors() {
	delete(o.Params, "error")
}

func (o *Option) Errors() error {
	if _, ok := o.Params["error"]; ok {
		return o.Params["error"].(error)
	}
	return nil
}
