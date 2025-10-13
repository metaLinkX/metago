package sys

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"server/internal/consts"
	"server/internal/dao"
	"server/internal/model"
	"server/internal/model/entity"
	"server/internal/model/input/sysin"
	"server/internal/service"
)

type sSysConfig struct{}

func NewSysConfig() *sSysConfig {
	return &sSysConfig{}
}

func init() {
	service.RegisterSysConfig(NewSysConfig())
}

// GetLogin 获取登录配置
func (s *sSysConfig) GetLogin(ctx context.Context) (conf *model.LoginConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &sysin.GetConfigInp{Group: "login"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

// GetWechat 获取微信配置
func (s *sSysConfig) GetWechat(ctx context.Context) (conf *model.WechatConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &sysin.GetConfigInp{Group: "wechat"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

// GetPay 获取支付配置
func (s *sSysConfig) GetPay(ctx context.Context) (conf *model.PayConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &sysin.GetConfigInp{Group: "pay"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

// GetSms 获取短信配置
func (s *sSysConfig) GetSms(ctx context.Context) (conf *model.SmsConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &sysin.GetConfigInp{Group: "sms"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

// GetUpload 获取上传配置
func (s *sSysConfig) GetUpload(ctx context.Context) (conf *model.UploadConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &sysin.GetConfigInp{Group: "upload"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

// GetSmtp 获取邮件配置
func (s *sSysConfig) GetSmtp(ctx context.Context) (conf *model.EmailConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &sysin.GetConfigInp{Group: "smtp"})
	if err != nil {
		return
	}
	if err = gconv.Scan(models.List, &conf); err != nil {
		return
	}

	conf.Addr = fmt.Sprintf("%s:%d", conf.Host, conf.Port)

	return
}

// GetBasic 获取基础配置
func (s *sSysConfig) GetBasic(ctx context.Context) (conf *model.BasicConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &sysin.GetConfigInp{Group: "basic"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

// GetApiJwtToken 获取本地token配置
func (s *sSysConfig) GetApiJwtToken(ctx context.Context) (conf *model.TokenConfig, err error) {
	err = g.Cfg().MustGet(ctx, "jwt.api").Scan(&conf)
	return
}

// GetAdminJwtToken 获取本地token配置
func (s *sSysConfig) GetAdminJwtToken(ctx context.Context) (conf *model.TokenConfig, err error) {
	err = g.Cfg().MustGet(ctx, "jwt.admin").Scan(&conf)
	return
}

// GetConfigByGroup 获取指定分组的配置
func (s *sSysConfig) GetConfigByGroup(ctx context.Context, in *sysin.GetConfigInp) (res *sysin.GetConfigModel, err error) {
	if in.Group == "" {
		err = gerror.New("分组不能为空")
		return
	}

	var models []*entity.SysConfig
	cols := dao.SysConfig.Columns()
	if err = dao.SysConfig.Ctx(ctx).Fields(cols.Key, cols.Value, cols.Type).Where(cols.Group, in.Group).Scan(&models); err != nil {
		err = gerror.Wrapf(err, "获取配置分组[ %v ]失败，请稍后重试！", in.Group)
		return
	}

	res = new(sysin.GetConfigModel)
	if len(models) > 0 {
		res.List = make(g.Map, len(models))
		for _, v := range models {
			val, err := s.ConversionType(ctx, v)
			if err != nil {
				return nil, err
			}
			res.List[v.Key] = val
		}
	}
	return
}

// ConversionType 转换类型
func (s *sSysConfig) ConversionType(ctx context.Context, models *entity.SysConfig) (value interface{}, err error) {
	if models == nil {
		err = gerror.New("数据不存在")
		return
	}
	return consts.ConvType(models.Value, models.Type), nil
}

func (s *sSysConfig) getConfigByKey(key string, models []*entity.SysConfig) *entity.SysConfig {
	if len(models) == 0 {
		return nil
	}
	for _, v := range models {
		if key == v.Key {
			return v
		}
	}
	return nil
}

// GetLoadGenerate 获取本地生成配置
func (s *sSysConfig) GetLoadGenerate(ctx context.Context) (conf *model.GenerateConfig, err error) {
	err = g.Cfg().MustGet(ctx, "genCode").Scan(&conf)
	return
}
