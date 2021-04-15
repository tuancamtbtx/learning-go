protoc
  --proto_path=. \
  --include_imports \
  --include_source_info \
  --descriptor_set_out=api_descriptor.pb \
  bookstore.proto