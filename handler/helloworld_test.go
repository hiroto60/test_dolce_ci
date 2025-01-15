package handler

import (
	"context"
	"reflect"
	"testing"

	"connectrpc.com/connect"
	greet "github.com/hiroto60/test_dolce_ci/gen/greet"
)

func TestGreetHandler_SayHello(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[greet.HelloRequest]
	}
	tests := []struct {
		name    string
		s       *GreetHandler
		args    args
		want    *connect.Response[greet.HelloResponse]
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "successful response with headers",
			s:    &GreetHandler{},
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&greet.HelloRequest{
					Name: "Alice",
				}),
			},
			want: func() *connect.Response[greet.HelloResponse] {
				resp := connect.NewResponse(&greet.HelloResponse{
					Message: "Hello, Alice!",
				})
				resp.Header().Set("Greet-Version", "v1")
				return resp
			}(),
			wantErr: false,
		},
		{
			name: "empty name in request with headers",
			s:    &GreetHandler{},
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&greet.HelloRequest{
					Name: "",
				}),
			},
			want: func() *connect.Response[greet.HelloResponse] {
				resp := connect.NewResponse(&greet.HelloResponse{
					Message: "Hello, !",
				})
				resp.Header().Set("Greet-Version", "v1")
				return resp
			}(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &GreetHandler{}
			got, err := s.SayHello(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GreetHandler.SayHello() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GreetHandler.SayHello() = %v, want %v", got, tt.want)
			}
		})
	}
}
