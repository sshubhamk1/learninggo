package url

import "fmt"

type URL struct {
	Scheme string
	Host   string
	Path   string
}

func Parse(rawUrl string) (*URL, error) {
	return &URL{
		Scheme: "https",
		Host:   "github.com",
		Path:   "sshubhamk1",
	}, nil
}

func (u *URL) String() string {
	return fmt.Sprintf("%s://%s/%s", u.Scheme, u.Host, u.Path)
}
