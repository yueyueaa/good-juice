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

# Basic

<a id="opIdBasic_GetCommentsList"></a>

## GET /v1/basic/Comments/list

GET /v1/basic/Comments/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|vid|query|string| 否 |none|
|token|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "status_code": "string",
  "status_msg": "string",
  "comment_list": [
    {
      "cid": "string",
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
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1CommentsListResponse](#schemav1commentslistresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdBasic_CollectionsAction"></a>

## POST /v1/basic/collections/action

POST /v1/basic/collections/action

> 返回示例

> 200 Response

```json
{
  "status_code": "string",
  "status_msg": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1Response](#schemav1response)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdBasic_DeleteVideos"></a>

## POST /v1/basic/delete/Videos

POST /v1/basic/delete/Videos

> 返回示例

> 200 Response

```json
{
  "status_code": "string",
  "status_msg": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1Response](#schemav1response)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdBasic_DeleteComments"></a>

## POST /v1/basic/delete/comments

POST /v1/basic/delete/comments

> 返回示例

> 200 Response

```json
{
  "status_code": "string",
  "status_msg": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1Response](#schemav1response)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdBasic_FavoriteAction"></a>

## POST /v1/basic/favorite/action

POST /v1/basic/favorite/action

> 返回示例

> 200 Response

```json
{
  "status_code": "string",
  "status_msg": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1Response](#schemav1response)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdBasic_GetFavoriteList"></a>

## GET /v1/basic/favorite/list

GET /v1/basic/favorite/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|uid|query|string| 否 |none|
|token|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "status_code": "string",
  "status_msg": "string",
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
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1FavoriteListResponse](#schemav1favoritelistresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdBasic_PublishComments"></a>

## POST /v1/basic/publish/comments

POST /v1/basic/publish/comments

> 返回示例

> 200 Response

```json
{
  "status_code": "string",
  "status_msg": "string",
  "comment": {
    "cid": "string",
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
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1PublishCommentsResponse](#schemav1publishcommentsresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdBasic_PublishImages"></a>

## POST /v1/basic/publish/images

POST /v1/basic/publish/images

> 返回示例

> 200 Response

```json
{
  "status_code": "string",
  "status_msg": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1Response](#schemav1response)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdBasic_PublishVideos"></a>

## POST /v1/basic/publish/videos

POST /v1/basic/publish/videos

> 返回示例

> 200 Response

```json
{
  "status_code": "string",
  "status_msg": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1Response](#schemav1response)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdBasic_RelationAction"></a>

## POST /v1/basic/relation/action

POST /v1/basic/relation/action

> 返回示例

> 200 Response

```json
{
  "status_code": "string",
  "status_msg": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1Response](#schemav1response)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdBasic_GetRelationFollowerList"></a>

## GET 粉丝列表

GET /v1/basic/relation/follower/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|uid|query|string| 否 |none|
|token|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "status_code": "string",
  "status_msg": "string",
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
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1RelationFollowerListResponse](#schemav1relationfollowerlistresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdBasic_GetCollectionsList"></a>

## GET /v1/basic/relation/list

GET /v1/basic/relation/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|uid|query|string| 否 |none|
|token|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "status_code": "string",
  "status_msg": "string",
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
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1CollectionsListResponse](#schemav1collectionslistresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdBasic_GetVideosList"></a>

## GET /v1/basic/videos/list

GET /v1/basic/videos/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|uid|query|string| 否 |none|
|token|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "status_code": "string",
  "status_msg": "string",
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
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1VideosListResponse](#schemav1videoslistresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

# User

<a id="opIdUser_UserLogin"></a>

## POST /v1/user/login

POST /v1/user/login

> 返回示例

> 200 Response

```json
{
  "status_code": "string",
  "status_msg": "string",
  "uid": "string",
  "token": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1LoginResponse](#schemav1loginresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdUser_UserRegister"></a>

## POST /v1/user/register

POST /v1/user/register

> 返回示例

> 200 Response

```json
{
  "status_code": "string",
  "status_msg": "string",
  "uid": "string",
  "token": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1RegisterResponse](#schemav1registerresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdUser_UserRegisterIntro"></a>

## POST /v1/user/register/intro

POST /v1/user/register/intro

> 返回示例

> 200 Response

```json
{
  "status_code": "string",
  "status_msg": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1RegisterIntroResponse](#schemav1registerintroresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdUser_GetUserInfo"></a>

## GET /v1/user/userinfo

GET /v1/user/userinfo

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|uid|query|string| 否 |none|
|token|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "status_code": "string",
  "status_msg": "string",
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

<a id="opIdUser_GetUsersInfo"></a>

## GET /v1/user/usersinfo

GET /v1/user/usersinfo

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|users_keywords|query|string| 否 |none|
|token|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "status_code": "string",
  "status_msg": "string",
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
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1UsersInfoResponse](#schemav1usersinforesponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

# Video

<a id="opIdVideo_ClassifiedVideos"></a>

## GET /v1/video/classified/videos

GET /v1/video/classified/videos

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|video_type|query|string| 否 |none|
|token|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "status_code": "string",
  "status_msg": "string",
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
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1ClassifiedVideosResponse](#schemav1classifiedvideosresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdVideo_GetVideosFeed"></a>

## GET /v1/video/videos/feed

GET /v1/video/videos/feed

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|last_time|query|string| 否 |none|
|token|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "status_code": "string",
  "status_msg": "string",
  "next_time": "string",
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
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1VideosFeedResponse](#schemav1videosfeedresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

<a id="opIdVideo_GetVideosInfo"></a>

## GET /v1/video/videosinfo

GET /v1/video/videosinfo

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|videos_keywords|query|string| 否 |none|
|token|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "status_code": "string",
  "status_msg": "string",
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
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1VideosInfoResponse](#schemav1videosinforesponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|An unexpected error response.|[runtimeError](#schemaruntimeerror)|

# 数据模型

<h2 id="tocS_v1VideosListResponse">v1VideosListResponse</h2>

<a id="schemav1videoslistresponse"></a>
<a id="schema_v1VideosListResponse"></a>
<a id="tocSv1videoslistresponse"></a>
<a id="tocsv1videoslistresponse"></a>

```json
{
  "status_code": "string",
  "status_msg": "string",
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
|status_code|string(uint64)|false|none||none|
|status_msg|string|false|none||none|
|video_list|[[v1Videos](#schemav1videos)]|false|none||none|

<h2 id="tocS_v1VideosInfoResponse">v1VideosInfoResponse</h2>

<a id="schemav1videosinforesponse"></a>
<a id="schema_v1VideosInfoResponse"></a>
<a id="tocSv1videosinforesponse"></a>
<a id="tocsv1videosinforesponse"></a>

```json
{
  "status_code": "string",
  "status_msg": "string",
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
|status_code|string|false|none||none|
|status_msg|string|false|none||none|
|video_list|[[v1Videos](#schemav1videos)]|false|none||none|

<h2 id="tocS_v1VideosFeedResponse">v1VideosFeedResponse</h2>

<a id="schemav1videosfeedresponse"></a>
<a id="schema_v1VideosFeedResponse"></a>
<a id="tocSv1videosfeedresponse"></a>
<a id="tocsv1videosfeedresponse"></a>

```json
{
  "status_code": "string",
  "status_msg": "string",
  "next_time": "string",
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
|status_code|string(int64)|false|none||none|
|status_msg|string|false|none||none|
|next_time|string(int64)|false|none||none|
|video_list|[[v1Videos](#schemav1videos)]|false|none||none|

<h2 id="tocS_v1Videos">v1Videos</h2>

<a id="schemav1videos"></a>
<a id="schema_v1Videos"></a>
<a id="tocSv1videos"></a>
<a id="tocsv1videos"></a>

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

<h2 id="tocS_v1UsersInfoResponse">v1UsersInfoResponse</h2>

<a id="schemav1usersinforesponse"></a>
<a id="schema_v1UsersInfoResponse"></a>
<a id="tocSv1usersinforesponse"></a>
<a id="tocsv1usersinforesponse"></a>

```json
{
  "status_code": "string",
  "status_msg": "string",
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
|status_code|string|false|none||none|
|status_msg|string|false|none||none|
|user_list|[[v1Users](#schemav1users)]|false|none||none|

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

<h2 id="tocS_v1UserInfoResponse">v1UserInfoResponse</h2>

<a id="schemav1userinforesponse"></a>
<a id="schema_v1UserInfoResponse"></a>
<a id="tocSv1userinforesponse"></a>
<a id="tocsv1userinforesponse"></a>

```json
{
  "status_code": "string",
  "status_msg": "string",
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
|status_code|string(uint64)|false|none||none|
|status_msg|string|false|none||none|
|user|[v1Users](#schemav1users)|false|none||none|

<h2 id="tocS_v1Response">v1Response</h2>

<a id="schemav1response"></a>
<a id="schema_v1Response"></a>
<a id="tocSv1response"></a>
<a id="tocsv1response"></a>

```json
{
  "status_code": "string",
  "status_msg": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|status_code|string(uint64)|false|none||none|
|status_msg|string|false|none||none|

<h2 id="tocS_v1RelationListResponse">v1RelationListResponse</h2>

<a id="schemav1relationlistresponse"></a>
<a id="schema_v1RelationListResponse"></a>
<a id="tocSv1relationlistresponse"></a>
<a id="tocsv1relationlistresponse"></a>

```json
{
  "status_code": "string",
  "status_msg": "string",
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
|status_code|string|false|none||none|
|status_msg|string|false|none||none|
|user_list|[[v1Users](#schemav1users)]|false|none||none|

<h2 id="tocS_v1RelationFollowerListResponse">v1RelationFollowerListResponse</h2>

<a id="schemav1relationfollowerlistresponse"></a>
<a id="schema_v1RelationFollowerListResponse"></a>
<a id="tocSv1relationfollowerlistresponse"></a>
<a id="tocsv1relationfollowerlistresponse"></a>

```json
{
  "status_code": "string",
  "status_msg": "string",
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
|status_code|string|false|none||none|
|status_msg|string|false|none||none|
|user_list|[[v1Users](#schemav1users)]|false|none||none|

<h2 id="tocS_v1RegisterResponse">v1RegisterResponse</h2>

<a id="schemav1registerresponse"></a>
<a id="schema_v1RegisterResponse"></a>
<a id="tocSv1registerresponse"></a>
<a id="tocsv1registerresponse"></a>

```json
{
  "status_code": "string",
  "status_msg": "string",
  "uid": "string",
  "token": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|status_code|string(uint64)|false|none||none|
|status_msg|string|false|none||none|
|uid|string(uint64)|false|none||none|
|token|string|false|none||none|

<h2 id="tocS_v1RegisterIntroResponse">v1RegisterIntroResponse</h2>

<a id="schemav1registerintroresponse"></a>
<a id="schema_v1RegisterIntroResponse"></a>
<a id="tocSv1registerintroresponse"></a>
<a id="tocsv1registerintroresponse"></a>

```json
{
  "status_code": "string",
  "status_msg": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|status_code|string(uint64)|false|none||none|
|status_msg|string|false|none||none|

<h2 id="tocS_v1PublishCommentsResponse">v1PublishCommentsResponse</h2>

<a id="schemav1publishcommentsresponse"></a>
<a id="schema_v1PublishCommentsResponse"></a>
<a id="tocSv1publishcommentsresponse"></a>
<a id="tocsv1publishcommentsresponse"></a>

```json
{
  "status_code": "string",
  "status_msg": "string",
  "comment": {
    "cid": "string",
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
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|status_code|string(uint64)|false|none||none|
|status_msg|string|false|none||none|
|comment|[v1Comment](#schemav1comment)|false|none||none|

<h2 id="tocS_v1LoginResponse">v1LoginResponse</h2>

<a id="schemav1loginresponse"></a>
<a id="schema_v1LoginResponse"></a>
<a id="tocSv1loginresponse"></a>
<a id="tocsv1loginresponse"></a>

```json
{
  "status_code": "string",
  "status_msg": "string",
  "uid": "string",
  "token": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|status_code|string(uint64)|false|none||none|
|status_msg|string|false|none||none|
|uid|string(uint64)|false|none||none|
|token|string|false|none||none|

<h2 id="tocS_v1FavoriteListResponse">v1FavoriteListResponse</h2>

<a id="schemav1favoritelistresponse"></a>
<a id="schema_v1FavoriteListResponse"></a>
<a id="tocSv1favoritelistresponse"></a>
<a id="tocsv1favoritelistresponse"></a>

```json
{
  "status_code": "string",
  "status_msg": "string",
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
|status_code|string|false|none||none|
|status_msg|string|false|none||none|
|video_list|[[v1Videos](#schemav1videos)]|false|none||none|

<h2 id="tocS_v1CommentsListResponse">v1CommentsListResponse</h2>

<a id="schemav1commentslistresponse"></a>
<a id="schema_v1CommentsListResponse"></a>
<a id="tocSv1commentslistresponse"></a>
<a id="tocsv1commentslistresponse"></a>

```json
{
  "status_code": "string",
  "status_msg": "string",
  "comment_list": [
    {
      "cid": "string",
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
|status_code|string(uint64)|false|none||none|
|status_msg|string|false|none||none|
|comment_list|[[v1Comment](#schemav1comment)]|false|none||none|

<h2 id="tocS_v1Comment">v1Comment</h2>

<a id="schemav1comment"></a>
<a id="schema_v1Comment"></a>
<a id="tocSv1comment"></a>
<a id="tocsv1comment"></a>

```json
{
  "cid": "string",
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
|cid|string(uint64)|false|none||none|
|user|[v1Users](#schemav1users)|false|none||none|
|content|string|false|none||none|
|create_date|string|false|none||none|

<h2 id="tocS_v1CollectionsListResponse">v1CollectionsListResponse</h2>

<a id="schemav1collectionslistresponse"></a>
<a id="schema_v1CollectionsListResponse"></a>
<a id="tocSv1collectionslistresponse"></a>
<a id="tocsv1collectionslistresponse"></a>

```json
{
  "status_code": "string",
  "status_msg": "string",
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
|status_code|string|false|none||none|
|status_msg|string|false|none||none|
|video_list|[[v1Videos](#schemav1videos)]|false|none||none|

<h2 id="tocS_v1ClassifiedVideosResponse">v1ClassifiedVideosResponse</h2>

<a id="schemav1classifiedvideosresponse"></a>
<a id="schema_v1ClassifiedVideosResponse"></a>
<a id="tocSv1classifiedvideosresponse"></a>
<a id="tocsv1classifiedvideosresponse"></a>

```json
{
  "status_code": "string",
  "status_msg": "string",
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
|status_code|string|false|none||none|
|status_msg|string|false|none||none|
|video_list|[[v1Videos](#schemav1videos)]|false|none||none|

<h2 id="tocS_v1Author">v1Author</h2>

<a id="schemav1author"></a>
<a id="schema_v1Author"></a>
<a id="tocSv1author"></a>
<a id="tocsv1author"></a>

```json
{
  "uid": "string",
  "name": "string",
  "follow_count": "string",
  "follower_count": "string",
  "is_follow": true
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|uid|string(uint64)|false|none||none|
|name|string|false|none||none|
|follow_count|string(uint64)|false|none||none|
|follower_count|string(uint64)|false|none||none|
|is_follow|boolean|false|none||none|

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

