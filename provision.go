package network

import (
	"regexp"
	"strings"
)

type Switch struct {
	Hostname string
	Username string
	Password string
	Ports    []Port
	SSHConfig *ssh.ClientConfig
}

func (s *Switch) addEntry(name, mac, id string) {
	for i, x := range s.Ports {
		if x.Name == name {
			s.Ports[i].Entries = append(x.Entries, PortEntry{id, mac})
		}
	}
}

func (s *Switch) addPort(name string) {
	for _, x := range s.Ports {
		if x.Name == name {
			return
		}
	}
	s.Ports = append(s.Ports, Port{
		Name:    name,
		Entries: []PortEntry{},
	})
}

type PortEntry struct {
	VLAN       string
	MACAddress string
}

type Port struct {
	Name    string
	Entries []PortEntry
}

type Connection struct {
}

var cisID = `(?m:^\s{1,5}\d{1,5})`
var cisMAC = `.{4}\..{4}\..{4}`
var cisInt = `(?m:\b\w{0,4}\/.*$)`

//parseMAT will return the Mac address table as  list of Ports with their entries per port
func (sw *Switch) ParseMAT(in string) {
	// regex id (start of string, whitespace)
	idr, _ := regexp.Compile(cisID)
	// regex string mac address
	addr, _ := regexp.Compile(cisMAC)
	// regex string port
	portR, _ := regexp.Compile(cisInt)

	vlan := idr.FindAllString(in, -1)
	add := addr.FindAllString(in, -1)[20:]
	port := portR.FindAllString(in, -1)
	if len(vlan) == len(add) && len(add) == len(port) {
		for _, x := range port {
			sw.addPort(x)
		}
		for i := range vlan {
			sw.addEntry(port[i], add[i], strings.Trim(vlan[i], " "))
		}
	}
}

func (s *Switch) LoadConfiguration() {
	s.SSHConfig = &ssh.ClientConfig{
		User: s.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(s.Password)
		},
		HostKeyCallback: ClientConfig.InsecureIgnoreHostKey,
		Timeout: (time.Seconds * 60) // 1 min timeout
	}
}

func (s *Switch) Execute() {
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:22", s.Hostname), s.SSHConfig)
	if err != nil {
		log.Fatal("Failed to connect to ssh")
	}
	defer client.Close()

	session, err := conn.NewSession()

	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	
}