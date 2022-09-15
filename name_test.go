package Ginkgo_test

import (
	"Ginkgo_test/names"
	"Ginkgo_test/requests"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Verify data for names", func() {
	It("Verify data for pepe", func() {
		a := names.AgeData{
			Name: "pepe",
    		Age: 51,
    		Count: 48187,
		}
		result, body := requests.GetResult("", "?name=pepe", "GET", "") //IF no body needed can replace with _
		Expect(result.StatusCode).To(Equal(200))
		age := names.GetAgeData(body)
		Expect(age).To(Equal(a))
	})
	It("Verify data for Ramón", func() {
		a := names.AgeData{
			Name: "ramon",
    		Age: 59,
    		Count: 26378,
		}
		result, body := requests.GetResult("", "?name=ramon", "GET", "") //Tilde...
		Expect(result.StatusCode).To(Equal(200))
		age := names.GetAgeData(body)
		Expect(age).To(Equal(a))
	})
	It("Verify data for Catalina", func() {
		a := names.AgeData{
			Name: "catalina",
    		Age: 32,
    		Count: 7830,
		}
		result, body := requests.GetResult("", "?name=catalina", "GET", "") 
		Expect(result.StatusCode).To(Equal(200))
		age := names.GetAgeData(body)
		Expect(age).To(Equal(a))
	})
	It("Verify data for No name", func() {
		result, _ := requests.GetResult("", "?name=", "GET", "")
		Expect(result.StatusCode).To(Equal(422)) //Weird but good answer.
	})
	It("Verify data an incorrect name", func() {
		a := names.AgeData{
			Name: "%$",
    		Age: 0,
    		Count: 0,
		}
		result, body := requests.GetResult("", "?name=%$", "GET", "") 
		Expect(result.StatusCode).To(Equal(200))
		age := names.GetAgeData(body)
		Expect(age).To(Equal(a))
	})
})