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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Blog is an object representing the database table.
type Blog struct {
	BlogID        int    `boil:"blog_id" json:"blog_id" toml:"blog_id" yaml:"blog_id"`
	MemberID      int    `boil:"member_id" json:"member_id" toml:"member_id" yaml:"member_id"`
	BlogURL       string `boil:"blog_url" json:"blog_url" toml:"blog_url" yaml:"blog_url"`
	LastBlogImg   string `boil:"last_blog_img" json:"last_blog_img" toml:"last_blog_img" yaml:"last_blog_img"`
	LastUpdatedAt string `boil:"last_updated_at" json:"last_updated_at" toml:"last_updated_at" yaml:"last_updated_at"`

	R *blogR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L blogL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BlogColumns = struct {
	BlogID        string
	MemberID      string
	BlogURL       string
	LastBlogImg   string
	LastUpdatedAt string
}{
	BlogID:        "blog_id",
	MemberID:      "member_id",
	BlogURL:       "blog_url",
	LastBlogImg:   "last_blog_img",
	LastUpdatedAt: "last_updated_at",
}

var BlogTableColumns = struct {
	BlogID        string
	MemberID      string
	BlogURL       string
	LastBlogImg   string
	LastUpdatedAt string
}{
	BlogID:        "blogs.blog_id",
	MemberID:      "blogs.member_id",
	BlogURL:       "blogs.blog_url",
	LastBlogImg:   "blogs.last_blog_img",
	LastUpdatedAt: "blogs.last_updated_at",
}

// Generated where

var BlogWhere = struct {
	BlogID        whereHelperint
	MemberID      whereHelperint
	BlogURL       whereHelperstring
	LastBlogImg   whereHelperstring
	LastUpdatedAt whereHelperstring
}{
	BlogID:        whereHelperint{field: "\"blogs\".\"blog_id\""},
	MemberID:      whereHelperint{field: "\"blogs\".\"member_id\""},
	BlogURL:       whereHelperstring{field: "\"blogs\".\"blog_url\""},
	LastBlogImg:   whereHelperstring{field: "\"blogs\".\"last_blog_img\""},
	LastUpdatedAt: whereHelperstring{field: "\"blogs\".\"last_updated_at\""},
}

// BlogRels is where relationship names are stored.
var BlogRels = struct {
	Member string
}{
	Member: "Member",
}

// blogR is where relationships are stored.
type blogR struct {
	Member *Member `boil:"Member" json:"Member" toml:"Member" yaml:"Member"`
}

// NewStruct creates a new relationship struct
func (*blogR) NewStruct() *blogR {
	return &blogR{}
}

// blogL is where Load methods for each relationship are stored.
type blogL struct{}

var (
	blogAllColumns            = []string{"blog_id", "member_id", "blog_url", "last_blog_img", "last_updated_at"}
	blogColumnsWithoutDefault = []string{"member_id", "blog_url", "last_blog_img", "last_updated_at"}
	blogColumnsWithDefault    = []string{"blog_id"}
	blogPrimaryKeyColumns     = []string{"blog_id"}
)

type (
	// BlogSlice is an alias for a slice of pointers to Blog.
	// This should almost always be used instead of []Blog.
	BlogSlice []*Blog
	// BlogHook is the signature for custom Blog hook methods
	BlogHook func(context.Context, boil.ContextExecutor, *Blog) error

	blogQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	blogType                 = reflect.TypeOf(&Blog{})
	blogMapping              = queries.MakeStructMapping(blogType)
	blogPrimaryKeyMapping, _ = queries.BindMapping(blogType, blogMapping, blogPrimaryKeyColumns)
	blogInsertCacheMut       sync.RWMutex
	blogInsertCache          = make(map[string]insertCache)
	blogUpdateCacheMut       sync.RWMutex
	blogUpdateCache          = make(map[string]updateCache)
	blogUpsertCacheMut       sync.RWMutex
	blogUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var blogBeforeInsertHooks []BlogHook
var blogBeforeUpdateHooks []BlogHook
var blogBeforeDeleteHooks []BlogHook
var blogBeforeUpsertHooks []BlogHook

var blogAfterInsertHooks []BlogHook
var blogAfterSelectHooks []BlogHook
var blogAfterUpdateHooks []BlogHook
var blogAfterDeleteHooks []BlogHook
var blogAfterUpsertHooks []BlogHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Blog) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range blogBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Blog) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range blogBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Blog) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range blogBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Blog) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range blogBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Blog) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range blogAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Blog) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range blogAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Blog) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range blogAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Blog) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range blogAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Blog) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range blogAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBlogHook registers your hook function for all future operations.
func AddBlogHook(hookPoint boil.HookPoint, blogHook BlogHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		blogBeforeInsertHooks = append(blogBeforeInsertHooks, blogHook)
	case boil.BeforeUpdateHook:
		blogBeforeUpdateHooks = append(blogBeforeUpdateHooks, blogHook)
	case boil.BeforeDeleteHook:
		blogBeforeDeleteHooks = append(blogBeforeDeleteHooks, blogHook)
	case boil.BeforeUpsertHook:
		blogBeforeUpsertHooks = append(blogBeforeUpsertHooks, blogHook)
	case boil.AfterInsertHook:
		blogAfterInsertHooks = append(blogAfterInsertHooks, blogHook)
	case boil.AfterSelectHook:
		blogAfterSelectHooks = append(blogAfterSelectHooks, blogHook)
	case boil.AfterUpdateHook:
		blogAfterUpdateHooks = append(blogAfterUpdateHooks, blogHook)
	case boil.AfterDeleteHook:
		blogAfterDeleteHooks = append(blogAfterDeleteHooks, blogHook)
	case boil.AfterUpsertHook:
		blogAfterUpsertHooks = append(blogAfterUpsertHooks, blogHook)
	}
}

