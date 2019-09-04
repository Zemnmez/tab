package tab

type Resolver struct{}

type Query struct{}

func (Resolver) Query() QueryResolver { return Query{} }

type Mutation struct{}

func (Resolver) Mutation() MutationResolver { return Mutation{} }
