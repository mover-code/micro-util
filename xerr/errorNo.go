package xerr

var (

	// 系统错误, 前缀为 100
	InternalServerErrno = New(10001, "内部服务器错误")
	ErrBind             = New(10002, "请求参数错误")

	// Verify email

	ErrEmailSend     = New(30001, "邮件发送失败")     // 邮件发送失败
	ErrEmail         = New(30002, "无效邮箱")       // 无效邮箱
	ErrMobile        = New(30003, "无效手机号")      // 无效手机号
	ErrVerify        = New(30004, "邮件验证失败")     // 邮件验证失败
	ErrEmailOrMobile = New(30005, "邮箱或手机号不能为空") // 邮箱或手机号不能为空

	// 数据库错误, 前缀为 201
	ErrSearch           = New(20100, "查询出错")
	ErrUpdate           = New(20101, "更新出错")
	ErrInsert           = New(20102, "添加出错")
	ErrDelete           = New(20103, "删除出错")
	ErrInsertRepeatedly = New(20104, "重复添加")
	EmptyData           = New(20105, "不能为空")
	SendErr             = New(20106, "发送失败")
	SendSucc            = New(20107, "发送成功")
	Empty               = New(20108, "无数据")

	// 认证错误, 前缀是 202
	ErrValidToken  = New(20201, "Token校验失败")  // token校验失败
	TokenExpired   = New(20203, "token已过期")   // token已过期
	TokenInvalid   = New(20206, "无效token")    // token无效
	ErrTokenSigned = New(20207, "token生成失败")  // token生成失败
	NotToken       = New(20208, "请求未携带token") // 请求未携带token
	ErrGetToken    = New(20209, "token查找失败")  // token查找失败
	NoToken        = New(20210, "token值为空")   // token值为空

	// 用户错误, 前缀为 203

	ErrUserNotFound        = New(20301, "用户不存在")    //用户不存在
	ErrPasswordIncorrect   = New(20302, "账号或密码错误")  //账号或密码错误
	ErrGetUserList         = New(20303, "获取用户列表失败") //获取用户列表失败
	ErrUserExistByUserName = New(20304, "用户名已存在")   //用户名已存在
	ErrUserExistByMobile   = New(20305, "手机号已存在")   //手机号已存在
	ErrGetUserInfo         = New(20306, "获取用户数据失败") //获取用户数据失败
	ErrEditUserInfo        = New(20307, "编辑用户失败")   //编辑用户失败
	ErrDeleteUser          = New(20308, "删除用户失败")   //删除用户失败
	ErrValidCode           = New(20309, "验证码错误")    //验证码错误
	ErrLogout              = New(20310, "退出失败")
	ErrForceLogout         = New(20311, "强制退出失败")    //强制退出失败
	ErrEmptyPhoneNumber    = New(20312, "手机号码为空")    //手机号码为空
	ErrCreateUser          = New(20313, "创建用户失败")    //创建用户失败
	ErrSearchUserExist     = New(20314, "查询用户存在失败")  //查询用户存在失败
	ErrPasswordLength      = New(20315, "密码长度不能小于8") //密码长度不能小于8位
	// 店铺相关
	ErrApply   = New(20400, "申请失败")
	ErrNoShop  = New(20401, "您还未开通店铺")
	ErrAddCart = New(0, "购物车添加失败")
)
