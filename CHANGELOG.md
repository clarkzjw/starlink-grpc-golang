# CHANGELOG

+ 2025.01.14.mr48087.1 / 461310f4-7bb6-43e5-8607-15f480713604.uterm_manifest.release

[Wi-Fi] `UnbridgedEthPorts` and `ApplyUnbridgedEthPorts` in `WifiConfig`. https://github.com/clarkzjw/starlink-grpc-golang/commit/a1eb95e2820bae11893e211321e0b303fcc59de3

+ 2025.01.10.mr47961.1 / 7046351c-be8d-4f57-a475-7c0485a04ee8.uterm_manifest.release

[Wi-Fi] `UnixTimestampNs` in `WifiClientSandboxRequest`

+ 2025.01.10.cr47961.1 / 92220b6b-7815-4477-ac69-22e3a0b02e20.uterm_manifest.release

[WiFi] (`ConfigHttpsContentHostingEnabled`](https://github.com/clarkzjw/starlink-grpc-golang/commit/0bc26348e4db54b8785dd5dd189d258d927fff8f)

+ 2024.12.28.mr47554.3 / 729bfb05-8e38-4857-9092-cb2d6955b5d7.uterm_manifest.release

[Dish] `INVALID_COUNTRY` was removed from dish Disablement Code

[WiFi] `SecsSinceLastPublicIpv4Change` is added to `WifiGetStatusResponse`.

[WiFi] `PopIpv4PingDropRateLast_15S`, `PopIpv6PingDropRateLast_15S`, `GoogleIpv4PingDropRateLast_15S`, `GoogleIpv6PingDropRateLast_15S`, `CloudflareIpv4PingDropRateLast_15S`, `CloudflareIpv6PingDropRateLast_15S`, `DnsResolverDropRate` in `WifiGetHistoryResponse`.

```
type WifiGetHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Current                            uint64                                                `protobuf:"varint,1,opt,name=current,proto3" json:"current,omitempty"`
	PingDropRate                       []float32                                             `protobuf:"fixed32,1001,rep,packed,name=ping_drop_rate,json=pingDropRate,proto3" json:"ping_drop_rate,omitempty"`
	PingLatencyMs                      []float32                                             `protobuf:"fixed32,1002,rep,packed,name=ping_latency_ms,json=pingLatencyMs,proto3" json:"ping_latency_ms,omitempty"`
	CurrentIndex_15S                   uint64                                                `protobuf:"varint,2,opt,name=current_index_15s,json=currentIndex15s,proto3" json:"current_index_15s,omitempty"`
	PopIpv4PingDropRateLast_15S        []float32                                             `protobuf:"fixed32,1003,rep,packed,name=pop_ipv4_ping_drop_rate_last_15s,json=popIpv4PingDropRateLast15s,proto3" json:"pop_ipv4_ping_drop_rate_last_15s,omitempty"`
	PopIpv6PingDropRateLast_15S        []float32                                             `protobuf:"fixed32,1004,rep,packed,name=pop_ipv6_ping_drop_rate_last_15s,json=popIpv6PingDropRateLast15s,proto3" json:"pop_ipv6_ping_drop_rate_last_15s,omitempty"`
	GoogleIpv4PingDropRateLast_15S     []float32                                             `protobuf:"fixed32,1005,rep,packed,name=google_ipv4_ping_drop_rate_last_15s,json=googleIpv4PingDropRateLast15s,proto3" json:"google_ipv4_ping_drop_rate_last_15s,omitempty"`
	GoogleIpv6PingDropRateLast_15S     []float32                                             `protobuf:"fixed32,1006,rep,packed,name=google_ipv6_ping_drop_rate_last_15s,json=googleIpv6PingDropRateLast15s,proto3" json:"google_ipv6_ping_drop_rate_last_15s,omitempty"`
	CloudflareIpv4PingDropRateLast_15S []float32                                             `protobuf:"fixed32,1007,rep,packed,name=cloudflare_ipv4_ping_drop_rate_last_15s,json=cloudflareIpv4PingDropRateLast15s,proto3" json:"cloudflare_ipv4_ping_drop_rate_last_15s,omitempty"`
	CloudflareIpv6PingDropRateLast_15S []float32                                             `protobuf:"fixed32,1008,rep,packed,name=cloudflare_ipv6_ping_drop_rate_last_15s,json=cloudflareIpv6PingDropRateLast15s,proto3" json:"cloudflare_ipv6_ping_drop_rate_last_15s,omitempty"`
	DnsResolverDropRate                map[string]*WifiGetHistoryResponse_DnsResolverHistory `protobuf:"bytes,1009,rep,name=dns_resolver_drop_rate,json=dnsResolverDropRate,proto3" json:"dns_resolver_drop_rate,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}
```

+ 2024.12.17.mr47156.1 / 3b5ef5f2-dd64-4e29-a33b-3f6e46d060e3.uterm_manifest.release

[Dish] A `gpsTimeS` field is added to the output of `get_diagnostics`.

```
grpcurl -plaintext -d {\"get_diagnostics\":{}} 192.168.100.1:9200 SpaceX.API.Device.Device/Handle

{
  "apiVersion": "28",
  "dishGetDiagnostics": {
    "id": "ut01000000-00000000-xxxx",
    "hardwareVersion": "rev3_proto2",
    "softwareVersion": "2024.12.17.mr47156.1",
    "utcOffsetS": -28799,
    "alerts": {

    },
    "disablementCode": "OKAY",
    "hardwareSelfTest": "PASSED",
    "location": {
      "enabled": true,
      "latitude": xxxx,
      "longitude": xxxx,
      "altitudeMeters": xxxx,
      "uncertaintyMetersValid": true,
      "uncertaintyMeters": 5,
      "gpsTimeS": 1.4191529686605e+09
    },
    "alignmentStats": {

    }
  }
}
```

[Dish] Two additional optional values in the output of `get_status`.

```
IsMovingFastPersisted              bool                                   `protobuf:"varint,1042,opt,name=is_moving_fast_persisted,json=isMovingFastPersisted,proto3" json:"is_moving_fast_persisted,omitempty"`
UpsuStats                          *DishUpsuStats                         `protobuf:"bytes,1043,opt,name=upsu_stats,json=upsuStats,proto3" json:"upsu_stats,omitempty"`
```

where `DishUpsuStats` is defined as

```
type DishUpsuStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppVersion  uint64  `protobuf:"varint,2,opt,name=app_version,json=appVersion,proto3" json:"app_version,omitempty"`
	BootVersion uint64  `protobuf:"varint,3,opt,name=boot_version,json=bootVersion,proto3" json:"boot_version,omitempty"`
	RomVersion  uint64  `protobuf:"varint,4,opt,name=rom_version,json=romVersion,proto3" json:"rom_version,omitempty"`
	Uptime      int64   `protobuf:"varint,5,opt,name=uptime,proto3" json:"uptime,omitempty"`
	DishPower   float32 `protobuf:"fixed32,6,opt,name=dish_power,json=dishPower,proto3" json:"dish_power,omitempty"`
	RouterPower float32 `protobuf:"fixed32,7,opt,name=router_power,json=routerPower,proto3" json:"router_power,omitempty"`
}
```


+ 2024.12.12.mr46969.2 / 28279e47-2013-4102-8b09-3541b2000cfa.uterm_manifest.release

[Dish] A `PLCPortStats` and related structures are added.

```golang
type PLCPortStats_PortStatus int32

const (
	PLCPortStats_INACTIVE          PLCPortStats_PortStatus = 0
	PLCPortStats_CHARGING          PLCPortStats_PortStatus = 1
	PLCPortStats_DISCHARGING       PLCPortStats_PortStatus = 2
	PLCPortStats_MOISTURE_DETECTED PLCPortStats_PortStatus = 3
)

// Enum value maps for PLCPortStats_PortStatus.
var (
	PLCPortStats_PortStatus_name = map[int32]string{
		0: "INACTIVE",
		1: "CHARGING",
		2: "DISCHARGING",
		3: "MOISTURE_DETECTED",
	}
	PLCPortStats_PortStatus_value = map[string]int32{
		"INACTIVE":          0,
		"CHARGING":          1,
		"DISCHARGING":       2,
		"MOISTURE_DETECTED": 3,
	}
)
```

+ 2024.12.04.mr46620 / 99ed8940-eb65-42e7-80a6-b77b76ebccf3.uterm_manifest.release

[Dish] A new disablement code `UNSUPPORTED_VERSION` was added.

```golang
// Enum value maps for UtDisablementCode.
var (
	UtDisablementCode_name = map[int32]string{
		0:  "UNKNOWN_STATE",
		1:  "OKAY",
		2:  "NO_ACTIVE_ACCOUNT",
		3:  "TOO_FAR_FROM_SERVICE_ADDRESS",
		4:  "IN_OCEAN",
		5:  "INVALID_COUNTRY",
		6:  "BLOCKED_COUNTRY",
		7:  "DATA_OVERAGE_SANDBOX_POLICY",
		8:  "CELL_IS_DISABLED",
		10: "ROAM_RESTRICTED",
		11: "UNKNOWN_LOCATION",
		12: "ACCOUNT_DISABLED",
		13: "UNSUPPORTED_VERSION",
    }
```

+ 2024.11.25.mr46348.1 / abebf685-9154-4f4b-8fbf-19471ac93420.uterm_manifest.release

[Dish] Add `Stowed` status to `get_diagnostics`.

+ 2024.11.21.mr46196.1 / 98664e73-3b73-4b2f-a1a4-b0946a502e66.uterm_manifest.release

[Dish] Add `AlignmentStats` to `DishGetDiagnosticsResponse`

```
grpcurl -plaintext -d {\"get_diagnostics\":{}} 192.168.100.1:9200 SpaceX.API.Device.Device/Handle
{
  "apiVersion": "28",
  "dishGetDiagnostics": {
    "id": "ut01000000-00000000-[...]",
    "hardwareVersion": "rev3_proto2",
    "softwareVersion": "2024.11.21.mr46196.1",
    "utcOffsetS": -28799,
    "alerts": {

    },
    "disablementCode": "OKAY",
    "hardwareSelfTest": "PASSED",
    "location": {
      "enabled": true,
      "latitude": [...],
      "longitude": [...],
      "altitudeMeters": [...],
      "uncertaintyMetersValid": true,
      "uncertaintyMeters": 5
    },
    "alignmentStats": {

    }
  }
}
```

Although it's empty while `get_status` shows the correct `alignmentStats`:

```
[...]
    "alignmentStats": {
      "tiltAngleDeg": 27.518227,
      "boresightAzimuthDeg": -1.3881348,
      "boresightElevationDeg": 62.642834,
      "attitudeEstimationState": "FILTER_CONVERGED",
      "attitudeUncertaintyDeg": 0.48990536,
      "desiredBoresightAzimuthDeg": 0.019135801,
      "desiredBoresightElevationDeg": 62.995728
    },
[...]
```

+ 2024.10.30.mr45184.1 / 61290a56-7774-44a7-8520-197b2a6b524c.uterm_manifest.release

A bunch of WiFi updates since `2024.10.20.cr44742`.

+ 2024.10.20.cr44742 / d3c7648a-a112-4b7c-aefc-86d465e5377f.uterm_manifest.release

When compared with `2024.10.15.mr44519`, it seems it lacks the fields new in `2024.10.15.mr44519`.

Looks almost the same as `2024.10.13.mr44429`.

+ 2024.10.15.mr44519 / fb2c588e-a2de-4b97-bafb-9a64e2155f11.uterm_manifest.release

https://github.com/clarkzjw/starlink-grpc-golang/commit/75ec2a3115cdde9c91441578c79cc3783370f659

```
[Wifi]

1. `VsnsVin` in `PoeStats`
2. `Alerts` in `WifiClient`
```

+ 2024.10.13.mr44429 / 6cc6306a-c5a6-4f78-934f-21f218dbbf47.uterm_manifest.release

https://github.com/clarkzjw/starlink-grpc-golang/commit/ba911c67e23c71a75c26f0b8d532e37e76e2a479

+ d96d1637-57b4-4bb2-950c-0d30c61a38ae.uterm_manifest.release

https://github.com/clarkzjw/starlink-grpc-golang/commit/082c1396fafc6da44eac96f1d10acd381007638a

+ 1a654eb0-56bc-4df9-a1f7-8627e1f4f9ce.uterm_manifest.release

https://github.com/clarkzjw/starlink-grpc-golang/commit/d2180255428dfbf615044bdb7e28cf6b7554e4f9

+ 37af8779-6499-4528-a5bd-8ab544a01acc.uterm_manifest.release

https://github.com/clarkzjw/starlink-grpc-golang/commit/5a80926b7265af504773ff66208c70cfed93bebb
