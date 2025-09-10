package tcp_packet

import (
	"testing"
)

func TestSerializePacket(t *testing.T) {
	pkt := PacketPressed{
		PacketBase: PacketBase{
			PacketType: PacketTypePressed,
		},
		InputData: 1,
	}

	data, err := SerializePacket(pkt)
	if err != nil {
		t.Fatalf("failed to serialize packet: %v", err)
	}

	t.Logf("serialized data: %v\n", data)
}

func TestDeserializePacket(t *testing.T) {
	pkt := PacketPressed{
		PacketBase: PacketBase{
			PacketType: PacketTypePressed,
		},
		InputData: 1,
	}

	data, err := SerializePacket(pkt)
	if err != nil {
		t.Fatalf("failed to serialize packet: %v", err)
	}

	deserializedPktIntf, err := DeserializePacket(data)
	if err != nil {
		t.Fatalf("failed to deserialize packet: %v", err)
	}

	if deserializedPkt := deserializedPktIntf.(PacketPressed); deserializedPkt.InputData != pkt.InputData {
		t.Errorf("expected InputData %d, got %d", pkt.InputData, deserializedPkt.InputData)
	} else {
		t.Logf("deserialized packet: %+v", deserializedPkt)
	}
}
