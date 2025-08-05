package transport

import (
	"log"
	"net"
)

type UDPReceiver struct {
	Port    int
	GroupIP string
	Encoder Encoder
	Handler MessageHandler
}

func (r *UDPReceiver) Start() {
	addr := &net.UDPAddr{
		IP:   net.ParseIP(r.GroupIP),
		Port: r.Port,
	}
	conn, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		log.Fatalf("Error joining multicast group: %v", err)
	}
	err = conn.SetReadBuffer(1024)
	if err != nil {
		panic(err)
	}
	log.Printf("Listening on multicast group %s:%d...\n", r.GroupIP, r.Port)
	buf := make([]byte, 1024)
	for {
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println("Read error:", err)
			continue
		}
		msg, err := r.Encoder.Decode(buf[:n])
		if err != nil {
			log.Println("Decode error:", err)
			continue
		}
		r.Handler.Handle(msg)
	}
}
