package types

func (s SpecialUser) GetName() string { return s.Id.String() }