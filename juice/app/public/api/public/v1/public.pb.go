// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: public.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PublicIdentityMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token      string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Timestampe uint64 `protobuf:"varint,3,opt,name=timestampe,proto3" json:"timestampe,omitempty"`
}

func (x *PublicIdentityMsg) Reset() {
	*x = PublicIdentityMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicIdentityMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicIdentityMsg) ProtoMessage() {}

func (x *PublicIdentityMsg) ProtoReflect() protoreflect.Message {
	mi := &file_public_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicIdentityMsg.ProtoReflect.Descriptor instead.
func (*PublicIdentityMsg) Descriptor() ([]byte, []int) {
	return file_public_proto_rawDescGZIP(), []int{0}
}

func (x *PublicIdentityMsg) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PublicIdentityMsg) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *PublicIdentityMsg) GetTimestampe() uint64 {
	if x != nil {
		return x.Timestampe
	}
	return 0
}

type PublicStatusMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code       uint64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg        string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Timestampe uint64 `protobuf:"varint,3,opt,name=timestampe,proto3" json:"timestampe,omitempty"`
}

func (x *PublicStatusMsg) Reset() {
	*x = PublicStatusMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicStatusMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicStatusMsg) ProtoMessage() {}

func (x *PublicStatusMsg) ProtoReflect() protoreflect.Message {
	mi := &file_public_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicStatusMsg.ProtoReflect.Descriptor instead.
func (*PublicStatusMsg) Descriptor() ([]byte, []int) {
	return file_public_proto_rawDescGZIP(), []int{1}
}

func (x *PublicStatusMsg) GetCode() uint64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PublicStatusMsg) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *PublicStatusMsg) GetTimestampe() uint64 {
	if x != nil {
		return x.Timestampe
	}
	return 0
}

type PublicUsers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid           uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	UserName      string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	FollowCount   uint64 `protobuf:"varint,3,opt,name=follow_count,json=followCount,proto3" json:"follow_count,omitempty"`
	FollowerCount uint64 `protobuf:"varint,4,opt,name=follower_count,json=followerCount,proto3" json:"follower_count,omitempty"`
	IsFollow      bool   `protobuf:"varint,5,opt,name=is_follow,json=isFollow,proto3" json:"is_follow,omitempty"` // False for oneself
	UserIntro     string `protobuf:"bytes,6,opt,name=user_intro,json=userIntro,proto3" json:"user_intro,omitempty"`
	HeadImageURL  []byte `protobuf:"bytes,7,opt,name=head_imageURL,json=headImageURL,proto3" json:"head_imageURL,omitempty"`
}

func (x *PublicUsers) Reset() {
	*x = PublicUsers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicUsers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicUsers) ProtoMessage() {}

func (x *PublicUsers) ProtoReflect() protoreflect.Message {
	mi := &file_public_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicUsers.ProtoReflect.Descriptor instead.
func (*PublicUsers) Descriptor() ([]byte, []int) {
	return file_public_proto_rawDescGZIP(), []int{2}
}

func (x *PublicUsers) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *PublicUsers) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *PublicUsers) GetFollowCount() uint64 {
	if x != nil {
		return x.FollowCount
	}
	return 0
}

func (x *PublicUsers) GetFollowerCount() uint64 {
	if x != nil {
		return x.FollowerCount
	}
	return 0
}

func (x *PublicUsers) GetIsFollow() bool {
	if x != nil {
		return x.IsFollow
	}
	return false
}

func (x *PublicUsers) GetUserIntro() string {
	if x != nil {
		return x.UserIntro
	}
	return ""
}

func (x *PublicUsers) GetHeadImageURL() []byte {
	if x != nil {
		return x.HeadImageURL
	}
	return nil
}

type PublicComment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentId  uint64       `protobuf:"varint,1,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
	PcommentId uint64       `protobuf:"varint,2,opt,name=pcomment_id,json=pcommentId,proto3" json:"pcomment_id,omitempty"`
	VideoId    uint64       `protobuf:"varint,3,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
	User       *PublicUsers `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	Content    string       `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	CreateDate string       `protobuf:"bytes,6,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty"`
}

