package ifpb.grpc.server;

import ifpb.grpc.HelloRequest;
import ifpb.grpc.HelloResponse;
import ifpb.grpc.HelloServiceGrpc.HelloServiceImplBase;
import io.grpc.stub.StreamObserver;

public class HelloServiceGrpcImpl extends HelloServiceImplBase {

	@Override
	public void hello(HelloRequest request, StreamObserver<HelloResponse> responseObserver) {
		//
		String text = request.getText();
		//
		HelloResponse.Builder builder = HelloResponse.newBuilder();
		builder.setText("Hello " + text);
		HelloResponse response = builder.build();
		//
		responseObserver.onNext(response);
		responseObserver.onCompleted();
	}
}
