/*
   Velociraptor - Hunting Evil
   Copyright (C) 2019 Velocidex Innovations.

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/
package functions

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/binary"
	"strconv"
	"strings"
	"time"
	"unicode/utf16"

	vql_subsystem "www.velocidex.com/golang/velociraptor/vql"
	vfilter "www.velocidex.com/golang/vfilter"
)

type _Base64DecodeArgs struct {
	String string `vfilter:"required,field=string,doc=A string to decode"`
}

type _Base64Decode struct{}

func (self _Base64Decode) Call(
	ctx context.Context,
	scope *vfilter.Scope,
	args *vfilter.Dict) vfilter.Any {
	arg := &_Base64DecodeArgs{}
	err := vfilter.ExtractArgs(scope, args, arg)
	if err != nil {
		scope.Log("base64decode: %s", err.Error())
		return vfilter.Null{}
	}

	result, err := base64.StdEncoding.DecodeString(arg.String)
	if err != nil {
		return vfilter.Null{}
	}
	return string(result)
}

func (self _Base64Decode) Info(scope *vfilter.Scope, type_map *vfilter.TypeMap) *vfilter.FunctionInfo {
	return &vfilter.FunctionInfo{
		Name:    "base64decode",
		ArgType: type_map.AddType(scope, &_Base64DecodeArgs{}),
	}
}

type _Base64EncodeArgs struct {
	String string `vfilter:"required,field=string,doc=A string to decode"`
}

type _Base64Encode struct{}

func (self _Base64Encode) Call(
	ctx context.Context,
	scope *vfilter.Scope,
	args *vfilter.Dict) vfilter.Any {
	arg := &_Base64EncodeArgs{}
	err := vfilter.ExtractArgs(scope, args, arg)
	if err != nil {
		scope.Log("base64encode: %s", err.Error())
		return vfilter.Null{}
	}

	result := base64.StdEncoding.EncodeToString(
		[]byte(arg.String))
	return string(result)
}

func (self _Base64Encode) Info(scope *vfilter.Scope, type_map *vfilter.TypeMap) *vfilter.FunctionInfo {
	return &vfilter.FunctionInfo{
		Name:    "base64encode",
		ArgType: type_map.AddType(scope, &_Base64EncodeArgs{}),
	}
}

type _ToLowerArgs struct {
	String string `vfilter:"required,field=string,doc=A string to lower"`
}

type _ToLower struct{}

func (self _ToLower) Call(
	ctx context.Context,
	scope *vfilter.Scope,
	args *vfilter.Dict) vfilter.Any {
	arg := &_ToLowerArgs{}
	err := vfilter.ExtractArgs(scope, args, arg)
	if err != nil {
		scope.Log("lowcase: %s", err.Error())
		return vfilter.Null{}
	}

	return strings.ToLower(arg.String)
}

func (self _ToLower) Info(scope *vfilter.Scope, type_map *vfilter.TypeMap) *vfilter.FunctionInfo {
	return &vfilter.FunctionInfo{
		Name:    "lowcase",
		ArgType: type_map.AddType(scope, &_ToLowerArgs{}),
	}
}

type _ToUpper struct{}

func (self _ToUpper) Call(
	ctx context.Context,
	scope *vfilter.Scope,
	args *vfilter.Dict) vfilter.Any {
	arg := &_ToLowerArgs{}
	err := vfilter.ExtractArgs(scope, args, arg)
	if err != nil {
		scope.Log("upcase: %s", err.Error())
		return vfilter.Null{}
	}

	return strings.ToUpper(arg.String)
}

func (self _ToUpper) Info(scope *vfilter.Scope, type_map *vfilter.TypeMap) *vfilter.FunctionInfo {
	return &vfilter.FunctionInfo{
		Name:    "upcase",
		ArgType: type_map.AddType(scope, &_ToLowerArgs{}),
	}
}

type _ToIntArgs struct {
	String string `vfilter:"required,field=string,doc=A string to convert to int"`
}

type _ToInt struct{}

func (self _ToInt) Call(
	ctx context.Context,
	scope *vfilter.Scope,
	args *vfilter.Dict) vfilter.Any {
	arg := &_ToIntArgs{}
	err := vfilter.ExtractArgs(scope, args, arg)
	if err != nil {
		scope.Log("atoi: %s", err.Error())
		return vfilter.Null{}
	}

	result, _ := strconv.Atoi(arg.String)
	return result
}

func (self _ToInt) Info(scope *vfilter.Scope, type_map *vfilter.TypeMap) *vfilter.FunctionInfo {
	return &vfilter.FunctionInfo{
		Name:    "atoi",
		Doc:     "Convert a string to an int.",
		ArgType: type_map.AddType(scope, &_ToIntArgs{}),
	}
}

type _ParseFloat struct{}

func (self _ParseFloat) Call(
	ctx context.Context,
	scope *vfilter.Scope,
	args *vfilter.Dict) vfilter.Any {
	arg := &_ToIntArgs{}
	err := vfilter.ExtractArgs(scope, args, arg)
	if err != nil {
		scope.Log("atoi: %s", err.Error())
		return vfilter.Null{}
	}

	result, _ := strconv.ParseFloat(arg.String, 64)
	return result
}

func (self _ParseFloat) Info(scope *vfilter.Scope, type_map *vfilter.TypeMap) *vfilter.FunctionInfo {
	return &vfilter.FunctionInfo{
		Name:    "parse_float",
		Doc:     "Convert a string to a float.",
		ArgType: type_map.AddType(scope, &_ToIntArgs{}),
	}
}

type _Now struct{}

func (self _Now) Call(
	ctx context.Context,
	scope *vfilter.Scope,
	args *vfilter.Dict) vfilter.Any {
	return time.Now().Unix()
}

func (self _Now) Info(scope *vfilter.Scope, type_map *vfilter.TypeMap) *vfilter.FunctionInfo {
	return &vfilter.FunctionInfo{
		Name:    "now",
		Doc:     "Returns current time in seconds since epoch.",
		ArgType: type_map.AddType(scope, &_ToIntArgs{}),
	}
}

type _UTF16 struct{}

func (self _UTF16) Call(
	ctx context.Context,
	scope *vfilter.Scope,
	args *vfilter.Dict) vfilter.Any {

	arg := &_Base64DecodeArgs{}
	err := vfilter.ExtractArgs(scope, args, arg)
	if err != nil {
		scope.Log("utf16: %s", err.Error())
		return vfilter.Null{}
	}

	ints := make([]uint16, len(arg.String)/2)
	if err := binary.Read(bytes.NewReader([]byte(arg.String)), binary.LittleEndian, &ints); err != nil {
		scope.Log("utf16: %s", err.Error())
		return vfilter.Null{}
	}

	return string(utf16.Decode(ints))
}

func (self _UTF16) Info(scope *vfilter.Scope, type_map *vfilter.TypeMap) *vfilter.FunctionInfo {
	return &vfilter.FunctionInfo{
		Name:    "utf16",
		Doc:     "Parse input from utf16.",
		ArgType: type_map.AddType(scope, &_Base64DecodeArgs{}),
	}
}

type _Scope struct{}

func (self _Scope) Call(
	ctx context.Context,
	scope *vfilter.Scope,
	args *vfilter.Dict) vfilter.Any {

	return scope
}

func (self _Scope) Info(scope *vfilter.Scope, type_map *vfilter.TypeMap) *vfilter.FunctionInfo {
	return &vfilter.FunctionInfo{
		Name: "scope",
		Doc:  "return the scope.",
	}
}

func init() {
	vql_subsystem.RegisterFunction(&_Base64Decode{})
	vql_subsystem.RegisterFunction(&_Base64Encode{})
	vql_subsystem.RegisterFunction(&_Scope{})
	vql_subsystem.RegisterFunction(&_ToInt{})
	vql_subsystem.RegisterFunction(&_ParseFloat{})
	vql_subsystem.RegisterFunction(&_Now{})
	vql_subsystem.RegisterFunction(&_ToLower{})
	vql_subsystem.RegisterFunction(&_ToUpper{})
	vql_subsystem.RegisterFunction(&_UTF16{})
}
