package fm_test

import (
	"bytes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"fm"
)

var _ = Describe("RunWithFile", func() {
	var (
		reverse fm.ReverseCmd

		err error

		in  *bytes.Buffer
		out *bytes.Buffer = &bytes.Buffer{}

		res string
	)
	JustBeforeEach(func() {
		err = reverse.RunWithFile(in, out)
	})
	Context("inが空", func() {
		BeforeEach(func() {
			in = bytes.NewBufferString("")
			out.Reset()
			res = ""
		})
		It("outは空", func() {
			Expect(out.String()).To(Equal(res))
		})
		It("errはnil", func() {
			Expect(err).To(BeNil())
		})
	})
	Context("inは改行なし文字列", func() {
		BeforeEach(func() {
			in = bytes.NewBufferString("test")
			out.Reset()
			res = "tset"
		})
		It("inを逆にした文字が返る", func() {
			Expect(out.String()).To(Equal(res))
		})
		It("errはnil", func() {
			Expect(err).To(BeNil())
		})
	})
	Context("改行を含む文字列", func() {
		BeforeEach(func() {
			in = bytes.NewBufferString(`line1
line2`)
			out.Reset()
			res = `2enil
1enil`
		})
		It("inを逆にした文字が返る", func() {
			Expect(out.String()).To(Equal(res))
		})
		It("errはnil", func() {
			Expect(err).To(BeNil())
		})
	})
	Context("日本語を含む", func() {
		BeforeEach(func() {
			in = bytes.NewBufferString("インプット")
			out.Reset()
			res = "トップンイ"
		})
		FIt("inを逆にした文字が返る", func() {
			Expect(out.String()).To(Equal(res))
		})
		It("errはnil", func() {
			Expect(err).To(BeNil())
		})
	})
})
