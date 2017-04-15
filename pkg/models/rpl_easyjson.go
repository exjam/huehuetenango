// AUTOGENERATED FILE: easyjson marshaler/unmarshalers.

package models

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

func easyjsonD0dbfd34DecodeGithubComDecafEmuHuehuetenangoPkgModels(in *jlexer.Lexer, out *RPL) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ID":
			out.ID = RPLID(in.String())
		case "Name":
			out.Name = string(in.String())
		case "IsRPX":
			out.IsRPX = bool(in.Bool())
		case "TitleID":
			out.TitleID = TitleID(in.Uint64())
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
func easyjsonD0dbfd34EncodeGithubComDecafEmuHuehuetenangoPkgModels(out *jwriter.Writer, in RPL) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"ID\":")
	out.String(string(in.ID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"Name\":")
	out.String(string(in.Name))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"IsRPX\":")
	out.Bool(bool(in.IsRPX))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"TitleID\":")
	out.Uint64(uint64(in.TitleID))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RPL) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD0dbfd34EncodeGithubComDecafEmuHuehuetenangoPkgModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RPL) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD0dbfd34EncodeGithubComDecafEmuHuehuetenangoPkgModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RPL) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD0dbfd34DecodeGithubComDecafEmuHuehuetenangoPkgModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RPL) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD0dbfd34DecodeGithubComDecafEmuHuehuetenangoPkgModels(l, v)
}
