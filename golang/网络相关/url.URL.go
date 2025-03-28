package url // import "net/url"

type URL struct {
	Scheme      string
	Opaque      string    // encoded opaque data
	User        *Userinfo // username and password information
	Host        string    // host or host:port (see Hostname and Port methods)
	Path        string    // path (relative paths may omit leading slash)
	RawPath     string    // encoded path hint (see EscapedPath method)
	OmitHost    bool      // do not emit empty host (authority)
	ForceQuery  bool      // append a query ('?') even if RawQuery is empty
	RawQuery    string    // encoded query values, without '?'
	Fragment    string    // fragment for references, without '#'
	RawFragment string    // encoded fragment hint (see EscapedFragment method)
}
    // A URL represents a parsed URL (technically, a URI reference).

    // The general form represented is:

    //     [scheme:][//[userinfo@]host][/]path[?query][#fragment]

    // URLs that do not start with a slash after the scheme are interpreted as:

    //     scheme:opaque[?query][#fragment]

    // The Host field contains the host and port subcomponents of the URL.
    // When the port is present, it is separated from the host with a colon.
    // When the host is an IPv6 address, it must be enclosed in square brackets:
    // "[fe80::1]:80". The net.JoinHostPort function combines a host and port into
    // a string suitable for the Host field, adding square brackets to the host
    // when necessary.

    // Note that the Path field is stored in decoded form: /%47%6f%2f becomes /Go/.
    // A consequence is that it is impossible to tell which slashes in the Path
    // were slashes in the raw URL and which were %2f. This distinction is rarely
    // important, but when it is, the code should use the URL.EscapedPath method,
    // which preserves the original encoding of Path.

    // The RawPath field is an optional field which is only set when the default
    // encoding of Path is different from the escaped path. See the EscapedPath
    // method for more details.

    // URL's String method uses the EscapedPath method to obtain the path.
