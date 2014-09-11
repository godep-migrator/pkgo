package anoa

type User struct {
	Provider string
	UID      string
	Email    string
	Nickname string
	Name     string
	Location string
	Image    string
	URLS     map[string]string
	Extra    map[string]string
	Raw      string
}
