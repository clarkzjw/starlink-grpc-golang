# CHANGELOG

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
