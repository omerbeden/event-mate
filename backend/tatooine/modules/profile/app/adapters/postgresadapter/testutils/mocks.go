package testutils

import (
	"context"
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/db"
)

type MockRow struct {
	ScanFunc func(dest ...interface{}) error
}

func (m *MockRow) Scan(dest ...any) error {
	if m.ScanFunc != nil {
		return m.ScanFunc(dest...)
	}
	return fmt.Errorf("ScanFunc not set")
}

type MockRows struct {
	Users   []model.UserProfile
	Current int
}

func (m *MockRows) Close() {
	panic("unimplemented")
}
func (m *MockRows) Err() error {
	panic("unimplemented")
}
func (m *MockRows) Next() bool {
	if len(m.Users) > 0 {
		return m.Current < len(m.Users)
	}

	return false
}
func (m *MockRows) Scan(dest ...any) error {

	if len(m.Users) > 0 {
		user := m.Users[m.Current]
		*dest[0].(*int64) = user.Id
		*dest[1].(*string) = user.Header.Name
		*dest[2].(*string) = user.Header.LastName
		*dest[3].(*string) = user.Header.ProfileImageUrl
		*dest[4].(*string) = user.Email
		*dest[5].(*float32) = user.Stat.Point
		*dest[6].(*string) = user.Adress.City

	}
	m.Current++
	return nil
}

type MockDBExecuter struct {
	QueryRowFunc func(ctx context.Context, sql string, args ...interface{}) db.Row
	QueryFunc    func(ctx context.Context, sql string, args ...any) (db.Rows, error)
	CopyFromFunc func(ctx context.Context, tableName db.Identifier, columnNames []string, rowSrc db.CopyFromSource) (int64, error)
	ExecFunc     func(ctx context.Context, sql string, arguments ...any) (db.CommandTag, error)
	BeginFunc    func(ctx context.Context) (db.Tx, error)
}

type MockTx struct {
	CommitFunc   func(ctx context.Context) error
	RollbackFunc func(ctx context.Context) error
	CopyFromFunc func(ctx context.Context, tableName db.Identifier, columnNames []string, rowSrc db.CopyFromSource) (int64, error)
	ExecFunc     func(ctx context.Context, sql string, arguments ...any) (commandTag db.CommandTag, err error)
	QueryFunc    func(ctx context.Context, sql string, args ...any) (db.Rows, error)
	QueryRowFunc func(ctx context.Context, sql string, args ...any) db.Row
}

func (m *MockTx) Commit(ctx context.Context) error {
	if m.CommitFunc != nil {
		return m.CommitFunc(ctx)
	}

	return nil
}
func (m *MockTx) Rollback(ctx context.Context) error {
	if m.RollbackFunc != nil {
		return m.RollbackFunc(ctx)
	}
	return nil
}
func (m *MockTx) CopyFrom(ctx context.Context, tableName db.Identifier, columnNames []string, rowSrc db.CopyFromSource) (int64, error) {
	if m.CopyFromFunc != nil {
		return m.CopyFromFunc(ctx, tableName, columnNames, rowSrc)
	}

	return 0, nil
}
func (m *MockTx) Exec(ctx context.Context, sql string, arguments ...any) (commandTag db.CommandTag, err error) {
	if m.ExecFunc != nil {
		return m.ExecFunc(ctx, sql, arguments...)
	}
	fmt.Println("tx exec not set")
	return db.CommandTag{}, nil
}
func (m *MockTx) Query(ctx context.Context, sql string, args ...any) (db.Rows, error) {
	if m.QueryFunc != nil {
		return m.QueryFunc(ctx, sql, args...)
	}
	return &MockRows{}, nil
}
func (m *MockTx) QueryRow(ctx context.Context, sql string, args ...any) db.Row {
	if m.QueryRowFunc != nil {
		return m.QueryRowFunc(ctx, sql, args)
	}

	return &MockRow{}
}

func (m *MockDBExecuter) Begin(ctx context.Context) (db.Tx, error) {
	if m.BeginFunc != nil {
		return m.BeginFunc(ctx)
	}

	return &MockTx{}, fmt.Errorf("begin not set")
}

func (m *MockDBExecuter) CopyFrom(ctx context.Context, tableName db.Identifier, columnNames []string, rowSrc db.CopyFromSource) (int64, error) {
	if m.CopyFromFunc != nil {
		return m.CopyFromFunc(ctx, tableName, columnNames, rowSrc)
	}
	return 0, fmt.Errorf("CopyFrom not set")
}

func (m *MockDBExecuter) Exec(ctx context.Context, sql string, arguments ...any) (db.CommandTag, error) {
	if m.ExecFunc != nil {
		return m.ExecFunc(ctx, sql, arguments)
	}

	return db.CommandTag{}, fmt.Errorf("CopyFrom not set")
}

func (*MockDBExecuter) Ping(ctx context.Context) error {
	panic("unimplemented")
}

func (m *MockDBExecuter) Query(ctx context.Context, sql string, args ...any) (db.Rows, error) {
	if m.QueryFunc != nil {
		return m.QueryFunc(ctx, sql, args...)
	}
	return nil, fmt.Errorf("QueryFunc not set")
}

func (m *MockDBExecuter) QueryRow(ctx context.Context, sql string, args ...any) db.Row {
	if m.QueryRowFunc != nil {
		return m.QueryRowFunc(ctx, sql, args...)
	}
	return &MockRow{}
}
