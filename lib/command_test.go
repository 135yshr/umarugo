package lib

import (
	"fmt"
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestCommnad(t *testing.T) {
	var (
		cmd    string
		subcmd string
		param  int
		sut    *Command
	)
	Describe(t, "初期化", func() {
		Context("cmdを指定するとき", func() {
			Before(func() {
				cmd = "cmd"
				sut = newCommand(cmd)
			})
			It("インスタンスが作成されること", func() {
				Expect(sut).To(Exist)
			})
			It("コマンドがcmdになっていること", func() {
				Expect(sut.cmd).To(Equal, cmd)
			})
			It("サブコマンドが空文字になっていること", func() {
				Expect(sut.subcmd).To(Equal, "")
			})
			It("パラメータが0になっていること", func() {
				Expect(sut.param).To(Equal, 0)
			})
		})
		Context("cmd と subcmdを指定したとき", func() {
			Before(func() {
				cmd, subcmd = "cmd", "subcmd"
				sut = newSubCommand(cmd, subcmd)
			})
			It("インスタンスが作成されること", func() {
				Expect(sut).To(Exist)
			})
			It("コマンドがcmdになっていること", func() {
				Expect(sut.cmd).To(Equal, cmd)
			})
			It("サブコマンドが空文字になっていること", func() {
				Expect(sut.subcmd).To(Equal, subcmd)
			})
			It("パラメータが0になっていること", func() {
				Expect(sut.param).To(Equal, 0)
			})
		})
		Context("cmd と 1 を指定したとき", func() {
			Before(func() {
				cmd = "cmd"
				param = 1
				sut = newCommandWithParam(cmd, param)
			})
			It("インスタンスが作成されること", func() {
				Expect(sut).To(Exist)
			})
			It("コマンドがcmdになっていること", func() {
				Expect(sut.cmd).To(Equal, cmd)
			})
			It("サブコマンドが空文字になっていること", func() {
				Expect(sut.subcmd).To(Equal, "")
			})
			It("パラメータが1になっていること", func() {
				Expect(sut.param).To(Equal, 1)
			})
		})
		Context("cmd と subcmd と 1 を指定したとき", func() {
			Before(func() {
				cmd, subcmd = "cmd", "subcmd"
				param = 1
				sut = newSubCommandWithParam(cmd, subcmd, param)
			})
			It("インスタンスが作成されること", func() {
				Expect(sut).To(Exist)
			})
			It("コマンドがcmdになっていること", func() {
				Expect(sut.cmd).To(Equal, cmd)
			})
			It("サブコマンドが空文字になっていること", func() {
				Expect(sut.subcmd).To(Equal, subcmd)
			})
			It("パラメータが0になっていること", func() {
				Expect(sut.param).To(Equal, param)
			})
		})
	})
	Describe(t, "文字列に変換", func() {
		Context("cmd subcmd 1 で初期化されたとき", func() {
			Before(func() {
				cmd, subcmd = "cmd", "subcmd"
				param = 1
				sut = newSubCommandWithParam(cmd, subcmd, param)
			})
			It("cmd subcmd 1になっている", func() {
				Expect(fmt.Sprint(sut)).To(Equal, "cmd subcmd 1")
			})
		})
		Context("cmd subcmd 2 で初期化されたとき", func() {
			Before(func() {
				cmd, subcmd = "cmd", "subcmd"
				param = 2
				sut = newSubCommandWithParam(cmd, subcmd, param)
			})
			It("cmd subcmd 1になっている", func() {
				Expect(fmt.Sprint(sut)).To(Equal, "cmd subcmd 2")
			})
		})
	})
}

func TestCommandList(t *testing.T) {
	var (
		sut *CommandList
		cmd *Command
		key int
		err error
	)
	Describe(t, "初期化", func() {
		Context("インスタンスを作成するとき", func() {
			Before(func() {
				sut = newCommandList()
			})
			It("インスタンスが作成されていること", func() {
				Expect(sut).To(Exist)
			})
		})
	})
	Describe(t, "コマンドを追加する", func() {
		Context("コマンドを１つ追加したとき", func() {
			Before(func() {
				sut = newCommandList()
				sut.Add(newCommand("test"))
			})
			It("コマンドの数が１になっていること", func() {
				Expect(sut.Len()).To(Equal, 1)
			})
		})
		Context("コマンドを2つ追加したとき", func() {
			Before(func() {
				sut = newCommandList()
				sut.Add(newCommand("test"))
				sut.Add(newCommand("test2"))
			})
			It("コマンドの数が2になっていること", func() {
				Expect(sut.Len()).To(Equal, 2)
			})
		})
	})
	Describe(t, "クリアする", func() {
		Context("コマンドを２つ追加して削除したとき", func() {
			Before(func() {
				sut = newCommandList()
				sut.Add(newCommand("test"))
				sut.Add(newCommand("test2"))
				sut.Clear()
			})
			It("コマンドの数が0になること", func() {
				Expect(sut.Len()).To(Equal, 0)
			})
		})
		Context("コマンドを４つ追加して削除したとき", func() {
			Before(func() {
				sut = newCommandList()
				sut.Add(newCommand("test"))
				sut.Add(newCommand("test2"))
				sut.Add(newCommand("test3"))
				sut.Add(newCommand("test4"))
				sut.Clear()
			})
			It("コマンドの数が0になること", func() {
				Expect(sut.Len()).To(Equal, 0)
			})
		})
	})
	Describe(t, "取得する", func() {
		Before(func() {
			sut = newCommandList()
			sut.Add(newCommand("test"))
			sut.Add(newCommand("test2"))
		})
		Context("コマンドを２つ追加して引数に１を渡したとき", func() {
			Before(func() {
				cmd = sut.Get(1)
			})
			It("コマンドが取得できること", func() {
				Expect(cmd).To(Exist)
			})
			It("cmdがtestになっていること", func() {
				Expect(cmd.cmd).To(Equal, "test")
			})
		})
		Context("コマンドを２つ追加して引数に2を渡したとき", func() {
			Before(func() {
				cmd = sut.Get(2)
			})
			It("コマンドが取得できること", func() {
				Expect(cmd).To(Exist)
			})
			It("cmdがtestになっていること", func() {
				Expect(cmd.cmd).To(Equal, "test2")
			})
		})
		Context("コマンドを２つ追加して引数に3を渡したとき", func() {
			Before(func() {
				cmd = sut.Get(3)
			})
			It("コマンドが存在しないこと", func() {
				Expect(cmd).To(NotExist)
			})
		})
		Context("コマンドを２つ追加して引数に0を渡したとき", func() {
			Before(func() {
				cmd = sut.Get(0)
			})
			It("コマンドが存在しないこと", func() {
				Expect(cmd).To(NotExist)
			})
		})
		Context("コマンドを２つ追加して引数に-1を渡したとき", func() {
			Before(func() {
				cmd = sut.Get(-1)
			})
			It("コマンドが存在しないこと", func() {
				Expect(cmd).To(NotExist)
			})
		})
	})
	Describe(t, "検索する", func() {
		Before(func() {
			sut = newCommandList()
			sut.Add(newCommand("test"))
			sut.Add(newCommand("test2"))
			sut.Add(newCommand("test3"))
		})
		Context("検索ワードにtestを指定したとき", func() {
			Before(func() {
				key, err = sut.Search(newCommand("test"))
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("keyが2になること", func() {
				Expect(key).To(Equal, 1)
			})
		})
		Context("検索ワードにtest2を指定したとき", func() {
			Before(func() {
				key, err = sut.Search(newCommand("test2"))
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("keyが2になること", func() {
				Expect(key).To(Equal, 2)
			})
		})
		Context("検索ワードにtest3を指定したとき", func() {
			Before(func() {
				key, err = sut.Search(newCommand("test3"))
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("keyが2になること", func() {
				Expect(key).To(Equal, 3)
			})
		})
		Context("検索ワードにtest4を指定したとき", func() {
			Before(func() {
				key, err = sut.Search(newCommand("test4"))
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(Exist)
			})
			It("keyが2になること", func() {
				Expect(key).To(Equal, -1)
			})
		})
	})
}
