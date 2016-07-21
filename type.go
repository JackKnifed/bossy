package bossy

type propagation struct {
	parentNode []string
	childNode  []string
}

type restriction struct {
	address     []string
	restriction func(interface{}) bool
}

type config struct {
	structure   interface{}
	defaults    interface{}
	data        interface{}
	limits      []restriction
	propagation []propagation
}
