package repo_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	pgx "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/adapters/repo"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/adapters/repo/testutils"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateActivity(t *testing.T) {
	startAt := time.Now()
	endAt := time.Now().Add(2 * time.Hour)

	test := []struct {
		name        string
		setupMock   func(*testutils.MockDBExecuter)
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
				Quota:     3,
			},
			expectError: false,
			setupMock: func(md *testutils.MockDBExecuter) {
				md.QueryRowFunc = func(ctx context.Context, sql string, args ...any) pgx.Row {
					return &testutils.MockRow{
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
				Quota:     3,
			},
			expectError: true,
			setupMock: func(md *testutils.MockDBExecuter) {
				md.QueryRowFunc = func(ctx context.Context, sql string, args ...any) pgx.Row {
					return &testutils.MockRow{
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
			mockDB := new(testutils.MockDBExecuter)
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
		setupMock   func(*testutils.MockDBExecuter)
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
				Quota:        3,
				Participants: []model.User{{ID: 1}, {ID: 2}, {ID: 3}},
			},
			expectError: false,
			setupMock: func(md *testutils.MockDBExecuter) {
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
				Quota:     3,
				Content:   "Sample Content",
			},
			expectError: true,
			setupMock: func(md *testutils.MockDBExecuter) {
				md.CopyFromFunc = func(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
					return 0, errors.New("database error")
				}
			},
		},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s,%d", tc.name, tc.id), func(t *testing.T) {
			mockDB := new(testutils.MockDBExecuter)
			if tc.setupMock != nil {
				tc.setupMock(mockDB)
			}

			activityRepository := repo.NewActivityRepo(mockDB)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			err := activityRepository.AddParticipants(ctx, *&tc.activity.ID, tc.activity.Participants)

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
			setupMock   func(*testutils.MockDBExecuter)
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
				setupMock: func(md *testutils.MockDBExecuter) {
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
				setupMock: func(md *testutils.MockDBExecuter) {
					md.ExecFunc = func(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
						return pgconn.NewCommandTag(""), errors.New("database error")
					}
				},
			},
		}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s,%d", tc.name, tc.id), func(t *testing.T) {
			mockDB := new(testutils.MockDBExecuter)
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
		setupMock   func(*testutils.MockDBExecuter)
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
				Content:   "Sample Content",
				Quota:     3,
				Location:  model.Location{City: "London"},
			},
			expectError: false,
			setupMock: func(md *testutils.MockDBExecuter) {
				md.QueryRowFunc = func(ctx context.Context, sql string, args ...any) pgx.Row {
					return &testutils.MockRow{
						ScanFunc: func(dest ...any) error {
							*dest[0].(*int64) = int64(1)
							*dest[1].(*string) = "Sample Activity"
							*dest[2].(*string) = "Outdoor"
							*dest[3].(*int64) = int64(1)
							*dest[4].(*time.Time) = startAt
							*dest[5].(*time.Time) = endAt
							*dest[6].(*string) = "Sample Content"
							*dest[7].(*int) = 3
							*dest[8].(*string) = "London"
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
			setupMock: func(md *testutils.MockDBExecuter) {
				md.QueryRowFunc = func(ctx context.Context, sql string, args ...any) pgx.Row {
					return &testutils.MockRow{
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
			mockDB := new(testutils.MockDBExecuter)
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
			Quota:    3,
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
			Quota:    3,
		},
	}

	test := []struct {
		name        string
		setupMock   func(*testutils.MockDBExecuter)
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
			setupMock: func(md *testutils.MockDBExecuter) {
				md.QueryFunc = func(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
					return &testutils.MockRows{
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
			setupMock: func(md *testutils.MockDBExecuter) {
				md.QueryFunc = func(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
					return nil, errors.New("database error")
				}
			},
		},
	}

	for _, tc := range test {
		t.Run(fmt.Sprintf("%s,%d", tc.name, tc.id), func(t *testing.T) {
			mockDB := new(testutils.MockDBExecuter)
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
			setupMock   func(*testutils.MockDBExecuter)
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
				setupMock: func(md *testutils.MockDBExecuter) {
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
				setupMock: func(md *testutils.MockDBExecuter) {
					md.ExecFunc = func(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
						return pgconn.NewCommandTag(""), errors.New("database error")
					}
				},
			},
		}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s,%d", tc.name, tc.id), func(t *testing.T) {
			mockDB := new(testutils.MockDBExecuter)
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
			setupMock   func(*testutils.MockDBExecuter)
			id          int
			activityId  int64
			expectError bool
		}{
			{
				name:        "should update activity",
				id:          1,
				activityId:  1,
				expectError: false,
				setupMock: func(md *testutils.MockDBExecuter) {
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
				setupMock: func(md *testutils.MockDBExecuter) {
					md.ExecFunc = func(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
						return pgconn.NewCommandTag(""), errors.New("database error")
					}
				},
			},
		}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s,%d", tc.name, tc.id), func(t *testing.T) {
			mockDB := new(testutils.MockDBExecuter)
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
