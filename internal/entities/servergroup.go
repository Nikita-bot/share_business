package entities

import (
	uuid "github.com/satori/go.uuid"
)

type ServerGroup struct {
	ID             uuid.UUID
	Name           string
	Description    string
	OrganizationID uuid.UUID
	IsDefault      bool
	Deleted        bool
}

type ServerGroupCreation struct {
	Name           string
	Description    string
	OrganizationID uuid.UUID
}

type ServerGroupUpdate struct {
	Name        *string
	Description *string
}

type ServerGroupFilter struct {
	Pagination
	OrganizationID *uuid.UUID
}
