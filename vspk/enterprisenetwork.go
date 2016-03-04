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

// EnterpriseNetworkIdentity represents the Identity of the object
var EnterpriseNetworkIdentity = bambou.Identity{
	Name:     "enterprisenetwork",
	Category: "enterprisenetworks",
}

// EnterpriseNetworksList represents a list of EnterpriseNetworks
type EnterpriseNetworksList []*EnterpriseNetwork

// EnterpriseNetworksAncestor is the interface of an ancestor of a EnterpriseNetwork must implement.
type EnterpriseNetworksAncestor interface {
	EnterpriseNetworks(*bambou.FetchingInfo) (EnterpriseNetworksList, *bambou.Error)
	CreateEnterpriseNetworks(*EnterpriseNetwork) *bambou.Error
}

// EnterpriseNetwork represents the model of a enterprisenetwork
type EnterpriseNetwork struct {
	ID            string `json:"ID,omitempty"`
	ParentID      string `json:"parentID,omitempty"`
	ParentType    string `json:"parentType,omitempty"`
	Owner         string `json:"owner,omitempty"`
	IPType        string `json:"IPType,omitempty"`
	Address       string `json:"address,omitempty"`
	EntityScope   string `json:"entityScope,omitempty"`
	ExternalID    string `json:"externalID,omitempty"`
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	Name          string `json:"name,omitempty"`
	Netmask       string `json:"netmask,omitempty"`
}

// NewEnterpriseNetwork returns a new *EnterpriseNetwork
func NewEnterpriseNetwork() *EnterpriseNetwork {

	return &EnterpriseNetwork{
		IPType: "IPV4",
	}
}

// Identity returns the Identity of the object.
func (o *EnterpriseNetwork) Identity() bambou.Identity {

	return EnterpriseNetworkIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EnterpriseNetwork) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EnterpriseNetwork) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the EnterpriseNetwork from the server
func (o *EnterpriseNetwork) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the EnterpriseNetwork into the server
func (o *EnterpriseNetwork) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the EnterpriseNetwork from the server
func (o *EnterpriseNetwork) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// EventLogs retrieves the list of child EventLogs of the EnterpriseNetwork
func (o *EnterpriseNetwork) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}

// CreateEventLog creates a new child EventLog under the EnterpriseNetwork
func (o *EnterpriseNetwork) CreateEventLog(child *EventLog) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the EnterpriseNetwork
func (o *EnterpriseNetwork) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the EnterpriseNetwork
func (o *EnterpriseNetwork) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the EnterpriseNetwork
func (o *EnterpriseNetwork) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the EnterpriseNetwork
func (o *EnterpriseNetwork) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// NetworkMacroGroups retrieves the list of child NetworkMacroGroups of the EnterpriseNetwork
func (o *EnterpriseNetwork) NetworkMacroGroups(info *bambou.FetchingInfo) (NetworkMacroGroupsList, *bambou.Error) {

	var list NetworkMacroGroupsList
	err := bambou.CurrentSession().FetchChildren(o, NetworkMacroGroupIdentity, &list, info)
	return list, err
}

// AssignNetworkMacroGroups assigns the list of NetworkMacroGroups to the EnterpriseNetwork
func (o *EnterpriseNetwork) AssignNetworkMacroGroups(children NetworkMacroGroupsList) *bambou.Error {

	list := []bambou.Identifiable{}
	for _, c := range children {
		list = append(list, c)
	}

	return bambou.CurrentSession().AssignChildren(o, list, NetworkMacroGroupIdentity)
}
