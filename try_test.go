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
		val = "5"
		panic(io.EOF)
	}).Catch(io.EOF, func(caught Caught) {
		assert.EqualValues(t, "5", val)
	}).Finally(func() {
		assert.EqualValues(t, "5", val)
	})
}
