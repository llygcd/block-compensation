package opb_client

import (
	"fmt"
	"google.golang.org/grpc"
	"testing"
)

func TestName(t *testing.T) {
	conn, err := grpc.Dial("47.100.192.234:9090", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	Conn = conn

	/*	client := nft.NewQueryClient(conn)
		denom, err := client.Denom(context.Background(), &nft.QueryDenomRequest{DenomId: "avata8232930613442183168"})
		if err != nil {
			panic(err)
		}

		fmt.Println(denom)*/

	/*	owner, err := nft.NewQueryClient(conn).Owner(context.Background(), &nft.QueryOwnerRequest{Owner: "iaa1zdves3x9wqml4rlgqyzyx3er8xlnnjdquzfgqe", DenomId: "iaa1zdves3x9wqml4rlgqyzyx3er8xlnnjdquzfgqe"})
		if err != nil {
			panic(err)
		}
		fmt.Println(owner.Owner.Address)*/

	queryOwner, err := QueryOwner("iaa1lxvmp9h0v0dhzetmhstrmw3ecpplp5tljnr35f", "ycsc23")
	if err != nil {
		panic(err)
	}
	fmt.Println(queryOwner.Owner.Address)
}
