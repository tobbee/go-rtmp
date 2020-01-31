//
// Copyright (c) 2018- yutopp (yutopp@gmail.com)
//
// Distributed under the Boost Software License, Version 1.0. (See accompanying
// file LICENSE_1_0.txt or copy at  https://www.boost.org/LICENSE_1_0.txt)
//

package message

import (
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

//
type NetStreamPublish struct {
	CommandObject  interface{}
	PublishingName string
	PublishingType string
}

func (t *NetStreamPublish) FromArgs(args ...interface{}) error {
	//command := args[0] // will be nil
	t.PublishingName = args[1].(string)
	t.PublishingType = args[2].(string)

	return nil
}

func (t *NetStreamPublish) ToArgs(ty EncodingType) ([]interface{}, error) {
	return []interface{}{
		nil, // Always nil
		t.PublishingName,
		t.PublishingType,
	}, nil
}

//
type NetStreamPlay struct {
	CommandObject interface{}
	StreamName    string
	Start         int64
}

func (t *NetStreamPlay) FromArgs(args ...interface{}) error {
	//command := args[0] // will be nil
	t.StreamName = args[1].(string)
	t.Start = args[2].(int64)

	return nil
}

func (t *NetStreamPlay) ToArgs(ty EncodingType) ([]interface{}, error) {
	panic("Not implemented")
}

//
type NetStreamOnStatusLevel string

const (
	NetStreamOnStatusLevelStatus NetStreamOnStatusLevel = "status"
	NetStreamOnStatusLevelError  NetStreamOnStatusLevel = "error"
)

type NetStreamOnStatusCode string

const (
	NetStreamOnStatusCodeConnectSuccess      NetStreamOnStatusCode = "NetStream.Connect.Success"
	NetStreamOnStatusCodeConnectFailed       NetStreamOnStatusCode = "NetStream.Connect.Failed"
	NetStreamOnStatusCodeMuticastStreamReset NetStreamOnStatusCode = "NetStream.MulticastStream.Reset"
	NetStreamOnStatusCodePlayStart           NetStreamOnStatusCode = "NetStream.Play.Start"
	NetStreamOnStatusCodePlayStop            NetStreamOnStatusCode = "NetStream.Play.Stop"
	NetStreamOnStatusCodePlayReset           NetStreamOnStatusCode = "NetStream.Play.Reset"
	NetStreamOnStatusCodePlayFailed          NetStreamOnStatusCode = "NetStream.Play.Failed"
	NetStreamOnStatusCodePlayComplete        NetStreamOnStatusCode = "NetStream.Play.Complete"
	NetStreamOnStatusCodePlayUnpublishNotify NetStreamOnStatusCode = "NetStream.Play.UnpublishNotify"
	NetStreamOnStatusCodePlayStreamNotFound  NetStreamOnStatusCode = "NetStream.Play.StreamNotFound"
	NetStreamOnStatusCodePublishBadName      NetStreamOnStatusCode = "NetStream.Publish.BadName"
	NetStreamOnStatusCodePublishFailed       NetStreamOnStatusCode = "NetStream.Publish.Failed"
	NetStreamOnStatusCodePublishStart        NetStreamOnStatusCode = "NetStream.Publish.Start"
	NetStreamOnStatusCodeUnpublishSuccess    NetStreamOnStatusCode = "NetStream.Unpublish.Success"
	NetStreamOnStatusCodeDataStart           NetStreamOnStatusCode = "NetStream.Data.Start"
)

type NetStreamOnStatus struct {
	InfoObject NetStreamOnStatusInfoObject
}

type NetStreamOnStatusInfoObject struct {
	Level          NetStreamOnStatusLevel `mapstructure:"level" amf0:"level"`
	Code           NetStreamOnStatusCode  `mapstructure:"code" amf0:"code"`
	Description    string                 `mapstructure:"description" amf0:"description"`
	ClientID       int                    `mapstructure:"clientid" amf0:"clientid"`
	IsFastPlay     bool                   `mapstructure:"isFastPlay" amf0:"isFastPlay"`
	TimeCodeOffset string                 `mapstructure:"timecodeOffset" amf0:"timecodeOffset"`
}

func (t *NetStreamOnStatus) FromArgs(args ...interface{}) error {
	properties, ok := args[0].(map[string]interface{})
	if !ok {
		return errors.New("NetStreamOnStatus arg[0] incorrect type")
	}
	if err := mapstructure.Decode(properties, &t.InfoObject); err != nil {
		return errors.Wrapf(err, "Failed to map NetStreamOnStatus")
	}
	return nil
}

func (t *NetStreamOnStatus) ToArgs(ty EncodingType) ([]interface{}, error) {
	return []interface{}{
		nil, // Always nil
		t.InfoObject,
	}, nil
}

//
type NetStreamDeleteStream struct {
	StreamID uint32
}

func (t *NetStreamDeleteStream) FromArgs(args ...interface{}) error {
	// args[0] is unknown, ignore
	t.StreamID = args[1].(uint32)

	return nil
}

func (t *NetStreamDeleteStream) ToArgs(ty EncodingType) ([]interface{}, error) {
	panic("Not implemented")
}

//
type NetStreamFCPublish struct {
	StreamName string
}

func (t *NetStreamFCPublish) FromArgs(args ...interface{}) error {
	// args[0] is unknown, ignore
	t.StreamName = args[1].(string)

	return nil
}

func (t *NetStreamFCPublish) ToArgs(ty EncodingType) ([]interface{}, error) {
	return []interface{}{
		nil, // no command object
		t.StreamName,
	}, nil
}

//
type NetStreamFCUnpublish struct {
	StreamName string
}

func (t *NetStreamFCUnpublish) FromArgs(args ...interface{}) error {
	// args[0] is unknown, ignore
	t.StreamName = args[1].(string)

	return nil
}

func (t *NetStreamFCUnpublish) ToArgs(ty EncodingType) ([]interface{}, error) {
	return []interface{}{
		nil, // no command object
		t.StreamName,
	}, nil
}

//
type NetStreamSetDataFrame struct {
	Payload []byte
}

func (t *NetStreamSetDataFrame) FromArgs(args ...interface{}) error {
	t.Payload = args[0].([]byte)

	return nil
}

func (t *NetStreamSetDataFrame) ToArgs(ty EncodingType) ([]interface{}, error) {
	panic("Not implemented")
}

//
type NetStreamGetStreamLength struct {
	StreamName string
}

func (t *NetStreamGetStreamLength) FromArgs(args ...interface{}) error {
	// args[0] is unknown, ignore
	t.StreamName = args[1].(string)

	return nil
}

func (t *NetStreamGetStreamLength) ToArgs(ty EncodingType) ([]interface{}, error) {
	return []interface{}{
		nil, // no command object
		t.StreamName,
	}, nil
}

//
type NetStreamPing struct {
}

func (t *NetStreamPing) FromArgs(args ...interface{}) error {
	// args[0] is unknown, ignore

	return nil
}

func (t *NetStreamPing) ToArgs(ty EncodingType) ([]interface{}, error) {
	return []interface{}{
		nil, // no command object
	}, nil
}

//
type NetStreamCloseStream struct {
}

func (t *NetStreamCloseStream) FromArgs(args ...interface{}) error {
	// args[0] is unknown, ignore

	return nil
}

func (t *NetStreamCloseStream) ToArgs(ty EncodingType) ([]interface{}, error) {
	return []interface{}{
		nil, // no command object
	}, nil
}

//
type NetStreamClientPlayMessage struct {
	Filename string
}

func (t *NetStreamClientPlayMessage) FromArgs(args ...interface{}) error {
	panic("Not implemented")
}

func (t *NetStreamClientPlayMessage) ToArgs(ty EncodingType) ([]interface{}, error) {
	return []interface{}{nil, t.Filename}, nil
}

//
type NetStreamRtmpSampleAccess struct {
	AudioAllowed bool
	VideoAllowed bool
}

func (t *NetStreamRtmpSampleAccess) FromArgs(args ...interface{}) error {
	audioAllowed, ok := args[0].(bool)
	if !ok {
		return errors.New("NetStreamRtmpSampleAccess args[0] incorrect type")
	}
	videoAllowed, ok := args[1].(bool)
	if !ok {
		return errors.New("NetStreamRtmpSampleAccess args[1] incorrect type")
	}
	t.AudioAllowed = audioAllowed
	t.VideoAllowed = videoAllowed
	return nil
}

func (t *NetStreamRtmpSampleAccess) ToArgs(ty EncodingType) ([]interface{}, error) {
	return []interface{}{t.AudioAllowed, t.VideoAllowed}, nil
}

//
type NetStreamOnMetaData struct {
	InfoObject NetStreamOnMetaDataInfoObject
}

type NetStreamOnMetaDataInfoObject struct {
	Duration        int                  `mapstructure:"duration" amf0:"duration"`
	Width           int                  `mapstructure:"width" amf0:"width"`
	Height          int                  `mapstructure:"height" amf0:"height"`
	VideoDataRate   float64              `mapstructure:"videodatarate" amf0:"videodatarate"`
	FrameRate       int                  `mapstructure:"framerate" amf0:"framerate"`
	VideoCodecID    string               `mapstructure:"videocodecid" amf0:"videocodecid"`
	AudioSampleSize int                  `mapstructure:"audiosamplesize" amf0:"audiosamplesize"`
	Stereo          bool                 `mapstructure:"stereo" amf0:"stereo"`
	Encoder         string               `mapstructure:"encoder" amf0:"encoder"`
	FileSize        int64                `mapstructure:"filesize" amf0:"filesize"`
	TrackInfo       []NetStreamTrackInfo `mapstructure:"trackinfo" amf0:"trackinfo"`
	AudioChannels   int                  `mapstructure:"audiochannels" amf0:"audiochannels"`
	AudioSampleRate int                  `mapstructure:"audiosamplerate" amf0:"audiosamplerate"`
	AudioCodecID    string               `mapstructure:"audiocodecid" amf0:"audiocodecid"`
	AudioDataRate   int                  `mapstructure:"audiodatarate" amf0:"audiodatarate"`
	Transcoder      NetStreamTranscoder  `mapstructure:"transcoder" amf0:"transcoder"`
}

type NetStreamTrackInfo struct {
	Type              string                       `mapstructure:"type" amf0:"type"` // TODO audio, video, ?
	Timescale         int                          `mapstructure:"timescale" amf0:"timescale"`
	Config            string                       `mapstructure:"config" amf0:"config"`
	Description       string                       `mapstructure:"description" amf0:"description"`
	SampleDescription []NetStreamSampleDescription `mapstructure:"sampledescription" amf0:"sampledescription"`
}

type NetStreamSampleDescription struct {
	SampleType string `mapstructure:"sampletype" amf0:"sampletype"`
}

type NetStreamTranscoder struct {
	AudioCodec          string `mapstructure:"audioCodec" amf0:"audioCodec"`
	AudioBitRate        int    `mapstructure:"audioBitrate" amf0:"audioBitrate"`
	AudioEncodingParams bool   `mapstructure:"audioEncodingParams" amf0:"audioEncodingParams"`
	VideoCodec          string `mapstructure:"videoCodec" amf0:"videoCodec"`
	VideoBitRate        string `mapstructure:"videoBitrate" amf0:"videoBitrate"`
}

func (t *NetStreamOnMetaData) FromArgs(args ...interface{}) error {
	properties, ok := args[0].(map[string]interface{})
	if !ok {
		return errors.New("NetStreamOnMetaData args[0] incorrect type")
	}
	if err := mapstructure.Decode(properties, &t.InfoObject); err != nil {
		return errors.Wrapf(err, "Failed to map NetStreamOnMetaData")
	}
	return nil
}

func (t *NetStreamOnMetaData) ToArgs(ty EncodingType) ([]interface{}, error) {
	return []interface{}{t.InfoObject}, nil
}

//
type NetStreamOnSDES struct {
	InfoObject NetStreamOnSDESInfoObject
}

type NetStreamOnSDESInfoObject struct {
	CNAME string
}

func (t *NetStreamOnSDES) FromArgs(args ...interface{}) error {
	properties, ok := args[0].(map[string]interface{})
	if !ok {
		return errors.New("NetStreamOnSDES args[0] incorrect type")
	}
	if err := mapstructure.Decode(properties, &t.InfoObject); err != nil {
		return errors.Wrapf(err, "Failed to map NetStreamOnSDES")
	}
	return nil
}

func (t *NetStreamOnSDES) ToArgs(ty EncodingType) ([]interface{}, error) {
	return []interface{}{t.InfoObject}, nil
}
