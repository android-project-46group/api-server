// Code generated by SQLBoiler 4.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Formation is an object representing the database table.
type Formation struct {
	FormationID  int        `boil:"formation_id" json:"formation_id" toml:"formation_id" yaml:"formation_id"`
	FirstRowNum  null.Int16 `boil:"first_row_num" json:"first_row_num,omitempty" toml:"first_row_num" yaml:"first_row_num,omitempty"`
	SecondRowNum null.Int16 `boil:"second_row_num" json:"second_row_num,omitempty" toml:"second_row_num" yaml:"second_row_num,omitempty"`
	ThirdRowNum  null.Int16 `boil:"third_row_num" json:"third_row_num,omitempty" toml:"third_row_num" yaml:"third_row_num,omitempty"`

	R *formationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L formationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var FormationColumns = struct {
	FormationID  string
	FirstRowNum  string
	SecondRowNum string
	ThirdRowNum  string
}{
	FormationID:  "formation_id",
	FirstRowNum:  "first_row_num",
	SecondRowNum: "second_row_num",
	ThirdRowNum:  "third_row_num",
}

var FormationTableColumns = struct {
	FormationID  string
	FirstRowNum  string
	SecondRowNum string
	ThirdRowNum  string
}{
	FormationID:  "formations.formation_id",
	FirstRowNum:  "formations.first_row_num",
	SecondRowNum: "formations.second_row_num",
	ThirdRowNum:  "formations.third_row_num",
}

// Generated where

type whereHelpernull_Int16 struct{ field string }

func (w whereHelpernull_Int16) EQ(x null.Int16) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int16) NEQ(x null.Int16) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int16) LT(x null.Int16) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int16) LTE(x null.Int16) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int16) GT(x null.Int16) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int16) GTE(x null.Int16) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Int16) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int16) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var FormationWhere = struct {
	FormationID  whereHelperint
	FirstRowNum  whereHelpernull_Int16
	SecondRowNum whereHelpernull_Int16
	ThirdRowNum  whereHelpernull_Int16
}{
	FormationID:  whereHelperint{field: "\"formations\".\"formation_id\""},
	FirstRowNum:  whereHelpernull_Int16{field: "\"formations\".\"first_row_num\""},
	SecondRowNum: whereHelpernull_Int16{field: "\"formations\".\"second_row_num\""},
	ThirdRowNum:  whereHelpernull_Int16{field: "\"formations\".\"third_row_num\""},
}

// FormationRels is where relationship names are stored.
var FormationRels = struct {
	Songs string
}{
	Songs: "Songs",
}

// formationR is where relationships are stored.
type formationR struct {
	Songs SongSlice `boil:"Songs" json:"Songs" toml:"Songs" yaml:"Songs"`
}

// NewStruct creates a new relationship struct
func (*formationR) NewStruct() *formationR {
	return &formationR{}
}

// formationL is where Load methods for each relationship are stored.
type formationL struct{}

var (
	formationAllColumns            = []string{"formation_id", "first_row_num", "second_row_num", "third_row_num"}
	formationColumnsWithoutDefault = []string{"first_row_num", "second_row_num", "third_row_num"}
	formationColumnsWithDefault    = []string{"formation_id"}
	formationPrimaryKeyColumns     = []string{"formation_id"}
)

type (
	// FormationSlice is an alias for a slice of pointers to Formation.
	// This should almost always be used instead of []Formation.
	FormationSlice []*Formation
	// FormationHook is the signature for custom Formation hook methods
	FormationHook func(context.Context, boil.ContextExecutor, *Formation) error

	formationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	formationType                 = reflect.TypeOf(&Formation{})
	formationMapping              = queries.MakeStructMapping(formationType)
	formationPrimaryKeyMapping, _ = queries.BindMapping(formationType, formationMapping, formationPrimaryKeyColumns)
	formationInsertCacheMut       sync.RWMutex
	formationInsertCache          = make(map[string]insertCache)
	formationUpdateCacheMut       sync.RWMutex
	formationUpdateCache          = make(map[string]updateCache)
	formationUpsertCacheMut       sync.RWMutex
	formationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var formationBeforeInsertHooks []FormationHook
var formationBeforeUpdateHooks []FormationHook
var formationBeforeDeleteHooks []FormationHook
var formationBeforeUpsertHooks []FormationHook

var formationAfterInsertHooks []FormationHook
var formationAfterSelectHooks []FormationHook
var formationAfterUpdateHooks []FormationHook
var formationAfterDeleteHooks []FormationHook
var formationAfterUpsertHooks []FormationHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Formation) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range formationBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Formation) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range formationBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Formation) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range formationBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Formation) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range formationBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Formation) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range formationAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Formation) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range formationAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Formation) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range formationAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Formation) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range formationAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Formation) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range formationAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFormationHook registers your hook function for all future operations.
func AddFormationHook(hookPoint boil.HookPoint, formationHook FormationHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		formationBeforeInsertHooks = append(formationBeforeInsertHooks, formationHook)
	case boil.BeforeUpdateHook:
		formationBeforeUpdateHooks = append(formationBeforeUpdateHooks, formationHook)
	case boil.BeforeDeleteHook:
		formationBeforeDeleteHooks = append(formationBeforeDeleteHooks, formationHook)
	case boil.BeforeUpsertHook:
		formationBeforeUpsertHooks = append(formationBeforeUpsertHooks, formationHook)
	case boil.AfterInsertHook:
		formationAfterInsertHooks = append(formationAfterInsertHooks, formationHook)
	case boil.AfterSelectHook:
		formationAfterSelectHooks = append(formationAfterSelectHooks, formationHook)
	case boil.AfterUpdateHook:
		formationAfterUpdateHooks = append(formationAfterUpdateHooks, formationHook)
	case boil.AfterDeleteHook:
		formationAfterDeleteHooks = append(formationAfterDeleteHooks, formationHook)
	case boil.AfterUpsertHook:
		formationAfterUpsertHooks = append(formationAfterUpsertHooks, formationHook)
	}
}