func (x *PublicComment) Reset() {
	*x = PublicComment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicComment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicComment) ProtoMessage() {}

func (x *PublicComment) ProtoReflect() protoreflect.Message {
	mi := &file_public_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicComment.ProtoReflect.Descriptor instead.
func (*PublicComment) Descriptor() ([]byte, []int) {
	return file_public_proto_rawDescGZIP(), []int{3}
}

func (x *PublicComment) GetCommentId() uint64 {
	if x != nil {
		return x.CommentId
	}
	return 0
}

func (x *PublicComment) GetPcommentId() uint64 {
	if x != nil {
		return x.PcommentId
	}
	return 0
}

func (x *PublicComment) GetVideoId() uint64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *PublicComment) GetUser() *PublicUsers {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *PublicComment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *PublicComment) GetCreateDate() string {
	if x != nil {
		return x.CreateDate
	}
	return ""
}

type PublicVideo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vid           uint64       `protobuf:"varint,1,opt,name=vid,proto3" json:"vid,omitempty"`
	User          *PublicUsers `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	PlayUrl       []byte       `protobuf:"bytes,3,opt,name=play_url,json=playUrl,proto3" json:"play_url,omitempty"`
	CoverUrl      []byte       `protobuf:"bytes,4,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url,omitempty"`
	FavoriteCount uint64       `protobuf:"varint,5,opt,name=favorite_count,json=favoriteCount,proto3" json:"favorite_count,omitempty"`
	CommentCount  uint64       `protobuf:"varint,6,opt,name=comment_count,json=commentCount,proto3" json:"comment_count,omitempty"`
	IsFavorite    bool         `protobuf:"varint,7,opt,name=is_favorite,json=isFavorite,proto3" json:"is_favorite,omitempty"`
	Title         string       `protobuf:"bytes,8,opt,name=title,proto3" json:"title,omitempty"`
	PublishDate   uint64       `protobuf:"varint,9,opt,name=publish_date,json=publishDate,proto3" json:"publish_date,omitempty"`
}

func (x *PublicVideo) Reset() {
	*x = PublicVideo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicVideo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicVideo) ProtoMessage() {}

func (x *PublicVideo) ProtoReflect() protoreflect.Message {
	mi := &file_public_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicVideo.ProtoReflect.Descriptor instead.
func (*PublicVideo) Descriptor() ([]byte, []int) {
	return file_public_proto_rawDescGZIP(), []int{4}
}

func (x *PublicVideo) GetVid() uint64 {
	if x != nil {
		return x.Vid
	}
	return 0
}

func (x *PublicVideo) GetUser() *PublicUsers {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *PublicVideo) GetPlayUrl() []byte {
	if x != nil {
		return x.PlayUrl
	}
	return nil
}

func (x *PublicVideo) GetCoverUrl() []byte {
	if x != nil {
		return x.CoverUrl
	}
	return nil
}

func (x *PublicVideo) GetFavoriteCount() uint64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

func (x *PublicVideo) GetCommentCount() uint64 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

func (x *PublicVideo) GetIsFavorite() bool {
	if x != nil {
		return x.IsFavorite
	}
	return false
}

func (x *PublicVideo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PublicVideo) GetPublishDate() uint64 {
	if x != nil {
		return x.PublishDate
	}
	return 0
}

type VerifyUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity *PublicIdentityMsg `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
}

func (x *VerifyUserRequest) Reset() {
	*x = VerifyUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyUserRequest) ProtoMessage() {}

func (x *VerifyUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_public_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyUserRequest.ProtoReflect.Descriptor instead.
func (*VerifyUserRequest) Descriptor() ([]byte, []int) {
	return file_public_proto_rawDescGZIP(), []int{5}
}

func (x *VerifyUserRequest) GetIdentity() *PublicIdentityMsg {
	if x != nil {
		return x.Identity
	}
	return nil
}

type VerifyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity bool `protobuf:"varint,1,opt,name=identity,proto3" json:"identity,omitempty"`
}

func (x *VerifyResponse) Reset() {
	*x = VerifyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyResponse) ProtoMessage() {}

