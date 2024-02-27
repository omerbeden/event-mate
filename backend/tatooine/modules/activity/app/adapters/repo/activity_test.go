package repo_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	pgx "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/adapters/repo"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/stretchr/testify/assert"
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
	return m.Current < len(m.Activities)
}
func (m *MockRows) Scan(dest ...any) error {
	activity := m.Activities[m.Current]
	fmt.Printf("activities: %+v\n", activity)
	*dest[0].(*int64) = activity.ID
	*dest[1].(*string) = activity.Title
	*dest[2].(*string) = activity.Category
	*dest[3].(*time.Time) = activity.StartAt
	*dest[4].(*time.Time) = activity.EndAt
	*dest[5].(*int64) = activity.CreatedBy.ID
	*dest[6].(*string) = activity.CreatedBy.Name
	*dest[7].(*string) = activity.CreatedBy.LastName
	*dest[8].(*string) = activity.CreatedBy.ProfileImageUrl
	*dest[9].(*float64) = activity.CreatedBy.ProfilePoint
	*dest[10].(*string) = activity.Location.City

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

func TestCreateActivity(t *testing.T) {
	startAt := time.Now()
	endAt := time.Now().Add(2 * time.Hour)

	test := []struct {
		name        string
		setupMock   func(*MockDBExecuter)
		id          int64
		activity    *model.Activity
		expectError bool
	}{
		{
			name: "should create an activity",
			id:   1,
			activity: &model.Activity{
				Title:     "Sample Activity",
				Category:  "Outdoor",
				CreatedBy: model.User{ID: 1},
				StartAt:   startAt,
				EndAt:     endAt,
				Content:   "Sample Content",
			},
			expectError: false,
			setupMock: func(md *MockDBExecuter) {
				md.QueryRowFunc = func(ctx context.Context, sql string, args ...any) pgx.Row {
					return &MockRow{
						ScanFunc: func(dest ...any) error {
							*dest[0].(*int64) = int64(1)
							return nil
						},
					}
				}
			},
		},
		{
			name: "should return an error",
			id:   2,
			activity: &model.Activity{
				Title:     "Sample Activity",
				Category:  "Outdoor",
				CreatedBy: model.User{ID: 1},
				StartAt:   startAt,
				EndAt:     endAt,
				Content:   "Sample Content",
			},
			expectError: true,
			setupMock: func(md *MockDBExecuter) {
				md.QueryRowFunc = func(ctx context.Context, sql string, args ...any) pgx.Row {
					return &MockRow{
						ScanFunc: func(dest ...any) error {
							return errors.New("database error")
						},
					}
				}
			},
		},
	}
	for _, tc := range test {
		t.Run(fmt.Sprintf("%s,%d", tc.name, tc.id), func(t *testing.T) {
			mockDB := new(MockDBExecuter)
			if tc.setupMock != nil {
				tc.setupMock(mockDB)
			}

			activityRepository := repo.NewActivityRepo(mockDB)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			res, err := activityRepository.Create(ctx, *tc.activity)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, int64(1), res.ID)
			}
		})
	}
}

