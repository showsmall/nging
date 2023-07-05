// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"github.com/webx-top/com"
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/db/lib/factory/pagination"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/param"
)

type Slice_NgingConfig []*NgingConfig

func (s Slice_NgingConfig) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingConfig) RangeRaw(fn func(m *NgingConfig) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingConfig) GroupBy(keyField string) map[string][]*NgingConfig {
	r := map[string][]*NgingConfig{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*NgingConfig{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (s Slice_NgingConfig) KeyBy(keyField string) map[string]*NgingConfig {
	r := map[string]*NgingConfig{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (s Slice_NgingConfig) AsKV(keyField string, valueField string) param.Store {
	r := param.Store{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (s Slice_NgingConfig) Transform(transfers map[string]param.Transfer) []param.Store {
	r := make([]param.Store, len(s))
	for idx, row := range s {
		r[idx] = row.AsMap().Transform(transfers)
	}
	return r
}

func (s Slice_NgingConfig) FromList(data interface{}) Slice_NgingConfig {
	values, ok := data.([]*NgingConfig)
	if !ok {
		for _, value := range data.([]interface{}) {
			row := &NgingConfig{}
			row.FromRow(value.(map[string]interface{}))
			s = append(s, row)
		}
		return s
	}
	s = append(s, values...)

	return s
}

func NewNgingConfig(ctx echo.Context) *NgingConfig {
	m := &NgingConfig{}
	m.SetContext(ctx)
	return m
}

// NgingConfig 配置
type NgingConfig struct {
	base    factory.Base
	objects []*NgingConfig

	Key         string `db:"key,pk" bson:"key" comment:"键" json:"key" xml:"key"`
	Group       string `db:"group,pk" bson:"group" comment:"组" json:"group" xml:"group"`
	Label       string `db:"label" bson:"label" comment:"选项名称" json:"label" xml:"label"`
	Value       string `db:"value" bson:"value" comment:"值" json:"value" xml:"value"`
	Description string `db:"description" bson:"description" comment:"简介" json:"description" xml:"description"`
	Type        string `db:"type" bson:"type" comment:"值类型(list-以半角逗号分隔的值列表)" json:"type" xml:"type"`
	Sort        int    `db:"sort" bson:"sort" comment:"排序" json:"sort" xml:"sort"`
	Disabled    string `db:"disabled" bson:"disabled" comment:"是否禁用" json:"disabled" xml:"disabled"`
	Encrypted   string `db:"encrypted" bson:"encrypted" comment:"是否加密" json:"encrypted" xml:"encrypted"`
}

// - base function

func (a *NgingConfig) Trans() factory.Transactioner {
	return a.base.Trans()
}

func (a *NgingConfig) Use(trans factory.Transactioner) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *NgingConfig) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *NgingConfig) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *NgingConfig) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *NgingConfig) Context() echo.Context {
	return a.base.Context()
}

func (a *NgingConfig) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *NgingConfig) ConnID() int {
	return a.base.ConnID()
}

func (a *NgingConfig) SetNamer(namer func(factory.Model) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *NgingConfig) Namer() func(factory.Model) string {
	return a.base.Namer()
}

func (a *NgingConfig) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *NgingConfig) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

func (a *NgingConfig) New(structName string, connID ...int) factory.Model {
	return a.base.New(structName, connID...)
}

func (a *NgingConfig) Base_() factory.Baser {
	return &a.base
}

// - current function

func (a *NgingConfig) Objects() []*NgingConfig {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *NgingConfig) XObjects() Slice_NgingConfig {
	return Slice_NgingConfig(a.Objects())
}

func (a *NgingConfig) NewObjects() factory.Ranger {
	return &Slice_NgingConfig{}
}

func (a *NgingConfig) InitObjects() *[]*NgingConfig {
	a.objects = []*NgingConfig{}
	return &a.objects
}

func (a *NgingConfig) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *NgingConfig) Short_() string {
	return "nging_config"
}

func (a *NgingConfig) Struct_() string {
	return "NgingConfig"
}

func (a *NgingConfig) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *NgingConfig) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.SetConnID(source.ConnID())
	a.SetNamer(source.Namer())
	return a
}

func (a *NgingConfig) Get(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	base := a.base
	if !a.base.Eventable() {
		err = a.Param(mw, args...).SetRecv(a).One()
		a.base = base
		return
	}
	queryParam := a.Param(mw, args...).SetRecv(a)
	if err = DBI.FireReading(a, queryParam); err != nil {
		return
	}
	err = queryParam.One()
	a.base = base
	if err == nil {
		err = DBI.FireReaded(a, queryParam)
	}
	return
}

