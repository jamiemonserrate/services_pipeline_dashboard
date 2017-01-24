package integration_test

import (
	"io/ioutil"
	"net/http"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Integration", func() {

	var session *gexec.Session
	var resp *http.Response

	BeforeEach(func() {
		command := exec.Command(pathToPipelineDashboardCLI)
		session, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).ToNot(HaveOccurred())
	})

	It("should return a valid response", func() {
		Eventually(func() int {
			resp, err = http.Get("http://localhost:8080/")
			if err != nil {
				return 0
			}

			return resp.StatusCode
		}).Should(Equal(200))

		responseBody, err := ioutil.ReadAll(resp.Body)
		Expect(err).ToNot(HaveOccurred())
		Expect(string(responseBody)).To(ContainSubstring("Jamie"))
	})

	AfterEach(func() {
		session.Kill()
	})

})
