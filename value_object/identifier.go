package value_object

import (
	"github.com/oklog/ulid/v2"
)

type Identifier ulid.ULID

func NewIdentifier() Identifier {
	return Identifier(ulid.Make())
}

func NewIdentifierFrom(rawId string) (id Identifier, err error) {
	parsedUlid, err := ulid.ParseStrict(rawId)
	return Identifier(parsedUlid), err
}

func (id Identifier) Equals(otherId Identifier) bool {
	return otherId.String() == id.String()
}

func (id Identifier) String() string {
	return ulid.ULID(id).String()
}

func (id Identifier) MarshalText() ([]byte, error) {
	return []byte(id.String()), nil
}

func (id *Identifier) UnmarshalText(text []byte) error {
	parsedId, err := NewIdentifierFrom(string(text))
	if err != nil {
		return err
	}

	*id = parsedId
	return nil
}
