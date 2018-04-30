package ifpb.grpc.server;

import java.io.IOException;

import ifpb.grpc.HelloRequest;
import ifpb.grpc.HelloServiceGrpc;
import io.grpc.Server;
import io.grpc.ServerBuilder;

/**
 * Hello world!
 *
 */
public class App {
	public static void main(String[] args) throws IOException, InterruptedException {
		//
		HelloRequest.Builder builder = HelloRequest.newBuilder();
		HelloRequest req = builder.setText("Ari Garcia")
				.build();
		//
		ServerBuilder serverBuilder = ServerBuilder.forPort(8080);
		serverBuilder.addService(new HelloServiceGrpcImpl());
		Server server = serverBuilder.build();
		server.start();
	    System.out.println("Server started!");
	    server.awaitTermination();
		
	}
}
