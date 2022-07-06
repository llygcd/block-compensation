package opb

import (
	"fmt"
	"testing"
)

func Test_client(t *testing.T) {
	owner, err := QueryOwner("iaa1zdves3x9wqml4rlgqyzyx3er8xlnnjdquzfgqe", "nft79te5stv68qfqv7holep3i345q", &PageRequest{})
	fmt.Print(err)
	fmt.Print(owner)
}
