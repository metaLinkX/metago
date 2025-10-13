// Package model

package model

// TokenConfig 登录令牌配置
type TokenConfig struct {
	SecretKey       string `json:"secretKey"`
	Expires         int64  `json:"expires"`
	AutoRefresh     bool   `json:"autoRefresh"`
	RefreshInterval int64  `json:"refreshInterval"`
	MaxRefreshTimes int64  `json:"maxRefreshTimes"`
	MultiLogin      bool   `json:"multiLogin"`
}

// GenerateAppCrudTemplate curd模板
type GenerateAppCrudTemplate struct {
	Group          string `json:"group"`
	IsAddon        bool   `json:"isAddon"`
	MasterPackage  string `json:"masterPackage"`
	TemplatePath   string `json:"templatePath"`
	ApiPath        string `json:"apiPath"`
	InputPath      string `json:"inputPath"`
	ControllerPath string `json:"controllerPath"`
	LogicPath      string `json:"logicPath"`
	RouterPath     string `json:"routerPath"`
	SqlPath        string `json:"sqlPath"`
	WebApiPath     string `json:"webApiPath"`
	WebViewsPath   string `json:"webViewsPath"`
}

// GenerateAppQueueTemplate 消息队列模板
type GenerateAppQueueTemplate struct {
	Group        string `json:"group"`
	TemplatePath string `json:"templatePath"`
}

// GenerateAppTreeTemplate 关系树列表模板
type GenerateAppTreeTemplate struct {
	Group        string `json:"group"`
	TemplatePath string `json:"templatePath"`
}

// GenerateConfig 生成代码配置
type GenerateConfig struct {
	AllowedIPs  []string `json:"allowedIPs"`
	Application struct {
		Crud struct {
			Templates []*GenerateAppCrudTemplate `json:"templates"`
		} `json:"crud"`
		Queue struct {
			Templates []*GenerateAppQueueTemplate `json:"templates"`
		} `json:"queue"`
		Tree struct {
			Templates []*GenerateAppTreeTemplate `json:"templates"`
		} `json:"tree"`
	} `json:"application"`
	Delimiters    []string `json:"delimiters"`
	DevPath       string   `json:"devPath"`
	DisableTables []string `json:"disableTables"`
	SelectDbs     []string `json:"selectDbs"`
}