func (a *NgingConfig) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv).List()
	}
	queryParam := a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv)
	if err := DBI.FireReading(a, queryParam); err != nil {
		return nil, err
	}
	cnt, err := queryParam.List()
	if err == nil {
		switch v := recv.(type) {
		case *[]*NgingConfig:
			err = DBI.FireReaded(a, queryParam, Slice_NgingConfig(*v))
		case []*NgingConfig:
			err = DBI.FireReaded(a, queryParam, Slice_NgingConfig(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingConfig) GroupBy(keyField string, inputRows ...[]*NgingConfig) map[string][]*NgingConfig {
	var rows Slice_NgingConfig
	if len(inputRows) > 0 {
		rows = Slice_NgingConfig(inputRows[0])
	} else {
		rows = Slice_NgingConfig(a.Objects())
	}
	return rows.GroupBy(keyField)
}

func (a *NgingConfig) KeyBy(keyField string, inputRows ...[]*NgingConfig) map[string]*NgingConfig {
	var rows Slice_NgingConfig
	if len(inputRows) > 0 {
		rows = Slice_NgingConfig(inputRows[0])
	} else {
		rows = Slice_NgingConfig(a.Objects())
	}
	return rows.KeyBy(keyField)
}

func (a *NgingConfig) AsKV(keyField string, valueField string, inputRows ...[]*NgingConfig) param.Store {
	var rows Slice_NgingConfig
	if len(inputRows) > 0 {
		rows = Slice_NgingConfig(inputRows[0])
	} else {
		rows = Slice_NgingConfig(a.Objects())
	}
	return rows.AsKV(keyField, valueField)
}

func (a *NgingConfig) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv).List()
	}
	queryParam := a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv)
	if err := DBI.FireReading(a, queryParam); err != nil {
		return nil, err
	}
	cnt, err := queryParam.List()
	if err == nil {
		switch v := recv.(type) {
		case *[]*NgingConfig:
			err = DBI.FireReaded(a, queryParam, Slice_NgingConfig(*v))
		case []*NgingConfig:
			err = DBI.FireReaded(a, queryParam, Slice_NgingConfig(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingConfig) Insert() (pk interface{}, err error) {

	if len(a.Type) == 0 {
		a.Type = "text"
	}
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if len(a.Encrypted) == 0 {
		a.Encrypted = "N"
	}
	if a.base.Eventable() {
		err = DBI.Fire("creating", a, nil)
		if err != nil {
			return
		}
	}
	pk, err = a.Param(nil).SetSend(a).Insert()

	if err == nil && a.base.Eventable() {
		err = DBI.Fire("created", a, nil)
	}
	return
}

func (a *NgingConfig) Update(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if len(a.Type) == 0 {
		a.Type = "text"
	}
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if len(a.Encrypted) == 0 {
		a.Encrypted = "N"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(a).Update()
	}
	if err = DBI.Fire("updating", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(a).Update(); err != nil {
		return
	}
	return DBI.Fire("updated", a, mw, args...)
}

func (a *NgingConfig) Updatex(mw func(db.Result) db.Result, args ...interface{}) (affected int64, err error) {

	if len(a.Type) == 0 {
		a.Type = "text"
	}
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if len(a.Encrypted) == 0 {
		a.Encrypted = "N"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(a).Updatex()
	}
	if err = DBI.Fire("updating", a, mw, args...); err != nil {
		return
	}
	if affected, err = a.Param(mw, args...).SetSend(a).Updatex(); err != nil {
		return
	}
	err = DBI.Fire("updated", a, mw, args...)
	return
}

func (a *NgingConfig) UpdateByFields(mw func(db.Result) db.Result, fields []string, args ...interface{}) (err error) {

	if len(a.Type) == 0 {
		a.Type = "text"
	}
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if len(a.Encrypted) == 0 {
		a.Encrypted = "N"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).UpdateByStruct(a, fields...)
	}
	editColumns := make([]string, len(fields))
	for index, field := range fields {
		editColumns[index] = com.SnakeCase(field)
	}
	if err = DBI.FireUpdate("updating", a, editColumns, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).UpdateByStruct(a, fields...); err != nil {
		return
	}
	err = DBI.FireUpdate("updated", a, editColumns, mw, args...)
	return
}

func (a *NgingConfig) UpdatexByFields(mw func(db.Result) db.Result, fields []string, args ...interface{}) (affected int64, err error) {

	if len(a.Type) == 0 {
		a.Type = "text"
	}
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if len(a.Encrypted) == 0 {
		a.Encrypted = "N"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).UpdatexByStruct(a, fields...)
	}
	editColumns := make([]string, len(fields))
	for index, field := range fields {
		editColumns[index] = com.SnakeCase(field)
	}
	if err = DBI.FireUpdate("updating", a, editColumns, mw, args...); err != nil {
		return
	}
	if affected, err = a.Param(mw, args...).UpdatexByStruct(a, fields...); err != nil {
		return
	}
	err = DBI.FireUpdate("updated", a, editColumns, mw, args...)
	return
}

func (a *NgingConfig) UpdateField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.UpdateFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingConfig) UpdatexField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (affected int64, err error) {
	return a.UpdatexFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingConfig) UpdateFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

	if val, ok := kvset["type"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["type"] = "text"
		}
	}
	if val, ok := kvset["disabled"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["disabled"] = "N"
		}
	}
	if val, ok := kvset["encrypted"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["encrypted"] = "N"
		}
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(kvset).Update()
	}
	m := *a
	m.FromRow(kvset)
	var editColumns []string
	for column := range kvset {
		editColumns = append(editColumns, column)
	}
	if err = DBI.FireUpdate("updating", &m, editColumns, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(kvset).Update(); err != nil {
		return
	}
	return DBI.FireUpdate("updated", &m, editColumns, mw, args...)
}