// One returns a single formation record from the query.
func (q formationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Formation, error) {
	o := &Formation{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for formations")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Formation records from the query.
func (q formationQuery) All(ctx context.Context, exec boil.ContextExecutor) (FormationSlice, error) {
	var o []*Formation

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Formation slice")
	}

	if len(formationAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Formation records in the query.
func (q formationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count formations rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q formationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if formations exists")
	}

	return count > 0, nil
}

// Songs retrieves all the song's Songs with an executor.
func (o *Formation) Songs(mods ...qm.QueryMod) songQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"songs\".\"formation_id\"=?", o.FormationID),
	)

	query := Songs(queryMods...)
	queries.SetFrom(query.Query, "\"songs\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"songs\".*"})
	}

	return query
}

// LoadSongs allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (formationL) LoadSongs(ctx context.Context, e boil.ContextExecutor, singular bool, maybeFormation interface{}, mods queries.Applicator) error {
	var slice []*Formation
	var object *Formation

	if singular {
		object = maybeFormation.(*Formation)
	} else {
		slice = *maybeFormation.(*[]*Formation)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &formationR{}
		}
		args = append(args, object.FormationID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &formationR{}
			}

			for _, a := range args {
				if a == obj.FormationID {
					continue Outer
				}
			}

			args = append(args, obj.FormationID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`songs`),
		qm.WhereIn(`songs.formation_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load songs")
	}

	var resultSlice []*Song
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice songs")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on songs")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for songs")
	}

	if len(songAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Songs = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &songR{}
			}
			foreign.R.Formation = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FormationID == foreign.FormationID {
				local.R.Songs = append(local.R.Songs, foreign)
				if foreign.R == nil {
					foreign.R = &songR{}
				}
				foreign.R.Formation = local
				break
			}
		}
	}

	return nil
}

// AddSongs adds the given related objects to the existing relationships
// of the formation, optionally inserting them as new records.
// Appends related to o.R.Songs.
// Sets related.R.Formation appropriately.
func (o *Formation) AddSongs(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Song) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.FormationID = o.FormationID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"songs\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"formation_id"}),
				strmangle.WhereClause("\"", "\"", 2, songPrimaryKeyColumns),
			)
			values := []interface{}{o.FormationID, rel.SongID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.FormationID = o.FormationID
		}
	}

	if o.R == nil {
		o.R = &formationR{
			Songs: related,
		}
	} else {
		o.R.Songs = append(o.R.Songs, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &songR{
				Formation: o,
			}
		} else {
			rel.R.Formation = o
		}
	}
	return nil
}

// Formations retrieves all the records using an executor.
func Formations(mods ...qm.QueryMod) formationQuery {
	mods = append(mods, qm.From("\"formations\""))
	return formationQuery{NewQuery(mods...)}
}

// FindFormation retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFormation(ctx context.Context, exec boil.ContextExecutor, formationID int, selectCols ...string) (*Formation, error) {
	formationObj := &Formation{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"formations\" where \"formation_id\"=$1", sel,
	)

	q := queries.Raw(query, formationID)

	err := q.Bind(ctx, exec, formationObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from formations")
	}

	if err = formationObj.doAfterSelectHooks(ctx, exec); err != nil {
		return formationObj, err
	}

	return formationObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Formation) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no formations provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(formationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	formationInsertCacheMut.RLock()
	cache, cached := formationInsertCache[key]
	formationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			formationAllColumns,
			formationColumnsWithDefault,
			formationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(formationType, formationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(formationType, formationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"formations\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"formations\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into formations")
	}

	if !cached {
		formationInsertCacheMut.Lock()
		formationInsertCache[key] = cache
		formationInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Formation.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Formation) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	formationUpdateCacheMut.RLock()
	cache, cached := formationUpdateCache[key]
	formationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			formationAllColumns,
			formationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update formations, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"formations\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, formationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(formationType, formationMapping, append(wl, formationPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update formations row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for formations")
	}

	if !cached {
		formationUpdateCacheMut.Lock()
		formationUpdateCache[key] = cache
		formationUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q formationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for formations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for formations")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FormationSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), formationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"formations\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, formationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in formation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all formation")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Formation) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no formations provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(formationColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	formationUpsertCacheMut.RLock()
	cache, cached := formationUpsertCache[key]
	formationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			formationAllColumns,
			formationColumnsWithDefault,
			formationColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			formationAllColumns,
			formationPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert formations, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(formationPrimaryKeyColumns))
			copy(conflict, formationPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"formations\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(formationType, formationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(formationType, formationMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert formations")
	}

	if !cached {
		formationUpsertCacheMut.Lock()
		formationUpsertCache[key] = cache
		formationUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Formation record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Formation) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Formation provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), formationPrimaryKeyMapping)
	sql := "DELETE FROM \"formations\" WHERE \"formation_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from formations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for formations")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q formationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no formationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from formations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for formations")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FormationSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(formationBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), formationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"formations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, formationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from formation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for formations")
	}

	if len(formationAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Formation) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindFormation(ctx, exec, o.FormationID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FormationSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := FormationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), formationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"formations\".* FROM \"formations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, formationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FormationSlice")
	}

	*o = slice

	return nil
}

// FormationExists checks if the Formation row exists.
func FormationExists(ctx context.Context, exec boil.ContextExecutor, formationID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"formations\" where \"formation_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, formationID)
	}
	row := exec.QueryRowContext(ctx, sql, formationID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if formations exists")
	}

	return exists, nil
}
