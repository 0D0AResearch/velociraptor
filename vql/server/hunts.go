// +build server_vql

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
// VQL plugins for running on the server.

package server

import (
	"context"
	"path"

	api_proto "www.velocidex.com/golang/velociraptor/api/proto"
	"www.velocidex.com/golang/velociraptor/artifacts"
	config_proto "www.velocidex.com/golang/velociraptor/config/proto"
	"www.velocidex.com/golang/velociraptor/constants"
	"www.velocidex.com/golang/velociraptor/datastore"
	"www.velocidex.com/golang/velociraptor/file_store"
	"www.velocidex.com/golang/velociraptor/file_store/csv"
	"www.velocidex.com/golang/velociraptor/flows"
	"www.velocidex.com/golang/velociraptor/services"
	vql_subsystem "www.velocidex.com/golang/velociraptor/vql"
	"www.velocidex.com/golang/vfilter"
)

type HuntsPluginArgs struct{}

type HuntsPlugin struct{}

func (self HuntsPlugin) Call(
	ctx context.Context,
	scope *vfilter.Scope,
	args *vfilter.Dict) <-chan vfilter.Row {
	output_chan := make(chan vfilter.Row)
	go func() {
		defer close(output_chan)

		arg := &HuntsPluginArgs{}
		err := vfilter.ExtractArgs(scope, args, arg)
		if err != nil {
			scope.Log("hunts: %v", err)
			return
		}

		any_config_obj, _ := scope.Resolve("server_config")
		config_obj, ok := any_config_obj.(*config_proto.Config)
		if !ok {
			scope.Log("Command can only run on the server")
			return
		}

		db, err := datastore.GetDB(config_obj)
		if err != nil {
			scope.Log("Error: %v", err)
			return
		}

		hunts, err := db.ListChildren(config_obj, constants.HUNTS_URN, 0, 100)
		if err != nil {
			scope.Log("Error: %v", err)
			return
		}

		for _, hunt_urn := range hunts {
			hunt_obj := &api_proto.Hunt{}
			err = db.GetSubject(config_obj, hunt_urn, hunt_obj)
			if err != nil {
				continue
			}

			// Re-read the stats into the hunt object.
			hunt_stats := &api_proto.HuntStats{}
			err := db.GetSubject(config_obj, hunt_urn+"/stats", hunt_stats)
			if err == nil {
				hunt_obj.Stats = hunt_stats
			}

			output_chan <- hunt_obj
		}
	}()

	return output_chan
}

func (self HuntsPlugin) Info(scope *vfilter.Scope, type_map *vfilter.TypeMap) *vfilter.PluginInfo {
	return &vfilter.PluginInfo{
		Name:    "hunts",
		Doc:     "Retrieve the list of hunts.",
		RowType: type_map.AddType(scope, &api_proto.ApiClient{}),
		ArgType: type_map.AddType(scope, &HuntsPluginArgs{}),
	}
}

type HuntResultsPluginArgs struct {
	Artifact string `vfilter:"required,field=artifact,doc=The artifact to retrieve"`
	Source   string `vfilter:"optional,field=source,doc=An optional source within the artifact."`
	HuntId   string `vfilter:"required,field=hunt_id,doc=The hunt id to read."`
	Brief    bool   `vfilter:"optional,field=brief,doc=If set we return less columns."`
}

type HuntResultsPlugin struct{}

