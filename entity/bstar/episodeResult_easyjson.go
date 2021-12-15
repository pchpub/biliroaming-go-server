// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package bstar

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

func easyjsonE6b67dbaDecodeGithubComJasonKhew96BiliroamingGoServerEntityBstar(in *jlexer.Lexer, out *EpisodeResultData) {
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
		case "subtitle_suggest_key":
			out.SubtitleSuggestKey = string(in.String())
		case "jump":
			if m, ok := out.Jump.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.Jump.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.Jump = in.Interface()
			}
		case "subtitles":
			if in.IsNull() {
				in.Skip()
				out.Subtitles = nil
			} else {
				in.Delim('[')
				if out.Subtitles == nil {
					if !in.IsDelim(']') {
						out.Subtitles = make([]Subtitles, 0, 1)
					} else {
						out.Subtitles = []Subtitles{}
					}
				} else {
					out.Subtitles = (out.Subtitles)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Subtitles
					(v1).UnmarshalEasyJSON(in)
					out.Subtitles = append(out.Subtitles, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonE6b67dbaEncodeGithubComJasonKhew96BiliroamingGoServerEntityBstar(out *jwriter.Writer, in EpisodeResultData) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"subtitle_suggest_key\":"
		out.RawString(prefix[1:])
		out.String(string(in.SubtitleSuggestKey))
	}
	{
		const prefix string = ",\"jump\":"
		out.RawString(prefix)
		if m, ok := in.Jump.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.Jump.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.Jump))
		}
	}
	{
		const prefix string = ",\"subtitles\":"
		out.RawString(prefix)
		if in.Subtitles == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Subtitles {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EpisodeResultData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE6b67dbaEncodeGithubComJasonKhew96BiliroamingGoServerEntityBstar(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EpisodeResultData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE6b67dbaEncodeGithubComJasonKhew96BiliroamingGoServerEntityBstar(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EpisodeResultData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE6b67dbaDecodeGithubComJasonKhew96BiliroamingGoServerEntityBstar(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EpisodeResultData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE6b67dbaDecodeGithubComJasonKhew96BiliroamingGoServerEntityBstar(l, v)
}
func easyjsonE6b67dbaDecodeGithubComJasonKhew96BiliroamingGoServerEntityBstar1(in *jlexer.Lexer, out *EpisodeResult) {
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
		case "code":
			out.Code = int(in.Int())
		case "message":
			out.Message = string(in.String())
		case "ttl":
			out.TTL = int(in.Int())
		case "data":
			(out.Data).UnmarshalEasyJSON(in)
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
func easyjsonE6b67dbaEncodeGithubComJasonKhew96BiliroamingGoServerEntityBstar1(out *jwriter.Writer, in EpisodeResult) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"code\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Code))
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"ttl\":"
		out.RawString(prefix)
		out.Int(int(in.TTL))
	}
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix)
		(in.Data).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EpisodeResult) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE6b67dbaEncodeGithubComJasonKhew96BiliroamingGoServerEntityBstar1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EpisodeResult) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE6b67dbaEncodeGithubComJasonKhew96BiliroamingGoServerEntityBstar1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EpisodeResult) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE6b67dbaDecodeGithubComJasonKhew96BiliroamingGoServerEntityBstar1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EpisodeResult) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE6b67dbaDecodeGithubComJasonKhew96BiliroamingGoServerEntityBstar1(l, v)
}
