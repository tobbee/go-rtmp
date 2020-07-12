//
// Copyright (c) 2018- yutopp (yutopp@gmail.com)
//
// Distributed under the Boost Software License, Version 1.0. (See accompanying
// file LICENSE_1_0.txt or copy at  https://www.boost.org/LICENSE_1_0.txt)
//

package rtmp

import (
	"github.com/tobbee/go-rtmp/internal"
	"github.com/tobbee/go-rtmp/message"
)

var _ stateHandler = (*clientControlNotConnectedHandler)(nil)

// clientControlNotConnectedHandler Handle control messages from a server in flow of connecting.
//   transitions:
//     | "_result" -> controlStreamStateConnected
//     | _         -> self
//
type clientControlNotConnectedHandler struct {
	sh *streamHandler
}

func (h *clientControlNotConnectedHandler) onMessage(
	chunkStreamID int,
	timestamp uint32,
	msg message.Message,
) error {
	switch msg := msg.(type) {
	case *message.UserCtrl:
		return h.sh.stream.userHandler().OnUserCtrlEvent(timestamp, msg.Event)

	default:
		return internal.ErrPassThroughMsg
	}
}

func (h *clientControlNotConnectedHandler) onData(
	chunkStreamID int,
	timestamp uint32,
	dataMsg *message.DataMessage,
	body interface{},
) error {
	return internal.ErrPassThroughMsg
}

func (h *clientControlNotConnectedHandler) onCommand(
	chunkStreamID int,
	timestamp uint32,
	cmdMsg *message.CommandMessage,
	body interface{},
) error {
	l := h.sh.Logger()

	switch cmd := body.(type) {
	case *message.NetConnectionConnectResult:
		l.Info("ConnectResult")
		l.Infof("Result: Info = %+v, Props = %+v", cmd.Information, cmd.Properties)

		return nil

	default:
		return internal.ErrPassThroughMsg
	}
}

func (h *clientControlNotConnectedHandler) onWinAckSize(
	chunkStreamID int,
	timestamp uint32,
	ackMsg *message.WinAckSize,
) error {
	return h.sh.stream.WriteWinAckSize(chunkStreamID, 0, ackMsg)
}
