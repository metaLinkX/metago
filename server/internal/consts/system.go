package consts

const VersionApp = "1.0.0"
const (
	HTTPContentTypeXml    = "text/xml"
	HTTPContentTypeHtml   = "text/html"
	HTTPContentTypeStream = "text/event-stream"
	HTTPContentTypeJson   = "application/json"
)

const (
	ApiPrefix   = "/api"
	AdminPrefix = "/admin"
)

const (
	ContextHTTPKey  = "httpContext" // http上下文变量名称
	SuperRoleKey    = "super"       // 超管角色唯一标识符，通过角色验证超管
	DefaultPageSize = 20            //  默认分页
	NilJsonToString = "{}"          // 空json初始化值
	MaxSortIncr     = 10            // 最大排序值增量
)
