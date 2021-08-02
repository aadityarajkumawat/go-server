package cookies

/**
1. user registers
2. we take user_id and create a session
3. using the same user_id KEY for session -> `sess:<user_id>`
4. and we also store this user_id in client(cookies)
5. on any request we can get the cookie and check if user is authenticated
*/

type CookieI struct {
	UserSessionID string
	Username      string
	Password      string
}

type CookieStore struct {
	Cookies []CookieI
}

var CookieStoreI = &CookieStore{Cookies: make([]CookieI, 0, 10)}

func (store *CookieStore) GetCookie(key string) string {
	for i := 0; i < len(store.Cookies); i++ {
		if key == store.Cookies[i].UserSessionID {
			return key
		}
	}
	return ""
}

func (store *CookieStore) SetCookie(cookie CookieI) {
	store.Cookies = append(store.Cookies, cookie)
}
