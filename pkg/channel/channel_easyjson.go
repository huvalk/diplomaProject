// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package channel

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson5af9c81fDecodeDiplomaProjectPkgChannel(in *jlexer.Lexer, out *Notification) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "type":
			out.Type = string(in.String())
		case "status":
			out.Status = string(in.String())
		case "message":
			out.Message = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson5af9c81fEncodeDiplomaProjectPkgChannel(out *jwriter.Writer, in Notification) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Type != "" {
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type))
	}
	if in.Status != "" {
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Status))
	}
	{
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Notification) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5af9c81fEncodeDiplomaProjectPkgChannel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Notification) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5af9c81fEncodeDiplomaProjectPkgChannel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Notification) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5af9c81fDecodeDiplomaProjectPkgChannel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Notification) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5af9c81fDecodeDiplomaProjectPkgChannel(l, v)
}
