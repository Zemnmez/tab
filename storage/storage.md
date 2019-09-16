# Storage
Storage is implemented as an abstraction on a key-value type store. As you'll note, the actual `IStorage`, `ITxn` etc bear a striking resemblance to the `os.File` interfaces. This is for two reasons: (1) this interface is extremely compatiable with other parts of the go libraries and (2) it means that the storage can be potentially reduced to an on-disk database for debugging and portability purposes.

## Tables

Tables fit into an interesting niche in key-value stores. The idea of collecting kinds of data to iterate through is indispensible in searching and indexing of such stores, but it's usually not the case that the kv storage provides first party access to an API that performs this action -- especially since a 'table' in kv storage is not so much a separate logical unit, but a prefix to the keys of some type of thing. That is to say, if I have a table of users, it is expected that my users' keys will all have the same prefix that an iterator can use to determine what table the user is in.

This creates somewhat of a small conundrum: most of the time, we want to retrieve an item by its unique uuid, but a small percentage of the time we may want to retrieve it by its type. Unless we make several layers of database indirection we can't easily have both these things.

My solution to this problem is to change the way identifiers are defined for our database stored types. Instead of having the ID of a record being a 1:1 mapping of entropy, we instead take `type_identifier + entropy`. This maintains the uniqueness and mappability of keys while also allowing iteration through types, or tables of values.