func TestAddParticipants(t *testing.T) {
	startAt := time.Now()
	endAt := time.Now().Add(2 * time.Hour)

	tests := []struct {
		name        string
		setupMock   func(*MockDBExecuter)
		id          int64
		activity    *model.Activity
		expectError bool
	}{
		{
			name: "should add participants",
			id:   1,
			activity: &model.Activity{
				Title:        "Sample Activity",
				Category:     "Outdoor",
				CreatedBy:    model.User{ID: 1},
				StartAt:      startAt,
				EndAt:        endAt,
				Content:      "Sample Content",
				Participants: []model.User{{ID: 1}, {ID: 2}, {ID: 3}},
			},
			expectError: false,
			setupMock: func(md *MockDBExecuter) {
				md.CopyFromFunc = func(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
					return 3, nil
				}
			},
		},
		{
			name: "should return an error",
			id:   2,
			activity: &model.Activity{
				Title:     "Sample Activity",
				Category:  "Outdoor",
				CreatedBy: model.User{ID: 1},
				StartAt:   startAt,
				EndAt:     endAt,
				Content:   "Sample Content",
			},
			expectError: true,
			setupMock: func(md *MockDBExecuter) {
				md.QueryRowFunc = func(ctx context.Context, sql string, args ...any) pgx.Row {
					return &MockRow{
						ScanFunc: func(dest ...any) error {
							return errors.New("database error")
						},
					}
				}
			},
		},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s,%d", tc.name, tc.id), func(t *testing.T) {
			mockDB := new(MockDBExecuter)
			if tc.setupMock != nil {
				tc.setupMock(mockDB)
			}

			activityRepository := repo.NewActivityRepo(mockDB)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			err := activityRepository.AddParticipants(ctx, *tc.activity)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestAddParticipant(t *testing.T) {
	tests :=
		[]struct {
			name        string
			setupMock   func(*MockDBExecuter)
			id          int
			activityId  int64
			user        model.User
			expectError bool
		}{
			{
				name:        "should add participant",
				id:          1,
				activityId:  int64(1),
				user:        model.User{},
				expectError: false,
				setupMock: func(md *MockDBExecuter) {
					md.ExecFunc = func(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
						return pgconn.NewCommandTag("INSERT 0 1"), nil
					}
				},
			},
			{
				name:        "should return error",
				id:          2,
				activityId:  int64(1),
				user:        model.User{},
				expectError: true,
				setupMock: func(md *MockDBExecuter) {
					md.ExecFunc = func(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
						return pgconn.NewCommandTag(""), errors.New("database error")
					}
				},
			},
		}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s,%d", tc.name, tc.id), func(t *testing.T) {
			mockDB := new(MockDBExecuter)
			if tc.setupMock != nil {
				tc.setupMock(mockDB)
			}

			activityRepository := repo.NewActivityRepo(mockDB)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			err := activityRepository.AddParticipant(ctx, tc.activityId, tc.user)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetActivityByID(t *testing.T) {
	startAt := time.Now()
	endAt := time.Now().Add(2 * time.Hour)

	test := []struct {
		name        string
		setupMock   func(*MockDBExecuter)
		id          int64
		expected    *model.Activity
		expectError bool
	}{
		{
			name: "TestGetActivityByID Success",
			id:   int64(1),
			expected: &model.Activity{
				ID:        1,
				Title:     "Sample Activity",
				Category:  "Outdoor",
				CreatedBy: model.User{ID: 1},
				StartAt:   startAt,
				EndAt:     endAt,
				Location:  model.Location{City: "London"},
			},
			expectError: false,
			setupMock: func(md *MockDBExecuter) {
				md.QueryRowFunc = func(ctx context.Context, sql string, args ...any) pgx.Row {
					return &MockRow{
						ScanFunc: func(dest ...any) error {
							*dest[0].(*int64) = int64(1)
							*dest[1].(*string) = "Sample Activity"
							*dest[2].(*string) = "Outdoor"
							*dest[3].(*int64) = int64(1)
							*dest[4].(*time.Time) = startAt
							*dest[5].(*time.Time) = endAt
							*dest[6].(*string) = "London"
							return nil
						},
					}
				}
			},
		},
		{
			name:        "TestGetActivityByID Success",
			id:          int64(2),
			expected:    nil,
			expectError: true,
			setupMock: func(md *MockDBExecuter) {
				md.QueryRowFunc = func(ctx context.Context, sql string, args ...any) pgx.Row {
					return &MockRow{
						ScanFunc: func(dest ...any) error {
							return errors.New("database error")
						},
					}
				}
			},
		},
	}

	for _, tc := range test {
		t.Run(fmt.Sprintf("%s,%d", tc.name, tc.id), func(t *testing.T) {
			mockDB := new(MockDBExecuter)
			if tc.setupMock != nil {
				tc.setupMock(mockDB)
			}

			activityRepository := repo.NewActivityRepo(mockDB)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			res, err := activityRepository.GetByID(ctx, tc.id)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, res)
			}
		})
	}
}

func TestGetActivitiesByLocation(t *testing.T) {
	startAt := time.Now()
	endAt := time.Now().Add(2 * time.Hour)
	activities := []model.Activity{
		{
			ID:       1,
			Title:    "Sample Activity",
			Category: "Outdoor",
			StartAt:  startAt,
			EndAt:    endAt,
			CreatedBy: model.User{ID: 1,
				Name:            "jack",
				LastName:        "sparrow",
				ProfileImageUrl: "imageurl.png",
				ProfilePoint:    3},
			Location: model.Location{City: "London"},
		},
		{
			ID:       2,
			Title:    "Sample Activity",
			Category: "travel",
			CreatedBy: model.User{ID: 1,
				Name:            "palpatine",
				LastName:        "senator",
				ProfileImageUrl: "imageurl.png",
				ProfilePoint:    7},
			StartAt:  startAt,
			EndAt:    endAt,
			Location: model.Location{City: "London"},
		},
	}

	test := []struct {
		name        string
		setupMock   func(*MockDBExecuter)
		id          int64
		expected    []model.Activity
		location    model.Location
		expectError bool
	}{
		{
			name:        "should return activities by location",
			id:          int64(1),
			location:    model.Location{City: "London"},
			expected:    activities,
			expectError: false,
			setupMock: func(md *MockDBExecuter) {
				md.QueryFunc = func(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
					return &MockRows{
						Activities: activities,
						Current:    0,
					}, nil

				}
			},
		},
		{
			name:        "should return database error",
			id:          int64(2),
			expected:    activities,
			location:    model.Location{City: "London"},
			expectError: true,
			setupMock: func(md *MockDBExecuter) {
				md.QueryFunc = func(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
					return nil, errors.New("database error")
				}
			},
		},
	}

	for _, tc := range test {
		t.Run(fmt.Sprintf("%s,%d", tc.name, tc.id), func(t *testing.T) {
			mockDB := new(MockDBExecuter)
			if tc.setupMock != nil {
				tc.setupMock(mockDB)
			}

			activityRepository := repo.NewActivityRepo(mockDB)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			res, err := activityRepository.GetByLocation(ctx, &tc.location)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, res)
				assert.Equal(t, tc.expected, res)
			}
		})
	}
}

