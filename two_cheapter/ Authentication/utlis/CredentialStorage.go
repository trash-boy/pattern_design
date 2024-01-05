package utlis

type CredentialStorage interface {
	getPasswordByAppId(appId int)string
}

type MyStorage struct {
	count map[int]string
}



func (s *MyStorage)getPasswordByAppId(appId int)string{
	if val,ok := s.count[appId];ok{
		return val
	}
	return "liuyang"
}
