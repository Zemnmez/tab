package types

func (s SpecialUser) GetName() string { return s.Name() }
func (s SpecialUser) Name() string    { return s.Id.String() }
