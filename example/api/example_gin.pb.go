// Code generated by protoc-gen-go-gin. DO NOT EDIT.
// versions:
// - protoc-gen-go-gin v1.0.0
// - protoc             v3.21.1
// source: example.proto

package api

import (
	context "context"
	gin "github.com/gin-gonic/gin"
	gincontext "github.com/go-kenka/kratos-gin/gincontext"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = new(gincontext.GinData)
var _ = new(gin.H)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationBlogServiceCreateArticle = "/example.BlogService/CreateArticle"
const OperationBlogServiceGetArticles = "/example.BlogService/GetArticles"

type BlogServiceHTTPServer interface {
	CreateArticle(context.Context, *Article) (*Article, error)
	GetArticles(context.Context, *GetArticlesReq) (*GetArticlesResp, error)
}

func RegisterBlogServiceHTTPServer(r gin.IRouter, srv BlogServiceHTTPServer) {
	r.GET("/v1/author/:author_id/articles", _BlogService_GetArticles0_HTTP_Handler(srv))
	r.GET("/v1/articles", _BlogService_GetArticles1_HTTP_Handler(srv))
	r.POST("/v1/author/:author_id/articles", _BlogService_CreateArticle0_HTTP_Handler(srv))
}

func _BlogService_GetArticles0_HTTP_Handler(srv BlogServiceHTTPServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var in GetArticlesReq
		// query
		if err := ctx.BindQuery(&in); err != nil {
			ctx.Error(err)
			return
		}
		// params
		if err := ctx.BindUri(&in); err != nil {
			ctx.Error(err)
			return
		}
		// header,ip等常用信息, form表单信息,包括上传文件
		newCtx := gincontext.NewContext(ctx)
		reply, err := srv.GetArticles(newCtx, &in)
		if err != nil {
			ctx.Error(err)
			return
		}
		ctx.JSON(200, reply)
	}
}

func _BlogService_GetArticles1_HTTP_Handler(srv BlogServiceHTTPServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var in GetArticlesReq
		// query
		if err := ctx.BindQuery(&in); err != nil {
			ctx.Error(err)
			return
		}
		// header,ip等常用信息, form表单信息,包括上传文件
		newCtx := gincontext.NewContext(ctx)
		reply, err := srv.GetArticles(newCtx, &in)
		if err != nil {
			ctx.Error(err)
			return
		}
		ctx.JSON(200, reply)
	}
}

func _BlogService_CreateArticle0_HTTP_Handler(srv BlogServiceHTTPServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var in Article
		// body
		if err := ctx.BindJSON(&in); err != nil {
			ctx.Error(err)
			return
		}
		// params
		if err := ctx.BindUri(&in); err != nil {
			ctx.Error(err)
			return
		}
		// header,ip等常用信息, form表单信息,包括上传文件
		newCtx := gincontext.NewContext(ctx)
		reply, err := srv.CreateArticle(newCtx, &in)
		if err != nil {
			ctx.Error(err)
			return
		}
		ctx.JSON(200, reply)
	}
}

type BlogServiceHTTPClient interface {
	CreateArticle(ctx context.Context, req *Article, opts ...http.CallOption) (rsp *Article, err error)
	GetArticles(ctx context.Context, req *GetArticlesReq, opts ...http.CallOption) (rsp *GetArticlesResp, err error)
}

type BlogServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewBlogServiceHTTPClient(client *http.Client) BlogServiceHTTPClient {
	return &BlogServiceHTTPClientImpl{client}
}

func (c *BlogServiceHTTPClientImpl) CreateArticle(ctx context.Context, in *Article, opts ...http.CallOption) (*Article, error) {
	var out Article
	pattern := "/v1/author/{author_id}/articles"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBlogServiceCreateArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BlogServiceHTTPClientImpl) GetArticles(ctx context.Context, in *GetArticlesReq, opts ...http.CallOption) (*GetArticlesResp, error) {
	var out GetArticlesResp
	pattern := "/v1/articles"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBlogServiceGetArticles))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}