// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             v4.24.4
// source: app/video/api/video/v1/video.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationVideoManageDeleteVideo = "/v1.VideoManage/DeleteVideo"
const OperationVideoManagePublishVideo = "/v1.VideoManage/PublishVideo"

type VideoManageHTTPServer interface {
	// DeleteVideo 删除视频
	DeleteVideo(context.Context, *DeleteVideoRequest) (*DeleteVideoResponse, error)
	// PublishVideo 发布视频
	PublishVideo(context.Context, *PublishVideoRequest) (*PublishVideoResponse, error)
}

func RegisterVideoManageHTTPServer(s *http.Server, srv VideoManageHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/video/manege/publish", _VideoManage_PublishVideo0_HTTP_Handler(srv))
	r.DELETE("/v1/video/manege/delete", _VideoManage_DeleteVideo0_HTTP_Handler(srv))
}

func _VideoManage_PublishVideo0_HTTP_Handler(srv VideoManageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PublishVideoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationVideoManagePublishVideo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PublishVideo(ctx, req.(*PublishVideoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PublishVideoResponse)
		return ctx.Result(200, reply)
	}
}

func _VideoManage_DeleteVideo0_HTTP_Handler(srv VideoManageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteVideoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationVideoManageDeleteVideo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteVideo(ctx, req.(*DeleteVideoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteVideoResponse)
		return ctx.Result(200, reply)
	}
}

type VideoManageHTTPClient interface {
	DeleteVideo(ctx context.Context, req *DeleteVideoRequest, opts ...http.CallOption) (rsp *DeleteVideoResponse, err error)
	PublishVideo(ctx context.Context, req *PublishVideoRequest, opts ...http.CallOption) (rsp *PublishVideoResponse, err error)
}

type VideoManageHTTPClientImpl struct {
	cc *http.Client
}

func NewVideoManageHTTPClient(client *http.Client) VideoManageHTTPClient {
	return &VideoManageHTTPClientImpl{client}
}

func (c *VideoManageHTTPClientImpl) DeleteVideo(ctx context.Context, in *DeleteVideoRequest, opts ...http.CallOption) (*DeleteVideoResponse, error) {
	var out DeleteVideoResponse
	pattern := "/v1/video/manege/delete"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationVideoManageDeleteVideo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *VideoManageHTTPClientImpl) PublishVideo(ctx context.Context, in *PublishVideoRequest, opts ...http.CallOption) (*PublishVideoResponse, error) {
	var out PublishVideoResponse
	pattern := "/v1/video/manege/publish"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationVideoManagePublishVideo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

const OperationVideoInteractDeleteComment = "/v1.VideoInteract/DeleteComment"
const OperationVideoInteractGetCommentsList = "/v1.VideoInteract/GetCommentsList"
const OperationVideoInteractGetLikeList = "/v1.VideoInteract/GetLikeList"
const OperationVideoInteractGetSaveList = "/v1.VideoInteract/GetSaveList"
const OperationVideoInteractPublishComment = "/v1.VideoInteract/PublishComment"
const OperationVideoInteractSetLike = "/v1.VideoInteract/SetLike"
const OperationVideoInteractSetSave = "/v1.VideoInteract/SetSave"

type VideoInteractHTTPServer interface {
	DeleteComment(context.Context, *DeleteCommentRequest) (*DeleteCommentResponse, error)
	GetCommentsList(context.Context, *GetCommentListRequest) (*CommentListResponse, error)
	GetLikeList(context.Context, *GetLikeListRequest) (*LikeListResponse, error)
	GetSaveList(context.Context, *GetSaveListRequest) (*SaveListResponse, error)
	// PublishComment 视频评论
	PublishComment(context.Context, *PublishCommentRequest) (*PublishCommentResponse, error)
	// SetLike 视频点赞
	SetLike(context.Context, *SetLikeRequest) (*LikeResponse, error)
	// SetSave 视频收藏
	SetSave(context.Context, *SetSaveRequest) (*SaveResponse, error)
}

func RegisterVideoInteractHTTPServer(s *http.Server, srv VideoInteractHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/video_interact/like/set", _VideoInteract_SetLike0_HTTP_Handler(srv))
	r.GET("/v1/video_interact/like/list", _VideoInteract_GetLikeList0_HTTP_Handler(srv))
	r.POST("/v1/video_interact/save/set", _VideoInteract_SetSave0_HTTP_Handler(srv))
	r.GET("/v1/video_interact/save/list", _VideoInteract_GetSaveList0_HTTP_Handler(srv))
	r.POST("/v1/video_interact/comment/publish", _VideoInteract_PublishComment0_HTTP_Handler(srv))
	r.DELETE("/v1/video_interact/comment/delete", _VideoInteract_DeleteComment0_HTTP_Handler(srv))
	r.GET("/v1/video_interact/comment/list", _VideoInteract_GetCommentsList0_HTTP_Handler(srv))
}

func _VideoInteract_SetLike0_HTTP_Handler(srv VideoInteractHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetLikeRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationVideoInteractSetLike)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetLike(ctx, req.(*SetLikeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LikeResponse)
		return ctx.Result(200, reply)
	}
}