func (self HuntResultsPlugin) Call(
	ctx context.Context,
	scope *vfilter.Scope,
	args *vfilter.Dict) <-chan vfilter.Row {
	output_chan := make(chan vfilter.Row)
	go func() {
		defer close(output_chan)

		arg := &HuntResultsPluginArgs{}
		err := vfilter.ExtractArgs(scope, args, arg)
		if err != nil {
			scope.Log("hunt_results: %v", err)
			return
		}

		any_config_obj, _ := scope.Resolve("server_config")
		config_obj, ok := any_config_obj.(*config_proto.Config)
		if !ok {
			scope.Log("Command can only run on the server")
			return
		}

		// Backwards compatibility.
		hunt_id := path.Base(arg.HuntId)
		file_path := path.Join("hunts", hunt_id+".csv")
		file_store_factory := file_store.GetFileStore(config_obj)
		fd, err := file_store_factory.ReadFile(file_path)
		if err != nil {
			scope.Log("Error %v: %v\n", err, file_path)
			return
		}
		defer fd.Close()

		// Read each CSV file and emit it with
		// some extra columns for context.
		for row := range csv.GetCSVReader(fd) {
			participation_row := &services.ParticipationRecord{}
			err := vfilter.ExtractArgs(scope, row, participation_row)
			if err != nil {
				return
			}

			if participation_row.Participate {
				flow_obj, err := flows.GetAFF4FlowObject(
					config_obj, participation_row.FlowId)
				if err != nil {
					continue
				}

				// Read individual flow's
				// results. Artifacts are by
				// definition client artifacts - hunts
				// only run on client artifacts.
				flow_id := path.Base(participation_row.FlowId)
				result_path := artifacts.GetCSVPath(
					participation_row.ClientId, "",
					flow_id, arg.Artifact, arg.Source,
					artifacts.MODE_CLIENT)
				fd, err := file_store_factory.ReadFile(result_path)
				if err != nil {
					continue
				}
				defer fd.Close()

				// Read each CSV file and emit it with
				// some extra columns for context.
				for row := range csv.GetCSVReader(fd) {
					value := row.
						Set("FlowId", flow_id).
						Set("ClientId",
							participation_row.ClientId).
						Set("Fqdn",
							participation_row.Fqdn)

					if !arg.Brief {
						value.
							Set("HuntId",
								participation_row.HuntId).
							Set("FlowContext",
								flow_obj.FlowContext)
					}
					output_chan <- value
				}
			}
		}
	}()

	return output_chan
}

func (self HuntResultsPlugin) Info(scope *vfilter.Scope, type_map *vfilter.TypeMap) *vfilter.PluginInfo {
	return &vfilter.PluginInfo{
		Name:    "hunt_results",
		Doc:     "Retrieve the results of a hunt.",
		RowType: type_map.AddType(scope, &api_proto.ApiClient{}),
		ArgType: type_map.AddType(scope, &HuntResultsPluginArgs{}),
	}
}

type HuntFlowsPluginArgs struct {
	HuntId string `vfilter:"required,field=hunt_id,doc=The hunt id to inspect."`
}

type HuntFlowsPlugin struct{}

func (self HuntFlowsPlugin) Call(
	ctx context.Context,
	scope *vfilter.Scope,
	args *vfilter.Dict) <-chan vfilter.Row {
	output_chan := make(chan vfilter.Row)
	go func() {
		defer close(output_chan)

		arg := &HuntFlowsPluginArgs{}
		err := vfilter.ExtractArgs(scope, args, arg)
		if err != nil {
			scope.Log("hunt_flows: %v", err)
			return
		}

		any_config_obj, _ := scope.Resolve("server_config")
		config_obj, ok := any_config_obj.(*config_proto.Config)
		if !ok {
			scope.Log("Command can only run on the server")
			return
		}

		file_path := path.Join("hunts", arg.HuntId+".csv")
		file_store_factory := file_store.GetFileStore(config_obj)
		fd, err := file_store_factory.ReadFile(file_path)
		if err != nil {
			scope.Log("Error %v: %v\n", err, file_path)
			return
		}
		defer fd.Close()

		// Read each CSV file and emit it with
		// some extra columns for context.
		for row := range csv.GetCSVReader(fd) {
			participation_row := &services.ParticipationRecord{}
			err := vfilter.ExtractArgs(scope, row, participation_row)
			if err != nil {
				return
			}

			result := vfilter.NewDict().
				Set("HuntId", participation_row.HuntId).
				Set("ClientId", participation_row.ClientId).
				Set("Fqdn", participation_row.Fqdn).
				Set("Participate", participation_row.Participate).
				Set("Flow", vfilter.Null{})

			if participation_row.Participate {
				flow_obj, err := flows.GetAFF4FlowObject(
					config_obj, participation_row.FlowId)
				if err == nil {
					result.Set("Flow", flow_obj)
				}
			}

			output_chan <- result
		}
	}()

	return output_chan
}

func (self HuntFlowsPlugin) Info(scope *vfilter.Scope, type_map *vfilter.TypeMap) *vfilter.PluginInfo {
	return &vfilter.PluginInfo{
		Name:    "hunt_flows",
		Doc:     "Retrieve the flows launched by a hunt.",
		ArgType: type_map.AddType(scope, &HuntFlowsPluginArgs{}),
	}
}

func init() {
	vql_subsystem.RegisterPlugin(&HuntsPlugin{})
	vql_subsystem.RegisterPlugin(&HuntResultsPlugin{})
	vql_subsystem.RegisterPlugin(&HuntFlowsPlugin{})
}
