package register_approval

import (
	"nashrul-be/crm/entities"
	mocks_repositories "nashrul-be/crm/repositories/mocks"
	mocks_db "nashrul-be/crm/utils/db/mocks"
	"reflect"
	"testing"
)

type fields struct {
	registerRepository mocks_repositories.RegisterApprovalRepositoryInterface
	actorRepository    mocks_repositories.ActorRepositoryInterface
}

func defaultUseCase(t *testing.T) fields {
	return fields{
		registerRepository: *mocks_repositories.NewRegisterApprovalRepositoryInterface(t),
		actorRepository:    *mocks_repositories.NewActorRepositoryInterface(t),
	}
}

func Test_registerApprovalUseCase_Approve(t *testing.T) {
	type args struct {
		username     []string
		superAdminID uint
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult map[string][]string
		wantErr    bool
	}{
		{
			name:   "Normal Case",
			fields: defaultUseCase(t),
			args: args{
				username:     []string{"siapa_aja"},
				superAdminID: 1,
			},
			wantResult: map[string][]string{
				"failed":  {},
				"success": {"siapa_aja"},
			},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := registerApprovalUseCase{
				registerRepository: &tt.fields.registerRepository,
				actorRepository:    &tt.fields.actorRepository,
			}
			actor := []entities.Actor{
				{
					ID:       2,
					Username: "siapa_aja",
					Password: "asdjkalsd",
					RoleID:   2,
				},
			}
			ids := []uint{actor[0].ID}
			approval := []entities.RegisterApproval{
				{
					ID:           1,
					AdminID:      2,
					SuperAdminID: 0,
					Status:       "pending",
				},
			}
			tx := mocks_db.NewTransactor(t)
			tt.fields.actorRepository.EXPECT().GetByUsernameBatch(tt.args.username).Return(actor, nil)
			tt.fields.registerRepository.EXPECT().GetByAdminIdBatch(ids).Return(approval, nil)
			tt.fields.registerRepository.EXPECT().InitTransaction().Return(tx, nil)
			tt.fields.actorRepository.EXPECT().Begin(tx).Return(&tt.fields.actorRepository)
			tt.fields.registerRepository.EXPECT().Begin(tx).Return(&tt.fields.registerRepository)
			approval[0] = approval[0].Approve(tt.args.superAdminID)
			tt.fields.registerRepository.EXPECT().Update(approval[0]).Return(nil)
			actor[0] = actor[0].Verify().Activate()
			tt.fields.actorRepository.EXPECT().Update(actor[0]).Return(nil)
			tx.EXPECT().Commit().Return(nil)
			gotResult, err := uc.Approve(tt.args.username, tt.args.superAdminID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Approve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Approve() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_registerApprovalUseCase_GetAllPendingApproval(t *testing.T) {
	tests := []struct {
		name          string
		fields        fields
		wantApprovals []entities.RegisterApproval
		wantErr       bool
	}{
		{
			name:   "Normal case",
			fields: defaultUseCase(t),
			wantApprovals: []entities.RegisterApproval{
				{
					ID:           1,
					AdminID:      2,
					Admin:        entities.Actor{},
					SuperAdminID: 1,
					Status:       "pending",
				},
			},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := registerApprovalUseCase{
				registerRepository: &tt.fields.registerRepository,
				actorRepository:    &tt.fields.actorRepository,
			}
			tt.fields.registerRepository.EXPECT().GetAllPendingApproval().Return(tt.wantApprovals, nil)
			gotApprovals, err := uc.GetAllPendingApproval()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllPendingApproval() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotApprovals, tt.wantApprovals) {
				t.Errorf("GetAllPendingApproval() gotApprovals = %v, want %v", gotApprovals, tt.wantApprovals)
			}
		})
	}
}

func Test_registerApprovalUseCase_Rejected(t *testing.T) {
	type args struct {
		username     []string
		superAdminID uint
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult map[string][]string
		wantErr    bool
	}{
		{
			name:   "Normal case",
			fields: defaultUseCase(t),
			args: args{
				username:     []string{"siapa_aja"},
				superAdminID: 1,
			},
			wantResult: map[string][]string{
				"failed":  {},
				"success": {"siapa_aja"},
			},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := registerApprovalUseCase{
				registerRepository: &tt.fields.registerRepository,
				actorRepository:    &tt.fields.actorRepository,
			}
			actor := []entities.Actor{
				{
					ID:       2,
					Username: "siapa_aja",
					Password: "asdjkalsd",
					RoleID:   2,
				},
			}
			ids := []uint{actor[0].ID}
			approval := []entities.RegisterApproval{
				{
					ID:           1,
					AdminID:      2,
					SuperAdminID: 0,
					Status:       "pending",
				},
			}
			tt.fields.actorRepository.EXPECT().GetByUsernameBatch(tt.args.username).Return(actor, nil)
			tt.fields.registerRepository.EXPECT().GetByAdminIdBatch(ids).Return(approval, nil)
			approval[0] = approval[0].Reject(tt.args.superAdminID)
			tt.fields.registerRepository.EXPECT().Update(approval[0]).Return(nil)
			gotResult, err := uc.Rejected(tt.args.username, tt.args.superAdminID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rejected() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Rejected() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
