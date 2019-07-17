package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Link Looker")
}

var _ = Describe("#getPage", func() {
	It("when given a valid url", func() {
		validURL := "https://www.google.com"
		_, err := getPage(validURL)
		Ω(err).Should(BeNil())
	})

	It("when given a bad URL", func() {
		validURL := "https://www.fdafsfdsafdsa.com"
		_, err := getPage(validURL)
		Ω(err).ShouldNot(BeNil())
	})

	It("when given an invalid URL", func() {
		validURL := "https://www.warrendouglas.com/jfdlasjskl"
		_, err := getPage(validURL)
		Ω(err).ShouldNot(BeNil())
	})
})

var _ = Describe("#getPageContent", func() {
	It("when given a valid page body", func() {
		validURL := "https://www.google.com"
		page, _ := getPage(validURL)
		_, err := getPageContent(page)
		Ω(err).Should(BeNil())
	})

	It("when given nil for a body", func() {
		_, err := getPageContent(nil)
		Ω(err).ShouldNot(BeNil())
	})
})

var _ = Describe("#getLinks", func() {
	It("when given a valid page body", func() {
		validURL := "https://www.warrendouglas.com"
		page, _ := getPage(validURL)
		doc, _ := getPageContent(page)
		links, err := getLinks(doc)
		Ω(err).Should(BeNil())
		expectedLinks := 18
		Ω(len(links)).Should(Equal(expectedLinks))
	})

	It("when given nil for a body", func() {
		_, err := getLinks(nil)
		Ω(err).ShouldNot(BeNil())
	})
})

var _ = Describe("#getPageLinks", func() {
	It("when given a valid url", func() {
		validURL := "https://www.warrendouglas.com"
		links, err := getPageLinks(validURL)
		expectedLinks := 18
		Ω(len(links)).Should(Equal(expectedLinks))
		Ω(err).Should(BeNil())
	})

	It("when given a valid url", func() {
		validURL := "https://www.fdafdsf.com"
		links, err := getPageLinks(validURL)
		Ω(links).Should(BeNil())
		Ω(err).ShouldNot(BeNil())
	})
})
