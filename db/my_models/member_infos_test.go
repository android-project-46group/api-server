// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testMemberInfos(t *testing.T) {
	t.Parallel()

	query := MemberInfos()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMemberInfosDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MemberInfo{}
	if err = randomize.Struct(seed, o, memberInfoDBTypes, true, memberInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MemberInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MemberInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMemberInfosQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MemberInfo{}
	if err = randomize.Struct(seed, o, memberInfoDBTypes, true, memberInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MemberInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := MemberInfos().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MemberInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMemberInfosSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MemberInfo{}
	if err = randomize.Struct(seed, o, memberInfoDBTypes, true, memberInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MemberInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MemberInfoSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MemberInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMemberInfosExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MemberInfo{}
	if err = randomize.Struct(seed, o, memberInfoDBTypes, true, memberInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MemberInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MemberInfoExists(ctx, tx, o.MemberInfoID)
	if err != nil {
		t.Errorf("Unable to check if MemberInfo exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MemberInfoExists to return true, but got false.")
	}
}

func testMemberInfosFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MemberInfo{}
	if err = randomize.Struct(seed, o, memberInfoDBTypes, true, memberInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MemberInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	memberInfoFound, err := FindMemberInfo(ctx, tx, o.MemberInfoID)
	if err != nil {
		t.Error(err)
	}

	if memberInfoFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMemberInfosBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MemberInfo{}
	if err = randomize.Struct(seed, o, memberInfoDBTypes, true, memberInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MemberInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = MemberInfos().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMemberInfosOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MemberInfo{}
	if err = randomize.Struct(seed, o, memberInfoDBTypes, true, memberInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MemberInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := MemberInfos().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMemberInfosAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	memberInfoOne := &MemberInfo{}
	memberInfoTwo := &MemberInfo{}
	if err = randomize.Struct(seed, memberInfoOne, memberInfoDBTypes, false, memberInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MemberInfo struct: %s", err)
	}
	if err = randomize.Struct(seed, memberInfoTwo, memberInfoDBTypes, false, memberInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MemberInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = memberInfoOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = memberInfoTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MemberInfos().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMemberInfosCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	memberInfoOne := &MemberInfo{}
	memberInfoTwo := &MemberInfo{}
	if err = randomize.Struct(seed, memberInfoOne, memberInfoDBTypes, false, memberInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MemberInfo struct: %s", err)
	}
	if err = randomize.Struct(seed, memberInfoTwo, memberInfoDBTypes, false, memberInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MemberInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = memberInfoOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = memberInfoTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MemberInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func memberInfoBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *MemberInfo) error {
	*o = MemberInfo{}
	return nil
}

func memberInfoAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *MemberInfo) error {
	*o = MemberInfo{}
	return nil
}

func memberInfoAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *MemberInfo) error {
	*o = MemberInfo{}
	return nil
}

func memberInfoBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MemberInfo) error {
	*o = MemberInfo{}
	return nil
}

func memberInfoAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MemberInfo) error {
	*o = MemberInfo{}
	return nil
}

func memberInfoBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MemberInfo) error {
	*o = MemberInfo{}
	return nil
}

func memberInfoAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MemberInfo) error {
	*o = MemberInfo{}
	return nil
}

func memberInfoBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MemberInfo) error {
	*o = MemberInfo{}
	return nil
}

func memberInfoAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MemberInfo) error {
	*o = MemberInfo{}
	return nil
}

func testMemberInfosHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &MemberInfo{}
	o := &MemberInfo{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, memberInfoDBTypes, false); err != nil {
		t.Errorf("Unable to randomize MemberInfo object: %s", err)
	}

	AddMemberInfoHook(boil.BeforeInsertHook, memberInfoBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	memberInfoBeforeInsertHooks = []MemberInfoHook{}

	AddMemberInfoHook(boil.AfterInsertHook, memberInfoAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	memberInfoAfterInsertHooks = []MemberInfoHook{}

	AddMemberInfoHook(boil.AfterSelectHook, memberInfoAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	memberInfoAfterSelectHooks = []MemberInfoHook{}

	AddMemberInfoHook(boil.BeforeUpdateHook, memberInfoBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	memberInfoBeforeUpdateHooks = []MemberInfoHook{}

	AddMemberInfoHook(boil.AfterUpdateHook, memberInfoAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	memberInfoAfterUpdateHooks = []MemberInfoHook{}

	AddMemberInfoHook(boil.BeforeDeleteHook, memberInfoBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	memberInfoBeforeDeleteHooks = []MemberInfoHook{}

	AddMemberInfoHook(boil.AfterDeleteHook, memberInfoAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	memberInfoAfterDeleteHooks = []MemberInfoHook{}

	AddMemberInfoHook(boil.BeforeUpsertHook, memberInfoBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	memberInfoBeforeUpsertHooks = []MemberInfoHook{}

	AddMemberInfoHook(boil.AfterUpsertHook, memberInfoAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	memberInfoAfterUpsertHooks = []MemberInfoHook{}
}

func testMemberInfosInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MemberInfo{}
	if err = randomize.Struct(seed, o, memberInfoDBTypes, true, memberInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MemberInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MemberInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMemberInfosInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MemberInfo{}
	if err = randomize.Struct(seed, o, memberInfoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MemberInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(memberInfoColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := MemberInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMemberInfoToOneLocaleUsingLocale(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local MemberInfo
	var foreign Locale

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, memberInfoDBTypes, false, memberInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MemberInfo struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, localeDBTypes, false, localeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Locale struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.LocaleID = foreign.LocaleID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Locale().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.LocaleID != foreign.LocaleID {
		t.Errorf("want: %v, got %v", foreign.LocaleID, check.LocaleID)
	}

	slice := MemberInfoSlice{&local}
	if err = local.L.LoadLocale(ctx, tx, false, (*[]*MemberInfo)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Locale == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Locale = nil
	if err = local.L.LoadLocale(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Locale == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testMemberInfoToOneMemberUsingMember(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local MemberInfo
	var foreign Member

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, memberInfoDBTypes, false, memberInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MemberInfo struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, memberDBTypes, false, memberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Member struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.MemberID = foreign.MemberID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Member().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.MemberID != foreign.MemberID {
		t.Errorf("want: %v, got %v", foreign.MemberID, check.MemberID)
	}

	slice := MemberInfoSlice{&local}
	if err = local.L.LoadMember(ctx, tx, false, (*[]*MemberInfo)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Member == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Member = nil
	if err = local.L.LoadMember(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Member == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testMemberInfoToOneSetOpLocaleUsingLocale(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MemberInfo
	var b, c Locale

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, memberInfoDBTypes, false, strmangle.SetComplement(memberInfoPrimaryKeyColumns, memberInfoColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, localeDBTypes, false, strmangle.SetComplement(localePrimaryKeyColumns, localeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, localeDBTypes, false, strmangle.SetComplement(localePrimaryKeyColumns, localeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Locale{&b, &c} {
		err = a.SetLocale(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Locale != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.MemberInfos[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.LocaleID != x.LocaleID {
			t.Error("foreign key was wrong value", a.LocaleID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.LocaleID))
		reflect.Indirect(reflect.ValueOf(&a.LocaleID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.LocaleID != x.LocaleID {
			t.Error("foreign key was wrong value", a.LocaleID, x.LocaleID)
		}
	}
}
func testMemberInfoToOneSetOpMemberUsingMember(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MemberInfo
	var b, c Member

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, memberInfoDBTypes, false, strmangle.SetComplement(memberInfoPrimaryKeyColumns, memberInfoColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, memberDBTypes, false, strmangle.SetComplement(memberPrimaryKeyColumns, memberColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, memberDBTypes, false, strmangle.SetComplement(memberPrimaryKeyColumns, memberColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Member{&b, &c} {
		err = a.SetMember(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Member != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.MemberInfos[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.MemberID != x.MemberID {
			t.Error("foreign key was wrong value", a.MemberID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.MemberID))
		reflect.Indirect(reflect.ValueOf(&a.MemberID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.MemberID != x.MemberID {
			t.Error("foreign key was wrong value", a.MemberID, x.MemberID)
		}
	}
}

func testMemberInfosReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MemberInfo{}
	if err = randomize.Struct(seed, o, memberInfoDBTypes, true, memberInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MemberInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMemberInfosReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MemberInfo{}
	if err = randomize.Struct(seed, o, memberInfoDBTypes, true, memberInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MemberInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MemberInfoSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMemberInfosSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MemberInfo{}
	if err = randomize.Struct(seed, o, memberInfoDBTypes, true, memberInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MemberInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MemberInfos().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	memberInfoDBTypes = map[string]string{`MemberInfoID`: `integer`, `MemberID`: `integer`, `Birthday`: `date`, `BloodType`: `character varying`, `HeightCM`: `double precision`, `Generation`: `character varying`, `BlogURL`: `character varying`, `ImgURL`: `character varying`, `LocaleID`: `integer`}
	_                 = bytes.MinRead
)

func testMemberInfosUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(memberInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(memberInfoAllColumns) == len(memberInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MemberInfo{}
	if err = randomize.Struct(seed, o, memberInfoDBTypes, true, memberInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MemberInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MemberInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, memberInfoDBTypes, true, memberInfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MemberInfo struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMemberInfosSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(memberInfoAllColumns) == len(memberInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MemberInfo{}
	if err = randomize.Struct(seed, o, memberInfoDBTypes, true, memberInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MemberInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MemberInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, memberInfoDBTypes, true, memberInfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MemberInfo struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(memberInfoAllColumns, memberInfoPrimaryKeyColumns) {
		fields = memberInfoAllColumns
	} else {
		fields = strmangle.SetComplement(
			memberInfoAllColumns,
			memberInfoPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := MemberInfoSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testMemberInfosUpsert(t *testing.T) {
	t.Parallel()

	if len(memberInfoAllColumns) == len(memberInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := MemberInfo{}
	if err = randomize.Struct(seed, &o, memberInfoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MemberInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MemberInfo: %s", err)
	}

	count, err := MemberInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, memberInfoDBTypes, false, memberInfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MemberInfo struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MemberInfo: %s", err)
	}

	count, err = MemberInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
