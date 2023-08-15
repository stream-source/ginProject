package session

type Provider interface {
	//SessionInit 函数实现 Session 的初始化，操作成功则返回此新的 Session 变量
	SessionInit(sid string) (Session, error)
	//SessionRead 函数返回 sid 所代表的 Session 变量，如果不存在，那么将以 sid 为参数调用SessionInit 函数创建并返回一个新的 Session 变量
	SessionRead(sid string) (Session, error)
	//SessionDestroy 函数用来销毁 sid 对应的 Session 变量
	SessionDestroy(sid string) error
	//SessionGC 根据 maxLifeTime 来删除过期的数据
	SessionGC(maxLifeTime int64)
}

type Session interface {
	Set(key, value interface{}) error // set session value
	Get(key interface{}) interface{}  // get session value
	Delete(key interface{}) error     // delete session value
	SessionID() string
	GetLastAccessTime() int64 // back current sessionID
}
