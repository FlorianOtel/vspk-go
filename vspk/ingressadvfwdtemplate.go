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

// IngressAdvFwdTemplateIdentity represents the Identity of the object
var IngressAdvFwdTemplateIdentity = bambou.Identity{
	Name:     "ingressadvfwdtemplate",
	Category: "ingressadvfwdtemplates",
}

// IngressAdvFwdTemplatesList represents a list of IngressAdvFwdTemplates
type IngressAdvFwdTemplatesList []*IngressAdvFwdTemplate

// IngressAdvFwdTemplatesAncestor is the interface of an ancestor of a IngressAdvFwdTemplate must implement.
type IngressAdvFwdTemplatesAncestor interface {
	IngressAdvFwdTemplates(*bambou.FetchingInfo) (IngressAdvFwdTemplatesList, *bambou.Error)
	CreateIngressAdvFwdTemplates(*IngressAdvFwdTemplate) *bambou.Error
}

// IngressAdvFwdTemplate represents the model of a ingressadvfwdtemplate
type IngressAdvFwdTemplate struct {
	ID                     string `json:"ID,omitempty"`
	ParentID               string `json:"parentID,omitempty"`
	ParentType             string `json:"parentType,omitempty"`
	Owner                  string `json:"owner,omitempty"`
	Name                   string `json:"name,omitempty"`
	LastUpdatedBy          string `json:"lastUpdatedBy,omitempty"`
	Active                 bool   `json:"active"`
	Description            string `json:"description,omitempty"`
	EntityScope            string `json:"entityScope,omitempty"`
	PolicyState            string `json:"policyState,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	PriorityType           string `json:"priorityType,omitempty"`
	AssociatedLiveEntityID string `json:"associatedLiveEntityID,omitempty"`
	ExternalID             string `json:"externalID,omitempty"`
}

// NewIngressAdvFwdTemplate returns a new *IngressAdvFwdTemplate
func NewIngressAdvFwdTemplate() *IngressAdvFwdTemplate {

	return &IngressAdvFwdTemplate{}
}

// Identity returns the Identity of the object.
func (o *IngressAdvFwdTemplate) Identity() bambou.Identity {

	return IngressAdvFwdTemplateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *IngressAdvFwdTemplate) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *IngressAdvFwdTemplate) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the IngressAdvFwdTemplate from the server
func (o *IngressAdvFwdTemplate) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the IngressAdvFwdTemplate into the server
func (o *IngressAdvFwdTemplate) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the IngressAdvFwdTemplate from the server
func (o *IngressAdvFwdTemplate) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the IngressAdvFwdTemplate
func (o *IngressAdvFwdTemplate) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the IngressAdvFwdTemplate
func (o *IngressAdvFwdTemplate) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the IngressAdvFwdTemplate
func (o *IngressAdvFwdTemplate) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the IngressAdvFwdTemplate
func (o *IngressAdvFwdTemplate) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// IngressAdvFwdEntryTemplates retrieves the list of child IngressAdvFwdEntryTemplates of the IngressAdvFwdTemplate
func (o *IngressAdvFwdTemplate) IngressAdvFwdEntryTemplates(info *bambou.FetchingInfo) (IngressAdvFwdEntryTemplatesList, *bambou.Error) {

	var list IngressAdvFwdEntryTemplatesList
	err := bambou.CurrentSession().FetchChildren(o, IngressAdvFwdEntryTemplateIdentity, &list, info)
	return list, err
}

// CreateIngressAdvFwdEntryTemplate creates a new child IngressAdvFwdEntryTemplate under the IngressAdvFwdTemplate
func (o *IngressAdvFwdTemplate) CreateIngressAdvFwdEntryTemplate(child *IngressAdvFwdEntryTemplate) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Jobs retrieves the list of child Jobs of the IngressAdvFwdTemplate
func (o *IngressAdvFwdTemplate) Jobs(info *bambou.FetchingInfo) (JobsList, *bambou.Error) {

	var list JobsList
	err := bambou.CurrentSession().FetchChildren(o, JobIdentity, &list, info)
	return list, err
}

// CreateJob creates a new child Job under the IngressAdvFwdTemplate
func (o *IngressAdvFwdTemplate) CreateJob(child *Job) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
