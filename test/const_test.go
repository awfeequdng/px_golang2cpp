package main

const (
	// StateNone means this schema element is absent and can't be used.
	StateNone SchemaState = iota
	// StateDeleteOnly means we can only delete items for this schema element.
	StateDeleteOnly
	// StateWriteOnly means we can use any write operation on this schema element,
	// but outer can't read the changed data.
	StateWriteOnly
	// StateWriteReorganization means we are re-organizing whole data after write only state.
	StateWriteReorganization
	// StateDeleteReorganization means we are re-organizing whole data after delete only state.
	StateDeleteReorganization
	// StatePublic means this schema element is ok for all write and read operations.
	StatePublic
)

const (
	// StateNone means this schema element is absent and can't be used.
	StateNone1 SchemaState = 1 << iota
	// StateDeleteOnly means we can only delete items for this schema element.
	StateDeleteOnly1
	// StateWriteOnly means we can use any write operation on this schema element,
	// but outer can't read the changed data.
	StateWriteOnly1
	// StateWriteReorganization means we are re-organizing whole data after write only state.
	StateWriteReorganization1
	// StateDeleteReorganization means we are re-organizing whole data after delete only state.
	StateDeleteReorganization1
	// StatePublic means this schema element is ok for all write and read operations.
	StatePublic1
)
