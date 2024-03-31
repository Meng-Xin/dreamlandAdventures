package shared

type AUTHORIZATION = string

const (
	// Token TODO 当前的Token值是死的，后续由HTTP登录服务生成，并存入Redis共享。
	Token AUTHORIZATION = "feature_filter_example"
	// AuthKey 是一个固定值，在元数据中传输。
	AuthKey AUTHORIZATION = "authorization"
)
