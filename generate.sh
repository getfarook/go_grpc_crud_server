#!/bin/bash

#this is a helper file to generate proto files for both go and dart and copy dart files to flutter project

#protoc calcpb/calc.proto --go_out=plugins=grpc:.

# generate files for go
#protoc prodpb/prod.proto --go_out=plugins=grpc:.
protoc prodpb/*.proto --go_out=plugins=grpc:.

# generate files for Dart
#protoc  prodpb/prod.proto  --dart_out=grpc:prodpb/dart_out  
protoc  prodpb/*.proto  --dart_out=grpc:prodpb/dart_out

# copy the dart out files to flutter project directory
cp -v /Users/farook/go/src/github.com/getfarook/salespro_backend/prodpb/dart_out/prodpb/* /Volumes/SamSSD/workspace/FlutterProjects/sales_pro/lib/protos_out