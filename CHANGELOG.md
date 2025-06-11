# CHANGELOG

+ 2025.06.03.mr56985 / 89c80a37-97db-4b91-a4f8-ba05369fd3d1.uterm_manifest.release

[WiFi] `WiredMeshNotUsingWanIface`, `ClientVlanFirst` in `WanStarlinkRouterPair`.

+ 2025.05.18.mr55915 / b255f9f4-3d9d-471c-a9ba-16149de5d478.uterm_manifest.release

[Dish] New `Request_DishAviationTest`

Request is defined as:

```golang
type DishAviationTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ThermalDemandFraction      float32                          `protobuf:"fixed32,1,opt,name=thermal_demand_fraction,json=thermalDemandFraction,proto3" json:"thermal_demand_fraction,omitempty"`
	ApplyThermalDemandFraction bool                             `protobuf:"varint,2,opt,name=apply_thermal_demand_fraction,json=applyThermalDemandFraction,proto3" json:"apply_thermal_demand_fraction,omitempty"`
	EthSpeed                   DishAviationTestRequest_EthSpeed `protobuf:"varint,3,opt,name=eth_speed,json=ethSpeed,proto3,enum=SpaceX.API.Device.DishAviationTestRequest_EthSpeed" json:"eth_speed,omitempty"`
	ApplyEthSpeed              bool                             `protobuf:"varint,4,opt,name=apply_eth_speed,json=applyEthSpeed,proto3" json:"apply_eth_speed,omitempty"`
	EthAmplitudeRegisters      uint32                           `protobuf:"varint,5,opt,name=eth_amplitude_registers,json=ethAmplitudeRegisters,proto3" json:"eth_amplitude_registers,omitempty"`
	ApplyEthAmplitudeRegisters bool                             `protobuf:"varint,6,opt,name=apply_eth_amplitude_registers,json=applyEthAmplitudeRegisters,proto3" json:"apply_eth_amplitude_registers,omitempty"`
}
```

`DishAviationTestRequest_EthSpeed` is defined as:

```golang
type DishAviationTestRequest_EthSpeed int32

const (
	DishAviationTestRequest_ETH_SPEED_100_MBPS  DishAviationTestRequest_EthSpeed = 0
	DishAviationTestRequest_ETH_SPEED_1000_MBPS DishAviationTestRequest_EthSpeed = 1
)

// Enum value maps for DishAviationTestRequest_EthSpeed.
var (
	DishAviationTestRequest_EthSpeed_name = map[int32]string{
		0: "ETH_SPEED_100_MBPS",
		1: "ETH_SPEED_1000_MBPS",
	}
	DishAviationTestRequest_EthSpeed_value = map[string]int32{
		"ETH_SPEED_100_MBPS":  0,
		"ETH_SPEED_1000_MBPS": 1,
	}
)
```

Response is current empty:

```golang
type DishAviationTestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}
```

[Dish] Two more possible alerts in `get_status` response:

```golang
DishWaterDetected          bool `protobuf:"varint,20,opt,name=dish_water_detected,json=dishWaterDetected,proto3" json:"dish_water_detected,omitempty"`
RouterWaterDetected        bool `protobuf:"varint,21,opt,name=router_water_detected,json=routerWaterDetected,proto3" json:"router_water_detected,omitempty"`
```



+ 2025.05.15.cr55718 / acced5ef-ab18-4f52-98ba-d0c232895cb6.uterm_manifest.release

A few proto files have been renamed from `spacex/api` to `spacex_api`.

```
spacex/api/common/status/status.proto
spacex_api/device/command.proto
spacex_api/device/common.proto
spacex_api/device/device.proto
spacex_api/device/dish_config.proto
spacex_api/device/dish.proto
spacex_api/device/rssi_scan.proto
spacex_api/device/services/unlock/service.proto
spacex_api/device/transceiver.proto
spacex_api/device/wifi_config.proto
spacex_api/device/wifi.proto
spacex_api/device/wifi_util.proto
spacex/api/satellites/network/ut_disablement_codes.proto
spacex/api/telemetron/public/common/time.proto
spacex/api/telemetron/public/integrations/ut_pop_link_report.proto
```

+ 2025.05.08.mr55289 / 4aa5e4be-7b7b-42ab-ac33-993b3ef1183a.uterm_manifest.release

[Dish] New `ned2dishQuaternion` in `get_status`

```
type Quaternion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QScalar float32 `protobuf:"fixed32,1,opt,name=q_scalar,json=qScalar,proto3" json:"q_scalar,omitempty"`
	QX      float32 `protobuf:"fixed32,2,opt,name=q_x,json=qX,proto3" json:"q_x,omitempty"`
	QY      float32 `protobuf:"fixed32,3,opt,name=q_y,json=qY,proto3" json:"q_y,omitempty"`
	QZ      float32 `protobuf:"fixed32,4,opt,name=q_z,json=qZ,proto3" json:"q_z,omitempty"`
}
```

[Dish] Possible `downstreamRouters` in `get_status`, e.g.,

```
  "downstreamRouters": {
      "Router-0100000000000000xxxx": {
        "lastSeen": "1747292896928698453"
      }
    },
```

