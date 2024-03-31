package main

import (
	"context"
	"dreamlandAdventures/shared"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/errs"
	"trpc.group/trpc-go/trpc-go/filter"
	"trpc.group/trpc-go/trpc-go/log"
)

func serverFilter(ctx context.Context, req interface{}, next filter.ServerHandleFunc) (rsp interface{}, err error) {
	log.InfoContext(ctx, "server filter start %s", trpc.GetMetaData(ctx, shared.AuthKey))
	// check token from context metadata
	if !valid(trpc.GetMetaData(ctx, shared.AuthKey)) {
		return nil, errs.Newf(errs.RetServerAuthFail, "auth fail")
	}
	// run business logic
	rsp, err = next(ctx, req)

	log.InfoContext(ctx, "server filter end")
	return rsp, err
}

// valid validates the authorization
func valid(authorization []byte) bool {
	return string(authorization) == shared.Token
}
