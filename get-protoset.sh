#!/usr/bin/env bash

# requires https://github.com/fullstorydev/grpcurl

grpc_endpoint="192.168.100.1:9200"

dish_version=$(grpcurl -plaintext -d {\"get_status\":{}} $grpc_endpoint SpaceX.API.Device.Device/Handle | grep softwareVersion | awk '{print $2}' | awk -F',' '{print $1}' | awk -F'\"' '{print $2}')
echo "$dish_version" > VERSION
mkdir -p "protoset/$dish_version"

# get dish protoset
grpcurl -plaintext -protoset-out protoset/"$dish_version"/dish.protoset $grpc_endpoint describe SpaceX.API.Device.Device

protoc --decode_raw < protoset/"$dish_version"/dish.protoset | grep "proto\"$" | awk '{print $2}' | awk -F\" '{print $2}' | grep spacex | sort | uniq > protoset/"$dish_version"/VERSION
