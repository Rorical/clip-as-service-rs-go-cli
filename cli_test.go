package main

import (
	"context"
	"testing"
)

func TestClipRsClient_EncodeText(t *testing.T) {
	ctx := context.Background()
	cli, err := NewClipRsClient("localhost:50051")
	if err != nil {
		t.Fatal(err)
	}
	vectors, err := cli.EncodeText(ctx, []string{
		"你好",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(vectors[0]))
	t.Logf("%v", vectors[0])
}

func BenchmarkClipRsClient_EncodeText(b *testing.B) {
	ctx := context.Background()
	cli, err := NewClipRsClient("localhost:50051")
	if err != nil {
		b.Fatal(err)
	}
	texts := []string{
		"资本主义的世界没钱还是别去",
		"MuriData是一个去中心化可索引数据库，通过区块链实现激励和监督，通过零知识证明等设计实现高性能的查询。",
	}
	for i := 0; i < b.N; i++ {
		_, err = cli.EncodeText(ctx, texts)
		if err != nil {
			b.Fatal(err)
		}
	}
}
