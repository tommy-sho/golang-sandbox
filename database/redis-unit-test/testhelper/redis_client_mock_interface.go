package redis_mock

// このコードではinterfaceを使ったRedisのテストをする
// 参考はhttps://github.com/go-redis/redis/issues/332

// 使用するメソッドをinterfaceとして実装する
type DB interface {
	SetToken()
	GetIDByToken()
}
