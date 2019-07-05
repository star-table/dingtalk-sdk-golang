# Introduce
DingDing SDK implemented with Golang.

DingDing Documents: [https://open-doc.dingtalk.com/microapp/serverapi3](https://open-doc.dingtalk.com/microapp/serverapi3)
# Features
Partial feature structure and completion:
 - [x] 应用授权
   - [x] 获取企业凭证
   - [x] 获取企业授权信息
   - [x] 获取授权应用信息
 - [ ] 身份验证
 - [ ] 通讯录管理
   - [ ] 用户管理
     - [ ] 获取部门用户userid列表
	 - [x] 获取部门用户
	 - [ ] 获取部门用户详情
	 - [ ] 获取管理员列表
	 - [ ] 获取管理员通讯录权限范围
	 - [ ] 查询管理员是否具备管理某个应用的权限
	 - [ ] 根据unionid获取userid
	 - [ ] 获取企业员工人数
   - [ ] 部门管理
     - [ ] 获取子部门ID列表
	 - [ ] 获取部门列表
	 - [ ] 获取部门详情
	 - [ ] 查询部门的所有上级父部门路径
	 - [ ] 查询指定用户的所有上级父部门路径
   - [ ] 角色管理
     - [ ] 获取角色列表
	 - [ ] 获取角色下的员工列表
	 - [ ] 获取角色组
	 - [ ] 获取角色详情
   - [ ] 外部联系人管理
     - [ ] 获取外部联系人标签列表
	 - [ ] 获取外部联系人列表
	 - [ ] 获取外部联系人详情
   - [ ] 通讯录权限范围
     - [ ] 获取通讯录权限范围
 - [x] 消息通知
   - [x] 工作消息通知
     - [x] 发送工作通知消息
	 - [x] 查询工作通知消息的发送进度
	 - [x] 查询工作通知消息的发送结果
	 - [x] 工作通知消息撤回
   - [x] 发送普通消息
     - [x] 发送普通消息
 - [ ] 文件存储
   - [x] 管理媒体文件
     - [x] 上传媒体文件
