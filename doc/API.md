# API 文档

### 获取服务器在线用户列表

#### 请求信息

- **URL**: `/api/v1/getOnlineUserCount/:id`
- **Method**: `GET`
- **Headers**:
  - `Content-Type: application/json`

#### 请求参数

- **Query Parameters**:
  - `id` (必填)：服务器索引 ID。

#### 响应

- **Success**:

  - **Status Code**: `200 OK`
  - **Body**:
    ```json
    {
      "count": 1,
      "data": [
        {
            "client_nickname": "serveradmin"
        }
        // ......
      ],
      "id": "4",
      "status": {
        "code": 200,
        "message": "Ok or Error"
      }
    }
    ```

### 添加服务器信息到数据库

#### 请求信息

- **URL**: `/api/v1/addServerInfo`
- **Method**: `POST`
- **Headers**:
  - `Content-Type: application/json`

#### 请求参数

- **Body**:
  ```json
  {
    "linksrv": "推荐使用SRV解析记录",
    "linkcity": "服务器地区",
    "apikey": "密钥",
    "email": "邮箱",
    "webquery": "webQuery地址"
  }
  ```

#### 响应

- **Success**
  - **Status Code**: `200 Ok`
  - **Body**
    ```json
    {
      "email": "邮箱",
      "linkcity": "服务器地区",
      "linksrv": "服务器地址",
      "status": {
        "code": 0,
        "message": "ok"
      }
    }
    ```

### 创建频道

#### 请求信息

- **URL**: `/api/v1/createChannel`
- **Method**: `POST`
- **Headers**:
  - `Content-Type: application/json`

#### 请求参数

- **Body**:
  ```json
  {
    "sid": "4",
    "channel_name": "channel#4",
    "channel_password": "123123"
  }
  ```

#### 响应

- **Success**
  - **Status Code**: `200 Ok`
  - **Body**
    ```json
    {
      "body": [
        {
          "cid": "13"
        }
      ],
      "status": {
        "code": 0,
        "message": "ok"
      }
    }
    ```
