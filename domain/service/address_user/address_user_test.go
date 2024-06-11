package address_user

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"m1-article-service/domain/entity"
	mock_log "m1-article-service/mock/infrastructure"
	mock_address "m1-article-service/mock/repository"
	"testing"
	"time"
)

func TestService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})
	err := errors.New("error")

	var tests = []struct {
		name            string
		address         *entity.Address
		loggerMock      func() *mock_log.MockLog
		addressRepoMock func() *mock_address.MockAddress
		error           error
		ctx             context.Context
	}{
		{
			name: "success",
			loggerMock: func() *mock_log.MockLog {
				loggerInfra := mock_log.NewMockLog(ctrl)
				return loggerInfra
			},
			addressRepoMock: func() *mock_address.MockAddress {
				repoLogMock := mock_address.NewMockAddress(ctrl)
				repoLogMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(int64(1), nil)
				return repoLogMock
			},
			address: entity.NewAddress("c", "s", "co", "str", "3tgdsgds"),
			error:   nil,
			ctx:     context.Background(),
		},
		{
			name: "RepoError",
			loggerMock: func() *mock_log.MockLog {
				loggerInfra := mock_log.NewMockLog(ctrl)
				loggerInfra.EXPECT().Error(err).Return()
				return loggerInfra
			},
			addressRepoMock: func() *mock_address.MockAddress {
				repoLogMock := mock_address.NewMockAddress(ctrl)
				repoLogMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(int64(0), err)
				return repoLogMock
			},
			address: entity.NewAddress("c", "s", "co", "str", "3tgdsgds"),
			error:   err,
			ctx:     context.Background(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			logRepoMock := test.addressRepoMock()
			loggerMock := test.loggerMock()
			service := NewService(loggerMock, logRepoMock)
			_, err := service.Create(test.ctx, test.address)
			if !errors.Is(err, test.error) {
				t.Error("error is not equal")
			}

			loggerMock.EXPECT()
			logRepoMock.EXPECT()
		})
	}
}

func BenchmarkService_Create(b *testing.B) {
	ctrl := gomock.NewController(b)
	addressRepoMock := mock_address.NewMockAddress(ctrl)
	addressRepoMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(int64(1), nil)
	loggerMock := mock_log.NewMockLog(ctrl)
	b.ResetTimer()
	service := NewService(loggerMock, addressRepoMock)
	service.Create(context.Background(), entity.NewAddress("c", "s", "co", "str", "3tgdsgds"))
	fmt.Println(b.Elapsed())
	if b.Elapsed() > 150*time.Microsecond {
		b.Error("address_user service-createAddress takes too long to run")
	}
	loggerMock.EXPECT()
	addressRepoMock.EXPECT()

}

func TestService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})
	err := errors.New("error")

	var tests = []struct {
		name            string
		address         *entity.Address
		loggerMock      func() *mock_log.MockLog
		addressRepoMock func() *mock_address.MockAddress
		error           error
		ctx             context.Context
	}{
		{
			name: "success",
			loggerMock: func() *mock_log.MockLog {
				loggerInfra := mock_log.NewMockLog(ctrl)
				return loggerInfra
			},
			addressRepoMock: func() *mock_address.MockAddress {
				repoLogMock := mock_address.NewMockAddress(ctrl)
				repoLogMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
				return repoLogMock
			},
			address: entity.NewAddress("c", "s", "co", "str", "3tgdsgds"),
			error:   nil,
			ctx:     context.Background(),
		},
		{
			name: "RepoError",
			loggerMock: func() *mock_log.MockLog {
				loggerInfra := mock_log.NewMockLog(ctrl)
				loggerInfra.EXPECT().Error(err).Return()
				return loggerInfra
			},
			addressRepoMock: func() *mock_address.MockAddress {
				repoLogMock := mock_address.NewMockAddress(ctrl)
				repoLogMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(err)
				return repoLogMock
			},
			address: entity.NewAddress("c", "s", "co", "str", "3tgdsgds"),
			error:   err,
			ctx:     context.Background(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			logRepoMock := test.addressRepoMock()
			loggerMock := test.loggerMock()
			service := NewService(loggerMock, logRepoMock)
			err := service.Update(test.ctx, test.address)
			if !errors.Is(err, test.error) {
				t.Error("error is not equal")
			}

			loggerMock.EXPECT()
			logRepoMock.EXPECT()
		})
	}

}

