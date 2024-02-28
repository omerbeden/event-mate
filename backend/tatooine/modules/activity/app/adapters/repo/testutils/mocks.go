package testutils

import (
	"context"
	"fmt"
	"time"

	pgx "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
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
	Activities []model.Activity
	Rules      []string
	Flow       []string
	Current    int
}

func (m *MockRows) Close() {
	panic("unimplemented")
}
func (m *MockRows) Err() error {
	panic("unimplemented")
}
func (m *MockRows) CommandTag() pgconn.CommandTag {
	panic("unimplemented")
}
func (m *MockRows) FieldDescriptions() []pgconn.FieldDescription {
	panic("unimplemented")
}
func (m *MockRows) Next() bool {
	if len(m.Activities) > 0 {
		return m.Current < len(m.Activities)
	}
	if len(m.Rules) > 0 {
		return m.Current < len(m.Rules)
	}
	if len(m.Flow) > 0 {
		return m.Current < len(m.Flow)
	}
	return false
}
func (m *MockRows) Scan(dest ...any) error {
	if len(m.Activities) > 0 {
		activity := m.Activities[m.Current]
		*dest[0].(*int64) = activity.ID
		*dest[1].(*string) = activity.Title
		*dest[2].(*string) = activity.Category
		*dest[3].(*time.Time) = activity.StartAt
		*dest[4].(*time.Time) = activity.EndAt
		*dest[5].(*string) = activity.Content
		*dest[6].(*int) = activity.Quota
		*dest[7].(*int64) = activity.CreatedBy.ID
		*dest[8].(*string) = activity.CreatedBy.Name
		*dest[9].(*string) = activity.CreatedBy.LastName
		*dest[10].(*string) = activity.CreatedBy.ProfileImageUrl
		*dest[11].(*float64) = activity.CreatedBy.ProfilePoint
		*dest[12].(*string) = activity.Location.City
	}
	if len(m.Rules) > 0 {
		rule := m.Rules[m.Current]
		*dest[0].(*string) = rule
	}
	if len(m.Flow) > 0 {
		flow := m.Flow[m.Current]
		*dest[0].(*string) = flow
	}
	m.Current++
	return nil
}
func (m *MockRows) Values() ([]any, error) {
	panic("unimplemented")
}
func (m *MockRows) RawValues() [][]byte {
	panic("unimplemented")
}
func (m *MockRows) Conn() *pgx.Conn {
	panic("unimplemented")
}

type MockDBExecuter struct {
	QueryRowFunc func(ctx context.Context, sql string, args ...interface{}) pgx.Row
	QueryFunc    func(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	CopyFromFunc func(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
	ExecFunc     func(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
}

func (*MockDBExecuter) Begin(ctx context.Context) (pgx.Tx, error) {
	panic("unimplemented")
}

func (*MockDBExecuter) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	panic("unimplemented")
}

func (m *MockDBExecuter) Close() {
	panic("unimplemented")
}

func (m *MockDBExecuter) Config() *pgxpool.Config {
	panic("unimplemented")
}

func (m *MockDBExecuter) CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	if m.CopyFromFunc != nil {
		return m.CopyFromFunc(ctx, tableName, columnNames, rowSrc)
	}
	return 0, fmt.Errorf("CopyFrom not set")
}

func (m *MockDBExecuter) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
	if m.ExecFunc != nil {
		return m.ExecFunc(ctx, sql, arguments)
	}
	return pgconn.NewCommandTag(""), fmt.Errorf("CopyFrom not set")
}

func (*MockDBExecuter) Ping(ctx context.Context) error {
	panic("unimplemented")
}

func (m *MockDBExecuter) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	if m.QueryFunc != nil {
		return m.QueryFunc(ctx, sql, args...)
	}
	return nil, fmt.Errorf("QueryFunc not set")
}

func (m *MockDBExecuter) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	if m.QueryRowFunc != nil {
		return m.QueryRowFunc(ctx, sql, args...)
	}
	return &MockRow{}
}
