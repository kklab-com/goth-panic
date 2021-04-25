package kkpanic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTry(t *testing.T) {
	var val = "1"
	Try(func() {
		val = "2"
		panic("!!!")
	}).Catch(func(caught Caught) {
		assert.EqualValues(t, "!!!", caught.Data())
	}).Finally(func() {
		assert.EqualValues(t, "2", val)
	})

	Try(func() {
		val = "3"
	}).Catch(func(caught Caught) {
		assert.Fail(t, "dont")
	}).Finally(func() {
		assert.EqualValues(t, "3", val)
	})
}
