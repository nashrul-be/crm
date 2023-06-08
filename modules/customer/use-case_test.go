package customer

import (
	"nashrul-be/crm/entities"
	mocks_repositories "nashrul-be/crm/repositories/mocks"
	"reflect"
	"testing"
)

type fields struct {
	customerRepository mocks_repositories.CustomerRepositoryInterface
}

func defaultUseCase(t *testing.T) fields {
	return fields{
		customerRepository: *mocks_repositories.NewCustomerRepositoryInterface(t),
	}
}

func Test_useCase_CreateCustomer(t *testing.T) {
	type args struct {
		customer *entities.Customer
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
			args: args{customer: &entities.Customer{
				ID:        1,
				FirstName: "John",
				LastName:  "Doe",
				Email:     "john.does@gmail.com",
				Avatar:    "",
			}},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCase{
				customerRepository: &tt.fields.customerRepository,
			}
			tt.fields.customerRepository.EXPECT().Create(*tt.args.customer).Return(*tt.args.customer, nil)
			if err := uc.CreateCustomer(tt.args.customer); (err != nil) != tt.wantErr {
				t.Errorf("CreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_useCase_DeleteCustomer(t *testing.T) {
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
			name:    "Normal Case",
			fields:  defaultUseCase(t),
			args:    args{id: 1},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCase{
				customerRepository: &tt.fields.customerRepository,
			}
			tt.fields.customerRepository.EXPECT().Delete(tt.args.id).Return(nil)
			if err := uc.DeleteCustomer(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteCustomer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_useCase_GetAll(t *testing.T) {
	type args struct {
		limit  uint
		offset uint
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantCustomers []entities.Customer
		wantErr       bool
	}{
		{
			name:   "Normal Case",
			fields: defaultUseCase(t),
			args: args{
				limit:  1,
				offset: 10,
			},
			wantCustomers: []entities.Customer{
				{
					ID:        1,
					FirstName: "akjsaslkdj",
					LastName:  "alksjdalkshdksajdh",
					Email:     "askjdhaksjdh@gmail.com",
				},
				{
					ID:        2,
					FirstName: "lashdalksdsajdbsada",
					LastName:  "lkasjdoiu2oeuo",
					Email:     "alksjdalksdklajsdljsd",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCase{
				customerRepository: &tt.fields.customerRepository,
			}
			tt.fields.customerRepository.EXPECT().GetAll(tt.args.limit, tt.args.offset).Return(tt.wantCustomers, nil)
			gotCustomers, err := uc.GetAll(tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCustomers, tt.wantCustomers) {
				t.Errorf("GetAll() gotCustomers = %v, want %v", gotCustomers, tt.wantCustomers)
			}
		})
	}
}

func Test_useCase_GetAllByEmail(t *testing.T) {
	type args struct {
		email  string
		limit  uint
		offset uint
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantCustomers []entities.Customer
		wantErr       bool
	}{
		{
			name:   "Normal Case",
			fields: defaultUseCase(t),
			args: args{
				email:  "mail.com",
				limit:  1,
				offset: 10,
			},
			wantCustomers: []entities.Customer{
				{
					ID:        1,
					FirstName: "apa_aja",
					LastName:  "lksajdk",
					Email:     "Asjdlks@gmail.com",
				},
				{
					ID:        2,
					FirstName: "apa_aja",
					LastName:  "lksajdk",
					Email:     "asasd@gmail.com",
				},
			},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCase{
				customerRepository: &tt.fields.customerRepository,
			}
			tt.fields.customerRepository.EXPECT().GetAllByEmail(tt.args.email, tt.args.limit, tt.args.offset).Return(tt.wantCustomers, nil)
			gotCustomers, err := uc.GetAllByEmail(tt.args.email, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCustomers, tt.wantCustomers) {
				t.Errorf("GetAllByEmail() gotCustomers = %v, want %v", gotCustomers, tt.wantCustomers)
			}
		})
	}
}

func Test_useCase_GetAllByName(t *testing.T) {
	type args struct {
		name   string
		limit  uint
		offset uint
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantCustomers []entities.Customer
		wantErr       bool
	}{
		{
			name:   "Normal Cae",
			fields: defaultUseCase(t),
			args: args{
				name:   "sukro",
				limit:  1,
				offset: 10,
			},
			wantCustomers: []entities.Customer{
				{
					ID:        1,
					FirstName: "sukro",
					LastName:  "polo",
					Email:     "sukro.polo@gmail.com",
				},
			},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCase{
				customerRepository: &tt.fields.customerRepository,
			}
			tt.fields.customerRepository.EXPECT().GetAllByName(tt.args.name, tt.args.limit, tt.args.offset).Return(tt.wantCustomers, nil)
			gotCustomers, err := uc.GetAllByName(tt.args.name, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCustomers, tt.wantCustomers) {
				t.Errorf("GetAllByName() gotCustomers = %v, want %v", gotCustomers, tt.wantCustomers)
			}
		})
	}
}

func Test_useCase_GetByID(t *testing.T) {
	type args struct {
		id uint
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantCustomer entities.Customer
		wantErr      bool
	}{
		{
			name:   "Normal case",
			fields: defaultUseCase(t),
			args:   args{id: 1},
			wantCustomer: entities.Customer{
				ID:        1,
				FirstName: "aklsjdlkasdj",
				LastName:  "alksdlkasdlkjad",
				Email:     "mail@gmail.com",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCase{
				customerRepository: &tt.fields.customerRepository,
			}
			tt.fields.customerRepository.EXPECT().GetByID(tt.args.id).Return(tt.wantCustomer, nil)
			gotCustomer, err := uc.GetByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCustomer, tt.wantCustomer) {
				t.Errorf("GetByID() gotCustomer = %v, want %v", gotCustomer, tt.wantCustomer)
			}
		})
	}
}

func Test_useCase_UpdateCustomer(t *testing.T) {
	type args struct {
		customer *entities.Customer
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
			args: args{
				customer: &entities.Customer{
					ID:        1,
					FirstName: "Apaaja",
					LastName:  "siapa",
					Email:     "apaaja@gmail.com",
				}},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCase{
				customerRepository: &tt.fields.customerRepository,
			}
			tt.fields.customerRepository.EXPECT().Update(*tt.args.customer).Return(nil)
			if err := uc.UpdateCustomer(tt.args.customer); (err != nil) != tt.wantErr {
				t.Errorf("UpdateCustomer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
