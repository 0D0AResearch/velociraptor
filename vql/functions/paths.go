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
	"context"

	"www.velocidex.com/golang/velociraptor/utils"
	vql_subsystem "www.velocidex.com/golang/velociraptor/vql"
	"www.velocidex.com/golang/vfilter"
)

type DirnameArgs struct {
	Path string `vfilter:"required,field=path,doc=Extract directory name of path"`
}

type DirnameFunction struct{}

func (self *DirnameFunction) Call(ctx context.Context,
	scope *vfilter.Scope,
	args *vfilter.Dict) vfilter.Any {
	arg := &DirnameArgs{}
	err := vfilter.ExtractArgs(scope, args, arg)
	if err != nil {
		scope.Log("dirname: %s", err.Error())
		return false
	}

	components := utils.SplitComponents(arg.Path)
	if len(components) > 0 {
		result := utils.JoinComponents(components[:len(components)-1], "/")
		if arg.Path[0] == '/' {
			result = "/" + result
		}

		return result
	}
	return vfilter.Null{}
}

func (self DirnameFunction) Info(scope *vfilter.Scope, type_map *vfilter.TypeMap) *vfilter.FunctionInfo {
	return &vfilter.FunctionInfo{
		Name:    "dirname",
		Doc:     "Return the directory path.",
		ArgType: type_map.AddType(scope, &DirnameArgs{}),
	}
}

type BasenameFunction struct{}

func (self *BasenameFunction) Call(ctx context.Context,
	scope *vfilter.Scope,
	args *vfilter.Dict) vfilter.Any {
	arg := &DirnameArgs{}
	err := vfilter.ExtractArgs(scope, args, arg)
	if err != nil {
		scope.Log("basename: %s", err.Error())
		return false
	}

	components := utils.SplitComponents(arg.Path)
	if len(components) > 0 {
		return components[len(components)-1]
	}

	return vfilter.Null{}
}

func (self BasenameFunction) Info(scope *vfilter.Scope, type_map *vfilter.TypeMap) *vfilter.FunctionInfo {
	return &vfilter.FunctionInfo{
		Name:    "basename",
		Doc:     "Return the basename of the path.",
		ArgType: type_map.AddType(scope, &DirnameArgs{}),
	}
}

func init() {
	vql_subsystem.RegisterFunction(&DirnameFunction{})
	vql_subsystem.RegisterFunction(&BasenameFunction{})
}
