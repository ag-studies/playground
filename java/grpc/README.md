# GRPC App Test

This is an app test for helloworld service gRPC.

## Project Structure

This project has 3 modules:
- client
- server
- shared

The client-project runs grpc sync, and async request.

The server-project runs grpc server.

The shared-project has '.proto' file.

## Running App

Compile and package this project:
```
$ mvn compile packag
```

For run server:
```
$ java -cp shared/target/shared-0.0.1-SNAPSHOT.jar:server/target/server-0.0.1-SNAPSHOT.jar:target/libs/* ifpb.grpc.server.App
```

For run client:
```
$ java -cp shared/target/shared-0.0.1-SNAPSHOT.jar:client/target/client-0.0.1-SNAPSHOT.jar:target/libs/* ag.ifpb.client.App
```

## Referencies

- [Protocol Buffer Basics: Java](https://developers.google.com/protocol-buffers/docs/javatutorial)
- [Java Quickstart for gRPC](https://grpc.io/docs/quickstart/java.html)
- [Introdução a gRPC](http://www.baeldung.com/grpc-introduction)


