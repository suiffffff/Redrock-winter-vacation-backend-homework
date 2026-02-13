
---

# API 接口使用说明文档

本文档旨在指导开发者如何接入和使用本系统 API。

## 1. 通用规范

### 1.1 请求头 (Headers)
所有接口请求需包含以下头部信息：
*   **Content-Type**: `application/json`
*   **Authorization**: `Bearer <access_token>` (仅限需要认证的接口)

### 1.2 统一响应格式
系统采用统一的 JSON 结构返回数据。

**成功响应示例** (`code: 0`):
```json
{
  "code": 0,
  "message": "success",
  "data": { ... }
}
```

**错误响应示例** (非 0):
```json
{
  "code": 10001,
  "message": "参数错误",
  "data": null
}
```

### 1.3 分页参数
列表接口均支持分页查询，默认参数如下：
*   `page`: 页码 (默认 `1`)
*   `page_size`: 每页数量 (默认 `10`，最大 `100`)

### 1.4 部门枚举值 (Department)
在注册及相关业务中，`department` 字段需使用以下枚举值：

| 枚举值 (Value) | 显示标签 (Label) | 说明 |
| :--- | :--- | :--- |
| `backend` | 后端 | 后端开发 |
| `frontend` | 前端 | 前端开发 |
| `sre` | SRE | 运维/SRE |
| `product` | 产品 | 产品经理 |
| `design` | 视觉设计 | UI/UX 设计 |
| `android` | Android | 安卓开发 |
| `ios` | iOS | iOS 开发 |

> **注意**：API 响应中通常会同时返回 `department` (枚举值) 和 `department_label` (中文标签)。

---

## 2. 用户模块 (User)

本模块处理用户的注册、登录及个人信息管理。

### 2.1 用户注册
*   **接口地址**: `POST /user/register`
*   **认证**: 否
*   **参数**:
    *   `username` (必填): 用户名
    *   `password` (必填): 密码
    *   `nickname` (必填): 昵称
    *   `department` (必填): 部门枚举值 (见 1.4)

### 2.2 用户登录
*   **接口地址**: `POST /user/login`
*   **认证**: 否
*   **参数**: `username`, `password`
*   **响应**: 返回 `access_token` (用于后续接口认证) 和 `refresh_token`。

### 2.3 刷新 Token
*   **接口地址**: `POST /user/refresh`
*   **认证**: 否
*   **参数**: `refresh_token` (必填)

### 2.4 获取当前用户信息
*   **接口地址**: `GET /user/profile`
*   **认证**: 是

### 2.5 注销账号
*   **接口地址**: `DELETE /user/account`
*   **认证**: 是
*   **参数**: `password` (必填，用于二次确认)

---

## 3. 作业模块 (Homework)

本模块涉及作业的发布、管理和查看。
*   **权限说明**:
    *   **老登 (Admin)**: 拥有发布、修改、删除作业的权限。
    *   **小登 (Student)**: 仅拥有查看权限。

### 3.1 发布作业 (Admin)
*   **接口地址**: `POST /homework`
*   **参数**:
    *   `title` (必填): 作业标题
    *   `description` (必填): 作业描述
    *   `department` (必填): 所属部门
    *   `deadline` (必填): 截止时间 (格式: `YYYY-MM-DD HH:mm:ss`)
    *   `allow_late` (选填): 是否允许补交 (默认 `false`)

### 3.2 获取作业列表
*   **接口地址**: `GET /homework`
*   **查询参数**: `department` (筛选), `page`, `page_size`

### 3.3 获取作业详情
*   **接口地址**: `GET /homework/:id`
*   **说明**: 如果是学生（小登）调用，响应中会包含 `my_submission` 字段，显示该学生的提交状态。

### 3.4 修改作业 (Admin)
*   **接口地址**: `PUT /homework/:id`
*   **权限**: 仅限同部门管理员。
*   **参数**: `title`, `description`, `deadline`, `allow_late` (均为选填)。

### 3.5 删除作业 (Admin)
*   **接口地址**: `DELETE /homework/:id`
*   **权限**: 仅限同部门管理员。

---

## 4. 作业提交模块 (Submission)

本模块处理学生的作业提交及管理员的批改。

