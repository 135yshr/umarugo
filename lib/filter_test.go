package lib

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestFilter(t *testing.T) {
	Describe(t, "初期化", func() {
		Context("インスタンスを作成ししたとき", func() {
			It("インスタンスを取得できること", func() {
				sut := NewFilter([]string{"食う"})
				Expect(sut).To(Exist)
			})
		})
	})
	Describe(t, "フィルタリングする", func() {
		var (
			sut *Filter
		)
		Before(func() {
			sut = NewFilter([]string{"食う", "寝る", "遊ぶ"})
		})
		Context("帰ったらおやつを食うを渡ししたとき", func() {
			It("食うが抽出できる", func() {
				actual := sut.Filter("帰ったらおやつを食う")
				Expect(actual).To(Equal, map[int]string{24: "食う"})
			})
		})
		Context("帰ったらおやつを食うそのあとカレーを食うを渡ししたとき", func() {
			It("食うが２つ抽出できる", func() {
				actual := sut.Filter("帰ったらおやつを食うそのあとカレーを食う")
				Expect(actual).To(Equal, map[int]string{24: "食う", 54: "食う"})
			})
		})
		Context("帰ったらおやつを食うを食べ終わったらゲームで遊ぶを渡ししたとき", func() {
			It("食うと遊ぶが抽出できる", func() {
				actual := sut.Filter("帰ったらおやつを食うを食べ終わったらゲームで遊ぶ")
				Expect(actual).To(Equal, map[int]string{24: "食う", 66: "遊ぶ"})
			})
		})
		Context("帰ったらおやつを食うを食べ終わったらゲームで遊ぶ。遊び終わったら寝るを渡ししたとき", func() {
			It("食う寝る遊ぶが抽出できる", func() {
				actual := sut.Filter("帰ったらおやつを食うを食べ終わったらゲームで遊ぶ。遊び終わったら寝る")
				Expect(actual).To(Equal, map[int]string{24: "食う", 66: "遊ぶ", 96: "寝る"})
			})
		})
	})
}
