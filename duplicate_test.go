package fm_test

import (
	"bytes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"fm"
)

var _ = Describe("RunWithFile", func() {
	var (
		duplicate fm.DuplicateCmd

		err error

		in    *bytes.Buffer
		count int

		res string
	)
	JustBeforeEach(func() {
		err = duplicate.RunWithFile(in, count)
	})
	Context("inが空", func() {
		BeforeEach(func() {
			in = bytes.NewBufferString("")
			res = ""
			count = 2
		})
		It("空", func() {
			Expect(in.String()).To(Equal(res))
		})
		It("errはnil", func() {
			Expect(err).To(BeNil())
		})
	})
	Context("countが0回", func() {
		BeforeEach(func() {
			in = bytes.NewBufferString("test")
			res = ""
			count = 0
		})
		It("in変わりなし", func() {
			Expect(in.String()).To(Equal(res))
		})
		It("errはnil", func() {
			Expect(err).To(BeNil())
		})
	})
	Context("複数回複製", func() {
		BeforeEach(func() {
			in = bytes.NewBufferString(`input1
input2`)
			res = `input1
input2`
			count = 1
		})
		It("1回複製される", func() {
			Expect(in.String()).To(Equal(res))
		})
		It("errはnil", func() {
			Expect(err).To(BeNil())
		})
	})
	Context("日本語複数回複製", func() {
		BeforeEach(func() {
			in = bytes.NewBufferString(`ライン1
ライン2`)
			res = `ライン1
ライン2ライン1
ライン2`
			count = 2
		})
		It("2回複製される", func() {
			Expect(in.String()).To(Equal(res))
		})
		It("errはnil", func() {
			Expect(err).To(BeNil())
		})
	})
})
