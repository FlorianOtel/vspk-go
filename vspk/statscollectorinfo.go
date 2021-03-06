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

import "github.com/FlorianOtel/go-bambou/bambou"

// StatsCollectorInfoIdentity represents the Identity of the object
var StatsCollectorInfoIdentity = bambou.Identity{
	Name:     "statisticscollector",
	Category: "statisticscollector",
}

// StatsCollectorInfosList represents a list of StatsCollectorInfos
type StatsCollectorInfosList []*StatsCollectorInfo

// StatsCollectorInfosAncestor is the interface of an ancestor of a StatsCollectorInfo must implement.
type StatsCollectorInfosAncestor interface {
	StatsCollectorInfos(*bambou.FetchingInfo) (StatsCollectorInfosList, *bambou.Error)
	CreateStatsCollectorInfos(*StatsCollectorInfo) *bambou.Error
}

// StatsCollectorInfo represents the model of a statisticscollector
type StatsCollectorInfo struct {
	ID            string `json:"ID,omitempty"`
	ParentID      string `json:"parentID,omitempty"`
	ParentType    string `json:"parentType,omitempty"`
	Owner         string `json:"owner,omitempty"`
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	AddressType   string `json:"addressType,omitempty"`
	EntityScope   string `json:"entityScope,omitempty"`
	Port          string `json:"port,omitempty"`
	IpAddress     string `json:"ipAddress,omitempty"`
	ProtoBufPort  string `json:"protoBufPort,omitempty"`
	ExternalID    string `json:"externalID,omitempty"`
}

// NewStatsCollectorInfo returns a new *StatsCollectorInfo
func NewStatsCollectorInfo() *StatsCollectorInfo {

	return &StatsCollectorInfo{}
}

// Identity returns the Identity of the object.
func (o *StatsCollectorInfo) Identity() bambou.Identity {

	return StatsCollectorInfoIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *StatsCollectorInfo) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *StatsCollectorInfo) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the StatsCollectorInfo from the server
func (o *StatsCollectorInfo) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the StatsCollectorInfo into the server
func (o *StatsCollectorInfo) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the StatsCollectorInfo from the server
func (o *StatsCollectorInfo) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the StatsCollectorInfo
func (o *StatsCollectorInfo) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the StatsCollectorInfo
func (o *StatsCollectorInfo) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the StatsCollectorInfo
func (o *StatsCollectorInfo) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the StatsCollectorInfo
func (o *StatsCollectorInfo) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
