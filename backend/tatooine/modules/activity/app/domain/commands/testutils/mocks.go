package testutils

import (
	"context"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/adapters/postgresadapter/testutils"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/db"
)

type MockActivityRepo struct {
	Activity            model.Activity
	Activities          []model.GetActivityCommandResult
	CreateFunc          func(ctx context.Context, tx db.Tx, activity model.Activity) (*model.Activity, error)
	GetByIDFunc         func(ctx context.Context, id int64) (*model.Activity, error)
	GetByLocationFunc   func(ctx context.Context, location *model.Location) ([]model.GetActivityCommandResult, error)
	UpdateByIDFunc      func(ctx context.Context, activityId int64, activity model.Activity) (bool, error)
	DeleteByIDFunc      func(ctx context.Context, activityId int64) (bool, error)
	AddParticipantsFunc func(ctx context.Context, activityId int64, participants []model.User) error
	AddParticipantFunc  func(ctx context.Context, activityId int64, participant model.User) error
	GetPartipantsFunc   func(ctx context.Context, activityId int64) ([]model.User, error)
}

func (m *MockActivityRepo) Create(ctx context.Context, tx db.Tx, activity model.Activity) (*model.Activity, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, tx, activity)
	}
	return &m.Activity, nil
}
func (m *MockActivityRepo) GetByID(ctx context.Context, activityId int64) (*model.Activity, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(ctx, activityId)
	}
	return &m.Activity, nil
}
func (m *MockActivityRepo) GetByLocation(ctx context.Context, location *model.Location) ([]model.GetActivityCommandResult, error) {
	if m.GetByLocationFunc != nil {
		return m.GetByLocationFunc(ctx, location)
	}
	return m.Activities, nil
}
func (m *MockActivityRepo) UpdateByID(ctx context.Context, activityId int64, activity model.Activity) (bool, error) {
	if m.UpdateByIDFunc != nil {
		return m.UpdateByIDFunc(ctx, activityId, activity)
	}
	return true, nil
}
func (m *MockActivityRepo) DeleteByID(ctx context.Context, activityId int64) (bool, error) {
	if m.DeleteByIDFunc != nil {
		return m.DeleteByIDFunc(ctx, activityId)
	}
	return true, nil
}
func (m *MockActivityRepo) AddParticipants(ctx context.Context, activityId int64, participants []model.User) error {
	if m.AddParticipantsFunc != nil {
		return m.AddParticipantsFunc(ctx, activityId, participants)
	}
	return nil
}
func (m *MockActivityRepo) AddParticipant(ctx context.Context, activityId int64, participant model.User) error {
	if m.AddParticipantFunc != nil {
		return m.AddParticipantFunc(ctx, activityId, participant)
	}
	return nil
}
func (m *MockActivityRepo) GetParticipants(ctx context.Context, activityId int64) ([]model.User, error) {
	if m.GetPartipantsFunc != nil {
		return m.GetPartipantsFunc(ctx, activityId)
	}
	return m.Activity.Participants, nil
}

type MockActivityRulesRepo struct {
	Rules                   []string
	GetActivityRulesFunc    func(ctx context.Context, activityId int64) ([]string, error)
	CreateActivityRulesFunc func(ctx context.Context, tx db.Tx, activityId int64, rules []string) error
}

func (m *MockActivityRulesRepo) CreateActivityRules(ctx context.Context, tx db.Tx, activityId int64, rules []string) error {
	if m.CreateActivityRulesFunc != nil {
		return m.CreateActivityRulesFunc(ctx, tx, activityId, rules)
	}
	return nil
}
func (m *MockActivityRulesRepo) GetActivityRules(ctx context.Context, activityId int64) ([]string, error) {
	if m.GetActivityRulesFunc != nil {
		return m.GetActivityRulesFunc(ctx, activityId)
	}
	return m.Rules, nil
}

type MockActivityFlowRepo struct {
	Flow                   []string
	CreateActivityFlowFunc func(ctx context.Context, tx db.Tx, activityId int64, flow []string) error
	GetActivityFlowFunc    func(context.Context, int64) ([]string, error)
}

func (m *MockActivityFlowRepo) CreateActivityFlow(ctx context.Context, tx db.Tx, activityId int64, flow []string) error {
	if m.CreateActivityFlowFunc != nil {
		return m.CreateActivityFlowFunc(ctx, tx, activityId, flow)
	}
	return nil
}
func (m *MockActivityFlowRepo) GetActivityFlow(ctx context.Context, activityId int64) ([]string, error) {
	if m.GetActivityFlowFunc != nil {
		return m.GetActivityFlowFunc(ctx, activityId)
	}
	return m.Flow, nil
}

type MockLocationRepo struct {
	CreateFunc     func(ctx context.Context, tx db.Tx, location *model.Location) (bool, error)
	UpdateByIDFunc func(ctx context.Context, activity model.Location) (bool, error)
}

func (m *MockLocationRepo) Create(ctx context.Context, tx db.Tx, location *model.Location) (bool, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, tx, location)
	}
	return true, nil
}
func (m *MockLocationRepo) UpdateByID(ctx context.Context, activity model.Location) (bool, error) {
	if m.UpdateByIDFunc != nil {
		return m.UpdateByIDFunc(ctx, activity)
	}
	return true, nil
}

type MockTxnManager struct{}

func (m *MockTxnManager) Begin(ctx context.Context) (db.Tx, error) {
	return &testutils.MockTx{}, nil
}
