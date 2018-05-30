# Calls Protoc Library: AKA: Compiles Protobuf code.
build:
	# Updated to use go-micro plugin instead of grpc plugin.
	protoc -I. --go_out=plugins=micro:. proto/auth/auth.proto 
	# Builds an image by the name user-service (Dot means that build process looks in current directory)
	docker build -t shippy-auth-service .
	#docker push dillonlpeterson/user:latest 
run: 
	docker run --net="host" \
		-p 50051 \
		-e DB_HOST=localhost \
		-e DB_PASS=password \
		-e DB_USER=postgres \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		-e DISABLE_AUTH=true \
		shippy-auth-service 



