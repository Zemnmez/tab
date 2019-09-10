package resolver

type Query struct{}
type Mutation struct{}

type Resolver struct{}

func (Resolver) Query() Query       { return Query{} }
func (Resolver) Mutation() Mutation { return Mutation{} }

func (Query) Ok() bool    { return true }
func (Mutation) Ok() bool { return true }

type Etc struct{}
