package network

type Switch struct {
	Hostname string
	Username string
	Password string
}

type PortEntry struct {
	VLAN int
	MACAddress string
}

type Port struct {
	Entries PortEntry[]
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

	// show mac address-table


}

//parseMAT will return a list of Ports with their entries
func parseMAT(in string) []Port {
	
}