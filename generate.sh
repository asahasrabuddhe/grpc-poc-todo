#!/bin/bash
protoc -I proto/ proto/errors/*.proto --go_out=plugins=grpc:../../
protoc -I proto/ proto/todo/*.proto --go_out=plugins=grpc:../../
protoc -I proto/ proto/user/*.proto --go_out=plugins=grpc:../../
