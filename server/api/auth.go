package api

type Auth struct {
	UserId string
	Email  string
}

func Verify(authorization string) Auth {
	return Auth{}
}
