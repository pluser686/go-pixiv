package pixiv

import (
	"fmt"
	"net/url"
	"strconv"
)

func parseNextPageOffset(s string) (uint64, error) {
	if s == "" {
		return 0, nil
	}
	u, err := url.Parse(s)
	if err != nil {
		return 0, fmt.Errorf("parse next_url error: %s {%s}", s, err)
	}

	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return 0, fmt.Errorf("parse next_url error: %s {%s}", s, err)
	}
	offset, err := strconv.ParseUint(m.Get("max_bookmark_id"), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parse next_url error: %s {%s}", s, err)
	}
	return offset, nil
}