func (x *VerifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_public_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyResponse.ProtoReflect.Descriptor instead.
func (*VerifyResponse) Descriptor() ([]byte, []int) {
	return file_public_proto_rawDescGZIP(), []int{6}
}

func (x *VerifyResponse) GetIdentity() bool {
	if x != nil {
		return x.Identity
	}
	return false
}

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_public_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_public_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserRequest) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *PublicUsers `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_public_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_public_proto_rawDescGZIP(), []int{8}
}

func (x *UserResponse) GetUser() *PublicUsers {
	if x != nil {
		return x.User
	}
	return nil
}

var File_public_proto protoreflect.FileDescriptor

var file_public_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x11, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x73, 0x67, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x65, 0x22, 0x57, 0x0a, 0x0f, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x65, 0x22, 0xe7, 0x01, 0x0a, 0x0b, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x73, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x12, 0x23, 0x0a, 0x0d, 0x68, 0x65, 0x61,
	0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0c, 0x68, 0x65, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x22, 0xd1,
	0x01, 0x0a, 0x0d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x70, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61,
	0x74, 0x65, 0x22, 0xa9, 0x02, 0x0a, 0x0b, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x03, 0x76, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x19, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x46, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x44, 0x61, 0x74, 0x65, 0x22, 0x4d,
	0x0a, 0x11, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x4d, 0x73, 0x67, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x2c, 0x0a,
	0x0e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x22, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22,
	0x3a, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2a, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x32, 0xc2, 0x01, 0x0a, 0x06,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x60, 0x0a, 0x0a, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x76, 0x31,
	0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x56, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22,
	0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x42, 0x23, 0x5a, 0x21, 0x6a, 0x75, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_public_proto_rawDescOnce sync.Once
	file_public_proto_rawDescData = file_public_proto_rawDesc
)

func file_public_proto_rawDescGZIP() []byte {
	file_public_proto_rawDescOnce.Do(func() {
		file_public_proto_rawDescData = protoimpl.X.CompressGZIP(file_public_proto_rawDescData)
	})
	return file_public_proto_rawDescData
}

var file_public_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_public_proto_goTypes = []interface{}{
	(*PublicIdentityMsg)(nil), // 0: public.v1.PublicIdentityMsg
	(*PublicStatusMsg)(nil),   // 1: public.v1.PublicStatusMsg
	(*PublicUsers)(nil),       // 2: public.v1.PublicUsers
	(*PublicComment)(nil),     // 3: public.v1.PublicComment
	(*PublicVideo)(nil),       // 4: public.v1.PublicVideo
	(*VerifyUserRequest)(nil), // 5: public.v1.VerifyUserRequest
	(*VerifyResponse)(nil),    // 6: public.v1.VerifyResponse
	(*GetUserRequest)(nil),    // 7: public.v1.GetUserRequest
	(*UserResponse)(nil),      // 8: public.v1.UserResponse
}
var file_public_proto_depIdxs = []int32{
	2, // 0: public.v1.PublicComment.user:type_name -> public.v1.PublicUsers
	2, // 1: public.v1.PublicVideo.user:type_name -> public.v1.PublicUsers
	0, // 2: public.v1.VerifyUserRequest.identity:type_name -> public.v1.PublicIdentityMsg
	2, // 3: public.v1.UserResponse.user:type_name -> public.v1.PublicUsers
	5, // 4: public.v1.Public.VerifyUser:input_type -> public.v1.VerifyUserRequest
	7, // 5: public.v1.Public.GetUser:input_type -> public.v1.GetUserRequest
	6, // 6: public.v1.Public.VerifyUser:output_type -> public.v1.VerifyResponse
	8, // 7: public.v1.Public.GetUser:output_type -> public.v1.UserResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_public_proto_init() }
func file_public_proto_init() {
	if File_public_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_public_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicIdentityMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_public_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicStatusMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_public_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicUsers); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_public_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicComment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_public_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicVideo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_public_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyUserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_public_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_public_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_public_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_public_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_public_proto_goTypes,
		DependencyIndexes: file_public_proto_depIdxs,
		MessageInfos:      file_public_proto_msgTypes,
	}.Build()
	File_public_proto = out.File
	file_public_proto_rawDesc = nil
	file_public_proto_goTypes = nil
	file_public_proto_depIdxs = nil
}
