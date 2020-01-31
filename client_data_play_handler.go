//
// Copyright (c) 2018- yutopp (yutopp@gmail.com)
//
// Distributed under the Boost Software License, Version 1.0. (See accompanying
// file LICENSE_1_0.txt or copy at  https://www.boost.org/LICENSE_1_0.txt)
//

package rtmp

import (
	"github.com/yutopp/go-rtmp/internal"
	"github.com/yutopp/go-rtmp/message"
)

var _ stateHandler = (*clientDataPlayHandler)(nil)

// clientDataPlayHandler Handle data messages from a server (NOT IMPLEMENTED).
//   transitions:
//     | _ -> self
type clientDataPlayHandler struct {
	sh *streamHandler
}

func (h *clientDataPlayHandler) onMessage(
	chunkStreamID int,
	timestamp uint32,
	msg message.Message,
) error {
	switch msg := msg.(type) {
	case *message.AudioMessage:
		return h.sh.stream.userHandler().OnAudio(timestamp, msg.Payload)

	case *message.VideoMessage:
		return h.sh.stream.userHandler().OnVideo(timestamp, msg.Payload)

	case *message.UserCtrl:
		return h.sh.stream.userHandler().OnUserCtrlEvent(timestamp, msg.Event)

	default:
		return internal.ErrPassThroughMsg
	}
}

func (h *clientDataPlayHandler) onData(
	chunkStreamID int,
	timestamp uint32,
	dataMsg *message.DataMessage,
	body interface{},
) error {
	return internal.ErrPassThroughMsg
}

func (h *clientDataPlayHandler) onCommand(
	chunkStreamID int,
	timestamp uint32,
	cmdMsg *message.CommandMessage,
	body interface{},
) error {
	return internal.ErrPassThroughMsg
}

func (h *clientDataPlayHandler) onWinAckSize(
	chunkStreamID int,
	timestamp uint32,
	ackMsg *message.WinAckSize,
) (err error) {
	return nil
}
