package url

import (
	"errors"
	"fmt"
	"strings"
)

type URL struct {
	Scheme string
	Host   string
	Path   string
}

func Parse(rawUrl string) (*URL, error) {
	schema, rest, ok := strings.Cut(rawUrl, "://")
	if !ok {
		return nil, errors.New("missing Scheme")
	}
	host, path, _ := strings.Cut(rest, "/")
	return &URL{
		Scheme: schema,
		Host:   host,
		Path:   path,
	}, nil
}

func (u *URL) String() string {
	return fmt.Sprintf("%s://%s/%s", u.Scheme, u.Host, u.Path)
}
