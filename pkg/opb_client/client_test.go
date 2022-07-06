package opb_client

import (
	"context"
	"fmt"
	"github.com/llygcd/block-compensation/pkg/opb_client/nft"
	"google.golang.org/grpc"
	"testing"
)

func TestName(t *testing.T) {
	conn, err := grpc.Dial("47.100.192.234:9090", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := nft.NewQueryClient(conn)
	denom, err := client.Denom(context.Background(), &nft.QueryDenomRequest{DenomId: "avata8232930613442183168"})
	if err != nil {
		panic(err)
	}

	fmt.Println(denom)
}