// One returns a single blog record from the query.
func (q blogQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Blog, error) {
	o := &Blog{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for blogs")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Blog records from the query.
func (q blogQuery) All(ctx context.Context, exec boil.ContextExecutor) (BlogSlice, error) {
	var o []*Blog

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Blog slice")
	}

	if len(blogAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Blog records in the query.
func (q blogQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count blogs rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q blogQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if blogs exists")
	}

	return count > 0, nil
}

// Member pointed to by the foreign key.
func (o *Blog) Member(mods ...qm.QueryMod) memberQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"member_id\" = ?", o.MemberID),
	}

	queryMods = append(queryMods, mods...)

	query := Members(queryMods...)
	queries.SetFrom(query.Query, "\"members\"")

	return query
}

// LoadMember allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (blogL) LoadMember(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBlog interface{}, mods queries.Applicator) error {
	var slice []*Blog
	var object *Blog

	if singular {
		object = maybeBlog.(*Blog)
	} else {
		slice = *maybeBlog.(*[]*Blog)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &blogR{}
		}
		args = append(args, object.MemberID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &blogR{}
			}

			for _, a := range args {
				if a == obj.MemberID {
					continue Outer
				}
			}

			args = append(args, obj.MemberID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`members`),
		qm.WhereIn(`members.member_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Member")
	}

	var resultSlice []*Member
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Member")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for members")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for members")
	}

	if len(blogAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Member = foreign
		if foreign.R == nil {
			foreign.R = &memberR{}
		}
		foreign.R.Blogs = append(foreign.R.Blogs, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.MemberID == foreign.MemberID {
				local.R.Member = foreign
				if foreign.R == nil {
					foreign.R = &memberR{}
				}
				foreign.R.Blogs = append(foreign.R.Blogs, local)
				break
			}
		}
	}

	return nil
}

// SetMember of the blog to the related item.
// Sets o.R.Member to related.
// Adds o to related.R.Blogs.
func (o *Blog) SetMember(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Member) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"blogs\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"member_id"}),
		strmangle.WhereClause("\"", "\"", 2, blogPrimaryKeyColumns),
	)
	values := []interface{}{related.MemberID, o.BlogID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.MemberID = related.MemberID
	if o.R == nil {
		o.R = &blogR{
			Member: related,
		}
	} else {
		o.R.Member = related
	}

	if related.R == nil {
		related.R = &memberR{
			Blogs: BlogSlice{o},
		}
	} else {
		related.R.Blogs = append(related.R.Blogs, o)
	}

	return nil
}

// Blogs retrieves all the records using an executor.
func Blogs(mods ...qm.QueryMod) blogQuery {
	mods = append(mods, qm.From("\"blogs\""))
	return blogQuery{NewQuery(mods...)}
}

// FindBlog retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBlog(ctx context.Context, exec boil.ContextExecutor, blogID int, selectCols ...string) (*Blog, error) {
	blogObj := &Blog{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"blogs\" where \"blog_id\"=$1", sel,
	)

	q := queries.Raw(query, blogID)

	err := q.Bind(ctx, exec, blogObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from blogs")
	}

	if err = blogObj.doAfterSelectHooks(ctx, exec); err != nil {
		return blogObj, err
	}

	return blogObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Blog) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no blogs provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(blogColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	blogInsertCacheMut.RLock()
	cache, cached := blogInsertCache[key]
	blogInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			blogAllColumns,
			blogColumnsWithDefault,
			blogColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(blogType, blogMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(blogType, blogMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"blogs\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"blogs\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into blogs")
	}

	if !cached {
		blogInsertCacheMut.Lock()
		blogInsertCache[key] = cache
		blogInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Blog.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Blog) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	blogUpdateCacheMut.RLock()
	cache, cached := blogUpdateCache[key]
	blogUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			blogAllColumns,
			blogPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update blogs, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"blogs\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, blogPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(blogType, blogMapping, append(wl, blogPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update blogs row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for blogs")
	}

	if !cached {
		blogUpdateCacheMut.Lock()
		blogUpdateCache[key] = cache
		blogUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q blogQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for blogs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for blogs")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BlogSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), blogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"blogs\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, blogPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in blog slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all blog")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Blog) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no blogs provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(blogColumnsWithDefault, o)

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

	blogUpsertCacheMut.RLock()
	cache, cached := blogUpsertCache[key]
	blogUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			blogAllColumns,
			blogColumnsWithDefault,
			blogColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			blogAllColumns,
			blogPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert blogs, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(blogPrimaryKeyColumns))
			copy(conflict, blogPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"blogs\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(blogType, blogMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(blogType, blogMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert blogs")
	}

	if !cached {
		blogUpsertCacheMut.Lock()
		blogUpsertCache[key] = cache
		blogUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Blog record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Blog) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Blog provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), blogPrimaryKeyMapping)
	sql := "DELETE FROM \"blogs\" WHERE \"blog_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from blogs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for blogs")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q blogQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no blogQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from blogs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for blogs")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BlogSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(blogBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), blogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"blogs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, blogPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from blog slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for blogs")
	}

	if len(blogAfterDeleteHooks) != 0 {
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
func (o *Blog) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBlog(ctx, exec, o.BlogID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BlogSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BlogSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), blogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"blogs\".* FROM \"blogs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, blogPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BlogSlice")
	}

	*o = slice

	return nil
}

// BlogExists checks if the Blog row exists.
func BlogExists(ctx context.Context, exec boil.ContextExecutor, blogID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"blogs\" where \"blog_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, blogID)
	}
	row := exec.QueryRowContext(ctx, sql, blogID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if blogs exists")
	}

	return exists, nil
}
