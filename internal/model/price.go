package model

import timestamp "github.com/golang/protobuf/ptypes/timestamp"

type PriceInfo struct {
	Symbol    string               `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Price     float32              `protobuf:"fixed64,2,opt,name=current_price,proto3" json:"price,omitempty"`
	Change24h float32              `protobuf:"fixed64,3,opt,name=change24h,proto3" json:"change,omitempty"`
	Timestamp *timestamp.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Change1h  float32              `protobuf:"fixed64,5,opt,name=change1h,proto3" json:"change1h,omitempty"`
	Change7d  float32              `protobuf:"fixed64,6,opt,name=change7d,proto3" json:"change7d,omitempty"`
	High24h   float32              `protobuf:"fixed64,7,opt,name=high24h,proto3" json:"high24h,omitempty"`
	Low24h    float32              `protobuf:"fixed64,8,opt,name=low24h,proto3" json:"low24h,omitempty"`
}
