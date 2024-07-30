package services

import (
	"context"
	"sync"

	pb "github.com/didil/simple-text-rag-go-frontend/grpc_gen/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
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
	addr            string
	qaServiceClient pb.QAServiceClient
	conn            *grpc.ClientConn
	mu              sync.Mutex
}

func (c *grpcClient) newQAServiceClient(addr string) (pb.QAServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}
	client := pb.NewQAServiceClient(conn)
	return client, conn, nil
}

func (c *grpcClient) ensureConnected() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn == nil || c.conn.GetState() == connectivity.Shutdown {
		if c.conn != nil {
			c.conn.Close()
		}
		var err error
		c.qaServiceClient, c.conn, err = c.newQAServiceClient(c.addr)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *grpcClient) CreateCollection(ctx context.Context, name string, fileUrl string) error {
	if err := c.ensureConnected(); err != nil {
		return err
	}

	_, err := c.qaServiceClient.CreateCollection(ctx, &pb.CreateCollectionRequest{Name: name, FileUrl: fileUrl})
	if err != nil {
		return err
	}

	return nil
}

func (c *grpcClient) GetAnswer(ctx context.Context, collectionName string, question string) (string, error) {
	if err := c.ensureConnected(); err != nil {
		return "", err
	}

	resp, err := c.qaServiceClient.GetAnswer(ctx, &pb.GetAnswerRequest{CollectionName: collectionName, Question: question})
	if err != nil {
		return "", err
	}

	return resp.Text, nil
}
