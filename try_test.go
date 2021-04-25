package kkpanic

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTry(t *testing.T) {
	var val = "1"
	Try(func() {
		val = "2"
		panic("!!!")
	}).Catch("!!!", func(caught Caught) {
		assert.EqualValues(t, "!!!", caught.Data())
	}).Finally(func() {
		assert.EqualValues(t, "2", val)
	})

	Try(func() {
		val = "3"
		panic(3)
	}).CatchAll(func(caught Caught) {
		assert.EqualValues(t, 3, caught.Data())
	}).Finally(func() {
		assert.EqualValues(t, "3", val)
	})

	Try(func() {
		val = "4"
	}).Catch("!!!", func(caught Caught) {
		assert.Fail(t, "dont")
	}).Finally(func() {
		assert.EqualValues(t, "4", val)
	})

	Try(func() {
		Try(func() {
			val = "5"
			panic(io.EOF)
		}).Catch(io.EOF, func(caught Caught) {
			assert.EqualValues(t, "5", val)
			val = "s"
		}).Finally(func() {
			assert.EqualValues(t, "s", val)
			val = "6"
			panic("!")
		}, func() {
			assert.EqualValues(t, "6", val)
			val = "7"
			panic("!")
		}, func() {
			assert.EqualValues(t, "7", val)
			val = "8"
		})
	}).Finally(func() {
		assert.EqualValues(t, "8", val)
		val = "9"
	})

	assert.EqualValues(t, "9", val)
}
