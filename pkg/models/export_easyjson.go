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

func easyjson4bb85eceDecodeGithubComDecafEmuHuehuetenangoPkgModels(in *jlexer.Lexer, out *Export) {
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
			out.ID = ExportID(in.String())
		case "TitleID":
			out.TitleID = TitleID(in.Uint64())
		case "RPLID":
			out.RPLID = RPLID(in.String())
		case "Type":
			out.Type = ObjectType(in.Int())
		case "Name":
			out.Name = string(in.String())
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
func easyjson4bb85eceEncodeGithubComDecafEmuHuehuetenangoPkgModels(out *jwriter.Writer, in Export) {
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
	out.RawString("\"TitleID\":")
	out.Uint64(uint64(in.TitleID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"RPLID\":")
	out.String(string(in.RPLID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"Type\":")
	out.Int(int(in.Type))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"Name\":")
	out.String(string(in.Name))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Export) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4bb85eceEncodeGithubComDecafEmuHuehuetenangoPkgModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Export) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4bb85eceEncodeGithubComDecafEmuHuehuetenangoPkgModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Export) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4bb85eceDecodeGithubComDecafEmuHuehuetenangoPkgModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Export) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4bb85eceDecodeGithubComDecafEmuHuehuetenangoPkgModels(l, v)
}
