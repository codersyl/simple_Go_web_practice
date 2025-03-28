// 向 Encoder 中写入 JSON 数据

package json // import "encoding/json"

func (enc *Encoder) Encode(v any) error
// Encode writes the JSON encoding of v to the stream, with insignificant space
// characters elided, followed by a newline character.

// See the documentation for Marshal for details about the conversion of Go
// values to JSON.