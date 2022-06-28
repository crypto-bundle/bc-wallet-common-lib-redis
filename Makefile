bc-order-events:
	protoc -I ./pkg/order_events/proto \
    		--go_out=./pkg/order_events/ \
    		--go_opt=paths=source_relative \
    		--go-grpc_out=./pkg/order_events/ \
    		--go-grpc_opt=paths=source_relative \
    		--grpc-gateway_out=./pkg/order_events/ \
    		--grpc-gateway_opt=logtostderr=true \
    		--grpc-gateway_opt=paths=source_relative \
    		--doc_out=./pkg/order_events/docs/ \
    		--doc_opt=markdown,$@.md \
    		./pkg/order_events/proto/*.proto