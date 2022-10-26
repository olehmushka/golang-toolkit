package randomtools

import "github.com/google/uuid"

func UUID() uuid.UUID {
	return uuid.New()
}

func UUIDString() string {
	return uuid.NewString()
}
