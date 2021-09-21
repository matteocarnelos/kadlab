package kademlia

import (
	"fmt"
	"net"
	"time"
)

type Kademlia struct {
	rpc map[KademliaID]chan string
	net Network
	RT *RoutingTable
}

func NewKademlia(me Contact) *Kademlia {
	return &Kademlia{
		rpc: make(map[KademliaID]chan string),
		RT: NewRoutingTable(me),
	}
}

func (k *Kademlia) StartListen(ip string, port int) {
	k.net.ListenIP = net.ParseIP(ip)
	k.net.ListenPort = port
	go k.net.listen(k)
}

func (k *Kademlia) handleRPC(id *KademliaID, cmd string, args []string) string {
	switch cmd {
	case "PING":
		return "PINGREPLY"
	case "PINGREPLY":
		k.rpc[*id] <- "PINGREPLY"
	}
	return ""
}

func (k *Kademlia) LookupContact(target *Contact) {
	for _, c := range k.RT.FindClosestContacts(target.ID, 3) {
		k.net.SendFindContactMessage(target, &c)
	}
}

func (k *Kademlia) LookupData(hash string) {
	// TODO (M2.b)
}

func (k *Kademlia) Store(data []byte) {
	// TODO (M2.a)
}
