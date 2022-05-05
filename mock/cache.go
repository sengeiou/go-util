package mock

import (
	"github.com/golang/mock/gomock"
	iface "red-packet/util/interface"
	"testing"
)

func NewCache(t *testing.T) iface.ICache {
	m := gomock.NewController(t)
	mock := NewMockICache(m)

	mock.EXPECT().Get(gomock.Any(), gomock.Any()).AnyTimes().Return("", nil)
	mock.EXPECT().SetNX(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(true, nil)
	mock.EXPECT().SetEX(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().LPush(gomock.Any(), gomock.Any()).AnyTimes().Return(1, nil)
	mock.EXPECT().RPop(gomock.Any(), gomock.Any()).AnyTimes().Return("", nil)
	mock.EXPECT().Expire(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil)

	return mock
}