func _VideoInteract_GetLikeList0_HTTP_Handler(srv VideoInteractHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetLikeListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationVideoInteractGetLikeList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetLikeList(ctx, req.(*GetLikeListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LikeListResponse)
		return ctx.Result(200, reply)
	}
}

func _VideoInteract_SetSave0_HTTP_Handler(srv VideoInteractHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetSaveRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationVideoInteractSetSave)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetSave(ctx, req.(*SetSaveRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SaveResponse)
		return ctx.Result(200, reply)
	}
}

func _VideoInteract_GetSaveList0_HTTP_Handler(srv VideoInteractHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetSaveListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationVideoInteractGetSaveList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSaveList(ctx, req.(*GetSaveListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SaveListResponse)
		return ctx.Result(200, reply)
	}
}

func _VideoInteract_PublishComment0_HTTP_Handler(srv VideoInteractHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PublishCommentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationVideoInteractPublishComment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PublishComment(ctx, req.(*PublishCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PublishCommentResponse)
		return ctx.Result(200, reply)
	}
}

func _VideoInteract_DeleteComment0_HTTP_Handler(srv VideoInteractHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteCommentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationVideoInteractDeleteComment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteComment(ctx, req.(*DeleteCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteCommentResponse)
		return ctx.Result(200, reply)
	}
}

func _VideoInteract_GetCommentsList0_HTTP_Handler(srv VideoInteractHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCommentListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationVideoInteractGetCommentsList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCommentsList(ctx, req.(*GetCommentListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CommentListResponse)
		return ctx.Result(200, reply)
	}
}

type VideoInteractHTTPClient interface {
	DeleteComment(ctx context.Context, req *DeleteCommentRequest, opts ...http.CallOption) (rsp *DeleteCommentResponse, err error)
	GetCommentsList(ctx context.Context, req *GetCommentListRequest, opts ...http.CallOption) (rsp *CommentListResponse, err error)
	GetLikeList(ctx context.Context, req *GetLikeListRequest, opts ...http.CallOption) (rsp *LikeListResponse, err error)
	GetSaveList(ctx context.Context, req *GetSaveListRequest, opts ...http.CallOption) (rsp *SaveListResponse, err error)
	PublishComment(ctx context.Context, req *PublishCommentRequest, opts ...http.CallOption) (rsp *PublishCommentResponse, err error)
	SetLike(ctx context.Context, req *SetLikeRequest, opts ...http.CallOption) (rsp *LikeResponse, err error)
	SetSave(ctx context.Context, req *SetSaveRequest, opts ...http.CallOption) (rsp *SaveResponse, err error)
}

type VideoInteractHTTPClientImpl struct {
	cc *http.Client
}

