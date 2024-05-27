package model

import "google.golang.org/protobuf/proto"

func UnmarshalProtoAlert(alert []byte) (*Alert, error) {
	var a *Alert
	err := proto.Unmarshal(alert, a)
	if err != nil {
		return &Alert{}, err
	}
	return a, nil
}