### 4.1 提交作业 (Student)
*   **接口地址**: `POST /submission`
*   **权限**: 仅限学生（小登）。
*   **参数**:
    *   `homework_id` (必填): 作业 ID
    *   `content` (必填): 提交内容 (文本或链接)
    *   `file_url` (选填): 附件地址
*   **业务逻辑**:
    1. 检查作业是否存在及是否属于当前用户部门。
    2. 检查是否过截止时间。
    3. 如果过截止时间，检查 `allow_late` 是否为 true。
    4. 自动记录 `is_late` 状态。

### 4.2 获取我的提交列表 (Student)
*   **接口地址**: `GET /submission/my`

### 4.3 获取某作业的所有提交 (Admin)
*   **接口地址**: `GET /submission/homework/:homework_id`
*   **权限**: 仅限同部门管理员。

### 4.4 批改作业 (Admin)
*   **接口地址**: `PUT /submission/:id/review`
*   **权限**: 仅限同部门管理员。
*   **参数**:
    *   `score` (选填): 分数 (0-100)
    *   `comment` (必填): 评语
    *   `is_excellent` (选填): 是否标记为优秀作业

### 4.5 标记/取消优秀作业 (Admin)
*   **接口地址**: `PUT /submission/:id/excellent`
*   **参数**: `is_excellent` (必填，bool)

### 4.6 获取优秀作业列表
*   **接口地址**: `GET /submission/excellent`
*   **查询参数**: `department`, `page`, `page_size`

---

## 5. 进阶接口 (选做)

### 5.1 绑定邮箱
*   **接口地址**: `POST /user/bindEmail`
*   **参数**: `email` (必填)

### 5.2 AI 作业评价 (Admin)
*   **接口地址**: `POST /submission/:id/aiReview`
*   **说明**: 调用 AI 对作业代码进行分析并生成建议分数。

---

## 6. 接口权限汇总表

| 模块 | 方法 | 路径 | 功能 | 认证 | 权限要求 |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **用户** | POST | `/user/register` | 用户注册 | 否 | - |
| | POST | `/user/login` | 用户登录 | 否 | - |
| | POST | `/user/refresh` | 刷新 Token | 否 | - |
| | GET | `/user/profile` | 获取个人信息 | **是** | - |
| | DELETE | `/user/account` | 注销账号 | **是** | - |
| **作业** | POST | `/homework` | 发布作业 | **是** | **老登 (Admin)** |
| | GET | `/homework` | 获取列表 | **是** | - |
| | GET | `/homework/:id` | 获取详情 | **是** | - |
| | PUT | `/homework/:id` | 修改作业 | **是** | **老登 + 同部门** |
| | DELETE | `/homework/:id` | 删除作业 | **是** | **老登 + 同部门** |
| **提交** | POST | `/submission` | 提交作业 | **是** | **小登 (Student)** |
| | GET | `/submission/my` | 我的提交列表 | **是** | **小登 (Student)** |
| | GET | `/submission/homework/:id` | 作业所有提交 | **是** | **老登 + 同部门** |
| | PUT | `/submission/:id/review` | 批改作业 | **是** | **老登 + 同部门** |
| | PUT | `/submission/:id/excellent`| 标记优秀 | **是** | **老登 + 同部门** |
| | GET | `/submission/excellent` | 优秀作业列表 | **是** | - |
# Go-Gin 作业提交管理系统

## 3. 项目简介 (Project Intro)

本项目是一个基于 **Go** 语言和 **Gin** 框架开发的后端作业管理系统。系统旨在解决高校或培训机构内部的作业发布、提交与批改流程。

系统设计了完善的权限管理体系，区分 **老登 (Admin/讲师)** 和 **小登 (Student/学生)** 两种角色，并引入了 **部门 (Department)** 概念，实现了基于部门的数据隔离（同部门讲师管理同部门作业）。

主要业务流程涵盖：用户注册登录、作业发布与管理、学生提交作业、讲师批改评分以及优秀作业展示。

## 4. 技术栈说明 (Tech Stack)

本项目采用经典的分层架构设计，确保代码的解耦与可维护性。