func NewVideoInteractHTTPClient(client *http.Client) VideoInteractHTTPClient {
	return &VideoInteractHTTPClientImpl{client}
}

func (c *VideoInteractHTTPClientImpl) DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...http.CallOption) (*DeleteCommentResponse, error) {
	var out DeleteCommentResponse
	pattern := "/v1/video_interact/comment/delete"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationVideoInteractDeleteComment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *VideoInteractHTTPClientImpl) GetCommentsList(ctx context.Context, in *GetCommentListRequest, opts ...http.CallOption) (*CommentListResponse, error) {
	var out CommentListResponse
	pattern := "/v1/video_interact/comment/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationVideoInteractGetCommentsList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *VideoInteractHTTPClientImpl) GetLikeList(ctx context.Context, in *GetLikeListRequest, opts ...http.CallOption) (*LikeListResponse, error) {
	var out LikeListResponse
	pattern := "/v1/video_interact/like/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationVideoInteractGetLikeList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *VideoInteractHTTPClientImpl) GetSaveList(ctx context.Context, in *GetSaveListRequest, opts ...http.CallOption) (*SaveListResponse, error) {
	var out SaveListResponse
	pattern := "/v1/video_interact/save/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationVideoInteractGetSaveList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *VideoInteractHTTPClientImpl) PublishComment(ctx context.Context, in *PublishCommentRequest, opts ...http.CallOption) (*PublishCommentResponse, error) {
	var out PublishCommentResponse
	pattern := "/v1/video_interact/comment/publish"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationVideoInteractPublishComment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *VideoInteractHTTPClientImpl) SetLike(ctx context.Context, in *SetLikeRequest, opts ...http.CallOption) (*LikeResponse, error) {
	var out LikeResponse
	pattern := "/v1/video_interact/like/set"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationVideoInteractSetLike))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *VideoInteractHTTPClientImpl) SetSave(ctx context.Context, in *SetSaveRequest, opts ...http.CallOption) (*SaveResponse, error) {
	var out SaveResponse
	pattern := "/v1/video_interact/save/set"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationVideoInteractSetSave))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

const OperationVideoFeedClassifyVideoList = "/v1.VideoFeed/ClassifyVideoList"
const OperationVideoFeedGetUserVideoList = "/v1.VideoFeed/GetUserVideoList"
const OperationVideoFeedGetVideoList = "/v1.VideoFeed/GetVideoList"
const OperationVideoFeedGetVideoStream = "/v1.VideoFeed/GetVideoStream"
const OperationVideoFeedSearchVideoList = "/v1.VideoFeed/SearchVideoList"

type VideoFeedHTTPServer interface {
	ClassifyVideoList(context.Context, *ClassifyVideoListRequest) (*VideoListResponse, error)
	GetUserVideoList(context.Context, *GetUserVideoListRequest) (*VideoListResponse, error)
	GetVideoList(context.Context, *GetVideoListRequest) (*VideoListResponse, error)
	GetVideoStream(context.Context, *GetVideoStreamRequest) (*VideoStreamResponse, error)
	SearchVideoList(context.Context, *SearchVideoListRequest) (*VideoListResponse, error)
}

func RegisterVideoFeedHTTPServer(s *http.Server, srv VideoFeedHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/video_feed/list", _VideoFeed_GetVideoList0_HTTP_Handler(srv))
	r.GET("/v1/video_feed/list", _VideoFeed_GetUserVideoList0_HTTP_Handler(srv))
	r.GET("/v1/video_feed/info", _VideoFeed_SearchVideoList0_HTTP_Handler(srv))
	r.GET("/v1/video_feed/video/classified_list", _VideoFeed_ClassifyVideoList0_HTTP_Handler(srv))
	r.GET("/v1/video_feed/feed", _VideoFeed_GetVideoStream0_HTTP_Handler(srv))
}

