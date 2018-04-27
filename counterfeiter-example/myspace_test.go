package counterfeiter_example

import (
	"testing"

	counterfeiterexamplefakes "./counterfeiterexamplefakes"
)

func TestTimeConsuming(t *testing.T) {
	var fake = new(counterfeiterexamplefakes.FakeMySpecialInterface)

	Expect(fake.DoThingsCallCount()).To(Equal(1))

	str, num := fake.DoThingsArgsForCall(0)
	Expect(str).To(Equal("stuff"))
	Expect(num).To(Equal(uint64(5)))
}
