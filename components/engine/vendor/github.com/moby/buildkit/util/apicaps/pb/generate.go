package moby_buildkit_v1_apicaps //nolint:golint

//go:generate protoc -I=. -I=../../../vendor/ -I=../../../../../../ --gogo_out=plugins=grpc:. caps.proto
