---
title: juice v1.0.0
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.17"

---

# juice

> v1.0.0

Base URLs:

# Authentication

# UserBasic

<a id="opIdUserBasic_GetUserInfo"></a>

## GET 获取用户信息

GET /v1/user_basic/info

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_identity.user_id|query|string| 否 |none|
|user_identity.token|query|string| 否 |none|
|user_identity.timestampe|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "status": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  },
  "user": {
    "uid": "string",
    "user_name": "string",
    "follow_count": "string",
    "follower_count": "string",
    "is_follow": true,
    "user_intro": "string",
    "head_imageURL": "string"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1UserInfoResponse](#schemav1userinforesponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdUserBasic_UserLogin"></a>

## POST 用户登陆

POST /v1/user_basic/login

> 返回示例

> 200 Response

```json
{
  "status": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  },
  "IdentityMsg": {
    "user_id": "string",
    "token": "string",
    "timestampe": "string"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1UserLoginResponse](#schemav1userloginresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdUserBasic_UserRegister"></a>

## POST 用户注册

POST /v1/user_basic/register

> 返回示例

> 200 Response

```json
{
  "status": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  },
  "IdentityMsg": {
    "user_id": "string",
    "token": "string",
    "timestampe": "string"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1UserRegisterResponse](#schemav1userregisterresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdUserBasic_SearchUserList"></a>

## GET 搜索用户

GET /v1/user_basic/search

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_identity.user_id|query|string| 否 |none|
|user_identity.token|query|string| 否 |none|
|user_identity.timestampe|query|string| 否 |none|
|users_keywords|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "status": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  },
  "user_list": [
    {
      "uid": "string",
      "user_name": "string",
      "follow_count": "string",
      "follower_count": "string",
      "is_follow": true,
      "user_intro": "string",
      "head_imageURL": "string"
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1UserListResponse](#schemav1userlistresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdUserBasic_UpdateUserInfo"></a>

## POST 用户信息更新接口

POST /v1/user_basic/update

> 返回示例

> 200 Response

```json
{
  "status": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1UserUpdateInfoResponse](#schemav1userupdateinforesponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

# UserRelation

<a id="opIdUserRelation_GetFollowedList"></a>

## GET 关注列表

GET /v1/user_relation/followed

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_identity.user_id|query|string| 否 |none|
|user_identity.token|query|string| 否 |none|
|user_identity.timestampe|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "msg": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  },
  "user_list": [
    {
      "uid": "string",
      "user_name": "string",
      "follow_count": "string",
      "follower_count": "string",
      "is_follow": true,
      "user_intro": "string",
      "head_imageURL": "string"
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1FollowedListResponse](#schemav1followedlistresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdUserRelation_GetFollowerList"></a>

## GET 粉丝列表

GET /v1/user_relation/follower

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_identity.user_id|query|string| 否 |none|
|user_identity.token|query|string| 否 |none|
|user_identity.timestampe|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "msg": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  },
  "user_list": [
    {
      "uid": "string",
      "user_name": "string",
      "follow_count": "string",
      "follower_count": "string",
      "is_follow": true,
      "user_intro": "string",
      "head_imageURL": "string"
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1FollowerListResponse](#schemav1followerlistresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdUserRelation_SetFollow"></a>

## POST 关注（互动）

POST /v1/user_relation/set

> 返回示例

> 200 Response

```json
{
  "msg": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1SetFollowResponse](#schemav1setfollowresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

# VideoManage

<a id="opIdVideoManage_DeleteVideo"></a>

## DELETE 删除视频

DELETE /v1/video/manege/delete

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_identity.user_id|query|string| 否 |none|
|user_identity.token|query|string| 否 |none|
|user_identity.timestampe|query|string| 否 |none|
|video_id|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "msg": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1DeleteVideoResponse](#schemav1deletevideoresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdVideoManage_PublishVideo"></a>

## POST 发布视频

POST /v1/video/manege/publish

> 返回示例

> 200 Response

```json
{
  "msg": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1PublishVideoResponse](#schemav1publishvideoresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

# VideoFeed

<a id="opIdVideoFeed_GetVideoStream"></a>

## GET /v1/video_feed/feed

GET /v1/video_feed/feed

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_identity.user_id|query|string| 否 |none|
|user_identity.token|query|string| 否 |none|
|user_identity.timestampe|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "msg": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  },
  "video_list": [
    {
      "vid": "string",
      "user": {
        "uid": "string",
        "user_name": "string",
        "follow_count": "string",
        "follower_count": "string",
        "is_follow": true,
        "user_intro": "string",
        "head_imageURL": "string"
      },
      "play_url": "string",
      "cover_url": "string",
      "favorite_count": "string",
      "comment_count": "string",
      "is_favorite": true,
      "title": "string",
      "publish_date": "string"
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1VideoStreamResponse](#schemav1videostreamresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdVideoFeed_SearchVideoList"></a>

## GET /v1/video_feed/info

GET /v1/video_feed/info

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_identity.user_id|query|string| 否 |none|
|user_identity.token|query|string| 否 |none|
|user_identity.timestampe|query|string| 否 |none|
|video_keywords|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "msg": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  },
  "video_list": [
    {
      "vid": "string",
      "user": {
        "uid": "string",
        "user_name": "string",
        "follow_count": "string",
        "follower_count": "string",
        "is_follow": true,
        "user_intro": "string",
        "head_imageURL": "string"
      },
      "play_url": "string",
      "cover_url": "string",
      "favorite_count": "string",
      "comment_count": "string",
      "is_favorite": true,
      "title": "string",
      "publish_date": "string"
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1VideoListResponse](#schemav1videolistresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdVideoFeed_GetUserVideoList"></a>

## GET /v1/video_feed/list

GET /v1/video_feed/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_identity.user_id|query|string| 否 |none|
|user_identity.token|query|string| 否 |none|
|user_identity.timestampe|query|string| 否 |none|
|user_id|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "msg": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  },
  "video_list": [
    {
      "vid": "string",
      "user": {
        "uid": "string",
        "user_name": "string",
        "follow_count": "string",
        "follower_count": "string",
        "is_follow": true,
        "user_intro": "string",
        "head_imageURL": "string"
      },
      "play_url": "string",
      "cover_url": "string",
      "favorite_count": "string",
      "comment_count": "string",
      "is_favorite": true,
      "title": "string",
      "publish_date": "string"
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1VideoListResponse](#schemav1videolistresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdVideoFeed_ClassifyVideoList"></a>

## GET /v1/video_feed/video/classified_list

GET /v1/video_feed/video/classified_list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_identity.user_id|query|string| 否 |none|
|user_identity.token|query|string| 否 |none|
|user_identity.timestampe|query|string| 否 |none|
|video_type|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "msg": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  },
  "video_list": [
    {
      "vid": "string",
      "user": {
        "uid": "string",
        "user_name": "string",
        "follow_count": "string",
        "follower_count": "string",
        "is_follow": true,
        "user_intro": "string",
        "head_imageURL": "string"
      },
      "play_url": "string",
      "cover_url": "string",
      "favorite_count": "string",
      "comment_count": "string",
      "is_favorite": true,
      "title": "string",
      "publish_date": "string"
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1VideoListResponse](#schemav1videolistresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

# VideoInteract

<a id="opIdVideoInteract_DeleteComment"></a>

## DELETE /v1/video_interact/comment/delete

DELETE /v1/video_interact/comment/delete

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_identity.user_id|query|string| 否 |none|
|user_identity.token|query|string| 否 |none|
|user_identity.timestampe|query|string| 否 |none|
|cid|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "msg": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1DeleteCommentResponse](#schemav1deletecommentresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdVideoInteract_GetCommentsList"></a>

## GET /v1/video_interact/comment/list

GET /v1/video_interact/comment/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_identity.user_id|query|string| 否 |none|
|user_identity.token|query|string| 否 |none|
|user_identity.timestampe|query|string| 否 |none|
|vid|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "msg": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  },
  "comment_list": [
    {
      "comment_id": "string",
      "pcomment_id": "string",
      "video_id": "string",
      "user": {
        "uid": "string",
        "user_name": "string",
        "follow_count": "string",
        "follower_count": "string",
        "is_follow": true,
        "user_intro": "string",
        "head_imageURL": "string"
      },
      "content": "string",
      "create_date": "string"
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1CommentListResponse](#schemav1commentlistresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdVideoInteract_PublishComment"></a>

## POST 视频评论

POST /v1/video_interact/comment/publish

> 返回示例

> 200 Response

```json
{
  "msg": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1PublishCommentResponse](#schemav1publishcommentresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdVideoInteract_GetLikeList"></a>

## GET /v1/video_interact/like/list

GET /v1/video_interact/like/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_identity.user_id|query|string| 否 |none|
|user_identity.token|query|string| 否 |none|
|user_identity.timestampe|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "status_msg": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  },
  "video_list": [
    {
      "vid": "string",
      "user": {
        "uid": "string",
        "user_name": "string",
        "follow_count": "string",
        "follower_count": "string",
        "is_follow": true,
        "user_intro": "string",
        "head_imageURL": "string"
      },
      "play_url": "string",
      "cover_url": "string",
      "favorite_count": "string",
      "comment_count": "string",
      "is_favorite": true,
      "title": "string",
      "publish_date": "string"
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1LikeListResponse](#schemav1likelistresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdVideoInteract_SetLike"></a>

## POST 视频点赞

POST /v1/video_interact/like/set

> 返回示例

> 200 Response

```json
{
  "status_msg": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1LikeResponse](#schemav1likeresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdVideoInteract_GetSaveList"></a>

## GET /v1/video_interact/save/list

GET /v1/video_interact/save/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_identity.user_id|query|string| 否 |none|
|user_identity.token|query|string| 否 |none|
|user_identity.timestampe|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "status_msg": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  },
  "video_list": [
    {
      "vid": "string",
      "user": {
        "uid": "string",
        "user_name": "string",
        "follow_count": "string",
        "follower_count": "string",
        "is_follow": true,
        "user_intro": "string",
        "head_imageURL": "string"
      },
      "play_url": "string",
      "cover_url": "string",
      "favorite_count": "string",
      "comment_count": "string",
      "is_favorite": true,
      "title": "string",
      "publish_date": "string"
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1SaveListResponse](#schemav1savelistresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdVideoInteract_SetSave"></a>

## POST 视频收藏

POST /v1/video_interact/save/set

> 返回示例

> 200 Response

```json
{
  "status_msg": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1SaveResponse](#schemav1saveresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

# 数据模型

<h2 id="tocS_v1SaveResponse">v1SaveResponse</h2>

<a id="schemav1saveresponse"></a>
<a id="schema_v1SaveResponse"></a>
<a id="tocSv1saveresponse"></a>
<a id="tocsv1saveresponse"></a>

```json
{
  "status_msg": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  }
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|status_msg|[v1StatusMsg](#schemav1statusmsg)|false|none||none|

<h2 id="tocS_v1StatusMsg">v1StatusMsg</h2>

<a id="schemav1statusmsg"></a>
<a id="schema_v1StatusMsg"></a>
<a id="tocSv1statusmsg"></a>
<a id="tocsv1statusmsg"></a>

```json
{
  "code": "string",
  "msg": "string",
  "timestampe": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|code|string(uint64)|false|none||none|
|msg|string|false|none||none|
|timestampe|string(uint64)|false|none||none|

<h2 id="tocS_v1SaveListResponse">v1SaveListResponse</h2>

<a id="schemav1savelistresponse"></a>
<a id="schema_v1SaveListResponse"></a>
<a id="tocSv1savelistresponse"></a>
<a id="tocsv1savelistresponse"></a>

```json
{
  "status_msg": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  },
  "video_list": [
    {
      "vid": "string",
      "user": {
        "uid": "string",
        "user_name": "string",
        "follow_count": "string",
        "follower_count": "string",
        "is_follow": true,
        "user_intro": "string",
        "head_imageURL": "string"
      },
      "play_url": "string",
      "cover_url": "string",
      "favorite_count": "string",
      "comment_count": "string",
      "is_favorite": true,
      "title": "string",
      "publish_date": "string"
    }
  ]
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|status_msg|[v1StatusMsg](#schemav1statusmsg)|false|none||none|
|video_list|[[v1Video](#schemav1video)]|false|none||none|

<h2 id="tocS_v1Video">v1Video</h2>

<a id="schemav1video"></a>
<a id="schema_v1Video"></a>
<a id="tocSv1video"></a>
<a id="tocsv1video"></a>

```json
{
  "vid": "string",
  "user": {
    "uid": "string",
    "user_name": "string",
    "follow_count": "string",
    "follower_count": "string",
    "is_follow": true,
    "user_intro": "string",
    "head_imageURL": "string"
  },
  "play_url": "string",
  "cover_url": "string",
  "favorite_count": "string",
  "comment_count": "string",
  "is_favorite": true,
  "title": "string",
  "publish_date": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|vid|string(uint64)|false|none||none|
|user|[v1Users](#schemav1users)|false|none||none|
|play_url|string(byte)|false|none||none|
|cover_url|string(byte)|false|none||none|
|favorite_count|string(uint64)|false|none||none|
|comment_count|string(uint64)|false|none||none|
|is_favorite|boolean|false|none||none|
|title|string|false|none||none|
|publish_date|string(uint64)|false|none||none|

<h2 id="tocS_v1Users">v1Users</h2>

<a id="schemav1users"></a>
<a id="schema_v1Users"></a>
<a id="tocSv1users"></a>
<a id="tocsv1users"></a>

```json
{
  "uid": "string",
  "user_name": "string",
  "follow_count": "string",
  "follower_count": "string",
  "is_follow": true,
  "user_intro": "string",
  "head_imageURL": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|uid|string(uint64)|false|none||none|
|user_name|string|false|none||none|
|follow_count|string(uint64)|false|none||none|
|follower_count|string(uint64)|false|none||none|
|is_follow|boolean|false|none||none|
|user_intro|string|false|none||none|
|head_imageURL|string(byte)|false|none||none|

<h2 id="tocS_v1LikeResponse">v1LikeResponse</h2>

<a id="schemav1likeresponse"></a>
<a id="schema_v1LikeResponse"></a>
<a id="tocSv1likeresponse"></a>
<a id="tocsv1likeresponse"></a>

```json
{
  "status_msg": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  }
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|status_msg|[v1StatusMsg](#schemav1statusmsg)|false|none||none|

<h2 id="tocS_v1UserUpdateInfoResponse">v1UserUpdateInfoResponse</h2>

<a id="schemav1userupdateinforesponse"></a>
<a id="schema_v1UserUpdateInfoResponse"></a>
<a id="tocSv1userupdateinforesponse"></a>
<a id="tocsv1userupdateinforesponse"></a>

```json
{
  "status": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  }
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|status|[v1StatusMsg](#schemav1statusmsg)|false|none||none|

<h2 id="tocS_v1LikeListResponse">v1LikeListResponse</h2>

<a id="schemav1likelistresponse"></a>
<a id="schema_v1LikeListResponse"></a>
<a id="tocSv1likelistresponse"></a>
<a id="tocsv1likelistresponse"></a>

```json
{
  "status_msg": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  },
  "video_list": [
    {
      "vid": "string",
      "user": {
        "uid": "string",
        "user_name": "string",
        "follow_count": "string",
        "follower_count": "string",
        "is_follow": true,
        "user_intro": "string",
        "head_imageURL": "string"
      },
      "play_url": "string",
      "cover_url": "string",
      "favorite_count": "string",
      "comment_count": "string",
      "is_favorite": true,
      "title": "string",
      "publish_date": "string"
    }
  ]
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|status_msg|[v1StatusMsg](#schemav1statusmsg)|false|none||none|
|video_list|[[v1Video](#schemav1video)]|false|none||none|

<h2 id="tocS_v1UserListResponse">v1UserListResponse</h2>

<a id="schemav1userlistresponse"></a>
<a id="schema_v1UserListResponse"></a>
<a id="tocSv1userlistresponse"></a>
<a id="tocsv1userlistresponse"></a>

```json
{
  "status": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  },
  "user_list": [
    {
      "uid": "string",
      "user_name": "string",
      "follow_count": "string",
      "follower_count": "string",
      "is_follow": true,
      "user_intro": "string",
      "head_imageURL": "string"
    }
  ]
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|status|[v1StatusMsg](#schemav1statusmsg)|false|none||none|
|user_list|[[v1Users](#schemav1users)]|false|none||none|

<h2 id="tocS_v1PublishCommentResponse">v1PublishCommentResponse</h2>

<a id="schemav1publishcommentresponse"></a>
<a id="schema_v1PublishCommentResponse"></a>
<a id="tocSv1publishcommentresponse"></a>
<a id="tocsv1publishcommentresponse"></a>

```json
{
  "msg": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  }
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|msg|[v1StatusMsg](#schemav1statusmsg)|false|none||none|

<h2 id="tocS_v1SetFollowResponse">v1SetFollowResponse</h2>

<a id="schemav1setfollowresponse"></a>
<a id="schema_v1SetFollowResponse"></a>
<a id="tocSv1setfollowresponse"></a>
<a id="tocsv1setfollowresponse"></a>

```json
{
  "msg": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  }
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|msg|[v1StatusMsg](#schemav1statusmsg)|false|none||none|

<h2 id="tocS_v1UserRegisterResponse">v1UserRegisterResponse</h2>

<a id="schemav1userregisterresponse"></a>
<a id="schema_v1UserRegisterResponse"></a>
<a id="tocSv1userregisterresponse"></a>
<a id="tocsv1userregisterresponse"></a>

```json
{
  "status": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  },
  "IdentityMsg": {
    "user_id": "string",
    "token": "string",
    "timestampe": "string"
  }
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|status|[v1StatusMsg](#schemav1statusmsg)|false|none||none|
|IdentityMsg|[v1IdentityMsg](#schemav1identitymsg)|false|none||none|

<h2 id="tocS_v1IdentityMsg">v1IdentityMsg</h2>

<a id="schemav1identitymsg"></a>
<a id="schema_v1IdentityMsg"></a>
<a id="tocSv1identitymsg"></a>
<a id="tocsv1identitymsg"></a>

```json
{
  "user_id": "string",
  "token": "string",
  "timestampe": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|user_id|string(uint64)|false|none||none|
|token|string|false|none||none|
|timestampe|string(uint64)|false|none||none|

<h2 id="tocS_v1CommentListResponse">v1CommentListResponse</h2>

<a id="schemav1commentlistresponse"></a>
<a id="schema_v1CommentListResponse"></a>
<a id="tocSv1commentlistresponse"></a>
<a id="tocsv1commentlistresponse"></a>

```json
{
  "msg": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  },
  "comment_list": [
    {
      "comment_id": "string",
      "pcomment_id": "string",
      "video_id": "string",
      "user": {
        "uid": "string",
        "user_name": "string",
        "follow_count": "string",
        "follower_count": "string",
        "is_follow": true,
        "user_intro": "string",
        "head_imageURL": "string"
      },
      "content": "string",
      "create_date": "string"
    }
  ]
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|msg|[v1StatusMsg](#schemav1statusmsg)|false|none||none|
|comment_list|[[v1Comment](#schemav1comment)]|false|none||none|

<h2 id="tocS_v1Comment">v1Comment</h2>

<a id="schemav1comment"></a>
<a id="schema_v1Comment"></a>
<a id="tocSv1comment"></a>
<a id="tocsv1comment"></a>

```json
{
  "comment_id": "string",
  "pcomment_id": "string",
  "video_id": "string",
  "user": {
    "uid": "string",
    "user_name": "string",
    "follow_count": "string",
    "follower_count": "string",
    "is_follow": true,
    "user_intro": "string",
    "head_imageURL": "string"
  },
  "content": "string",
  "create_date": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|comment_id|string(uint64)|false|none||none|
|pcomment_id|string(uint64)|false|none||none|
|video_id|string(uint64)|false|none||none|
|user|[v1Users](#schemav1users)|false|none||none|
|content|string|false|none||none|
|create_date|string|false|none||none|

<h2 id="tocS_v1VideoListResponse">v1VideoListResponse</h2>

<a id="schemav1videolistresponse"></a>
<a id="schema_v1VideoListResponse"></a>
<a id="tocSv1videolistresponse"></a>
<a id="tocsv1videolistresponse"></a>

```json
{
  "msg": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  },
  "video_list": [
    {
      "vid": "string",
      "user": {
        "uid": "string",
        "user_name": "string",
        "follow_count": "string",
        "follower_count": "string",
        "is_follow": true,
        "user_intro": "string",
        "head_imageURL": "string"
      },
      "play_url": "string",
      "cover_url": "string",
      "favorite_count": "string",
      "comment_count": "string",
      "is_favorite": true,
      "title": "string",
      "publish_date": "string"
    }
  ]
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|msg|[v1StatusMsg](#schemav1statusmsg)|false|none||none|
|video_list|[[v1Video](#schemav1video)]|false|none||none|

<h2 id="tocS_v1PublishVideoResponse">v1PublishVideoResponse</h2>

<a id="schemav1publishvideoresponse"></a>
<a id="schema_v1PublishVideoResponse"></a>
<a id="tocSv1publishvideoresponse"></a>
<a id="tocsv1publishvideoresponse"></a>

```json
{
  "msg": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  }
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|msg|[v1StatusMsg](#schemav1statusmsg)|false|none||none|

<h2 id="tocS_v1FollowerListResponse">v1FollowerListResponse</h2>

<a id="schemav1followerlistresponse"></a>
<a id="schema_v1FollowerListResponse"></a>
<a id="tocSv1followerlistresponse"></a>
<a id="tocsv1followerlistresponse"></a>

```json
{
  "msg": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  },
  "user_list": [
    {
      "uid": "string",
      "user_name": "string",
      "follow_count": "string",
      "follower_count": "string",
      "is_follow": true,
      "user_intro": "string",
      "head_imageURL": "string"
    }
  ]
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|msg|[v1StatusMsg](#schemav1statusmsg)|false|none||none|
|user_list|[[v1Users](#schemav1users)]|false|none||none|

<h2 id="tocS_v1UserLoginResponse">v1UserLoginResponse</h2>

<a id="schemav1userloginresponse"></a>
<a id="schema_v1UserLoginResponse"></a>
<a id="tocSv1userloginresponse"></a>
<a id="tocsv1userloginresponse"></a>

```json
{
  "status": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  },
  "IdentityMsg": {
    "user_id": "string",
    "token": "string",
    "timestampe": "string"
  }
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|status|[v1StatusMsg](#schemav1statusmsg)|false|none||none|
|IdentityMsg|[v1IdentityMsg](#schemav1identitymsg)|false|none||none|

<h2 id="tocS_v1DeleteCommentResponse">v1DeleteCommentResponse</h2>

<a id="schemav1deletecommentresponse"></a>
<a id="schema_v1DeleteCommentResponse"></a>
<a id="tocSv1deletecommentresponse"></a>
<a id="tocsv1deletecommentresponse"></a>

```json
{
  "msg": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  }
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|msg|[v1StatusMsg](#schemav1statusmsg)|false|none||none|

<h2 id="tocS_v1VideoStreamResponse">v1VideoStreamResponse</h2>

<a id="schemav1videostreamresponse"></a>
<a id="schema_v1VideoStreamResponse"></a>
<a id="tocSv1videostreamresponse"></a>
<a id="tocsv1videostreamresponse"></a>

```json
{
  "msg": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  },
  "video_list": [
    {
      "vid": "string",
      "user": {
        "uid": "string",
        "user_name": "string",
        "follow_count": "string",
        "follower_count": "string",
        "is_follow": true,
        "user_intro": "string",
        "head_imageURL": "string"
      },
      "play_url": "string",
      "cover_url": "string",
      "favorite_count": "string",
      "comment_count": "string",
      "is_favorite": true,
      "title": "string",
      "publish_date": "string"
    }
  ]
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|msg|[v1StatusMsg](#schemav1statusmsg)|false|none||none|
|video_list|[[v1Video](#schemav1video)]|false|none||none|

<h2 id="tocS_v1DeleteVideoResponse">v1DeleteVideoResponse</h2>

<a id="schemav1deletevideoresponse"></a>
<a id="schema_v1DeleteVideoResponse"></a>
<a id="tocSv1deletevideoresponse"></a>
<a id="tocsv1deletevideoresponse"></a>

```json
{
  "msg": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  }
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|msg|[v1StatusMsg](#schemav1statusmsg)|false|none||none|

<h2 id="tocS_v1FollowedListResponse">v1FollowedListResponse</h2>

<a id="schemav1followedlistresponse"></a>
<a id="schema_v1FollowedListResponse"></a>
<a id="tocSv1followedlistresponse"></a>
<a id="tocsv1followedlistresponse"></a>

```json
{
  "msg": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  },
  "user_list": [
    {
      "uid": "string",
      "user_name": "string",
      "follow_count": "string",
      "follower_count": "string",
      "is_follow": true,
      "user_intro": "string",
      "head_imageURL": "string"
    }
  ]
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|msg|[v1StatusMsg](#schemav1statusmsg)|false|none||none|
|user_list|[[v1Users](#schemav1users)]|false|none||none|

<h2 id="tocS_runtimeError">runtimeError</h2>

<a id="schemaruntimeerror"></a>
<a id="schema_runtimeError"></a>
<a id="tocSruntimeerror"></a>
<a id="tocsruntimeerror"></a>

```json
{
  "error": "string",
  "code": 0,
  "message": "string",
  "details": [
    {
      "type_url": "string",
      "value": "string"
    }
  ]
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|error|string|false|none||none|
|code|integer(int32)|false|none||none|
|message|string|false|none||none|
|details|[[protobufAny](#schemaprotobufany)]|false|none||none|

<h2 id="tocS_protobufAny">protobufAny</h2>

<a id="schemaprotobufany"></a>
<a id="schema_protobufAny"></a>
<a id="tocSprotobufany"></a>
<a id="tocsprotobufany"></a>

```json
{
  "type_url": "string",
  "value": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|type_url|string|false|none||none|
|value|string(byte)|false|none||none|

<h2 id="tocS_v1UserInfoResponse">v1UserInfoResponse</h2>

<a id="schemav1userinforesponse"></a>
<a id="schema_v1UserInfoResponse"></a>
<a id="tocSv1userinforesponse"></a>
<a id="tocsv1userinforesponse"></a>

```json
{
  "status": {
    "code": "string",
    "msg": "string",
    "timestampe": "string"
  },
  "user": {
    "uid": "string",
    "user_name": "string",
    "follow_count": "string",
    "follower_count": "string",
    "is_follow": true,
    "user_intro": "string",
    "head_imageURL": "string"
  }
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|status|[v1StatusMsg](#schemav1statusmsg)|false|none||none|
|user|[v1Users](#schemav1users)|false|none||none|

