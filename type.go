package bossy

type propagation struct {
	parentNode []string
	childNode  []string
}

// A restriction is a function that determines whether the current level is
// valid. You may only have one restriction per level.
type restriction func() bool

// a config is a merging of three linked trees.
// the structure tree defines the valid/allowed structure within the config.
//   anything that does not fit into the structure is rejected when read.
// the defaults is another structure containing all default values - where needed.
// finally, data contains the actual data read in.
type config struct {
	structure   interface{}
	defaults    interface{}
	data        interface{}
	limit       restriction
	propagation []propagation
}