func (a *NgingConfig) UpdatexFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (affected int64, err error) {

	if val, ok := kvset["type"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["type"] = "text"
		}
	}
	if val, ok := kvset["disabled"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["disabled"] = "N"
		}
	}
	if val, ok := kvset["encrypted"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["encrypted"] = "N"
		}
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(kvset).Updatex()
	}
	m := *a
	m.FromRow(kvset)
	var editColumns []string
	for column := range kvset {
		editColumns = append(editColumns, column)
	}
	if err = DBI.FireUpdate("updating", &m, editColumns, mw, args...); err != nil {
		return
	}
	if affected, err = a.Param(mw, args...).SetSend(kvset).Updatex(); err != nil {
		return
	}
	err = DBI.FireUpdate("updated", &m, editColumns, mw, args...)
	return
}

func (a *NgingConfig) UpdateValues(mw func(db.Result) db.Result, keysValues *db.KeysValues, args ...interface{}) (err error) {
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(keysValues).Update()
	}
	m := *a
	m.FromRow(keysValues.Map())
	if err = DBI.FireUpdate("updating", &m, keysValues.Keys(), mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(keysValues).Update(); err != nil {
		return
	}
	return DBI.FireUpdate("updated", &m, keysValues.Keys(), mw, args...)
}

func (a *NgingConfig) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		if len(a.Type) == 0 {
			a.Type = "text"
		}
		if len(a.Disabled) == 0 {
			a.Disabled = "N"
		}
		if len(a.Encrypted) == 0 {
			a.Encrypted = "N"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		if len(a.Type) == 0 {
			a.Type = "text"
		}
		if len(a.Disabled) == 0 {
			a.Disabled = "N"
		}
		if len(a.Encrypted) == 0 {
			a.Encrypted = "N"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("creating", a, nil)
	})

	if err == nil && a.base.Eventable() {
		if pk == nil {
			err = DBI.Fire("updated", a, mw, args...)
		} else {
			err = DBI.Fire("created", a, nil)
		}
	}
	return
}

func (a *NgingConfig) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if !a.base.Eventable() {
		return a.Param(mw, args...).Delete()
	}
	if err = DBI.Fire("deleting", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).Delete(); err != nil {
		return
	}
	return DBI.Fire("deleted", a, mw, args...)
}

func (a *NgingConfig) Deletex(mw func(db.Result) db.Result, args ...interface{}) (affected int64, err error) {

	if !a.base.Eventable() {
		return a.Param(mw, args...).Deletex()
	}
	if err = DBI.Fire("deleting", a, mw, args...); err != nil {
		return
	}
	if affected, err = a.Param(mw, args...).Deletex(); err != nil {
		return
	}
	err = DBI.Fire("deleted", a, mw, args...)
	return
}

func (a *NgingConfig) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *NgingConfig) Exists(mw func(db.Result) db.Result, args ...interface{}) (bool, error) {
	return a.Param(mw, args...).Exists()
}

