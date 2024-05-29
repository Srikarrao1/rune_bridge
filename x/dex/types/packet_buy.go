package types

// ValidateBasic is used for validating the packet
func (p BuyPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p BuyPacketData) GetBytes() ([]byte, error) {
	var modulePacket DexPacketData

	modulePacket.Packet = &DexPacketData_BuyPacket{&p}

	return modulePacket.Marshal()
}
