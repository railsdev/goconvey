package something

import (
	. "github.com/mdwhatcott/goconvey"
)

func Test(t *testing.T) {
	Convey("Subject: Integer incrementation and decrementation", t, func() {
		x := 0

		Convey("Given a starting integer value", func() {
			x = 42

			Convey("When incremented", func() {
				x++

				Convey("The value should be greater by one", func() {
					So(x, ShouldEqual, 43)
				})
				Convey("The value should NOT be what it used to be", func() {
					So(x, ShouldNotEqual, 42)
				})
			})
			Convey("When decremented", func() {
				x--

				Convey("The value should be lesser by one", func() {
					So(x, ShouldEqual, 41)
				})
				Convey("The value should NOT be what it used to be", func() {
					So(x, ShouldNotEqual, 42)
				})
			})
			Cleanup(func() {
				x = 0
			})
		})
	})
}

/*
Output:

????
*/
