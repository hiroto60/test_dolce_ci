package handler

import (
	"context"
	"fmt"
	"log"

	"connectrpc.com/connect"
	"github.com/hiroto60/test_dolce_ci/gen/hoge"
)

type HogeHandler struct{}

func NewHogeHandler() *HogeHandler {
	return &HogeHandler{}
}

func (s *HogeHandler) Hoge(ctx context.Context, req *connect.Request[hoge.HogeRequest]) (*connect.Response[hoge.HogeResponse], error) {
	log.Println("Request headers: ", req.Header())
	// 本来ならここから、別の層のメソッドを呼び出すなどする

	res := connect.NewResponse(&hoge.HogeResponse{
		Message: fmt.Sprintf("Hoge, %s!", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}
