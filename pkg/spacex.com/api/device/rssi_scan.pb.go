// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: spacex_api/device/rssi_scan.proto

package device

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RssiEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ThetaDegree     float64 `protobuf:"fixed64,1,opt,name=theta_degree,json=thetaDegree,proto3" json:"theta_degree,omitempty"`
	PhiDegree       float64 `protobuf:"fixed64,2,opt,name=phi_degree,json=phiDegree,proto3" json:"phi_degree,omitempty"`
	RssiDbf         float64 `protobuf:"fixed64,3,opt,name=rssi_dbf,json=rssiDbf,proto3" json:"rssi_dbf,omitempty"`
	ScanTimestampMs uint64  `protobuf:"varint,4,opt,name=scan_timestamp_ms,json=scanTimestampMs,proto3" json:"scan_timestamp_ms,omitempty"`
}

func (x *RssiEntry) Reset() {
	*x = RssiEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacex_api_device_rssi_scan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RssiEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RssiEntry) ProtoMessage() {}

func (x *RssiEntry) ProtoReflect() protoreflect.Message {
	mi := &file_spacex_api_device_rssi_scan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RssiEntry.ProtoReflect.Descriptor instead.
func (*RssiEntry) Descriptor() ([]byte, []int) {
	return file_spacex_api_device_rssi_scan_proto_rawDescGZIP(), []int{0}
}

func (x *RssiEntry) GetThetaDegree() float64 {
	if x != nil {
		return x.ThetaDegree
	}
	return 0
}

func (x *RssiEntry) GetPhiDegree() float64 {
	if x != nil {
		return x.PhiDegree
	}
	return 0
}

func (x *RssiEntry) GetRssiDbf() float64 {
	if x != nil {
		return x.RssiDbf
	}
	return 0
}

func (x *RssiEntry) GetScanTimestampMs() uint64 {
	if x != nil {
		return x.ScanTimestampMs
	}
	return 0
}

type DishActivateRssiScan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel uint32 `protobuf:"varint,1,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *DishActivateRssiScan) Reset() {
	*x = DishActivateRssiScan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacex_api_device_rssi_scan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DishActivateRssiScan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DishActivateRssiScan) ProtoMessage() {}

func (x *DishActivateRssiScan) ProtoReflect() protoreflect.Message {
	mi := &file_spacex_api_device_rssi_scan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DishActivateRssiScan.ProtoReflect.Descriptor instead.
func (*DishActivateRssiScan) Descriptor() ([]byte, []int) {
	return file_spacex_api_device_rssi_scan_proto_rawDescGZIP(), []int{1}
}

func (x *DishActivateRssiScan) GetChannel() uint32 {
	if x != nil {
		return x.Channel
	}
	return 0
}

type DishGetRssiScanResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success          bool         `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Channel          uint32       `protobuf:"varint,2,opt,name=channel,proto3" json:"channel,omitempty"`
	RequestTimestamp uint64       `protobuf:"varint,3,opt,name=request_timestamp,json=requestTimestamp,proto3" json:"request_timestamp,omitempty"`
	NumberSamples    uint32       `protobuf:"varint,4,opt,name=number_samples,json=numberSamples,proto3" json:"number_samples,omitempty"`
	RssiScanPoints   []*RssiEntry `protobuf:"bytes,5,rep,name=rssi_scan_points,json=rssiScanPoints,proto3" json:"rssi_scan_points,omitempty"`
}

func (x *DishGetRssiScanResult) Reset() {
	*x = DishGetRssiScanResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacex_api_device_rssi_scan_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DishGetRssiScanResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DishGetRssiScanResult) ProtoMessage() {}

func (x *DishGetRssiScanResult) ProtoReflect() protoreflect.Message {
	mi := &file_spacex_api_device_rssi_scan_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DishGetRssiScanResult.ProtoReflect.Descriptor instead.
func (*DishGetRssiScanResult) Descriptor() ([]byte, []int) {
	return file_spacex_api_device_rssi_scan_proto_rawDescGZIP(), []int{2}
}

func (x *DishGetRssiScanResult) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DishGetRssiScanResult) GetChannel() uint32 {
	if x != nil {
		return x.Channel
	}
	return 0
}

func (x *DishGetRssiScanResult) GetRequestTimestamp() uint64 {
	if x != nil {
		return x.RequestTimestamp
	}
	return 0
}

func (x *DishGetRssiScanResult) GetNumberSamples() uint32 {
	if x != nil {
		return x.NumberSamples
	}
	return 0
}

