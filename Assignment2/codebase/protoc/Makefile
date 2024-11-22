db_out_dir=../golang/gogrpc
db_proto_files=dbs.proto

gen-db:
	protoc -I=./ --go_out=${db_out_dir} --go-grpc_out=${db_out_dir} ${db_proto_files}


log_go_dir=../golang/gogrpc
log_py_dir=../python/logging_service
log_proto_files=glog.proto


gen-go-log:
	protoc -I=./ --go_out=${log_go_dir} --go-grpc_out=${log_go_dir} ${log_proto_files}

py-log-clean:
	rm -rf ${log_py_dir}/*

gen-py-log: py-log-clean
	python -m grpc_tools.protoc -I=./ --python_out=${log_py_dir} --grpc_python_out=${log_py_dir} ${log_proto_files}