+ 2025.04.29.mr54648.2 / 022173f2-f284-4b31-a438-47718a77e227.uterm_manifest.release

[Dish] New `ScheduleReboot` in `UpdateRequest`.

+ 2025.04.20.mr54040 / adee186d-0c2f-4dc7-a3bd-db1e7ae17766.uterm_manifest.release

[Dish] `DlBandwidthRestrictedReason` and `UlBandwidthRestrictedReason` in `get_status` response.

[Dish] New `DishApsStats`:

```golang
type DishApsStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppVersion            uint64  `protobuf:"varint,2,opt,name=app_version,json=appVersion,proto3" json:"app_version,omitempty"`
	BootVersion           uint64  `protobuf:"varint,3,opt,name=boot_version,json=bootVersion,proto3" json:"boot_version,omitempty"`
	RomVersion            uint64  `protobuf:"varint,4,opt,name=rom_version,json=romVersion,proto3" json:"rom_version,omitempty"`
	Uptime                int64   `protobuf:"varint,5,opt,name=uptime,proto3" json:"uptime,omitempty"`
	DishPower             float32 `protobuf:"fixed32,6,opt,name=dish_power,json=dishPower,proto3" json:"dish_power,omitempty"`
	ForceDevSignedAllowed bool    `protobuf:"varint,7,opt,name=force_dev_signed_allowed,json=forceDevSignedAllowed,proto3" json:"force_dev_signed_allowed,omitempty"`
	DebugPortLocked       bool    `protobuf:"varint,8,opt,name=debug_port_locked,json=debugPortLocked,proto3" json:"debug_port_locked,omitempty"`
	BoardRev              int32   `protobuf:"varint,9,opt,name=board_rev,json=boardRev,proto3" json:"board_rev,omitempty"`
}
```

+ 2025.04.08.mr53207.1 / 9e5d61d9-e441-43a4-907a-46577dbea7ce.uterm_manifest.release

[Dish] `Response_SetPerVehicleConfig`, related to dishes with mobility?

+ 2025.04.08.cr53207 / 05de8289-7bcc-476b-ad62-8cf8cc2a73fe.uterm_manifest.release

[Dish] New `phyRxBeamSnrAvg` and `tCenter` in `get_status` response

+ 2025.03.25.mr52256.1 / b555d5d1-8198-451e-a546-8e732b39a52b.uterm_manifest.release

[Dish] New `HardwareSelfTestCodes` in `DishGetDiagnostics` Response, https://github.com/clarkzjw/starlink-grpc-golang/blob/master/pkg/spacex.com/api/device/device.pb.go#L356-L427

```golang
type DishGetDiagnosticsResponse_TestResultCode int32

const (
	DishGetDiagnosticsResponse_GENERAL           DishGetDiagnosticsResponse_TestResultCode = 0
	DishGetDiagnosticsResponse_BOOT_UP           DishGetDiagnosticsResponse_TestResultCode = 1
	DishGetDiagnosticsResponse_CPU_VOLTAGE       DishGetDiagnosticsResponse_TestResultCode = 2
	DishGetDiagnosticsResponse_DBF_AAP_CS        DishGetDiagnosticsResponse_TestResultCode = 3
	DishGetDiagnosticsResponse_DBF_NUM_FEMS      DishGetDiagnosticsResponse_TestResultCode = 4
	DishGetDiagnosticsResponse_DBF_READ_ERRORS   DishGetDiagnosticsResponse_TestResultCode = 5
	DishGetDiagnosticsResponse_DBF_T_DIE_0       DishGetDiagnosticsResponse_TestResultCode = 6
	DishGetDiagnosticsResponse_DBF_T_DIE_1       DishGetDiagnosticsResponse_TestResultCode = 7
	DishGetDiagnosticsResponse_DBF_T_DIE_0_VALID DishGetDiagnosticsResponse_TestResultCode = 8
	DishGetDiagnosticsResponse_DBF_T_DIE_1_VALID DishGetDiagnosticsResponse_TestResultCode = 9
	DishGetDiagnosticsResponse_ETH_PRIME         DishGetDiagnosticsResponse_TestResultCode = 10
	DishGetDiagnosticsResponse_EIRP              DishGetDiagnosticsResponse_TestResultCode = 11
	DishGetDiagnosticsResponse_FEM_CUT           DishGetDiagnosticsResponse_TestResultCode = 12
	DishGetDiagnosticsResponse_FUSE_AVS          DishGetDiagnosticsResponse_TestResultCode = 13
	DishGetDiagnosticsResponse_GPS               DishGetDiagnosticsResponse_TestResultCode = 14
	DishGetDiagnosticsResponse_IMU               DishGetDiagnosticsResponse_TestResultCode = 15
	DishGetDiagnosticsResponse_PHY               DishGetDiagnosticsResponse_TestResultCode = 16
	DishGetDiagnosticsResponse_SCP_ERROR         DishGetDiagnosticsResponse_TestResultCode = 17
	DishGetDiagnosticsResponse_TEMPERATURE       DishGetDiagnosticsResponse_TestResultCode = 18
	DishGetDiagnosticsResponse_VTSENS            DishGetDiagnosticsResponse_TestResultCode = 19
)
```

