package test

import (
	"example/entity"
	"testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestWeight(t *testing.T){
	g:= NewGomegaWithT(t)

	t.Run(`weight is required` ,func(t *testing.T){
		takeahistory := entity.TakeAHistory{

		Email: "gmail@gmail.com",
		PhoneNumber: "1234567890",
		}
		ok, err := govalidator.ValidateStruct(takeahistory)
	
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Weight is required"))
	})
}

func TestAll(t *testing.T){
	g:=NewGomegaWithT(t)

	t.Run(`All correct` ,func(t *testing.T){
		takeahistory := entity.TakeAHistory{
			Weight: 50,
			Email: "gmail@gmail.com",
			PhoneNumber: "1234567890",
		}
		ok, err := govalidator.ValidateStruct(takeahistory)
		g.Expect(err).To(BeNil())
		g.Expect(ok).To(BeTrue())
	})
}

func TestEmail(t *testing.T){
	g:=NewGomegaWithT(t)

	t.Run(`Email is not correct` ,func(t *testing.T){
		takeahistory := entity.TakeAHistory{
			Weight: 50.0,
			Email: "addfg",
			PhoneNumber: "1234567890",
		}
		ok ,err := govalidator.ValidateStruct(takeahistory)

		g.Expect(err).NotTo(BeNil())
		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err.Error()).To(Equal("Email is correct"))
	})
}

func TestEmailis(t *testing.T){
	g:=NewGomegaWithT(t)

	t.Run(`Email is required` ,func(t *testing.T){
		takeahistory := entity.TakeAHistory{
			Weight: 5.55,
			PhoneNumber: "1233456890",
		}
		ok, err := govalidator.ValidateStruct(takeahistory)

		g.Expect(err).NotTo(BeNil())
		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err.Error()).To(Equal("Email is required"))
	})
}

func TestPhoneNumber(t *testing.T){
	g:=NewGomegaWithT(t)

	t.Run(`PhoneNumber is not correct` ,func(t *testing.T){
		takeahistory:= entity.TakeAHistory{
			Weight: 5.00,
			Email: "g@gmail.com",
			PhoneNumber: "ff555555",
		}

		ok, err := govalidator.ValidateStruct(takeahistory)

		g.Expect(err).NotTo(BeNil())
		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err.Error()).To(Equal("PhoneNumber is correct"))
	})
}

func TestPhone(t *testing.T){
	g:=NewGomegaWithT(t)

	t.Run(`PhoneNumber is required` ,func(t *testing.T){
		takeahistory:=entity.TakeAHistory{
			Weight: 5.2,
			Email: "g@gamil.com",
		}

		ok, err := govalidator.ValidateStruct(takeahistory)

		g.Expect(err).NotTo(BeNil())
		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err.Error()).To(Equal("PhoneNumber is required"))
	})
}