func TestUpdateActivity(t *testing.T) {

	tests :=
		[]struct {
			name        string
			setupMock   func(*MockDBExecuter)
			id          int
			activity    model.Activity
			activityId  int64
			expectError bool
		}{
			{
				name:        "should update activity",
				id:          1,
				activity:    model.Activity{},
				activityId:  1,
				expectError: false,
				setupMock: func(md *MockDBExecuter) {
					md.ExecFunc = func(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
						return pgconn.NewCommandTag(""), nil
					}
				},
			},
			{
				name:        "should return error",
				id:          2,
				activityId:  2,
				activity:    model.Activity{},
				expectError: true,
				setupMock: func(md *MockDBExecuter) {
					md.ExecFunc = func(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
						return pgconn.NewCommandTag(""), errors.New("database error")
					}
				},
			},
		}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s,%d", tc.name, tc.id), func(t *testing.T) {
			mockDB := new(MockDBExecuter)
			if tc.setupMock != nil {
				tc.setupMock(mockDB)
			}

			activityRepository := repo.NewActivityRepo(mockDB)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			res, err := activityRepository.UpdateByID(ctx, tc.activityId, tc.activity)

			if tc.expectError {
				assert.Error(t, err)
				assert.False(t, res)
			} else {
				assert.NoError(t, err)
				assert.True(t, res)
			}
		})
	}
}

func TestDeleteActivityByID(t *testing.T) {

	tests :=
		[]struct {
			name        string
			setupMock   func(*MockDBExecuter)
			id          int
			activityId  int64
			expectError bool
		}{
			{
				name:        "should update activity",
				id:          1,
				activityId:  1,
				expectError: false,
				setupMock: func(md *MockDBExecuter) {
					md.ExecFunc = func(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
						return pgconn.NewCommandTag(""), nil
					}
				},
			},
			{
				name:        "should return error",
				id:          2,
				activityId:  2,
				expectError: true,
				setupMock: func(md *MockDBExecuter) {
					md.ExecFunc = func(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
						return pgconn.NewCommandTag(""), errors.New("database error")
					}
				},
			},
		}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s,%d", tc.name, tc.id), func(t *testing.T) {
			mockDB := new(MockDBExecuter)
			if tc.setupMock != nil {
				tc.setupMock(mockDB)
			}

			activityRepository := repo.NewActivityRepo(mockDB)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			res, err := activityRepository.DeleteByID(ctx, tc.activityId)

			if tc.expectError {
				assert.Error(t, err)
				assert.False(t, res)
			} else {
				assert.NoError(t, err)
				assert.True(t, res)
			}
		})
	}
}
