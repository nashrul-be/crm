package actor

import (
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"nashrul-be/crm/entities"
	mocks_repositories "nashrul-be/crm/repositories/mocks"
	mocks_db "nashrul-be/crm/utils/db/mocks"
	"reflect"
	"testing"
)

type fields struct {
	actorRepository            mocks_repositories.ActorRepositoryInterface
	roleRepository             mocks_repositories.RoleRepositoryInterface
	registerApprovalRepository mocks_repositories.RegisterApprovalRepositoryInterface
}

func defaultUseCase(t *testing.T) fields {
	return fields{
		actorRepository:            *mocks_repositories.NewActorRepositoryInterface(t),
		roleRepository:             *mocks_repositories.NewRoleRepositoryInterface(t),
		registerApprovalRepository: *mocks_repositories.NewRegisterApprovalRepositoryInterface(t),
	}
}

func Test_useCase_ActivateActor(t *testing.T) {
	type args struct {
		username string
		actor    entities.Actor
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantErr      bool
		activeStatus bool
	}{
		{
			name:         "Test with verified actor",
			fields:       defaultUseCase(t),
			args:         args{username: "randomUser", actor: entities.Actor{Verified: true}},
			wantErr:      false,
			activeStatus: true,
		},
		{
			name:         "Test with unverified actor",
			fields:       defaultUseCase(t),
			args:         args{username: "otherRandom", actor: entities.Actor{Verified: false}},
			wantErr:      true,
			activeStatus: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCase{
				actorRepository:            &tt.fields.actorRepository,
				roleRepository:             &tt.fields.roleRepository,
				registerApprovalRepository: &tt.fields.registerApprovalRepository,
			}
			tt.fields.actorRepository.EXPECT().GetByUsername(tt.args.username).Return(tt.args.actor, nil)
			tt.args.actor.Active = tt.activeStatus
			tt.fields.actorRepository.EXPECT().Save(tt.args.actor).Return(nil)
			if err := uc.ActivateActor(tt.args.username); (err != nil) != tt.wantErr {
				t.Errorf("ActivateActor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_useCase_CreateActor(t *testing.T) {
	type args struct {
		actor *entities.Actor
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Normal Scenario",
			fields: defaultUseCase(t),
			args: args{actor: &entities.Actor{
				Username: "ntahlah_123",
				Password: "0909sadsdsdas099090090",
			}},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCase{
				actorRepository:            &tt.fields.actorRepository,
				roleRepository:             &tt.fields.roleRepository,
				registerApprovalRepository: &tt.fields.registerApprovalRepository,
			}
			tx := mocks_db.NewTransactor(t)
			createdActor := *tt.args.actor
			createdActor.ID = 11
			createdActor.RoleID = 12
			createdApproval := entities.RegisterApproval{
				AdminID: createdActor.ID,
				Status:  "pending",
			}
			sameUsername := func(actor entities.Actor) bool {
				return actor.Username == createdActor.Username
			}
			tt.fields.actorRepository.EXPECT().InitTransaction().Return(tx, nil).Once()
			tt.fields.actorRepository.EXPECT().Begin(tx).Return(uc.actorRepository).Once()
			tt.fields.registerApprovalRepository.EXPECT().Begin(tx).Return(uc.registerApprovalRepository).Once()
			tt.fields.actorRepository.On("Create", mock.MatchedBy(sameUsername)).Return(createdActor, nil).Once()
			tt.fields.registerApprovalRepository.EXPECT().Create(createdApproval).Return(entities.RegisterApproval{}, nil).Once()
			tx.EXPECT().Commit().Return(nil)
			tt.fields.actorRepository.EXPECT().GetByID(createdActor.ID).Return(createdActor, nil).Once()
			tt.fields.roleRepository.EXPECT().GetByID(createdActor.RoleID).Return(entities.Role{}, nil).Once()
			if err := uc.CreateActor(tt.args.actor); (err != nil) != tt.wantErr {
				t.Errorf("CreateActor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_useCase_DeactivateActor(t *testing.T) {
	type args struct {
		username string
		actor    entities.Actor
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantErr      bool
		activeStatus bool
	}{
		{
			name:         "Test with verified actor",
			fields:       defaultUseCase(t),
			args:         args{username: "randomUser", actor: entities.Actor{Verified: true}},
			wantErr:      false,
			activeStatus: false,
		},
		{
			name:         "Test with unverified actor",
			fields:       defaultUseCase(t),
			args:         args{username: "otherRandom", actor: entities.Actor{Verified: false}},
			wantErr:      true,
			activeStatus: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCase{
				actorRepository:            &tt.fields.actorRepository,
				roleRepository:             &tt.fields.roleRepository,
				registerApprovalRepository: &tt.fields.registerApprovalRepository,
			}
			tt.fields.actorRepository.EXPECT().GetByUsername(tt.args.username).Return(tt.args.actor, nil)
			tt.args.actor.Active = tt.activeStatus
			tt.fields.actorRepository.EXPECT().Save(tt.args.actor).Return(nil)
			if err := uc.DeactivateActor(tt.args.username); (err != nil) != tt.wantErr {
				t.Errorf("DeactivateActor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_useCase_DeleteActor(t *testing.T) {
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Normal case",
			fields:  defaultUseCase(t),
			args:    args{id: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCase{
				actorRepository:            &tt.fields.actorRepository,
				roleRepository:             &tt.fields.roleRepository,
				registerApprovalRepository: &tt.fields.registerApprovalRepository,
			}
			tx := mocks_db.NewTransactor(t)
			tt.fields.actorRepository.EXPECT().InitTransaction().Return(tx, nil)
			tt.fields.actorRepository.EXPECT().Begin(tx).Return(&tt.fields.actorRepository)
			tt.fields.registerApprovalRepository.EXPECT().Begin(tx).Return(&tt.fields.registerApprovalRepository)
			tt.fields.registerApprovalRepository.EXPECT().DeleteByAdminId(tt.args.id).Return(nil)
			tt.fields.actorRepository.EXPECT().Delete(tt.args.id).Return(nil)
			tx.EXPECT().Commit().Return(&gorm.DB{
				Error: nil,
			})
			if err := uc.DeleteActor(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteActor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_useCase_GetAllByUsername(t *testing.T) {
	type args struct {
		username string
		limit    uint
		offset   uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entities.Actor
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "Normal case",
			fields: defaultUseCase(t),
			args: args{
				username: "a",
				limit:    1,
				offset:   10,
			},
			want: []entities.Actor{
				{
					ID:       1,
					Username: "Apa aja",
					Password: "usahdfakjfakjfds",
					RoleID:   2,
					Verified: true,
					Active:   true,
				},
				{
					ID:       2,
					Username: "Siapa aja",
					Password: "usahdfakjfakjfdssdasdasd",
					RoleID:   2,
					Verified: true,
					Active:   true,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCase{
				actorRepository:            &tt.fields.actorRepository,
				roleRepository:             &tt.fields.roleRepository,
				registerApprovalRepository: &tt.fields.registerApprovalRepository,
			}
			tt.fields.actorRepository.EXPECT().GetAllByUsername(tt.args.username, tt.args.limit, tt.args.offset).Return(tt.want, nil)
			got, err := uc.GetAllByUsername(tt.args.username, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllByUsername() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_useCase_GetByID(t *testing.T) {
	type args struct {
		id uint
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantActor entities.Actor
		wantErr   bool
	}{
		{
			name:   "Normal case",
			fields: defaultUseCase(t),
			args:   args{id: 1},
			wantActor: entities.Actor{
				ID:       1,
				Username: "apapsdad",
				Password: "lkasjdalskdjaslkfd",
				RoleID:   2,
				Role: entities.Role{
					ID:       2,
					RoleName: "admin",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCase{
				actorRepository:            &tt.fields.actorRepository,
				roleRepository:             &tt.fields.roleRepository,
				registerApprovalRepository: &tt.fields.registerApprovalRepository,
			}
			tt.fields.actorRepository.EXPECT().GetByID(tt.args.id).Return(tt.wantActor, nil)
			tt.fields.roleRepository.EXPECT().GetByID(tt.wantActor.RoleID).Return(tt.wantActor.Role, nil)
			gotActor, err := uc.GetByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotActor, tt.wantActor) {
				t.Errorf("GetByID() gotActor = %v, want %v", gotActor, tt.wantActor)
			}
		})
	}
}

func Test_useCase_GetByUsername(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantActor entities.Actor
		wantErr   bool
	}{
		{
			name:   "Normal Case",
			fields: defaultUseCase(t),
			args:   args{username: "Hoompa lumpa"},
			wantActor: entities.Actor{
				ID:       1,
				Username: "Hoompa lumpa",
				Password: "Poasjdkajsd,jadaskjhaskh",
				RoleID:   2,
			},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCase{
				actorRepository:            &tt.fields.actorRepository,
				roleRepository:             &tt.fields.roleRepository,
				registerApprovalRepository: &tt.fields.registerApprovalRepository,
			}
			tt.fields.actorRepository.EXPECT().GetByUsername(tt.args.username).Return(tt.wantActor, nil)
			tt.fields.roleRepository.EXPECT().GetByID(tt.wantActor.RoleID).Return(tt.wantActor.Role, nil)
			gotActor, err := uc.GetByUsername(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotActor, tt.wantActor) {
				t.Errorf("GetByUsername() gotActor = %v, want %v", gotActor, tt.wantActor)
			}
		})
	}
}

func Test_useCase_UpdateActor(t *testing.T) {
	type args struct {
		actor *entities.Actor
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Normal Case",
			fields: defaultUseCase(t),
			args: args{actor: &entities.Actor{
				ID:       1,
				Username: "apa_aja",
				Password: "aslkdjaldjasd97123kj",
				RoleID:   2,
			}},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCase{
				actorRepository:            &tt.fields.actorRepository,
				roleRepository:             &tt.fields.roleRepository,
				registerApprovalRepository: &tt.fields.registerApprovalRepository,
			}
			sameUsername := func(actor entities.Actor) bool {
				return actor.Username == tt.args.actor.Username
			}
			tt.fields.actorRepository.On("Update", mock.MatchedBy(sameUsername)).Return(nil)
			tt.fields.actorRepository.EXPECT().GetByID(tt.args.actor.ID).Return(*tt.args.actor, nil)
			tt.fields.roleRepository.EXPECT().GetByID(tt.args.actor.RoleID).Return(tt.args.actor.Role, nil)
			if err := uc.UpdateActor(tt.args.actor); (err != nil) != tt.wantErr {
				t.Errorf("UpdateActor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