+ 2025.03.20.mr51925.1 / edc8b999-05f4-4dd2-99d6-53e0fbb6c3f0.uterm_manifest.release

[Dish] New `DisablementCode`: `MOVING_TOO_FAST_FOR_POLICY`.

+ 2025.03.09.mr51189 / bb84b35f-5718-4348-a5f7-06df03048731.uterm_manifest.release

[Wi-Fi] Added `WifiConfig_StaticRoute`.

+ 2025.03.02.mr50739 / a19b04ea-de59-4809-acd1-d8cd40e9a2e5.uterm_manifest.release

[Dish] Possible `DlBandwidthRestrictedReason` and `UlBandwidthRestrictedReason` values in `get_status` response.

```golang
// Enum value maps for RateLimitReason.
var (
	RateLimitReason_name = map[int32]string{
		0: "UNKNOWN",
		1: "NO_LIMIT",
		2: "POLICY_LIMIT",
		3: "USER_CUSTOM_LIMIT",
		5: "OVERAGE_LIMIT",
	}
	RateLimitReason_value = map[string]int32{
		"UNKNOWN":           0,
		"NO_LIMIT":          1,
		"POLICY_LIMIT":      2,
		"USER_CUSTOM_LIMIT": 3,
		"OVERAGE_LIMIT":     5,
	}
)
```

[Wi-Fi] `OnlyOverflightBlocked` and `OfflineNetworksDisabled` in `WifiAlerts`.

[Wi-Fi] New `DisableWhenOffline` option for `WifiConfig_Network`.

+ 2025.02.09.cr49354.1.18167 / 2bea9726-0752-4b7b-98d5-f5edeaecbce8.uterm_manifest.release

[Dish] `Request_GetGoroutineStackTraces` is removed, as well as `RebootScheduledUtcTime`.

+ 2025.02.18.mr49904 / b1a8e2de-0c8b-42fd-a607-77492f3ddff7.uterm_manifest.release

[Dish] Added `Request_GetGoroutineStackTraces`, but it seems it is internal or unimplemented, https://github.com/clarkzjw/starlink-grpc-golang/commit/5d47e368685b19a9fbc83656e8bf606b7d4632c7

+ 2025.02.05.mr49087 / d1223ab6-7fb9-490d-affc-dea80a6be413.uterm_manifest.release

[Wi-Fi] `DishDisablementCode` in `WifiGetStatusResponse`.

+ 2025.01.22.mr48363 / 79a168dc-3961-4b23-9697-b6eda232af4d.uterm_manifest.release

[Wi-Fi] Minor WiFi updates: https://github.com/clarkzjw/starlink-grpc-golang/commit/ff01422593efcbcbbc0c8739169a3e9f812a83a4

+ 2025.01.14.mr48087.1 / 461310f4-7bb6-43e5-8607-15f480713604.uterm_manifest.release

[Wi-Fi] `UnbridgedEthPorts` and `ApplyUnbridgedEthPorts` in `WifiConfig`. https://github.com/clarkzjw/starlink-grpc-golang/commit/a1eb95e2820bae11893e211321e0b303fcc59de3

+ 2025.01.10.mr47961.1 / 7046351c-be8d-4f57-a475-7c0485a04ee8.uterm_manifest.release

[Wi-Fi] `UnixTimestampNs` in `WifiClientSandboxRequest`

+ 2025.01.10.cr47961.1 / 92220b6b-7815-4477-ac69-22e3a0b02e20.uterm_manifest.release

[Wi-Fi] [`ConfigHttpsContentHostingEnabled`](https://github.com/clarkzjw/starlink-grpc-golang/commit/0bc26348e4db54b8785dd5dd189d258d927fff8f)

+ 2024.12.28.mr47554.3 / 729bfb05-8e38-4857-9092-cb2d6955b5d7.uterm_manifest.release

[Dish] `INVALID_COUNTRY` was removed from dish Disablement Code

[Wi-Fi] `SecsSinceLastPublicIpv4Change` is added to `WifiGetStatusResponse`.

[Wi-Fi] `PopIpv4PingDropRateLast_15S`, `PopIpv6PingDropRateLast_15S`, `GoogleIpv4PingDropRateLast_15S`, `GoogleIpv6PingDropRateLast_15S`, `CloudflareIpv4PingDropRateLast_15S`, `CloudflareIpv6PingDropRateLast_15S`, `DnsResolverDropRate` in `WifiGetHistoryResponse`.

```golang
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

```golang
IsMovingFastPersisted              bool                                   `protobuf:"varint,1042,opt,name=is_moving_fast_persisted,json=isMovingFastPersisted,proto3" json:"is_moving_fast_persisted,omitempty"`
UpsuStats                          *DishUpsuStats                         `protobuf:"bytes,1043,opt,name=upsu_stats,json=upsuStats,proto3" json:"upsu_stats,omitempty"`
```

where `DishUpsuStats` is defined as

```golang
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
[Wi-Fi]

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
