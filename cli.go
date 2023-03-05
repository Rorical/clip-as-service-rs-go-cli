package main

import (
	"context"
	pb "github.com/Rorical/clip-as-service-rs-go-cli/encoder"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClipRsClient struct {
	client pb.EncoderClient
}

func NewClipRsClient(uri string) (*ClipRsClient, error) {
	conn, err := grpc.Dial(uri, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	cli := pb.NewEncoderClient(conn)
	return &ClipRsClient{
		client: cli,
	}, nil
}

func (cli *ClipRsClient) EncodeText(ctx context.Context, texts []string) ([][]float32, error) {
	request := &pb.EncodeTextRequest{Texts: texts}
	response, err := cli.client.EncodeText(ctx, request)
	if err != nil {
		return nil, err
	}
	batch := make([][]float32, len(response.GetEmbedding()))
	for i, b := range response.GetEmbedding() {
		batch[i] = b.Point
	}
	return batch, nil
}
