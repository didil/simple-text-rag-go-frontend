package services

import (
	"context"
	"sync"

	pb "github.com/didil/simple-text-rag-go-frontend/grpc_gen/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClient interface {
	CreateCollection(ctx context.Context, name string, fileUrl string) error
	GetAnswer(ctx context.Context, collectionName string, question string) (string, error)
}

func NewGrpcClient(addr string) (GrpcClient, error) {
	return &grpcClient{
		addr: addr,
		mu:   sync.Mutex{},
	}, nil
}

type grpcClient struct {
	addr string
	mu   sync.Mutex
}

func (c *grpcClient) newQAServiceClient() (pb.QAServiceClient, error) {
	conn, err := grpc.NewClient(c.addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewQAServiceClient(conn)
	return client, nil
}

func (c *grpcClient) CreateCollection(ctx context.Context, name string, fileUrl string) error {
	client, err := c.newQAServiceClient()
	if err != nil {
		return err
	}

	_, err = client.CreateCollection(ctx, &pb.CreateCollectionRequest{Name: name, FileUrl: fileUrl})
	if err != nil {
		return err
	}

	return nil
}

func (c *grpcClient) GetAnswer(ctx context.Context, collectionName string, question string) (string, error) {
	client, err := c.newQAServiceClient()
	if err != nil {
		return "", err
	}

	resp, err := client.GetAnswer(ctx, &pb.GetAnswerRequest{CollectionName: collectionName, Question: question})
	if err != nil {
		return "", err
	}

	return resp.Text, nil
}
