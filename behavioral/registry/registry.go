package registry

import (
	"fmt"

	"github.com/google/uuid"
)

// IDNotFound is an error thrown when an id is not found in a Registry
type IDNotFound struct {
	id string
}

func (e IDNotFound) Error() string {
	return fmt.Sprintf("ID not found in registry: %s", e.id)
}

// NameAlreadyRegistered is an error thrown when a duplicate id is found durring registration
type NameAlreadyRegistered struct {
	id string
}

func (e NameAlreadyRegistered) Error() string {
	return fmt.Sprintf("Identifier was already in registry: %s", e.id)
}

// Registry acts a generic data registry
// Objects are registered with the underlying struct, which returns an ID string
// The ID string can then be used to retreive the object, or to remove it from the registry
type Registry interface {

	// Deregister removes an item identified by the passed id
	Deregister(id string) error

	// Get retreives an item identified by the passed id from the registry
	Get(id string) (interface{}, error)

	// Prune removes nil entries from the registry
	Prune()

	// Register adds an item to the registry and returns an identifying string
	Register(data interface{}) string

	// RegisterName adds an item to the registry under the suppplied identifier and returns an identifying string
	// Returns a NameAlreadyRegistered error if the identifier has already been registered
	RegisterName(data interface{}, identifier string) error
}

type registry struct {
	registry map[string]interface{}
}

func (reg *registry) Deregister(id string) error {
	if _, ok := reg.registry[id]; ok {
		delete(reg.registry, id)
		return nil
	}
	return IDNotFound{id}
}

func (reg *registry) Get(id string) (interface{}, error) {
	if data, ok := reg.registry[id]; ok {
		return data, nil
	}
	return nil, IDNotFound{id}
}

func (reg *registry) Prune() {
	for key, value := range reg.registry {
		if value == nil {
			delete(reg.registry, key)
		}
	}
}

func (reg *registry) Register(data interface{}) string {
	// Just in case the impossible happens and there is an id colision
	for {
		newUUID, _ := uuid.NewRandom()
		id := newUUID.String()
		if _, ok := reg.registry[id]; !ok {
			reg.registry[id] = data
			return id
		}
	}
}

func (reg *registry) RegisterName(data interface{}, identifier string) error {
	if _, ok := reg.registry[identifier]; !ok {
		reg.registry[identifier] = data
		return nil
	}
	return NameAlreadyRegistered{identifier}
}

// New returns a new Registry interface
func New() Registry {
	return &registry{}
}
