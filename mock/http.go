package mock

import (
	"github.com/golang/mock/gomock"
	iface "red-packet/util/interface"
	"testing"
)

func NewHttpClient(t *testing.T) iface.IHttpClient {
	m := gomock.NewController(t)
	mock := NewMockIHttpClient(m)

	mock.EXPECT().Send().AnyTimes().Return("", nil)
	mock.EXPECT().SetTimeout(gomock.Any()).AnyTimes().Return()
	mock.EXPECT().SetRequest(gomock.Any()).AnyTimes().Return()

	return mock
}
