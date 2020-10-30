package arangodb

import (
	"strconv"

	"github.com/sbezverk/gobmp/pkg/message"
)

type lsLinkArangoMessage struct {
	*message.LSLink
}

func (l *lsLinkArangoMessage) MakeKey() string {
	var localIP, remoteIP, localID, remoteID string
	switch l.MTID {
	case 0:
		localIP = "0.0.0.0"
		remoteIP = "0.0.0.0"
	case 2:
		localIP = "::"
		remoteIP = "::"
	default:
		localIP = "unknown-mt-id"
		remoteIP = "unknown-mt-id"
	}
	if len(l.LocalLinkIP) != 0 {
		localIP = l.LocalLinkIP[0]
	}
	if len(l.RemoteLinkIP) != 0 {
		remoteIP = l.RemoteLinkIP[0]
	}
	localID = strconv.Itoa(int(l.LocalLinkID))
	remoteID = strconv.Itoa(int(l.RemoteLinkID))

	// The LSLink Key uses ProtocolID, DomainID, and Multi-Topology ID
	// to create unique Keys for DB entries in multi-area / multi-topology scenarios
	return strconv.Itoa(int(l.ProtocolID)) + "_" + strconv.Itoa(int(l.DomainID)) + "_" + strconv.Itoa(int(l.MTID)) + "_" + l.AreaID + "_" + l.IGPRouterID + "_" + localIP + "_" + localID + "_" + l.RemoteIGPRouterID + "_" + remoteIP + "_" + remoteID
}
