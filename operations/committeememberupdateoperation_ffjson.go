// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: committeememberupdateoperation.go

package operations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bitcard/bitshares/types"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *CommitteeMemberUpdateOperation) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *CommitteeMemberUpdateOperation) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	/* Struct fall back. type=types.CommitteeMember kind=struct */
	buf.WriteString(`{ "committee_member":`)
	err = buf.Encode(&j.CommitteeMember)
	if err != nil {
		return err
	}
	buf.WriteString(`,"committee_member_account":`)

	{

		obj, err = j.CommitteeMemberAccount.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteByte(',')
	if j.NewURL != nil {
		if true {
			buf.WriteString(`"new_url":`)

			{

				obj, err = j.NewURL.MarshalJSON()
				if err != nil {
					return err
				}
				buf.Write(obj)

			}
			buf.WriteByte(',')
		}
	}
	if j.Fee != nil {
		if true {
			/* Struct fall back. type=types.AssetAmount kind=struct */
			buf.WriteString(`"fee":`)
			err = buf.Encode(j.Fee)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtCommitteeMemberUpdateOperationbase = iota
	ffjtCommitteeMemberUpdateOperationnosuchkey

	ffjtCommitteeMemberUpdateOperationCommitteeMember

	ffjtCommitteeMemberUpdateOperationCommitteeMemberAccount

	ffjtCommitteeMemberUpdateOperationNewURL

	ffjtCommitteeMemberUpdateOperationFee
)

var ffjKeyCommitteeMemberUpdateOperationCommitteeMember = []byte("committee_member")

var ffjKeyCommitteeMemberUpdateOperationCommitteeMemberAccount = []byte("committee_member_account")

var ffjKeyCommitteeMemberUpdateOperationNewURL = []byte("new_url")

var ffjKeyCommitteeMemberUpdateOperationFee = []byte("fee")

// UnmarshalJSON umarshall json - template of ffjson
func (j *CommitteeMemberUpdateOperation) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *CommitteeMemberUpdateOperation) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtCommitteeMemberUpdateOperationbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtCommitteeMemberUpdateOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'c':

					if bytes.Equal(ffjKeyCommitteeMemberUpdateOperationCommitteeMember, kn) {
						currentKey = ffjtCommitteeMemberUpdateOperationCommitteeMember
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyCommitteeMemberUpdateOperationCommitteeMemberAccount, kn) {
						currentKey = ffjtCommitteeMemberUpdateOperationCommitteeMemberAccount
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'f':

					if bytes.Equal(ffjKeyCommitteeMemberUpdateOperationFee, kn) {
						currentKey = ffjtCommitteeMemberUpdateOperationFee
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'n':

					if bytes.Equal(ffjKeyCommitteeMemberUpdateOperationNewURL, kn) {
						currentKey = ffjtCommitteeMemberUpdateOperationNewURL
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeyCommitteeMemberUpdateOperationFee, kn) {
					currentKey = ffjtCommitteeMemberUpdateOperationFee
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyCommitteeMemberUpdateOperationNewURL, kn) {
					currentKey = ffjtCommitteeMemberUpdateOperationNewURL
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyCommitteeMemberUpdateOperationCommitteeMemberAccount, kn) {
					currentKey = ffjtCommitteeMemberUpdateOperationCommitteeMemberAccount
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyCommitteeMemberUpdateOperationCommitteeMember, kn) {
					currentKey = ffjtCommitteeMemberUpdateOperationCommitteeMember
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtCommitteeMemberUpdateOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtCommitteeMemberUpdateOperationCommitteeMember:
					goto handle_CommitteeMember

				case ffjtCommitteeMemberUpdateOperationCommitteeMemberAccount:
					goto handle_CommitteeMemberAccount

				case ffjtCommitteeMemberUpdateOperationNewURL:
					goto handle_NewURL

				case ffjtCommitteeMemberUpdateOperationFee:
					goto handle_Fee

				case ffjtCommitteeMemberUpdateOperationnosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_CommitteeMember:

	/* handler: j.CommitteeMember type=types.CommitteeMember kind=struct quoted=false*/

	{
		/* Falling back. type=types.CommitteeMember kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.CommitteeMember)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CommitteeMemberAccount:

	/* handler: j.CommitteeMemberAccount type=types.AccountID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.CommitteeMemberAccount.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_NewURL:

	/* handler: j.NewURL type=types.String kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			j.NewURL = nil

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			if j.NewURL == nil {
				j.NewURL = new(types.String)
			}

			err = j.NewURL.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Fee:

	/* handler: j.Fee type=types.AssetAmount kind=struct quoted=false*/

	{
		/* Falling back. type=types.AssetAmount kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.Fee)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}
