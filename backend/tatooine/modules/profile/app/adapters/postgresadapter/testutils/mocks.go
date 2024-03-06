package testutils

import (
	"context"
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/db"
)

type MockRow struct {
	user     model.UserProfile
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
		*dest[1].(*string) = user.Name
		*dest[2].(*string) = user.LastName
		*dest[3].(*string) = user.ProfileImageUrl
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
}

func (*MockDBExecuter) Begin(ctx context.Context) (db.Tx, error) {
	panic("unimplemented")
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
