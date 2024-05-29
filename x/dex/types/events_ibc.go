package types

// IBC events
const (
	EventTypeTimeout          = "timeout"
	EventTypeCreatePairPacket = "createPair_packet"
	EventTypeSellPacket       = "sell_packet"
	EventTypeBuyPacket        = "buy_packet"
	// this line is used by starport scaffolding # ibc/packet/event

	AttributeKeyAckSuccess = "success"
	AttributeKeyAck        = "acknowledgement"
	AttributeKeyAckError   = "error"
)
