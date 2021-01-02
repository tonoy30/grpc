protoc --go_out=plugins=grpc:. chat/chatpb/chat.proto
protoc --go_out=plugins=grpc:. sum/sumpb/sum.proto
protoc --go_out=plugins=grpc:. greet/greetpb/greet.proto