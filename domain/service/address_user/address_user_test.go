package address_user

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"m1-article-service/domain/entity"
	mock_log "m1-article-service/mock/infrastructure"
	repository "m1-article-service/mock/repository"
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
		addresses       []*entity.Address
		loggerMock      func() *mock_log.MockLog
		addressRepoMock func() *repository.MockAddress
		userRepoMock    func() *repository.MockUser
		user            *entity.User
		ctx             context.Context
	}{
		{
			name: "success",
			loggerMock: func() *mock_log.MockLog {
				loggerInfra := mock_log.NewMockLog(ctrl)
				return loggerInfra
			},
			addressRepoMock: func() *repository.MockAddress {
				addrRepoMock := repository.NewMockAddress(ctrl)
				addrRepoMock.EXPECT().BatchCreate(gomock.Any(), gomock.Any()).Return(nil)
				return addrRepoMock
			},
			userRepoMock: func() *repository.MockUser {
				userRepoMock := repository.NewMockUser(ctrl)
				userRepoMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(int64(1), nil)
				return userRepoMock
			},
			addresses: []*entity.Address{entity.NewAddress("c", "s", "co", "str", "3tgdsgds")},
			user:      &entity.User{Name: "Dgsgds", Lastname: "sfafsasf"},
			ctx:       context.Background(),
		},
		{
			name: "AddrRepoError",
			loggerMock: func() *mock_log.MockLog {
				loggerInfra := mock_log.NewMockLog(ctrl)
				loggerInfra.EXPECT().Error(err).MinTimes(2).Return()
				return loggerInfra
			},
			addressRepoMock: func() *repository.MockAddress {
				addrRepoMock := repository.NewMockAddress(ctrl)
				addrRepoMock.EXPECT().BatchCreate(gomock.Any(), gomock.Any()).MinTimes(2).Return(err)
				return addrRepoMock
			},
			userRepoMock: func() *repository.MockUser {
				userRepoMock := repository.NewMockUser(ctrl)
				userRepoMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(int64(1), nil)
				return userRepoMock
			},
			addresses: []*entity.Address{entity.NewAddress("c", "s", "co", "str", "3tgdsgds")},
			ctx:       context.Background(),
		},
		{
			name: "UserRepoError",
			loggerMock: func() *mock_log.MockLog {
				loggerInfra := mock_log.NewMockLog(ctrl)
				loggerInfra.EXPECT().Error(err).MinTimes(2).Return()
				return loggerInfra
			},
			addressRepoMock: func() *repository.MockAddress {
				addrRepoMock := repository.NewMockAddress(ctrl)
				return addrRepoMock
			},
			userRepoMock: func() *repository.MockUser {
				userRepoMock := repository.NewMockUser(ctrl)
				userRepoMock.EXPECT().Create(gomock.Any(), gomock.Any()).MinTimes(2).Return(nil)
				return userRepoMock
			},
			addresses: []*entity.Address{entity.NewAddress("c", "s", "co", "str", "3tgdsgds")},
			ctx:       context.Background(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			addressRepoMock := test.addressRepoMock()
			userRepoMock := test.userRepoMock()
			loggerMock := test.loggerMock()
			service := NewService(loggerMock, addressRepoMock, userRepoMock)
			service.Create(test.addresses, test.user)

			time.Sleep(500 * time.Millisecond)
			loggerMock.EXPECT()
			addressRepoMock.EXPECT()
			userRepoMock.EXPECT()
		})
	}
}

func BenchmarkService_Create(b *testing.B) {
	ctrl := gomock.NewController(b)
	addrRepoMock := repository.NewMockAddress(ctrl)
	addrRepoMock.EXPECT().BatchCreate(gomock.Any(), gomock.Any()).Return(nil)
	userRepoMock := repository.NewMockUser(ctrl)
	userRepoMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(int64(1), nil)
	loggerMock := mock_log.NewMockLog(ctrl)
	b.ResetTimer()
	service := NewService(loggerMock, addrRepoMock, userRepoMock)
	service.Create([]*entity.Address{entity.NewAddress("c", "s", "co",
		"str", "3tgdsgds")}, &entity.User{Name: "Fsfa", Lastname: "fsafsa"})
	fmt.Println(b.Elapsed())
	//wait for queue and goroutines
	for {
		if len(service.queue) == 0 {
			time.Sleep(200 * time.Microsecond)
			break
		}
	}
	if b.Elapsed() > 350*time.Millisecond {
		b.Error("address_user service-createBatchAddresses takes too long to run")
	}
	loggerMock.EXPECT()
	addrRepoMock.EXPECT()
	userRepoMock.EXPECT()
}

//
//func TestService_Detail(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	t.Cleanup(func() {
//		ctrl.Finish()
//	})
//	err := errors.New("error")
//	address := entity.NewAddress("c", "s", "co", "str", "3tgdsgds")
//
//	var tests = []struct {
//		name            string
//		id              int64
//		loggerMock      func() *mock_log.MockLog
//		addressRepoMock func() *repository.MockAddress
//		error           error
//		ctx             context.Context
//		returnedAddress *entity.Address
//	}{
//		{
//			name: "success",
//			loggerMock: func() *mock_log.MockLog {
//				loggerInfra := mock_log.NewMockLog(ctrl)
//				return loggerInfra
//			},
//			addressRepoMock: func() *repository.MockAddress {
//				repoLogMock := repository.NewMockAddress(ctrl)
//				repoLogMock.EXPECT().Detail(gomock.Any(), gomock.Any()).Return(address, nil)
//				return repoLogMock
//			},
//			id:              1,
//			error:           nil,
//			ctx:             context.Background(),
//			returnedAddress: address,
//		},
//		{
//			name: "RepoError",
//			loggerMock: func() *mock_log.MockLog {
//				loggerInfra := mock_log.NewMockLog(ctrl)
//				loggerInfra.EXPECT().Error(err).Return()
//				return loggerInfra
//			},
//			addressRepoMock: func() *repository.MockAddress {
//				repoLogMock := repository.NewMockAddress(ctrl)
//				repoLogMock.EXPECT().Detail(gomock.Any(), gomock.Any()).Return(nil, err)
//				return repoLogMock
//			},
//			id:              1,
//			error:           err,
//			ctx:             context.Background(),
//			returnedAddress: nil,
//		},
//	}
//
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			logRepoMock := test.addressRepoMock()
//			loggerMock := test.loggerMock()
//			service := NewService(loggerMock, logRepoMock)
//			resAddress, err := service.Detail(test.ctx, test.id)
//			if !errors.Is(err, test.error) {
//				t.Error("error is not equal")
//			}
//
//			if !gomock.Eq(resAddress).Matches(test.returnedAddress) {
//				t.Error("returned address_user is not right")
//			}
//			loggerMock.EXPECT()
//			logRepoMock.EXPECT()
//		})
//	}
//}
//
//func BenchmarkService_Detail(b *testing.B) {
//	ctrl := gomock.NewController(b)
//	addressRepoMock := repository.NewMockAddress(ctrl)
//	address := entity.NewAddress("c", "s", "co", "str", "3tgdsgds")
//	addressRepoMock.EXPECT().Detail(gomock.Any(), int64(1)).Return(address, nil)
//	loggerMock := mock_log.NewMockLog(ctrl)
//	b.ResetTimer()
//	service := NewService(loggerMock, addressRepoMock)
//
//	service.Detail(context.Background(), int64(1))
//	if b.Elapsed() > 150*time.Microsecond {
//		b.Error("address_user service-detail takes too long to run")
//	}
//	loggerMock.EXPECT()
//	addressRepoMock.EXPECT()
//}
