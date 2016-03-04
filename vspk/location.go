/*
  Copyright (c) 2015, Alcatel-Lucent Inc
  All rights reserved.

  Redistribution and use in source and binary forms, with or without
  modification, are permitted provided that the following conditions are met:
      * Redistributions of source code must retain the above copyright
        notice, this list of conditions and the following disclaimer.
      * Redistributions in binary form must reproduce the above copyright
        notice, this list of conditions and the following disclaimer in the
        documentation and/or other materials provided with the distribution.
      * Neither the name of the copyright holder nor the names of its contributors
        may be used to endorse or promote products derived from this software without
        specific prior written permission.

  THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
  ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
  WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
  DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY
  DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
  (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
  LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
  ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
  (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
  SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// LocationIdentity represents the Identity of the object
var LocationIdentity = bambou.Identity{
	Name:     "location",
	Category: "locations",
}

// LocationsList represents a list of Locations
type LocationsList []*Location

// LocationsAncestor is the interface of an ancestor of a Location must implement.
type LocationsAncestor interface {
	Locations(*bambou.FetchingInfo) (LocationsList, *bambou.Error)
	CreateLocations(*Location) *bambou.Error
}

// Location represents the model of a location
type Location struct {
	ID            string  `json:"ID,omitempty"`
	ParentID      string  `json:"parentID,omitempty"`
	ParentType    string  `json:"parentType,omitempty"`
	Owner         string  `json:"owner,omitempty"`
	Address       string  `json:"address,omitempty"`
	Country       string  `json:"country,omitempty"`
	EntityScope   string  `json:"entityScope,omitempty"`
	ExternalID    string  `json:"externalID,omitempty"`
	IgnoreGeocode bool    `json:"ignoreGeocode"`
	LastUpdatedBy string  `json:"lastUpdatedBy,omitempty"`
	Latitude      float64 `json:"latitude,omitempty"`
	Locality      string  `json:"locality,omitempty"`
	Longitude     float64 `json:"longitude,omitempty"`
	State         string  `json:"state,omitempty"`
	TimeZoneID    string  `json:"timeZoneID,omitempty"`
}

// NewLocation returns a new *Location
func NewLocation() *Location {

	return &Location{
		TimeZoneID: "UTC",
	}
}

// Identity returns the Identity of the object.
func (o *Location) Identity() bambou.Identity {

	return LocationIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Location) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Location) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Location from the server
func (o *Location) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Location into the server
func (o *Location) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Location from the server
func (o *Location) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the Location
func (o *Location) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the Location
func (o *Location) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the Location
func (o *Location) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the Location
func (o *Location) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