| 类别 | 技术/库 | 说明 |
| :--- | :--- | :--- |
| **编程语言** | Golang | 1.20+ |
| **Web 框架** | Gin | 高性能 HTTP Web 框架 |
| **数据库** | MySQL 8.0 | 关系型数据库 |
| **ORM 框架** | GORM | 数据持久层操作 (DAO层) |
| **认证鉴权** | JWT (JSON Web Token) | 用户登录态管理与中间件鉴权 |
| **配置管理** | Viper | 配置文件读取 (YAML/JSON) |
| **日志管理** | Zap / Logrus | 高性能日志记录 |

## 5. 项目结构说明 (Project Structure)

项目遵循标准的 **Dao - Service - Controller (Handler)** 分层架构：

```text
├── cmd/
│   └── main.go           # 项目启动入口
├── config/               # 配置文件 (config.yaml)
├── internal/
│   ├── router/           # 路由层：定义 API 路径，注册中间件
│   ├── handler/          # 控制层：处理 HTTP 请求，参数校验 (DTO)，调用 Service，返回响应
│   ├── service/          # 业务逻辑层：处理核心业务逻辑，事务控制
│   ├── dao/              # 数据访问层：直接与数据库交互 (GORM 操作)
│   ├── model/            # 数据库模型定义 (Struct <-> Table)
│   ├── dto/              # 数据传输对象：定义请求参数结构体和响应结构体
│   └── middleware/       # 中间件：JWT认证、CORS跨域、日志记录
├── pkg/                  # 公共工具包 (Utils, Error codes)
└── README.md             # 项目说明文档
```

*   **Router**: 负责路由分发。
*   **Handler**: 解析 `Context` 中的参数，绑定到 DTO，将结果以统一格式 JSON 返回。
*   **Service**: 真正的业务逻辑所在地（如：判断作业是否逾期、计算分数等）。
*   **Dao**: 封装 CRUD 操作，Service 层通过 Dao 层访问数据库。

## 1. 已实现功能清单 (Implemented Features)

### 👤 用户模块 (User)
- [x] **用户注册**：支持用户名、密码、昵称及部门（后端/前端/SRE等）录入。
- [x] **用户登录**：基于 JWT 的 Token 签发（Access Token + Refresh Token）。
- [x] **Token 刷新**：支持无感刷新 Token。
- [x] **个人信息**：获取当前登录用户信息。
- [x] **注销账号**：需二次验证密码。

### 📝 作业模块 (Homework)
- [x] **发布作业**：仅限管理员（老登），支持设置截止时间、是否允许补交。
- [x] **作业列表**：支持按部门筛选分页查询。
- [x] **作业详情**：学生查看时包含自己的提交状态。
- [x] **修改/删除作业**：仅限同部门管理员操作。

### 📤 提交与批改模块 (Submission)
- [x] **提交作业**：仅限学生（小登），自动校验截止日期与补交权限。
- [x] **我的提交**：学生查看自己的提交历史。
- [x] **作业提交列表**：管理员查看某作业下的所有学生提交。
- [x] **批改作业**：管理员打分、写评语。
- [x] **优秀作业**：
    - [x] 标记/取消优秀作业。
    - [x] 公开展示优秀作业列表。

## 2. 进阶功能说明 (Advanced Features)

- [ ] **邮箱绑定** (TODO)：支持用户绑定邮箱，用于接收作业通知。
- [ ] **AI 作业评价** (TODO)：集成 LLM (大模型) 接口，对学生提交的代码内容进行自动分析并给出建议分数。
- [ ] **文件上传服务**：集成 OSS 或本地静态资源服务，支持作业附件上传。

## 6. 本地运行指南 (Local Run Guide)

### 前置要求
1.  安装 Go (1.18+)
2.  安装 MySQL 8.0+

### 步骤
1.  **克隆项目**
    ```bash
    git clone https://github.com/suiffffff/Reorock-winter-vacation-backend-homework
    cd your-project
    ```

2.  **配置数据库**
    *   在 MySQL 中创建一个新的数据库（例如 `homework_db`）。
    *   修改 `config/config.yaml` 文件中的数据库连接配置：
        ```yaml
        mysql:
          host: "127.0.0.1"
          port: 3306
          user: "root"
          password: "your_password"
          dbname: "homework_db"
        ```

3.  **下载依赖**
    ```bash
    go mod tidy
    ```

4.  **运行项目**
    ```bash
    go run cmd/main.go
    ```
    *   项目默认运行在 `http://localhost:8080` (取决于配置文件)。
    *   GORM 会自动迁移表结构 (AutoMigrate)。