func BenchmarkService_Update(b *testing.B) {
	ctrl := gomock.NewController(b)
	addressRepoMock := mock_address.NewMockAddress(ctrl)
	addressRepoMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
	loggerMock := mock_log.NewMockLog(ctrl)
	b.ResetTimer()

	service := NewService(loggerMock, addressRepoMock)
	service.Update(context.Background(), entity.NewAddress("c", "s", "co", "str", "3tgdsgds"))
	if b.Elapsed() > 150*time.Microsecond {
		b.Error("address_user service-update takes too long to run")
	}
	loggerMock.EXPECT()
	addressRepoMock.EXPECT()
}

func TestService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})
	err := errors.New("error")
	var tests = []struct {
		name            string
		id              int64
		loggerMock      func() *mock_log.MockLog
		addressRepoMock func() *mock_address.MockAddress
		error           error
		ctx             context.Context
	}{
		{
			name: "success",
			loggerMock: func() *mock_log.MockLog {
				loggerInfra := mock_log.NewMockLog(ctrl)
				return loggerInfra
			},
			addressRepoMock: func() *mock_address.MockAddress {
				repoLogMock := mock_address.NewMockAddress(ctrl)
				repoLogMock.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)
				return repoLogMock
			},
			id:    1,
			error: nil,
			ctx:   context.Background(),
		},
		{
			name: "RepoError",
			loggerMock: func() *mock_log.MockLog {
				loggerInfra := mock_log.NewMockLog(ctrl)
				loggerInfra.EXPECT().Error(err).Return()
				return loggerInfra
			},
			addressRepoMock: func() *mock_address.MockAddress {
				repoLogMock := mock_address.NewMockAddress(ctrl)
				repoLogMock.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(err)
				return repoLogMock
			},
			id:    1,
			error: err,
			ctx:   context.Background(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			logRepoMock := test.addressRepoMock()
			loggerMock := test.loggerMock()
			service := NewService(loggerMock, logRepoMock)
			err := service.Delete(test.ctx, test.id)
			if !errors.Is(err, test.error) {
				t.Error("error is not equal")
			}

			loggerMock.EXPECT()
			logRepoMock.EXPECT()
		})
	}
}

func BenchmarkService_Delete(b *testing.B) {
	ctrl := gomock.NewController(b)
	addressRepoMock := mock_address.NewMockAddress(ctrl)
	addressRepoMock.EXPECT().Delete(gomock.Any(), int64(1)).Return(nil)
	loggerMock := mock_log.NewMockLog(ctrl)
	b.ResetTimer()
	service := NewService(loggerMock, addressRepoMock)
	service.Delete(context.Background(), int64(1))
	if b.Elapsed() > 150*time.Microsecond {
		b.Error("address_user service-delete takes too long to run")
	}
	loggerMock.EXPECT()
	addressRepoMock.EXPECT()

}

func TestService_Detail(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})
	err := errors.New("error")
	address := entity.NewAddress("c", "s", "co", "str", "3tgdsgds")

	var tests = []struct {
		name            string
		id              int64
		loggerMock      func() *mock_log.MockLog
		addressRepoMock func() *mock_address.MockAddress
		error           error
		ctx             context.Context
		returnedAddress *entity.Address
	}{
		{
			name: "success",
			loggerMock: func() *mock_log.MockLog {
				loggerInfra := mock_log.NewMockLog(ctrl)
				return loggerInfra
			},
			addressRepoMock: func() *mock_address.MockAddress {
				repoLogMock := mock_address.NewMockAddress(ctrl)
				repoLogMock.EXPECT().Detail(gomock.Any(), gomock.Any()).Return(address, nil)
				return repoLogMock
			},
			id:              1,
			error:           nil,
			ctx:             context.Background(),
			returnedAddress: address,
		},
		{
			name: "RepoError",
			loggerMock: func() *mock_log.MockLog {
				loggerInfra := mock_log.NewMockLog(ctrl)
				loggerInfra.EXPECT().Error(err).Return()
				return loggerInfra
			},
			addressRepoMock: func() *mock_address.MockAddress {
				repoLogMock := mock_address.NewMockAddress(ctrl)
				repoLogMock.EXPECT().Detail(gomock.Any(), gomock.Any()).Return(nil, err)
				return repoLogMock
			},
			id:              1,
			error:           err,
			ctx:             context.Background(),
			returnedAddress: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			logRepoMock := test.addressRepoMock()
			loggerMock := test.loggerMock()
			service := NewService(loggerMock, logRepoMock)
			resAddress, err := service.Detail(test.ctx, test.id)
			if !errors.Is(err, test.error) {
				t.Error("error is not equal")
			}

			if !gomock.Eq(resAddress).Matches(test.returnedAddress) {
				t.Error("returned address_user is not right")
			}
			loggerMock.EXPECT()
			logRepoMock.EXPECT()
		})
	}
}

