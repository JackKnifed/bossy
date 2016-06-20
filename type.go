package bossy

type config interface {
	LoadJSON() Error
	Export()
}

func (c config) Export() {

}
