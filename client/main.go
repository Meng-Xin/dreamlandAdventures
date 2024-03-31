package client

import (
	"context"
	"dreamlandAdventures/shared"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/filter"
	"trpc.group/trpc-go/trpc-go/log"
)

func clientFilter(ctx context.Context, req interface{}, rsp interface{}, next filter.ClientHandleFunc) error {
	log.InfoContext(ctx, "client filter start")
	// filter start, set token
	msg := trpc.Message(ctx)
	msg.WithClientMetaData(map[string][]byte{shared.AuthKey: []byte(shared.Token)})
	// run business logic
	err := next(ctx, req, rsp)
	// filter end
	log.InfoContext(ctx, "client filter end")
	return err
}