func (a *NgingConfig) Reset() *NgingConfig {
	a.Key = ``
	a.Group = ``
	a.Label = ``
	a.Value = ``
	a.Description = ``
	a.Type = ``
	a.Sort = 0
	a.Disabled = ``
	a.Encrypted = ``
	return a
}

func (a *NgingConfig) AsMap(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["Key"] = a.Key
		r["Group"] = a.Group
		r["Label"] = a.Label
		r["Value"] = a.Value
		r["Description"] = a.Description
		r["Type"] = a.Type
		r["Sort"] = a.Sort
		r["Disabled"] = a.Disabled
		r["Encrypted"] = a.Encrypted
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "Key":
			r["Key"] = a.Key
		case "Group":
			r["Group"] = a.Group
		case "Label":
			r["Label"] = a.Label
		case "Value":
			r["Value"] = a.Value
		case "Description":
			r["Description"] = a.Description
		case "Type":
			r["Type"] = a.Type
		case "Sort":
			r["Sort"] = a.Sort
		case "Disabled":
			r["Disabled"] = a.Disabled
		case "Encrypted":
			r["Encrypted"] = a.Encrypted
		}
	}
	return r
}

func (a *NgingConfig) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "key":
			a.Key = param.AsString(value)
		case "group":
			a.Group = param.AsString(value)
		case "label":
			a.Label = param.AsString(value)
		case "value":
			a.Value = param.AsString(value)
		case "description":
			a.Description = param.AsString(value)
		case "type":
			a.Type = param.AsString(value)
		case "sort":
			a.Sort = param.AsInt(value)
		case "disabled":
			a.Disabled = param.AsString(value)
		case "encrypted":
			a.Encrypted = param.AsString(value)
		}
	}
}

func (a *NgingConfig) Set(key interface{}, value ...interface{}) {
	switch k := key.(type) {
	case map[string]interface{}:
		for kk, vv := range k {
			a.Set(kk, vv)
		}
	default:
		var (
			kk string
			vv interface{}
		)
		if k, y := key.(string); y {
			kk = k
		} else {
			kk = fmt.Sprint(key)
		}
		if len(value) > 0 {
			vv = value[0]
		}
		switch kk {
		case "Key":
			a.Key = param.AsString(vv)
		case "Group":
			a.Group = param.AsString(vv)
		case "Label":
			a.Label = param.AsString(vv)
		case "Value":
			a.Value = param.AsString(vv)
		case "Description":
			a.Description = param.AsString(vv)
		case "Type":
			a.Type = param.AsString(vv)
		case "Sort":
			a.Sort = param.AsInt(vv)
		case "Disabled":
			a.Disabled = param.AsString(vv)
		case "Encrypted":
			a.Encrypted = param.AsString(vv)
		}
	}
}

func (a *NgingConfig) AsRow(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["key"] = a.Key
		r["group"] = a.Group
		r["label"] = a.Label
		r["value"] = a.Value
		r["description"] = a.Description
		r["type"] = a.Type
		r["sort"] = a.Sort
		r["disabled"] = a.Disabled
		r["encrypted"] = a.Encrypted
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "key":
			r["key"] = a.Key
		case "group":
			r["group"] = a.Group
		case "label":
			r["label"] = a.Label
		case "value":
			r["value"] = a.Value
		case "description":
			r["description"] = a.Description
		case "type":
			r["type"] = a.Type
		case "sort":
			r["sort"] = a.Sort
		case "disabled":
			r["disabled"] = a.Disabled
		case "encrypted":
			r["encrypted"] = a.Encrypted
		}
	}
	return r
}

func (a *NgingConfig) ListPage(cond *db.Compounds, sorts ...interface{}) error {
	_, err := pagination.NewLister(a, nil, func(r db.Result) db.Result {
		return r.OrderBy(sorts...)
	}, cond.And()).Paging(a.Context())
	return err
}

func (a *NgingConfig) ListPageAs(recv interface{}, cond *db.Compounds, sorts ...interface{}) error {
	_, err := pagination.NewLister(a, recv, func(r db.Result) db.Result {
		return r.OrderBy(sorts...)
	}, cond.And()).Paging(a.Context())
	return err
}

func (a *NgingConfig) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return DBI.Fields.BatchValidate(a.Short_(), kvset)
}

func (a *NgingConfig) Validate(field string, value interface{}) error {
	return DBI.Fields.Validate(a.Short_(), field, value)
}
