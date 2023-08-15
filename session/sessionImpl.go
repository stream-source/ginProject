package session

import (
	"sync"
	"time"
)

type CustomSession struct {
	ID string
	//此处定义成map类型，方便扩展session数据信息
	Values         map[interface{}]interface{}
	LastAccessTime int64
}

func (s *CustomSession) Set(key, value interface{}) error {
	s.Values[key] = value
	return nil
}

func (s *CustomSession) Get(key interface{}) interface{} {
	if value, exists := s.Values[key]; exists {
		return value
	}
	return nil
}

func (s *CustomSession) Delete(key interface{}) error {
	delete(s.Values, key)
	return nil
}

func (s *CustomSession) SessionID() string {
	return s.ID
}

// CustomProvider 实现了 Provider 接口
type CustomProvider struct {
	//建立sessionId与session之关系
	Sessions map[string]Session
	lock     sync.Mutex
}

func (p *CustomProvider) SessionInit(sid string) (Session, error) {
	//p.lock.Lock()
	//defer p.lock.Unlock()
	//定义session value Map 类型
	sessionValue := make(map[interface{}]interface{})
	sessionValue[sid] = sid
	newSession := &CustomSession{
		ID:     sid,
		Values: sessionValue,
	}
	//或者可以优化调用session#Set方法
	p.Sessions[sid] = newSession

	return newSession, nil
}

func (p *CustomProvider) SessionRead(sid string) (Session, error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if session, exists := p.Sessions[sid]; exists {
		return session, nil
	}

	return p.SessionInit(sid)
}

func (p *CustomProvider) SessionDestroy(sid string) error {
	p.lock.Lock()
	defer p.lock.Unlock()

	if session, exists := p.Sessions[sid]; exists {
		delete(p.Sessions, sid)
		session.Delete(session.SessionID()) // 清除 session 中的数据
	}

	return nil
}

func (p *CustomProvider) SessionGC(maxLifeTime int64) {
	p.lock.Lock()
	defer p.lock.Unlock()

	for sid, session := range p.Sessions {
		// 检查 session 是否过期
		if time.Now().Unix()-session.GetLastAccessTime() > maxLifeTime {
			delete(p.Sessions, sid)
			session.Delete(session.SessionID()) // 清除 session 中的数据
		}
	}
}
func (s *CustomSession) GetLastAccessTime() int64 {
	// 这里假设 Session 结构体中有一个 LastAccessTime 字段
	return s.LastAccessTime
}
