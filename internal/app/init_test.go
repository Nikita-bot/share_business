// Code generated by mtgroup-generator.
package app

import (
	"github.com/go-openapi/strfmt"
	"time"

	"github.com/google/uuid"
)

// Make sure not to overwrite this file after you generated it because all your edits would be lost!

var (
	isolatedEntityID = uuid.New().String()
	profileID        = uuid.New().String()
	profile          = Profile{
		ID:    profileID,
		Authn: true,
		Authz: Authz{
			User:    true,
			Admin:   true,
			Manager: true,
		},
		IsolatedEntityID: isolatedEntityID,
	}
	profileUnAuth = Profile{
		ID:    profileID,
		Authn: false,
		Authz: Authz{
			User:    false,
			Admin:   false,
			Manager: false,
		},
		IsolatedEntityID: isolatedEntityID,
	}
	testUser1 = &User{
		Active:     false,
		CreatedAt:  mustParseTime("1985-04-17T19:22:19.032Z"),
		FirebaseId: "consequuntur",
		ID:         uuid.New().String(),
		ModifiedAt: mustParseTime("2010-02-06T12:05:03.373Z"),
	}
	testUser2 = &User{
		Active:     false,
		CreatedAt:  mustParseTime("1991-04-15T17:43:24.141Z"),
		FirebaseId: "quaerat",
		ID:         uuid.New().String(),
		ModifiedAt: mustParseTime("2017-09-24T16:52:40.598Z"),
	}
	testUsers       = []*User{testUser1, testUser2}
	testWashServer1 = &WashServer{
		CreatedAt:    mustParseTime("1919-08-25T20:17:11.920Z"),
		ID:           uuid.New().String(),
		Key:          "dolorem",
		LastUpdateAt: mustParseTime("1914-04-26T05:50:20.619Z"),
		ModifiedAt:   mustParseTime("1901-10-10T08:46:05.056Z"),
		Name:         "et",
	}
	testWashServer2 = &WashServer{
		CreatedAt:    mustParseTime("1928-07-27T22:12:16.045Z"),
		ID:           uuid.New().String(),
		Key:          "iure",
		LastUpdateAt: mustParseTime("1928-11-09T20:11:53.534Z"),
		ModifiedAt:   mustParseTime("1919-07-17T09:06:05.219Z"),
		Name:         "exercitationem",
	}
	testWashServers = []*WashServer{testWashServer1, testWashServer2}
	listParams      = &ListParams{
		Offset: 0,
		Limit:  5,
	}
)

func mustParseTime(t string) *time.Time {
	date, err := time.Parse(time.RFC3339, t)
	if err != nil {
		date, _ = time.Parse(strfmt.RFC3339FullDate, t)
	}
	return &date
}
