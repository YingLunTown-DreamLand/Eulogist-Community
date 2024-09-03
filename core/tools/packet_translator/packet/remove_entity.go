package packet

import (
	neteasePacket "Eulogist/core/minecraft/protocol/packet"

	standardPacket "github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type RemoveEntity struct{}

func (pk *RemoveEntity) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.RemoveEntity{}
	input := standard.(*standardPacket.RemoveEntity)

	p.EntityNetworkID = uint32(input.EntityNetworkID)

	return &p
}

func (pk *RemoveEntity) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.RemoveEntity{}
	input := netease.(*neteasePacket.RemoveEntity)

	p.EntityNetworkID = uint64(input.EntityNetworkID)

	return &p
}