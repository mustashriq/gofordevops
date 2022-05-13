package structure_methods

type Avatar struct {
	URL  string
	size int64
}

type Client struct {
	ID   int64
	Name string
	Age  int
	IMG  *Avatar
}

func (c Client) HasAvatar() bool {
	if c.IMG != nil && c.IMG.URL != "" {
		return true
	}
	return false
}
