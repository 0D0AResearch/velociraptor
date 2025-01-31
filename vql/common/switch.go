package common

import (
	"context"
	"sort"

	vql_subsystem "www.velocidex.com/golang/velociraptor/vql"
	vfilter "www.velocidex.com/golang/vfilter"
)

type _SwitchPlugin struct{}

func (self _SwitchPlugin) Call(ctx context.Context,
	scope *vfilter.Scope,
	args *vfilter.Dict) <-chan vfilter.Row {
	output_chan := make(chan vfilter.Row)

	go func() {
		defer close(output_chan)

		queries := []vfilter.StoredQuery{}
		members := scope.GetMembers(args)
		sort.Strings(members)

		for _, member := range members {
			member_obj, _ := args.Get(member)
			lazy_v, ok := member_obj.(vfilter.LazyExpr)
			if ok {
				member_obj = lazy_v.Reduce()
			}

			query, ok := member_obj.(vfilter.StoredQuery)
			if !ok {
				scope.Log("Parameter " + member +
					" should be a query")
				return
			}
			queries = append(queries, query)
		}

		// Evaluate each query - the first query that returns
		// results will be emitted. We do not evaluate the
		// other queries at all.
		for _, query := range queries {
			found := false
			new_scope := scope.Copy()
			for item := range query.Eval(ctx, new_scope) {
				found = true
				output_chan <- item
			}

			if found {
				return
			}
		}
	}()

	return output_chan

}

func (self _SwitchPlugin) Info(
	scope *vfilter.Scope,
	type_map *vfilter.TypeMap) *vfilter.PluginInfo {
	return &vfilter.PluginInfo{
		Name: "switch",
		Doc:  "Executes each query. The first query to return any rows will be emitted.",
	}
}

func init() {
	vql_subsystem.RegisterPlugin(&_SwitchPlugin{})
}
