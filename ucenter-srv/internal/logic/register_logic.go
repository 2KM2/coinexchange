package logic

import (
	"common/tools"
	"context"
	"errors"
	"grpc-common/ucenter/types/register"
	"time"
	"ucenter-srv/internal/domain"
	"ucenter-srv/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

const RegisterCacheKey = "REGISTER:"

// RegisterLogic 注册的成员变量
type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	CaptchaDomain *domain.CaptchaDomain
	MemberDomain  *domain.MemberDomain
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:           ctx,
		svcCtx:        svcCtx,
		Logger:        logx.WithContext(ctx),
		CaptchaDomain: domain.NewCaptchaDomain(),
		MemberDomain:  domain.NewMemberDomain(svcCtx.Db),
	}
}

func (rl *RegisterLogic) RegisterByPhone(in *register.RegReq) (*register.RegRes, error) {
	//1.人机校验
	logx.Info("[RegisterByPhone] 人机校验")
	isVerify := rl.CaptchaDomain.Verify(
		in.Captcha.Server,
		rl.svcCtx.Config.Captcha.Vid,
		rl.svcCtx.Config.Captcha.Key,
		in.Captcha.Token,
		2,
		in.Ip,
	)
	logx.Info("[RegisterByPhone] result ", isVerify)
	if !isVerify {
		return nil, errors.New("人机校验不通过")
	}
	logx.Info("[RegisterByPhone] 人机校验通过....")

	//2.校验验证码 redis中获取验证码
	redisValue := ""
	err := rl.svcCtx.Cache.GetCtx(context.Background(), RegisterCacheKey+in.Phone, &redisValue)
	if err != nil {
		return nil, errors.New("验证码获取错误")
	}
	if in.Code != redisValue {
		return nil, errors.New("验证码输入错误")
	}
	//3.验证码通过,进行注册
	// 1.检查手机号是否已注册
	mem, err := rl.MemberDomain.FindByPhone(context.Background(), in.Phone)
	if err != nil {
		return nil, errors.New("服务异常,稍后重试")
	}
	if mem != nil {
		return nil, errors.New("此手机号已经被注册")
	}
	//4.生成member模型,存入数据库
	//4. 生成member模型，存入数据库
	err = rl.MemberDomain.Register(
		context.Background(),
		in.Phone,
		in.Password,
		in.Username,
		in.Country,
		in.SuperPartner,
		in.Promotion)
	if err != nil {
		return nil, errors.New("注册失败")
	}
	return &register.RegRes{}, nil
}
func (rl *RegisterLogic) SendCode(req *register.CodeReq) (*register.NoRes, error) {
	//* 收到手机号和国家标识
	//* 生成验证码
	//* 根据对应的国家和手机号调用对应的短信平台发送验证码
	//* 将验证码存入redis，过期时间5分钟
	//* 返回成功
	code := tools.Rand4Num()
	//假设调用短信平台发送验证码成功
	go func() {
		logx.Info("调用短信平台发送验证码成功")
	}()
	logx.Infof("验证码为：%s \n", code)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := rl.svcCtx.Cache.SetWithExpireCtx(ctx, RegisterCacheKey+req.Phone, code, 15*time.Minute)
	if err != nil {
		return nil, errors.New("验证码发送失败")
	}
	return &register.NoRes{}, nil
}
