package tpshp

// Command is a command sent over TP-Link's Smart-Home Protocol
type Command interface {
	// Prepare adds the command request to the map
	Prepare(m map[string]interface{}) error

	// Response should return a receiver for the commands reponse
	Response() interface{}
}
