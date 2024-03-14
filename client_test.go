package main

import (
	"context"
	"errors"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"

	"fast-client/mocks"
	service2 "fast-client/service"
)

var _ = Describe("Client", func() {
	var (
		service *mocks.Service

		ctx     context.Context
		errTest error
	)

	BeforeEach(func() {
		service = mocks.NewService(GinkgoT())

		ctx = context.Background()
		errTest = errors.New("test error")
	})

	Describe("Send", func() {
		It("should return no error", func() {
			data := []service2.Item{{}, {}, {}}

			service.EXPECT().GetLimits().Return(10, 1*time.Second)
			service.EXPECT().Process(ctx, mock.Anything).Return(nil)

			client := NewClient(service)
			err := client.Send(ctx, data)

			Expect(err).To(BeNil())
		})
		It("should return error", func() {
			data := []service2.Item{{}, {}, {}}

			service.EXPECT().GetLimits().Return(10, 1*time.Second)
			service.EXPECT().Process(ctx, mock.Anything).Return(errTest)

			client := NewClient(service)
			err := client.Send(ctx, data)

			Expect(err).To(Equal(errTest))
		})
	})
})
