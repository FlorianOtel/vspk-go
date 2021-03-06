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

import (
	"crypto/tls"
	"fmt"
	"strings"

	"github.com/FlorianOtel/go-bambou/bambou"
)

var (
	urlpostfix string
)

// Returns a new Session -- authentication using username + password
func NewSession(username, password, organization, url string) (*bambou.Session, *Me) {

	root := NewMe()
	url += urlpostfix

	session := bambou.NewSession(username, password, organization, url, root)

	return session, root
}

// Returns a new Session -- authentication using X509 certificate.
func NewX509Session(cert *tls.Certificate, url string) (*bambou.Session, *Me) {

	root := NewMe()
	url += urlpostfix

	session := bambou.NewX509Session(cert, url, root)

	return session, root
}

func init() {

	urlpostfix = "/" + SDKAPIPrefix + "/v" + strings.Replace(fmt.Sprintf("%.1f", SDKAPIVersion), ".", "_", 100)
}
