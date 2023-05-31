package alipay

const (
	kProductionPublicAppAuthorize = "https://openauth.alipay.com/oauth2/publicAppAuthorize.htm"
	kSandboxPublicAppAuthorize    = "https://openauth.alipaydev.com/oauth2/publicAppAuthorize.htm"
)

const (
	kProductionAppToAppAuth = "https://openauth.alipay.com/oauth2/appToAppAuth.htm"
	kSandboxAppToAppAuth    = "https://openauth.alipaydev.com/oauth2/appToAppAuth.htm"
)

type SystemOauthToken struct {
	AppAuthToken string `json:"-"` // 可选
	GrantType    string `json:"-"` // 值为 authorization_code 时，代表用code换取；值为refresh_token时，代表用refresh_token换取
	Code         string `json:"-"`
	RefreshToken string `json:"-"`
}

func (this SystemOauthToken) APIName() string {
	return "alipay.system.oauth.token"
}

func (this SystemOauthToken) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	m["grant_type"] = this.GrantType
	if this.Code != "" {
		m["code"] = this.Code
	}
	if this.RefreshToken != "" {
		m["refresh_token"] = this.RefreshToken
	}
	return m
}

func (this SystemOauthToken) ExtJSONParamName() string {
	return "biz_content"
}

func (this SystemOauthToken) ExtJSONParamValue() string {
	return marshal(this)
}

func (this *AliPay) GetSystemOauthToken(param SystemOauthToken) (results *SystemOauthTokenRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// SystemOauthTokenRsp 换取授权访问令牌接口请求参数
type SystemOauthTokenRsp struct {
	SystemOauthTokenRsp struct {
		UserId       string `json:"user_id"`
		AccessToken  string `json:"access_token"`
		ExpiresIn    int64  `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		ReExpiresIn  int64  `json:"re_expires_in"`
		AuthStart    string `json:"auth_start"`
	} `json:"alipay_system_oauth_token_response"`
	Sign string `json:"sign"`
}
