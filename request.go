package tpshp

import (
	"encoding/json"
	"fmt"
)

// Request describes the overall request layout for TP-Link Smart-Home Protocol commands
type Request struct {
	commands  map[string]map[string]interface{}
	responses map[string]map[string]interface{}
}

// NewRequest returns a new request
func NewRequest() *Request {
	return &Request{
		commands:  make(map[string]map[string]interface{}),
		responses: make(map[string]map[string]interface{}),
	}
}

// Namespaces returns all namespaces configured in the request
func (r *Request) Namespaces() []string {
	var ns []string
	for key := range r.commands {
		ns = append(ns, key)
	}

	return ns
}

// ResponseExpected returns true if the requests expects a response from the device
func (r *Request) ResponseExpected() bool {
	return len(r.responses) > 0
}

// Namespace returns a map with all commands for the given namespace
func (r *Request) Namespace(name string) map[string]interface{} {
	return r.commands[name]
}

// Responses returns all responses for the request. Responses are grouped by
// namespace and command name
func (r *Request) Responses() map[string]map[string]interface{} {
	return r.responses
}

// AddCommand adds a command to the request
func (r *Request) AddCommand(namespace string, command string, payload interface{}, response interface{}) *Request {
	if r.commands[namespace] == nil {
		r.commands[namespace] = make(map[string]interface{})
	}

	if payload != nil {
		if r.responses[namespace] == nil {
			r.responses[namespace] = make(map[string]interface{})
		}

		r.responses[namespace][command] = response
	}

	r.commands[namespace][command] = payload

	return r
}

// MarshalJSON implements json.Marshaler and returns the JSON representation of the
// request
func (r *Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.commands)
}

// UnmarshalJSON implements json.Unmarshaler and parses all command responses
func (r *Request) UnmarshalJSON(blob []byte) error {
	var responses map[string]map[string]interface{}

	if err := json.Unmarshal(blob, &responses); err != nil {
		return err
	}

	for namespace, commands := range r.responses {
		for cmd, response := range commands {
			if responses[namespace] == nil || responses[namespace][cmd] == nil {
				return fmt.Errorf("expected response for %s.%s but got nil", namespace, cmd)
			}

			// fast-path, if the user passed *interface{} for the response we can set it directly
			if v, ok := response.(*interface{}); ok {
				*v = responses[namespace][cmd]
				continue
			}

			// otherwise convert it back to json and marshal it into the actual response struct
			blob, err := json.Marshal(responses[namespace][cmd])
			if err != nil {
				return err
			}

			if err := json.Unmarshal(blob, response); err != nil {
				return err
			}
		}
	}

	return nil
}
