package navcoind

import (
	//. "github.com/NavExplorer/go-navcoind"
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Helpers", func() {
	Describe("Parse errors", func() {
		Context("No error", func() {
			response := rpcResponse{
				Id:     1212,
				Result: []byte("{}"),
				Err:    nil,
			}
			It("should return nil", func() {
				Expect(handleError(nil, &response)).To(BeNil())
			})
		})

		// err
		Context("errPrev is not nil", func() {
			errPrev := errors.New("fake error")
			response := rpcResponse{
				Id:     1212,
				Result: []byte("{}"),
				Err:    nil,
			}
			It("should occur", func() {
				err := handleError(errPrev, &response)
				Expect(Ω(err).Should(HaveOccurred()))
			})

		})

		// RPC error is not null
		Context("RPC error", func() {
			rpcError := make(map[string]interface{})
			rpcError["code"] = 6.32
			rpcError["message"] = "Fake error"
			response := rpcResponse{
				Id:     1212,
				Result: []byte("{}"),
				Err:    rpcError,
			}
			It("should occur", func() {
				err := handleError(nil, &response)
				Expect(Ω(err).Should(HaveOccurred()))
			})

		})

	})
})
