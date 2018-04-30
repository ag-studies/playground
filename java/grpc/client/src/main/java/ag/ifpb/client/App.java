package ag.ifpb.client;

import java.util.concurrent.ExecutionException;

import com.google.common.util.concurrent.ListenableFuture;

import ifpb.grpc.HelloRequest;
import ifpb.grpc.HelloResponse;
import ifpb.grpc.HelloServiceGrpc;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

/**
 * Hello world!
 *
 */
public class App {
	public static void main(String[] args) throws InterruptedException, ExecutionException {
		//channel
		ManagedChannel channel = ManagedChannelBuilder.forAddress("localhost", 8080).usePlaintext().build();
		HelloServiceGrpc.HelloServiceBlockingStub stub = HelloServiceGrpc.newBlockingStub(channel);
		HelloServiceGrpc.HelloServiceFutureStub futureStub = HelloServiceGrpc.newFutureStub(channel);
		//request
		HelloRequest req = HelloRequest.newBuilder().setText("Ari Garcia").build();
		//future
		ListenableFuture<HelloResponse> fut = futureStub.hello(req);
		System.out.println("Aguardando future");
		//sync call
		HelloResponse resp = stub.hello(req);
		System.out.println("Sync resp: " + resp.getText());
		//async call
		System.out.println("Async resp: " + fut.get().getText());
		//shutdown
		channel.shutdown();
	}
}