func _VideoFeed_GetVideoList0_HTTP_Handler(srv VideoFeedHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetVideoListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationVideoFeedGetVideoList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetVideoList(ctx, req.(*GetVideoListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*VideoListResponse)
		return ctx.Result(200, reply)
	}
}

func _VideoFeed_GetUserVideoList0_HTTP_Handler(srv VideoFeedHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserVideoListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationVideoFeedGetUserVideoList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserVideoList(ctx, req.(*GetUserVideoListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*VideoListResponse)
		return ctx.Result(200, reply)
	}
}

func _VideoFeed_SearchVideoList0_HTTP_Handler(srv VideoFeedHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SearchVideoListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationVideoFeedSearchVideoList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SearchVideoList(ctx, req.(*SearchVideoListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*VideoListResponse)
		return ctx.Result(200, reply)
	}
}

func _VideoFeed_ClassifyVideoList0_HTTP_Handler(srv VideoFeedHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ClassifyVideoListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationVideoFeedClassifyVideoList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ClassifyVideoList(ctx, req.(*ClassifyVideoListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*VideoListResponse)
		return ctx.Result(200, reply)
	}
}

func _VideoFeed_GetVideoStream0_HTTP_Handler(srv VideoFeedHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetVideoStreamRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationVideoFeedGetVideoStream)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetVideoStream(ctx, req.(*GetVideoStreamRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*VideoStreamResponse)
		return ctx.Result(200, reply)
	}
}

type VideoFeedHTTPClient interface {
	ClassifyVideoList(ctx context.Context, req *ClassifyVideoListRequest, opts ...http.CallOption) (rsp *VideoListResponse, err error)
	GetUserVideoList(ctx context.Context, req *GetUserVideoListRequest, opts ...http.CallOption) (rsp *VideoListResponse, err error)
	GetVideoList(ctx context.Context, req *GetVideoListRequest, opts ...http.CallOption) (rsp *VideoListResponse, err error)
	GetVideoStream(ctx context.Context, req *GetVideoStreamRequest, opts ...http.CallOption) (rsp *VideoStreamResponse, err error)
	SearchVideoList(ctx context.Context, req *SearchVideoListRequest, opts ...http.CallOption) (rsp *VideoListResponse, err error)
}

type VideoFeedHTTPClientImpl struct {
	cc *http.Client
}

func NewVideoFeedHTTPClient(client *http.Client) VideoFeedHTTPClient {
	return &VideoFeedHTTPClientImpl{client}
}

func (c *VideoFeedHTTPClientImpl) ClassifyVideoList(ctx context.Context, in *ClassifyVideoListRequest, opts ...http.CallOption) (*VideoListResponse, error) {
	var out VideoListResponse
	pattern := "/v1/video_feed/video/classified_list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationVideoFeedClassifyVideoList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *VideoFeedHTTPClientImpl) GetUserVideoList(ctx context.Context, in *GetUserVideoListRequest, opts ...http.CallOption) (*VideoListResponse, error) {
	var out VideoListResponse
	pattern := "/v1/video_feed/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationVideoFeedGetUserVideoList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *VideoFeedHTTPClientImpl) GetVideoList(ctx context.Context, in *GetVideoListRequest, opts ...http.CallOption) (*VideoListResponse, error) {
	var out VideoListResponse
	pattern := "/v1/video_feed/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationVideoFeedGetVideoList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *VideoFeedHTTPClientImpl) GetVideoStream(ctx context.Context, in *GetVideoStreamRequest, opts ...http.CallOption) (*VideoStreamResponse, error) {
	var out VideoStreamResponse
	pattern := "/v1/video_feed/feed"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationVideoFeedGetVideoStream))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *VideoFeedHTTPClientImpl) SearchVideoList(ctx context.Context, in *SearchVideoListRequest, opts ...http.CallOption) (*VideoListResponse, error) {
	var out VideoListResponse
	pattern := "/v1/video_feed/info"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationVideoFeedSearchVideoList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}