package net

import (
	"bytes"
	"enet"
	"errors"
	"github.com/tedsta/fission/core"
)

const (
	incomingBandwidth = 0 // unlimited for now
	outgoingBandwidth = 0
)

type NetType uint8
type NetId uint
type PacketHandler func(*core.InPacket)

const (
	None NetType = iota
	Server
	Client
)

type Peer struct {
	Id        NetId
	IpAddress string
	peer      *enet.Peer
}

type Connection struct {
	netType           NetType    // Net type
	host              *enet.Host // The network host
	peer              *Peer      // If it's a client connection, the peer
	peers             []*Peer    // The list of clients if I'm a server
	rcvScene          bool       // Whether or not the scene is being received.
	rcvObjectCount    int        // Number of objects received if scene is being received.
	totalObjectsToRcv int        // Number of objects to receive if the scene is being received.
	nextId            NetId      // The Id for the next peer
	onConnect         func(NetId)
	handlers          []PacketHandler
	onDisconnect      func(NetId)
}

func NewFakeConnection() *Connection {
	return &Connection{netType: None}
}

func NewServer(port int, onConnect, onDisconnect func(NetId)) (*Connection, error) {

	c := &Connection{onConnect: onConnect, onDisconnect: onDisconnect}

	serverAddress := enet.HostAddr(port)

	var err error
	c.host, err = enet.New(serverAddress, 32, 2, incomingBandwidth, outgoingBandwidth)
	if err != nil {
		return nil, err
	}

	c.netType = Server

	return c, nil
}

func NewClient(ip string, port int, onConnect, onDisconnect func(NetId)) (*Connection, error) {

	c := &Connection{onConnect: onConnect, onDisconnect: onDisconnect}
	c.peer = &Peer{}

	var serverAddress *enet.Addr
	var err error
	serverAddress, err = enet.Address(ip, port)
	if err != nil {
		return nil, err
	}

	// Create the enet host
	//mHost = enet_host_create(NULL, 1, 2, 57600 / 8, 14400 / 8);
	c.host, err = enet.New(nil, 1, 2, incomingBandwidth, outgoingBandwidth)
	if err != nil {
		return nil, err
	}

	c.peer.peer, err = c.host.Connect(serverAddress, 2, 0)
	if err != nil {
		return nil, err
	}

	// Wait up to 5 seconds for the connection attempt to succeed.
	event := c.host.Wait(5000)
	switch event.(type) {
	case *enet.ConnectEvent:
		// Successful connection
	default:
		c.peer.peer.Reset()
		return nil, errors.New("Connection to " + serverAddress.String() + " timed out")
	}

	// Wait 5 seconds for network ID
	event = c.host.Wait(5000)
	switch event.(type) {
	case *enet.ReceiveEvent:
		e := event.(*enet.ReceiveEvent)
		packet := core.NewInPacket(bytes.NewBufferString(e.Data))
		packet.Read(&c.peer.Id) // Get packet ID

	default:
		c.peer.peer.Reset()
		return nil, errors.New("Connection to " + ip + " failed")
	}

	c.netType = Client

	return c, nil
}

func (c *Connection) Update() {
	if c.netType == None {
		return
	}

	for i := 0; i < 100; i++ {
		// Check for events, but don't wait
		event := c.host.Wait(0)

		switch event.(type) {
		case *enet.ConnectEvent:
			e := event.(*enet.ConnectEvent)
			ip := e.Peer.Addr().String()

			// Add the new peer
			c.nextId++
			peer := &Peer{c.nextId, ip, e.Peer}
			c.peers = append(c.peers, peer)

			e.Peer.SetData(enet.DataPointer(peer))

			// Send the client its ID
			idPacket := core.NewOutPacket(nil)
			idPacket.Write(peer.Id)
			peer.peer.Send(0, idPacket.String(), enet.RELIABLE)
			c.host.Flush()
			idPacket.Clear()

			if c.onConnect != nil {
				c.onConnect(peer.Id)
			}

		case *enet.ReceiveEvent:
			e := event.(*enet.ReceiveEvent)

			packet := core.NewInPacket(bytes.NewBufferString(e.Data))

			var hndId int
			packet.Read(&hndId)

			// Handle the packet if possible
			if len(c.handlers) > hndId {
				c.handlers[hndId](packet)
			}

		case *enet.DisconnectEvent:
			e := event.(*enet.DisconnectEvent)
			peer := (*Peer)(e.Peer.Data())

			if peer != nil {
				if c.onDisconnect != nil {
					c.onDisconnect(peer.Id)
				}

				peer.peer = nil
				c.removePeer(peer.Id)
				e.Peer.SetData(nil)
			}

		default:
		}
	}
}

func (c *Connection) Send(p *core.OutPacket, handler int, netId, excludeId NetId, rel bool) {
	if c.netType == None {
		return
	}

	// Create the enet packet
	var flags enet.Flag
	if rel {
		flags |= enet.RELIABLE
	}

	// Handler ID packet
	handPack := core.NewOutPacket(nil)
	handPack.Write(handler)
	handPack.Append(p)

	if c.netType == Client {
		c.peer.peer.Send(0, handPack.String()+p.String(), flags)
	} else if netId > 0 {
		// It's a server and the client is specified. Tell only that client!
		c.findPeer(netId).peer.Send(0, handPack.String()+p.String(), flags)
	} else {
		// It's a server and the peer is unspecified
		if excludeId > 0 {
			for _, peer := range c.peers {
				if peer.Id != excludeId {
					peer.peer.Send(0, handPack.String()+p.String(), flags)
				}
			}
		} else {
			c.host.SendToAll(0, handPack.String()+p.String(), flags)
		}
	}
}

// TODO?
/*func (c *Connection) SendEntity(e *core.Entity, netId, excludeId NetId) {
    packet := core.NewOutPacket(nil)
    // packet.Write() // Needs packet ID
    object.Serialize(packet)

    c.Send(packet, connectorID, excludeID, true)
}*/

func (c *Connection) RegisterHandlerAuto(ph PacketHandler) int {
	c.handlers = append(c.handlers, ph)
	return len(c.handlers) - 1
}

func (c *Connection) RegisterHandler(hndId int, ph PacketHandler) {
	if hndId >= len(c.handlers) {
		newHnds := make([]PacketHandler, hndId+1)
		copy(newHnds, c.handlers)
		c.handlers = newHnds
	}

	c.handlers[hndId] = ph
}

func (c *Connection) findPeer(id NetId) *Peer {
	for _, p := range c.peers {
		if p.Id == id {
			return p
		}
	}
	return nil
}

func (c *Connection) removePeer(id NetId) {
	for i, p := range c.peers {
		if p.Id == id {
			c.peers[i] = c.peers[len(c.peers)-1]
			c.peers = c.peers[:len(c.peers)-1]
			return
		}
	}
}

func (c *Connection) Type() NetType {
	return c.netType
}

func (c *Connection) NetId() NetId {
	if c.peer != nil {
		return c.peer.Id
	}
	return 0
}
