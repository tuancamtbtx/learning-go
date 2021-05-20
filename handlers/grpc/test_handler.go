package grpc

import (
	"fmt"
	"testing"

	"github.com/tuancamtbtx/learning-go/genproto/go/common/entities"
)

// type TestMapEnity struct {
// 	suite.Suite
// }

// func NewTesting() *TestMapEnity {
// 	return &TestMapEnity{}
// }
func Test(t *testing.T) {
	test := &entities.ItemRank{
		Id:   1,
		Rank: 22,
	}
	arr := make([]*entities.ItemRank, 2)
	arr[0] = test

	list := &entities.ListRankItems{
		Items: arr,
	}

	tmp := &entities.ListRankItems{
		Items: nil,
	}
	if len(tmp.Items) > 0 {
		t.Errorf("got %q, wanted %q", 1, 2)

	}
	fmt.Println(list)
}
func Add(a int, b int) int {
	return a + b
}
func TestAdd(t *testing.T) {

	got := Add(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
