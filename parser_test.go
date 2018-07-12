package sdptransform

import (
	"fmt"
	"testing"
)

var sdpStr = `v=0
o=- 20518 0 IN IP4 203.0.113.1
s= 
t=0 0
c=IN IP4 203.0.113.1
a=ice-ufrag:F7gI
a=ice-pwd:x9cml/YzichV2+XlhiMu8g
a=fingerprint:sha-1 42:89:c5:c6:55:9d:6e:c8:e8:83:55:2a:39:f9:b6:eb:e9:a3:a9:e7
m=audio 54400 RTP/SAVPF 0 96
a=rtpmap:0 PCMU/8000
a=rtpmap:96 opus/48000/2
a=ptime:20
a=sendrecv
a=candidate:0 1 UDP 2113667327 203.0.113.1 54400 typ host
a=candidate:1 2 UDP 2113667326 203.0.113.1 54401 typ host
m=video 55400 RTP/SAVPF 97 98
a=rtpmap:97 H264/90000
a=fmtp:97 profile-level-id=4d0028;packetization-mode=1
a=rtpmap:98 VP8/90000
a=sendrecv
a=candidate:0 1 UDP 2113667327 203.0.113.1 55400 typ host
a=candidate:1 2 UDP 2113667326 203.0.113.1 55401 typ host
`

var simulcastStr = `a=simulcast:send 1,~4;2;3 recv c`

func TestParse(t *testing.T) {

	session, err := Parse([]byte(sdpStr))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(session.Data())
}

func TestSimulcast(t *testing.T) {

	ret := ParseSimulcastStreamList([]byte(simulcastStr))

	if len(ret) != 3 {
		t.Error("Simulcast parse error")
	}

	fmt.Println(ret)
}
