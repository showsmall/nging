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

	DBI.FieldsRegister(map[string]map[string]*factory.FieldInfo{"nging_ftp_permission": {"created": {Name: "created", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "创建时间", GoType: "uint", MyType: "", GoName: "Created"}, "id": {Name: "id", DataType: "int", Unsigned: true, PrimaryKey: true, AutoIncrement: true, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "", Comment: "ID", GoType: "uint", MyType: "", GoName: "Id"}, "permission": {Name: "permission", DataType: "text", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "", Comment: "权限JSON", GoType: "string", MyType: "", GoName: "Permission"}, "target_id": {Name: "target_id", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "目标ID", GoType: "uint", MyType: "", GoName: "TargetId"}, "target_type": {Name: "target_type", DataType: "enum", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{"user", "group"}, DefaultValue: "group", Comment: "目标类型", GoType: "string", MyType: "", GoName: "TargetType"}, "updated": {Name: "updated", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "更新时间", GoType: "uint", MyType: "", GoName: "Updated"}}, "nging_ftp_user": {"banned": {Name: "banned", DataType: "enum", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{"Y", "N"}, DefaultValue: "N", Comment: "是否禁止连接", GoType: "string", MyType: "", GoName: "Banned"}, "created": {Name: "created", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "", Comment: "创建时间 ", GoType: "uint", MyType: "", GoName: "Created"}, "directory": {Name: "directory", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 500, Options: []string{}, DefaultValue: "", Comment: "授权目录(一行一个) ", GoType: "string", MyType: "", GoName: "Directory"}, "group_id": {Name: "group_id", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "用户组", GoType: "uint", MyType: "", GoName: "GroupId"}, "id": {Name: "id", DataType: "int", Unsigned: true, PrimaryKey: true, AutoIncrement: true, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "", Comment: "", GoType: "uint", MyType: "", GoName: "Id"}, "ip_blacklist": {Name: "ip_blacklist", DataType: "text", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "", Comment: "IP黑名单(一行一个) ", GoType: "string", MyType: "", GoName: "IpBlacklist"}, "ip_whitelist": {Name: "ip_whitelist", DataType: "text", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "", Comment: "IP白名单(一行一个) ", GoType: "string", MyType: "", GoName: "IpWhitelist"}, "modify": {Name: "modify", DataType: "enum", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{"Y", "N"}, DefaultValue: "N", Comment: "是否默认写权限", GoType: "string", MyType: "", GoName: "Modify"}, "password": {Name: "password", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 150, Options: []string{}, DefaultValue: "", Comment: "密码", GoType: "string", MyType: "", GoName: "Password"}, "updated": {Name: "updated", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "修改时间", GoType: "uint", MyType: "", GoName: "Updated"}, "username": {Name: "username", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 120, Options: []string{}, DefaultValue: "", Comment: "用户名", GoType: "string", MyType: "", GoName: "Username"}}, "nging_ftp_user_group": {"banned": {Name: "banned", DataType: "enum", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{"Y", "N"}, DefaultValue: "N", Comment: "是否禁止组内用户连接", GoType: "string", MyType: "", GoName: "Banned"}, "created": {Name: "created", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "", Comment: "创建时间", GoType: "uint", MyType: "", GoName: "Created"}, "directory": {Name: "directory", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 500, Options: []string{}, DefaultValue: "", Comment: "授权目录", GoType: "string", MyType: "", GoName: "Directory"}, "disabled": {Name: "disabled", DataType: "enum", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{"Y", "N"}, DefaultValue: "N", Comment: "是否禁用", GoType: "string", MyType: "", GoName: "Disabled"}, "id": {Name: "id", DataType: "int", Unsigned: true, PrimaryKey: true, AutoIncrement: true, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "", Comment: "", GoType: "uint", MyType: "", GoName: "Id"}, "ip_blacklist": {Name: "ip_blacklist", DataType: "text", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "", Comment: "IP黑名单(一行一个)", GoType: "string", MyType: "", GoName: "IpBlacklist"}, "ip_whitelist": {Name: "ip_whitelist", DataType: "text", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "", Comment: "IP白名单(一行一个)", GoType: "string", MyType: "", GoName: "IpWhitelist"}, "modify": {Name: "modify", DataType: "enum", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{"Y", "N"}, DefaultValue: "N", Comment: "是否默认写权限", GoType: "string", MyType: "", GoName: "Modify"}, "name": {Name: "name", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 255, Options: []string{}, DefaultValue: "", Comment: "组名称", GoType: "string", MyType: "", GoName: "Name"}, "updated": {Name: "updated", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "修改时间", GoType: "uint", MyType: "", GoName: "Updated"}}})

	DBI.ColumnsRegister(map[string][]string{"nging_ftp_permission": {"id", "target_type", "target_id", "permission", "created", "updated"}, "nging_ftp_user": {"id", "username", "password", "banned", "directory", "modify", "ip_whitelist", "ip_blacklist", "created", "updated", "group_id"}, "nging_ftp_user_group": {"id", "name", "created", "updated", "disabled", "banned", "directory", "modify", "ip_whitelist", "ip_blacklist"}})

	DBI.ModelsRegister(factory.ModelInstancers{`NgingFtpPermission`: factory.NewMI("nging_ftp_permission", func(connID int) factory.Model { return &NgingFtpPermission{base: *factory.NewBase(connID)} }, "用户权限"), `NgingFtpUser`: factory.NewMI("nging_ftp_user", func(connID int) factory.Model { return &NgingFtpUser{base: *factory.NewBase(connID)} }, "FTP用户"), `NgingFtpUserGroup`: factory.NewMI("nging_ftp_user_group", func(connID int) factory.Model { return &NgingFtpUserGroup{base: *factory.NewBase(connID)} }, "FTP用户组")})

}
