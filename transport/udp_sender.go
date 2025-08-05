package transport

import (
	"net"
	"udp_sim/core"
)

type UDPSender struct {
	Addr    string // e.g. "224.0.0.1:9999"
	Encoder Encoder
}

func (s *UDPSender) Send(msg core.PoseMessage) error {
	rAddr, err := net.ResolveUDPAddr("udp", s.Addr)
	if err != nil {
		return err
	}
	conn, err := net.DialUDP("udp", nil, rAddr)
	if err != nil {
		return err
	}
	defer conn.Close()
	data, err := s.Encoder.Encode(msg)
	if err != nil {
		return err
	}
	_, err = conn.Write(data)
	return err
}
