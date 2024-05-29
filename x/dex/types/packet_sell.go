package types

// ValidateBasic is used for validating the packet
func (p SellPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p SellPacketData) GetBytes() ([]byte, error) {
	var modulePacket DexPacketData

	modulePacket.Packet = &DexPacketData_SellPacket{&p}

	return modulePacket.Marshal()
}
