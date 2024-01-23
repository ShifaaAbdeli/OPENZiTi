/*
	Copyright NetFoundry Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package network

import (
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

// A simple test to check for failure of alignment on atomic operations for 64 bit variables in a struct
func Test64BitAlignment(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("One of the variables that was tested is not properly 64-bit aligned.")
		}
	}()

	link := Link{}

	atomic.LoadInt64(&link.SrcLatency)
	atomic.LoadInt64(&link.DstLatency)
	atomic.LoadInt64(&link.Cost)
}

func TestLifecycle(t *testing.T) {
	linkController := newLinkController(nil)

	r0 := NewRouter("r0", "", "", 0, true)
	r1 := NewRouter("r1", "", "", 0, true)
	l0 := &Link{
		Id:    "l0",
		Src:   r0,
		DstId: r1.Id,
	}
	l0.Dst.Store(r1)

	linkController.add(l0)
	assert.True(t, linkController.has(l0))

	links := r0.routerLinks.GetLinks()
	assert.Equal(t, 1, len(links))
	assert.Equal(t, l0, links[0])

	links = r1.routerLinks.GetLinks()
	assert.Equal(t, 1, len(links))
	assert.Equal(t, l0, links[0])

	linkController.remove(l0)
	assert.False(t, linkController.has(l0))

	links = r0.routerLinks.GetLinks()
	assert.Equal(t, 0, len(links))

	links = r1.routerLinks.GetLinks()
	assert.Equal(t, 0, len(links))
}

func TestNeighbors(t *testing.T) {
	linkController := newLinkController(nil)

	r0 := newRouterForTest("r0", "", nil, nil, 0, true)
	r1 := newRouterForTest("r1", "", nil, nil, 0, true)
	l0 := newTestLink("l0", r0, r1)
	l0.addState(newLinkState(Connected))
	linkController.add(l0)

	neighbors := linkController.connectedNeighborsOfRouter(r0)
	assert.Equal(t, 1, len(neighbors))
	assert.Equal(t, r1, neighbors[0])
}

func newTestLink(id string, src, dst *Router) *Link {
	l := newLink(id, "tls", "tcp:localhost:1234", 0)
	l.Src = src
	l.DstId = dst.Id
	l.Dst.Store(dst)
	src.Connected.Store(true)
	dst.Connected.Store(true)
	return l
}
