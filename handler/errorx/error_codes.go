package errorx

// Defines a set of common error codes and corresponding messages
var (
	OK                 = add(0, "OK")         // OK
	NoLogin            = add(101, "用户未登录")    // NOT_LOGIN
	RequestErr         = add(400, "请求参数错误")   // INVALID_ARGUMENT
	Unauthorized       = add(401, "未认证或认证失败") // UNAUTHENTICATED
	AccessDenied       = add(403, "权限拒绝")     // PERMISSION_DENIED
	NotFound           = add(404, "资源未找到")    // NOT_FOUND
	MethodNotAllowed   = add(405, "方法不被允许")   // METHOD_NOT_ALLOWED
	Canceled           = add(498, "请求已取消")    // CANCELED
	ServerErr          = add(500, "服务器内部错误")  // INTERNAL_ERROR
	ServiceUnavailable = add(503, "服务不可用")    // UNAVAILABLE
	Deadline           = add(504, "请求超时")     // DEADLINE_EXCEEDED
	LimitExceed        = add(509, "超出资源限制")   // RESOURCE_EXHAUSTED
)
