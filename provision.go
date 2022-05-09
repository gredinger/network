package network

type Switch struct {
	Hostname string

	runningConfig string //current config
	startupConfig string //init config at load
}

type Connection struct {
}

//applyConfig will write changes to the switch
func (s Switch) applyConfig() {

}

//welp, this is gonna be fun
func (s Switch) grabConfig() {
	//lets start with init config, and see if anything changed?
	//ensure enable access (for cisco)

}