func (x *DishGetRssiScanResult) GetRssiScanPoints() []*RssiEntry {
	if x != nil {
		return x.RssiScanPoints
	}
	return nil
}

var File_spacex_api_device_rssi_scan_proto protoreflect.FileDescriptor

var file_spacex_api_device_rssi_scan_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x70, 0x61, 0x63, 0x65, 0x78, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x72, 0x73, 0x73, 0x69, 0x5f, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x53, 0x70, 0x61, 0x63, 0x65, 0x58, 0x2e, 0x41, 0x50, 0x49, 0x2e,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x94, 0x01, 0x0a, 0x09, 0x52, 0x73, 0x73, 0x69, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x68, 0x65, 0x74, 0x61, 0x5f, 0x64, 0x65,
	0x67, 0x72, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x74, 0x68, 0x65, 0x74,
	0x61, 0x44, 0x65, 0x67, 0x72, 0x65, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x68, 0x69, 0x5f, 0x64,
	0x65, 0x67, 0x72, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x70, 0x68, 0x69,
	0x44, 0x65, 0x67, 0x72, 0x65, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x73, 0x73, 0x69, 0x5f, 0x64,
	0x62, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x72, 0x73, 0x73, 0x69, 0x44, 0x62,
	0x66, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x5f, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x73, 0x63,
	0x61, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x4d, 0x73, 0x22, 0x30, 0x0a,
	0x14, 0x44, 0x69, 0x73, 0x68, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x73, 0x73,
	0x69, 0x53, 0x63, 0x61, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22,
	0xe7, 0x01, 0x0a, 0x15, 0x44, 0x69, 0x73, 0x68, 0x47, 0x65, 0x74, 0x52, 0x73, 0x73, 0x69, 0x53,
	0x63, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x2b, 0x0a,
	0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x25, 0x0a, 0x0e, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0d, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x73, 0x12, 0x46, 0x0a, 0x10, 0x72, 0x73, 0x73, 0x69, 0x5f, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x53, 0x70,
	0x61, 0x63, 0x65, 0x58, 0x2e, 0x41, 0x50, 0x49, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x52, 0x73, 0x73, 0x69, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x72, 0x73, 0x73, 0x69, 0x53,
	0x63, 0x61, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x42, 0x17, 0x5a, 0x15, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spacex_api_device_rssi_scan_proto_rawDescOnce sync.Once
	file_spacex_api_device_rssi_scan_proto_rawDescData = file_spacex_api_device_rssi_scan_proto_rawDesc
)

func file_spacex_api_device_rssi_scan_proto_rawDescGZIP() []byte {
	file_spacex_api_device_rssi_scan_proto_rawDescOnce.Do(func() {
		file_spacex_api_device_rssi_scan_proto_rawDescData = protoimpl.X.CompressGZIP(file_spacex_api_device_rssi_scan_proto_rawDescData)
	})
	return file_spacex_api_device_rssi_scan_proto_rawDescData
}

var file_spacex_api_device_rssi_scan_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_spacex_api_device_rssi_scan_proto_goTypes = []any{
	(*RssiEntry)(nil),             // 0: SpaceX.API.Device.RssiEntry
	(*DishActivateRssiScan)(nil),  // 1: SpaceX.API.Device.DishActivateRssiScan
	(*DishGetRssiScanResult)(nil), // 2: SpaceX.API.Device.DishGetRssiScanResult
}
var file_spacex_api_device_rssi_scan_proto_depIdxs = []int32{
	0, // 0: SpaceX.API.Device.DishGetRssiScanResult.rssi_scan_points:type_name -> SpaceX.API.Device.RssiEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_spacex_api_device_rssi_scan_proto_init() }
func file_spacex_api_device_rssi_scan_proto_init() {
	if File_spacex_api_device_rssi_scan_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spacex_api_device_rssi_scan_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RssiEntry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spacex_api_device_rssi_scan_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DishActivateRssiScan); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spacex_api_device_rssi_scan_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DishGetRssiScanResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spacex_api_device_rssi_scan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spacex_api_device_rssi_scan_proto_goTypes,
		DependencyIndexes: file_spacex_api_device_rssi_scan_proto_depIdxs,
		MessageInfos:      file_spacex_api_device_rssi_scan_proto_msgTypes,
	}.Build()
	File_spacex_api_device_rssi_scan_proto = out.File
	file_spacex_api_device_rssi_scan_proto_rawDesc = nil
	file_spacex_api_device_rssi_scan_proto_goTypes = nil
	file_spacex_api_device_rssi_scan_proto_depIdxs = nil
}
