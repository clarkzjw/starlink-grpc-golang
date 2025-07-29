#!/usr/bin/env bash

rm -rf ./pkg
mkdir ./pkg

dish_version=$(cat VERSION)
echo "$dish_version"

protoset_file=protoset/"$dish_version"/dish.protoset

protoc --go_out=./pkg/ --go-grpc_out=./pkg/ --descriptor_set_in="$protoset_file" spacex_api/device/rssi_scan.proto

protoc --go_out=./pkg/ --go-grpc_out=./pkg/ --descriptor_set_in="$protoset_file" spacex_api/common/status/status.proto
protoc --go_out=./pkg/ --go-grpc_out=./pkg/ --descriptor_set_in="$protoset_file" spacex_api/device/command.proto
protoc --go_out=./pkg/ --go-grpc_out=./pkg/ --descriptor_set_in="$protoset_file" spacex_api/device/common.proto

protoc --go_out=./pkg/  --go-grpc_out=./pkg/ \
       --go_opt=Mspacex_api/common/protobuf/internal.proto=spacex.com/api/common/protobuf/internal \
       --go_opt=Mspacex_api/satellites/network/ut_disablement_codes.proto=spacex.com/api/satellites/network/ut_disablement_codes \
       --go_opt=Mspacex_api/telemetron/public/common/time.proto=spacex.com/api/telemetron/public/common/time \
       --go_opt=Mspacex_api/device/services/unlock/service.proto=spacex.com/api/device/services/unlock/service \
       --go-grpc_opt=Mspacex_api/common/protobuf/internal.proto=spacex.com/api/common/protobuf/internal \
       --go-grpc_opt=Mspacex_api/satellites/network/ut_disablement_codes.proto=spacex.com/api/satellites/network/ut_disablement_codes \
       --go-grpc_opt=Mspacex_api/telemetron/public/common/time.proto=spacex.com/api/telemetron/public/common/time \
       --go-grpc_opt=Mspacex_api/device/services/unlock/service.proto=spacex.com/api/device/services/unlock/service \
       --descriptor_set_in="$protoset_file" \
       spacex_api/device/device.proto

protoc --go_out=./pkg/  --go-grpc_out=./pkg/ \
       --go_opt=Mspacex_api/telemetron/public/common/time.proto=spacex.com/api/telemetron/public/common/time \
       --go-grpc_opt=Mspacex_api/telemetron/public/common/time.proto=spacex.com/api/telemetron/public/common/time \
       --go_opt=Mspacex_api/common/protobuf/internal.proto=spacex.com/api/common/protobuf/internal \
       --go_opt=Mspacex_api/satellites/network/ut_disablement_codes.proto=spacex.com/api/satellites/network/ut_disablement_codes \
       --go-grpc_opt=Mspacex_api/common/protobuf/internal.proto=spacex.com/api/common/protobuf/internal \
       --go-grpc_opt=Mspacex_api/satellites/network/ut_disablement_codes.proto=spacex.com/api/satellites/network/ut_disablement_codes \
       --descriptor_set_in="$protoset_file" \
       spacex_api/device/dish.proto

protoc --go_out=./pkg/ --go-grpc_out=./pkg/ --descriptor_set_in="$protoset_file" spacex_api/device/dish_config.proto
protoc --go_out=./pkg/ --go-grpc_out=./pkg/ --descriptor_set_in="$protoset_file" spacex_api/device/rssi_scan.proto

protoc --go_out=./pkg/ --go-grpc_out=./pkg/ \
       --go_opt=Mspacex_api/device/services/unlock/service.proto=spacex.com/api/device/services/unlock/service \
       --go-grpc_opt=Mspacex_api/device/services/unlock/service.proto=spacex.com/api/device/services/unlock/service \
       --descriptor_set_in="$protoset_file" \
       spacex_api/device/services/unlock/service.proto

protoc --go_out=./pkg/  --go-grpc_out=./pkg/ \
       --go_opt=Mspacex_api/telemetron/public/common/time.proto=spacex.com/api/telemetron/public/common/time \
       --go-grpc_opt=Mspacex_api/telemetron/public/common/time.proto=spacex.com/api/telemetron/public/common/time \
       --go_opt=Mspacex_api/common/protobuf/internal.proto=spacex.com/api/common/protobuf/internal \
       --go_opt=Mspacex_api/satellites/network/ut_disablement_codes.proto=spacex.com/api/satellites/network/ut_disablement_codes \
       --go-grpc_opt=Mspacex_api/common/protobuf/internal.proto=spacex.com/api/common/protobuf/internal \
       --go-grpc_opt=Mspacex_api/satellites/network/ut_disablement_codes.proto=spacex.com/api/satellites/network/ut_disablement_codes \
       --descriptor_set_in="$protoset_file" \
       spacex_api/device/transceiver.proto

protoc --go_out=./pkg/  --go-grpc_out=./pkg/ \
       --go_opt=Mspacex_api/common/protobuf/internal.proto=spacex.com/api/common/protobuf/internal \
       --go_opt=Mspacex_api/satellites/network/ut_disablement_codes.proto=spacex.com/api/satellites/network/ut_disablement_codes \
       --go-grpc_opt=Mspacex_api/common/protobuf/internal.proto=spacex.com/api/common/protobuf/internal \
       --go-grpc_opt=Mspacex_api/satellites/network/ut_disablement_codes.proto=spacex.com/api/satellites/network/ut_disablement_codes \
       --descriptor_set_in="$protoset_file" \
       spacex_api/satellites/network/ut_disablement_codes.proto

protoc --go_out=./pkg/  --go-grpc_out=./pkg/ \
       --go_opt=Mspacex_api/telemetron/public/common/time.proto=spacex.com/api/telemetron/public/common/time \
       --go-grpc_opt=Mspacex_api/telemetron/public/common/time.proto=spacex.com/api/telemetron/public/common/time \
       --descriptor_set_in="$protoset_file" \
       spacex_api/telemetron/public/integrations/rate_limit_reason.proto

protoc --go_out=./pkg/  --go-grpc_out=./pkg/ \
       --go_opt=Mspacex_api/telemetron/public/common/time.proto=spacex.com/api/telemetron/public/common/time \
       --go-grpc_opt=Mspacex_api/telemetron/public/common/time.proto=spacex.com/api/telemetron/public/common/time \
       --descriptor_set_in="$protoset_file" \
       spacex_api/telemetron/public/common/time.proto

protoc --go_out=./pkg/  --go-grpc_out=./pkg/ \
       --go_opt=Mspacex_api/telemetron/public/common/time.proto=spacex.com/api/telemetron/public/common/time \
       --go-grpc_opt=Mspacex_api/telemetron/public/common/time.proto=spacex.com/api/telemetron/public/common/time \
       --go_opt=Mspacex_api/satellites/network/ut_disablement_codes.proto=spacex.com/api/satellites/network/ut_disablement_codes \
       --go-grpc_opt=Mspacex_api/satellites/network/ut_disablement_codes.proto=spacex.com/api/satellites/network/ut_disablement_codes \
       --descriptor_set_in="$protoset_file" \
       spacex_api/device/wifi.proto

protoc --go_out=./pkg/  --go-grpc_out=./pkg/ --descriptor_set_in="$protoset_file" spacex_api/device/wifi_config.proto
protoc --go_out=./pkg/  --go-grpc_out=./pkg/ --descriptor_set_in="$protoset_file" spacex_api/device/wifi_util.proto

find ./pkg/spacex.com -name "*.go" -print0 | xargs -0 sed -i 's|spacex.com/api|github.com/clarkzjw/starlink-grpc-golang/pkg/spacex.com/api|g'
