// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

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

func easyjsonD2b7633eDecodeDiplomaProjectApplicationModels(in *jlexer.Lexer, out *UserArr) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(UserArr, 0, 1)
			} else {
				*out = UserArr{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 User
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeDiplomaProjectApplicationModels(out *jwriter.Writer, in UserArr) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v UserArr) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeDiplomaProjectApplicationModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserArr) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeDiplomaProjectApplicationModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserArr) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeDiplomaProjectApplicationModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserArr) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeDiplomaProjectApplicationModels(l, v)
}
func easyjsonD2b7633eDecodeDiplomaProjectApplicationModels1(in *jlexer.Lexer, out *User) {
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
		case "id":
			out.Id = int(in.Int())
		case "first_name":
			out.FirstName = string(in.String())
		case "last_name":
			out.LastName = string(in.String())
		case "email":
			out.Email = string(in.String())
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
func easyjsonD2b7633eEncodeDiplomaProjectApplicationModels1(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"first_name\":"
		out.RawString(prefix)
		out.String(string(in.FirstName))
	}
	{
		const prefix string = ",\"last_name\":"
		out.RawString(prefix)
		out.String(string(in.LastName))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeDiplomaProjectApplicationModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeDiplomaProjectApplicationModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeDiplomaProjectApplicationModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeDiplomaProjectApplicationModels1(l, v)
}
func easyjsonD2b7633eDecodeDiplomaProjectApplicationModels2(in *jlexer.Lexer, out *TeamArr) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(TeamArr, 0, 1)
			} else {
				*out = TeamArr{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v4 Team
			(v4).UnmarshalEasyJSON(in)
			*out = append(*out, v4)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeDiplomaProjectApplicationModels2(out *jwriter.Writer, in TeamArr) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v5, v6 := range in {
			if v5 > 0 {
				out.RawByte(',')
			}
			(v6).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v TeamArr) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeDiplomaProjectApplicationModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TeamArr) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeDiplomaProjectApplicationModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TeamArr) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeDiplomaProjectApplicationModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TeamArr) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeDiplomaProjectApplicationModels2(l, v)
}
func easyjsonD2b7633eDecodeDiplomaProjectApplicationModels3(in *jlexer.Lexer, out *Team) {
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
		case "id":
			out.Id = int(in.Int())
		case "name":
			out.Name = string(in.String())
		case "members":
			(out.Members).UnmarshalEasyJSON(in)
		case "eventid":
			out.EventID = int(in.Int())
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
func easyjsonD2b7633eEncodeDiplomaProjectApplicationModels3(out *jwriter.Writer, in Team) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"members\":"
		out.RawString(prefix)
		(in.Members).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"eventid\":"
		out.RawString(prefix)
		out.Int(int(in.EventID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Team) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeDiplomaProjectApplicationModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Team) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeDiplomaProjectApplicationModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Team) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeDiplomaProjectApplicationModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Team) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeDiplomaProjectApplicationModels3(l, v)
}
func easyjsonD2b7633eDecodeDiplomaProjectApplicationModels4(in *jlexer.Lexer, out *SkillsArr) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(SkillsArr, 0, 1)
			} else {
				*out = SkillsArr{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v7 Skills
			(v7).UnmarshalEasyJSON(in)
			*out = append(*out, v7)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeDiplomaProjectApplicationModels4(out *jwriter.Writer, in SkillsArr) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v8, v9 := range in {
			if v8 > 0 {
				out.RawByte(',')
			}
			(v9).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v SkillsArr) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeDiplomaProjectApplicationModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SkillsArr) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeDiplomaProjectApplicationModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SkillsArr) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeDiplomaProjectApplicationModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SkillsArr) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeDiplomaProjectApplicationModels4(l, v)
}
func easyjsonD2b7633eDecodeDiplomaProjectApplicationModels5(in *jlexer.Lexer, out *Skills) {
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
		case "id":
			out.Id = int(in.Int())
		case "description":
			out.Description = string(in.String())
		case "tags":
			if in.IsNull() {
				in.Skip()
				out.Tags = nil
			} else {
				in.Delim('[')
				if out.Tags == nil {
					if !in.IsDelim(']') {
						out.Tags = make([]string, 0, 4)
					} else {
						out.Tags = []string{}
					}
				} else {
					out.Tags = (out.Tags)[:0]
				}
				for !in.IsDelim(']') {
					var v10 string
					v10 = string(in.String())
					out.Tags = append(out.Tags, v10)
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
func easyjsonD2b7633eEncodeDiplomaProjectApplicationModels5(out *jwriter.Writer, in Skills) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"tags\":"
		out.RawString(prefix)
		if in.Tags == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.Tags {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.String(string(v12))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Skills) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeDiplomaProjectApplicationModels5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Skills) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeDiplomaProjectApplicationModels5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Skills) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeDiplomaProjectApplicationModels5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Skills) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeDiplomaProjectApplicationModels5(l, v)
}
func easyjsonD2b7633eDecodeDiplomaProjectApplicationModels6(in *jlexer.Lexer, out *NotificationArr) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(NotificationArr, 0, 2)
			} else {
				*out = NotificationArr{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v13 Notification
			(v13).UnmarshalEasyJSON(in)
			*out = append(*out, v13)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeDiplomaProjectApplicationModels6(out *jwriter.Writer, in NotificationArr) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v14, v15 := range in {
			if v14 > 0 {
				out.RawByte(',')
			}
			(v15).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v NotificationArr) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeDiplomaProjectApplicationModels6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NotificationArr) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeDiplomaProjectApplicationModels6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NotificationArr) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeDiplomaProjectApplicationModels6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NotificationArr) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeDiplomaProjectApplicationModels6(l, v)
}
func easyjsonD2b7633eDecodeDiplomaProjectApplicationModels7(in *jlexer.Lexer, out *Notification) {
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
		case "userID":
			out.UserID = int(in.Int())
		case "type":
			out.Type = int(in.Int())
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
func easyjsonD2b7633eEncodeDiplomaProjectApplicationModels7(out *jwriter.Writer, in Notification) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"userID\":"
		out.RawString(prefix[1:])
		out.Int(int(in.UserID))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.Int(int(in.Type))
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Notification) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeDiplomaProjectApplicationModels7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Notification) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeDiplomaProjectApplicationModels7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Notification) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeDiplomaProjectApplicationModels7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Notification) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeDiplomaProjectApplicationModels7(l, v)
}
func easyjsonD2b7633eDecodeDiplomaProjectApplicationModels8(in *jlexer.Lexer, out *FeedUser) {
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
		case "id":
			out.Id = int(in.Int())
		case "first_name":
			out.FirstName = string(in.String())
		case "last_name":
			out.LastName = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "team":
			(out.Tm).UnmarshalEasyJSON(in)
		case "job_name":
			out.JobName = string(in.String())
		case "skills":
			if in.IsNull() {
				in.Skip()
				out.Skills = nil
			} else {
				in.Delim('[')
				if out.Skills == nil {
					if !in.IsDelim(']') {
						out.Skills = make([]Skills, 0, 1)
					} else {
						out.Skills = []Skills{}
					}
				} else {
					out.Skills = (out.Skills)[:0]
				}
				for !in.IsDelim(']') {
					var v16 Skills
					(v16).UnmarshalEasyJSON(in)
					out.Skills = append(out.Skills, v16)
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
func easyjsonD2b7633eEncodeDiplomaProjectApplicationModels8(out *jwriter.Writer, in FeedUser) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"first_name\":"
		out.RawString(prefix)
		out.String(string(in.FirstName))
	}
	{
		const prefix string = ",\"last_name\":"
		out.RawString(prefix)
		out.String(string(in.LastName))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"team\":"
		out.RawString(prefix)
		(in.Tm).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"job_name\":"
		out.RawString(prefix)
		out.String(string(in.JobName))
	}
	{
		const prefix string = ",\"skills\":"
		out.RawString(prefix)
		if in.Skills == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v17, v18 := range in.Skills {
				if v17 > 0 {
					out.RawByte(',')
				}
				(v18).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FeedUser) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeDiplomaProjectApplicationModels8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FeedUser) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeDiplomaProjectApplicationModels8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FeedUser) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeDiplomaProjectApplicationModels8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FeedUser) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeDiplomaProjectApplicationModels8(l, v)
}
func easyjsonD2b7633eDecodeDiplomaProjectApplicationModels9(in *jlexer.Lexer, out *FeedArr) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(FeedArr, 0, 1)
			} else {
				*out = FeedArr{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v19 Feed
			(v19).UnmarshalEasyJSON(in)
			*out = append(*out, v19)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeDiplomaProjectApplicationModels9(out *jwriter.Writer, in FeedArr) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v20, v21 := range in {
			if v20 > 0 {
				out.RawByte(',')
			}
			(v21).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v FeedArr) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeDiplomaProjectApplicationModels9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FeedArr) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeDiplomaProjectApplicationModels9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FeedArr) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeDiplomaProjectApplicationModels9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FeedArr) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeDiplomaProjectApplicationModels9(l, v)
}
func easyjsonD2b7633eDecodeDiplomaProjectApplicationModels10(in *jlexer.Lexer, out *Feed) {
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
		case "id":
			out.Id = int(in.Int())
		case "users":
			(out.Users).UnmarshalEasyJSON(in)
		case "event":
			out.Event = int(in.Int())
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
func easyjsonD2b7633eEncodeDiplomaProjectApplicationModels10(out *jwriter.Writer, in Feed) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"users\":"
		out.RawString(prefix)
		(in.Users).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"event\":"
		out.RawString(prefix)
		out.Int(int(in.Event))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Feed) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeDiplomaProjectApplicationModels10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Feed) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeDiplomaProjectApplicationModels10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Feed) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeDiplomaProjectApplicationModels10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Feed) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeDiplomaProjectApplicationModels10(l, v)
}
func easyjsonD2b7633eDecodeDiplomaProjectApplicationModels11(in *jlexer.Lexer, out *EventDB) {
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
		case "id":
			out.Id = int(in.Int())
		case "name":
			out.Name = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "founder":
			out.Founder = int(in.Int())
		case "date_start":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.DateStart).UnmarshalJSON(data))
			}
		case "date_end":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.DateEnd).UnmarshalJSON(data))
			}
		case "state":
			out.State = string(in.String())
		case "place":
			out.Place = string(in.String())
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
func easyjsonD2b7633eEncodeDiplomaProjectApplicationModels11(out *jwriter.Writer, in EventDB) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"founder\":"
		out.RawString(prefix)
		out.Int(int(in.Founder))
	}
	{
		const prefix string = ",\"date_start\":"
		out.RawString(prefix)
		out.Raw((in.DateStart).MarshalJSON())
	}
	{
		const prefix string = ",\"date_end\":"
		out.RawString(prefix)
		out.Raw((in.DateEnd).MarshalJSON())
	}
	{
		const prefix string = ",\"state\":"
		out.RawString(prefix)
		out.String(string(in.State))
	}
	{
		const prefix string = ",\"place\":"
		out.RawString(prefix)
		out.String(string(in.Place))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventDB) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeDiplomaProjectApplicationModels11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventDB) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeDiplomaProjectApplicationModels11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventDB) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeDiplomaProjectApplicationModels11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventDB) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeDiplomaProjectApplicationModels11(l, v)
}
func easyjsonD2b7633eDecodeDiplomaProjectApplicationModels12(in *jlexer.Lexer, out *EventArr) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(EventArr, 0, 0)
			} else {
				*out = EventArr{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v22 Event
			(v22).UnmarshalEasyJSON(in)
			*out = append(*out, v22)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeDiplomaProjectApplicationModels12(out *jwriter.Writer, in EventArr) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v23, v24 := range in {
			if v23 > 0 {
				out.RawByte(',')
			}
			(v24).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v EventArr) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeDiplomaProjectApplicationModels12(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventArr) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeDiplomaProjectApplicationModels12(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventArr) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeDiplomaProjectApplicationModels12(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventArr) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeDiplomaProjectApplicationModels12(l, v)
}
func easyjsonD2b7633eDecodeDiplomaProjectApplicationModels13(in *jlexer.Lexer, out *Event) {
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
		case "id":
			out.Id = int(in.Int())
		case "name":
			out.Name = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "founder":
			out.Founder = int(in.Int())
		case "date_start":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.DateStart).UnmarshalJSON(data))
			}
		case "date_end":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.DateEnd).UnmarshalJSON(data))
			}
		case "state":
			out.State = string(in.String())
		case "place":
			out.Place = string(in.String())
		case "feed":
			(out.Feed).UnmarshalEasyJSON(in)
		case "participants_count":
			out.ParticipantsCount = int(in.Int())
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
func easyjsonD2b7633eEncodeDiplomaProjectApplicationModels13(out *jwriter.Writer, in Event) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"founder\":"
		out.RawString(prefix)
		out.Int(int(in.Founder))
	}
	{
		const prefix string = ",\"date_start\":"
		out.RawString(prefix)
		out.Raw((in.DateStart).MarshalJSON())
	}
	{
		const prefix string = ",\"date_end\":"
		out.RawString(prefix)
		out.Raw((in.DateEnd).MarshalJSON())
	}
	{
		const prefix string = ",\"state\":"
		out.RawString(prefix)
		out.String(string(in.State))
	}
	{
		const prefix string = ",\"place\":"
		out.RawString(prefix)
		out.String(string(in.Place))
	}
	{
		const prefix string = ",\"feed\":"
		out.RawString(prefix)
		(in.Feed).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"participants_count\":"
		out.RawString(prefix)
		out.Int(int(in.ParticipantsCount))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Event) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeDiplomaProjectApplicationModels13(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Event) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeDiplomaProjectApplicationModels13(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Event) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeDiplomaProjectApplicationModels13(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Event) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeDiplomaProjectApplicationModels13(l, v)
}
func easyjsonD2b7633eDecodeDiplomaProjectApplicationModels14(in *jlexer.Lexer, out *AddToUser) {
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
		case "uid1":
			out.UID1 = int(in.Int())
		case "uid2":
			out.UID2 = int(in.Int())
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
func easyjsonD2b7633eEncodeDiplomaProjectApplicationModels14(out *jwriter.Writer, in AddToUser) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"uid1\":"
		out.RawString(prefix[1:])
		out.Int(int(in.UID1))
	}
	{
		const prefix string = ",\"uid2\":"
		out.RawString(prefix)
		out.Int(int(in.UID2))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AddToUser) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeDiplomaProjectApplicationModels14(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AddToUser) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeDiplomaProjectApplicationModels14(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AddToUser) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeDiplomaProjectApplicationModels14(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AddToUser) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeDiplomaProjectApplicationModels14(l, v)
}
func easyjsonD2b7633eDecodeDiplomaProjectApplicationModels15(in *jlexer.Lexer, out *AddToTeam) {
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
		case "tid":
			out.TID = int(in.Int())
		case "uid":
			out.UID = int(in.Int())
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
func easyjsonD2b7633eEncodeDiplomaProjectApplicationModels15(out *jwriter.Writer, in AddToTeam) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"tid\":"
		out.RawString(prefix[1:])
		out.Int(int(in.TID))
	}
	{
		const prefix string = ",\"uid\":"
		out.RawString(prefix)
		out.Int(int(in.UID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AddToTeam) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeDiplomaProjectApplicationModels15(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AddToTeam) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeDiplomaProjectApplicationModels15(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AddToTeam) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeDiplomaProjectApplicationModels15(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AddToTeam) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeDiplomaProjectApplicationModels15(l, v)
}
