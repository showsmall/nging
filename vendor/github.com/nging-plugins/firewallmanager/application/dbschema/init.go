// @generated Do not edit a file, which is automatically generated by the generator.

package dbschema

import (
	"github.com/webx-top/db/lib/factory"
)

var WithPrefix = func(tableName string) string {
	return "" + tableName
}

var DBI = factory.DefaultDBI

func init() {

	DBI.FieldsRegister(map[string]map[string]*factory.FieldInfo{"nging_firewall_rule_dynamic": {"action_arg": {Name: "action_arg", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 20, Options: []string{}, DefaultValue: "", Comment: "操作参数", GoType: "string", MyType: "", GoName: "ActionArg"}, "action_type": {Name: "action_type", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 10, Options: []string{}, DefaultValue: "ban", Comment: "操作类型", GoType: "string", MyType: "", GoName: "ActionType"}, "aggregate_duration": {Name: "aggregate_duration", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 10, Options: []string{}, DefaultValue: "", Comment: "合计时长", GoType: "string", MyType: "", GoName: "AggregateDuration"}, "aggregate_regexp": {Name: "aggregate_regexp", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 1000, Options: []string{}, DefaultValue: "", Comment: "合计规格(JSON数组)", GoType: "string", MyType: "", GoName: "AggregateRegexp"}, "created": {Name: "created", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "创建时间", GoType: "uint", MyType: "", GoName: "Created"}, "disabled": {Name: "disabled", DataType: "enum", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{"Y", "N"}, DefaultValue: "N", Comment: "是否(Y/N)禁用", GoType: "string", MyType: "", GoName: "Disabled"}, "id": {Name: "id", DataType: "int", Unsigned: true, PrimaryKey: true, AutoIncrement: true, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "", Comment: "ID", GoType: "uint", MyType: "", GoName: "Id"}, "name": {Name: "name", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 255, Options: []string{}, DefaultValue: "", Comment: "规则名称", GoType: "string", MyType: "", GoName: "Name"}, "occurrence_duration": {Name: "occurrence_duration", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 10, Options: []string{}, DefaultValue: "", Comment: "持续时长(如1h)", GoType: "string", MyType: "", GoName: "OccurrenceDuration"}, "occurrence_num": {Name: "occurrence_num", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "在持续一段时间内的出现次数", GoType: "uint", MyType: "", GoName: "OccurrenceNum"}, "regexp": {Name: "regexp", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 1000, Options: []string{}, DefaultValue: "", Comment: "正则规格(JSON数组)", GoType: "string", MyType: "", GoName: "Regexp"}, "source_args": {Name: "source_args", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 1000, Options: []string{}, DefaultValue: "", Comment: "资源参数(JSON数组)", GoType: "string", MyType: "", GoName: "SourceArgs"}, "source_type": {Name: "source_type", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 10, Options: []string{}, DefaultValue: "file", Comment: "资源类型", GoType: "string", MyType: "", GoName: "SourceType"}, "updated": {Name: "updated", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "更新时间", GoType: "uint", MyType: "", GoName: "Updated"}}, "nging_firewall_rule_static": {"action": {Name: "action", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 10, Options: []string{}, DefaultValue: "ACCEPT", Comment: "操作", GoType: "string", MyType: "", GoName: "Action"}, "conn_limit": {Name: "conn_limit", DataType: "bigint", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "连接数限制", GoType: "uint64", MyType: "", GoName: "ConnLimit"}, "created": {Name: "created", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "创建时间", GoType: "uint", MyType: "", GoName: "Created"}, "direction": {Name: "direction", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 15, Options: []string{}, DefaultValue: "INPUT", Comment: "方向", GoType: "string", MyType: "", GoName: "Direction"}, "disabled": {Name: "disabled", DataType: "enum", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{"Y", "N"}, DefaultValue: "N", Comment: "是否(Y/N)禁用", GoType: "string", MyType: "", GoName: "Disabled"}, "extra": {Name: "extra", DataType: "text", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "", Comment: "其它扩展设置", GoType: "string", MyType: "", GoName: "Extra"}, "id": {Name: "id", DataType: "int", Unsigned: true, PrimaryKey: true, AutoIncrement: true, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "", Comment: "ID", GoType: "uint", MyType: "", GoName: "Id"}, "interface": {Name: "interface", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 255, Options: []string{}, DefaultValue: "", Comment: "入站网口", GoType: "string", MyType: "", GoName: "Interface"}, "ip_version": {Name: "ip_version", DataType: "enum", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{"4", "6", "all"}, DefaultValue: "4", Comment: "IP版本", GoType: "string", MyType: "", GoName: "IpVersion"}, "local_ip": {Name: "local_ip", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 255, Options: []string{}, DefaultValue: "", Comment: "本地IP", GoType: "string", MyType: "", GoName: "LocalIp"}, "local_port": {Name: "local_port", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 255, Options: []string{}, DefaultValue: "", Comment: "本地端口", GoType: "string", MyType: "", GoName: "LocalPort"}, "name": {Name: "name", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 255, Options: []string{}, DefaultValue: "", Comment: "规则名称", GoType: "string", MyType: "", GoName: "Name"}, "nat_ip": {Name: "nat_ip", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 255, Options: []string{}, DefaultValue: "", Comment: "NAT IP", GoType: "string", MyType: "", GoName: "NatIp"}, "nat_port": {Name: "nat_port", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 255, Options: []string{}, DefaultValue: "", Comment: "NAT 端口", GoType: "string", MyType: "", GoName: "NatPort"}, "outerface": {Name: "outerface", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 255, Options: []string{}, DefaultValue: "", Comment: "出站往口", GoType: "string", MyType: "", GoName: "Outerface"}, "position": {Name: "position", DataType: "int", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: -0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "位置", GoType: "int", MyType: "", GoName: "Position"}, "protocol": {Name: "protocol", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 10, Options: []string{}, DefaultValue: "all", Comment: "协议", GoType: "string", MyType: "", GoName: "Protocol"}, "rate_burst": {Name: "rate_burst", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "频率允许峰值", GoType: "uint", MyType: "", GoName: "RateBurst"}, "rate_expires": {Name: "rate_expires", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "过期时间(秒)", GoType: "uint", MyType: "", GoName: "RateExpires"}, "rate_limit": {Name: "rate_limit", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 100, Options: []string{}, DefaultValue: "", Comment: "频率限制", GoType: "string", MyType: "", GoName: "RateLimit"}, "remote_ip": {Name: "remote_ip", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 255, Options: []string{}, DefaultValue: "", Comment: "远程IP", GoType: "string", MyType: "", GoName: "RemoteIp"}, "remote_port": {Name: "remote_port", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 255, Options: []string{}, DefaultValue: "", Comment: "远程端口", GoType: "string", MyType: "", GoName: "RemotePort"}, "state": {Name: "state", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 100, Options: []string{}, DefaultValue: "", Comment: "状态(多个用逗号\",\"分隔)", GoType: "string", MyType: "", GoName: "State"}, "type": {Name: "type", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 10, Options: []string{}, DefaultValue: "filter", Comment: "类型(filter/nat/mangle/raw)", GoType: "string", MyType: "", GoName: "Type"}, "updated": {Name: "updated", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "更新时间", GoType: "uint", MyType: "", GoName: "Updated"}}})

	DBI.ColumnsRegister(map[string][]string{"nging_firewall_rule_dynamic": {"id", "name", "source_type", "source_args", "regexp", "action_type", "action_arg", "aggregate_duration", "aggregate_regexp", "occurrence_num", "occurrence_duration", "disabled", "created", "updated"}, "nging_firewall_rule_static": {"id", "type", "position", "name", "direction", "protocol", "remote_ip", "remote_port", "local_ip", "local_port", "nat_ip", "nat_port", "interface", "outerface", "state", "conn_limit", "rate_limit", "rate_burst", "rate_expires", "extra", "action", "ip_version", "disabled", "created", "updated"}})

	DBI.ModelsRegister(factory.ModelInstancers{`NgingFirewallRuleDynamic`: factory.NewMI("nging_firewall_rule_dynamic", func(connID int) factory.Model {
		return &NgingFirewallRuleDynamic{base: *((&factory.Base{}).SetConnID(connID))}
	}, "防火墙动态规则"), `NgingFirewallRuleStatic`: factory.NewMI("nging_firewall_rule_static", func(connID int) factory.Model {
		return &NgingFirewallRuleStatic{base: *((&factory.Base{}).SetConnID(connID))}
	}, "防火墙静态规则")})

}
