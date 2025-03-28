package url // import "net/url"

type Values map[string][]string
// Values maps a string key to a list of values. It is typically used for query
// parameters and form values. Unlike in the http.Header map, the keys in a
// Values map are case-sensitive.