func BenchmarkService_Detail(b *testing.B) {
	ctrl := gomock.NewController(b)
	addressRepoMock := mock_address.NewMockAddress(ctrl)
	address := entity.NewAddress("c", "s", "co", "str", "3tgdsgds")
	addressRepoMock.EXPECT().Detail(gomock.Any(), int64(1)).Return(address, nil)
	loggerMock := mock_log.NewMockLog(ctrl)
	b.ResetTimer()
	service := NewService(loggerMock, addressRepoMock)

	service.Detail(context.Background(), int64(1))
	if b.Elapsed() > 150*time.Microsecond {
		b.Error("address_user service-detail takes too long to run")
	}
	loggerMock.EXPECT()
	addressRepoMock.EXPECT()
}

func TestService_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})
	err := errors.New("error")
	addresss := []*entity.Address{
		entity.NewAddress("c1", "s", "co", "str", "3tgdsgds"),
		entity.NewAddress("c2", "s", "co", "str", "3tgdsgds"),
		entity.NewAddress("c3", "s", "co", "str", "3tgdsgds"),
	}

	var tests = []struct {
		name            string
		loggerMock      func() *mock_log.MockLog
		addressRepoMock func() *mock_address.MockAddress
		error           error
		ctx             context.Context
		addresss        []*entity.Address
	}{
		{
			name: "success",
			loggerMock: func() *mock_log.MockLog {
				loggerInfra := mock_log.NewMockLog(ctrl)
				return loggerInfra
			},
			addressRepoMock: func() *mock_address.MockAddress {
				repoLogMock := mock_address.NewMockAddress(ctrl)
				repoLogMock.EXPECT().List(gomock.Any(), gomock.Any()).Return(addresss, nil)
				return repoLogMock
			},
			error:    nil,
			ctx:      context.Background(),
			addresss: addresss,
		},
		{
			name: "RepoError",
			loggerMock: func() *mock_log.MockLog {
				loggerInfra := mock_log.NewMockLog(ctrl)
				loggerInfra.EXPECT().Error(err).Return()
				return loggerInfra
			},
			addressRepoMock: func() *mock_address.MockAddress {
				repoLogMock := mock_address.NewMockAddress(ctrl)
				repoLogMock.EXPECT().List(gomock.Any(), gomock.Any()).Return(nil, err)
				return repoLogMock
			},
			error:    err,
			ctx:      context.Background(),
			addresss: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			logRepoMock := test.addressRepoMock()
			loggerMock := test.loggerMock()
			service := NewService(loggerMock, logRepoMock)
			resAddress, err := service.List(test.ctx, 1)
			if !errors.Is(err, test.error) {
				t.Error("error is not equal")
			}

			if !gomock.Eq(resAddress).Matches(test.addresss) {
				t.Error("returned addresss is not the same")
			}
			loggerMock.EXPECT()
			logRepoMock.EXPECT()
		})
	}
}

func BenchmarkService_List(b *testing.B) {
	ctrl := gomock.NewController(b)
	addressRepoMock := mock_address.NewMockAddress(ctrl)
	addresss := []*entity.Address{
		entity.NewAddress("c1", "s", "co", "str", "3tgdsgds"),
		entity.NewAddress("c2", "s", "co", "str", "3tgdsgds"),
		entity.NewAddress("c3", "s", "co", "str", "3tgdsgds"),
	}
	addressRepoMock.EXPECT().List(gomock.Any(), uint16(1)).Return(addresss, nil)
	loggerMock := mock_log.NewMockLog(ctrl)
	b.ResetTimer()
	service := NewService(loggerMock, addressRepoMock)
	service.List(context.Background(), uint16(1))
	if b.Elapsed() > 150*time.Microsecond {
		b.Error("address_user service-detail takes too long to run")
	}
	loggerMock.EXPECT()
	addressRepoMock.EXPECT()
}
