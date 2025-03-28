// 给头部设置 k-v 键值对

package http // import "net/http"

func (h Header) Set(key, value string)
    Set sets the header entries associated with key to the single element value.
    It replaces any existing values associated with key. The key is case
    insensitive; it is canonicalized by textproto.CanonicalMIMEHeaderKey.
    To use non-canonical keys, assign to the map directly.