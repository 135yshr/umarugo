package lib

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestConerter(t *testing.T) {
	var (
		sut      func([]byte) (*Command, int, error)
		cmd      *Command
		seek     int
		expected []byte
		err      error
	)
	Describe(t, "コンバーターのインスタンスを作成する", func() {
		Context("スペースを渡したとき", func() {
			Before(func() {
				sut, err = createFunction(byte('U'))
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("インスタンスが返ってくること", func() {
				Expect(sut).To(Exist)
			})
		})
		Context("タブを渡したとき", func() {
			Before(func() {
				sut, err = createFunction(byte('M'))
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("インスタンスが返ってくること", func() {
				Expect(sut).To(Exist)
			})
		})
		Context("改行を渡したとき", func() {
			Before(func() {
				sut, err = createFunction(byte('R'))
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("インスタンスが返ってくること", func() {
				Expect(sut).To(Exist)
			})
		})
		Context("スペース・タブ・改行以外の文字が渡されたとき", func() {
			Before(func() {
				sut, err = createFunction(byte('A'))
			})
			It("インスタンスがnilなること", func() {
				Expect(sut).To(NotExist)
			})
			It("エラーが発生すること", func() {
				Expect(err).To(Exist)
			})
		})
	})
	Describe(t, "スタックを操作する", func() {
		Context("SSTLを渡したとき", func() {
			Before(func() {
				expected = []byte{'U', 'U', 'M', 'R'}
				cmd, seek, err = stackManipulation(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("stack push 1 コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newSubCommandWithParam("stack", "push", 1))
			})
		})
		Context("SSTSLを渡したとき", func() {
			Before(func() {
				expected = []byte{'U', 'U', 'M', 'U', 'R'}
				cmd, seek, err = stackManipulation(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("stack push 2 コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newSubCommandWithParam("stack", "push", 2))
			})
		})
		Context("SSTSSLを渡したとき", func() {
			Before(func() {
				expected = []byte{'U', 'U', 'M', 'U', 'U', 'R'}
				cmd, seek, err = stackManipulation(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("stack push 4 コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newSubCommandWithParam("stack", "push", 4))
			})
		})
		Context("STTLを渡したとき", func() {
			Before(func() {
				expected = []byte{'U', 'M', 'M', 'R'}
				cmd, seek, err = stackManipulation(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("stack push -1 コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newSubCommandWithParam("stack", "push", -1))
			})
		})
		Context("LSを渡したとき", func() {
			Before(func() {
				expected = []byte{'R', 'U'}
				cmd, seek, err = stackManipulation(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("stack copy コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newSubCommand("stack", "copy"))
			})
		})
		Context("LTを渡したとき", func() {
			Before(func() {
				expected = []byte{'R', 'M'}
				cmd, seek, err = stackManipulation(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("stack swap コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newSubCommand("stack", "swap"))
			})
		})
		Context("LLを渡したとき", func() {
			Before(func() {
				expected = []byte{'R', 'R'}
				cmd, seek, err = stackManipulation(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("stack remove コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newSubCommand("stack", "remove"))
			})
		})
		Context("TSSTLを渡したとき", func() {
			Before(func() {
				expected = []byte{'M', 'U', 'U', 'M', 'R'}
				cmd, seek, err = stackManipulation(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("stack ncopy 1 コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newSubCommandWithParam("stack", "ncopy", 1))
			})
		})
		Context("TSSTLを渡したとき", func() {
			Before(func() {
				expected = []byte{'M', 'R', 'U', 'M', 'R'}
				cmd, seek, err = stackManipulation(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("stack move 1 コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newSubCommandWithParam("stack", "move", 1))
			})
		})
		Context("TTを渡したとき", func() {
			Before(func() {
				expected = []byte{'M', 'M'}
				cmd, seek, err = stackManipulation(expected)
			})
			It("コマンドが存在しないこと", func() {
				Expect(cmd).To(NotExist)
			})
			It("エラーが発生すること", func() {
				Expect(err).To(Exist)
			})
			It("読み込んだ文字数が0であること", func() {
				Expect(seek).To(Equal, 0)
			})
		})
	})
	Describe(t, "制御に関する命令", func() {
		Context("SSTSSTLを渡したとき", func() {
			Before(func() {
				expected = []byte{'U', 'U', 'M', 'U', 'U', 'M', 'R'}
				cmd, seek, err = flowControl(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("label 1001 コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newSubCommand("label", "1001"))
			})
		})
		Context("STTSSTLを渡したとき", func() {
			Before(func() {
				expected = []byte{'U', 'M', 'M', 'U', 'U', 'M', 'R'}
				cmd, seek, err = flowControl(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("call 1001 コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newSubCommand("call", "1001"))
			})
		})
		Context("SLTSSTLを渡したとき", func() {
			Before(func() {
				expected = []byte{'U', 'R', 'M', 'U', 'U', 'M', 'R'}
				cmd, seek, err = flowControl(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("goto 1001 コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newSubCommand("goto", "1001"))
			})
		})
		Context("TSTSSTLを渡したとき", func() {
			Before(func() {
				expected = []byte{'M', 'U', 'M', 'U', 'U', 'M', 'R'}
				cmd, seek, err = flowControl(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("if 0 goto 1001 コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newSubCommand("if stack==0 then goto", "1001"))
			})
		})
		Context("TTTSSTLを渡したとき", func() {
			Before(func() {
				expected = []byte{'M', 'M', 'M', 'U', 'U', 'M', 'R'}
				cmd, seek, err = flowControl(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("if <0 goto 1001 コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newSubCommand("if stack<0 then goto", "1001"))
			})
		})
		Context("TLを渡したとき", func() {
			Before(func() {
				expected = []byte{'M', 'R'}
				cmd, seek, err = flowControl(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("return コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newCommand("return"))
			})
		})
		Context("LLを渡したとき", func() {
			Before(func() {
				expected = []byte{'R', 'R'}
				cmd, seek, err = flowControl(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("exit コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newCommand("exit"))
			})
		})
		Context("TTを渡したとき", func() {
			Before(func() {
				expected = []byte{'R', 'M'}
				cmd, seek, err = flowControl(expected)
			})
			It("コマンドが存在しないこと", func() {
				Expect(cmd).To(NotExist)
			})
			It("エラーが発生すること", func() {
				Expect(err).To(Exist)
			})
			It("読み込んだ文字数が0であること", func() {
				Expect(seek).To(Equal, 0)
			})
		})
	})
	Describe(t, "演算命令の関数", func() {
		Context("SSを渡したとき", func() {
			Before(func() {
				expected = []byte{'U', 'U'}
				cmd, seek, err = arithmetic(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("add コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newCommand("add"))
			})
		})
		Context("STを渡したとき", func() {
			Before(func() {
				expected = []byte{'U', 'M'}
				cmd, seek, err = arithmetic(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("sub コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newCommand("sub"))
			})
		})
		Context("SLを渡したとき", func() {
			Before(func() {
				expected = []byte{'U', 'R'}
				cmd, seek, err = arithmetic(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("mul コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newCommand("mul"))
			})
		})
		Context("TSを渡したとき", func() {
			Before(func() {
				expected = []byte{'M', 'U'}
				cmd, seek, err = arithmetic(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("div コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newCommand("div"))
			})
		})
		Context("TTを渡したとき", func() {
			Before(func() {
				expected = []byte{'M', 'M'}
				cmd, seek, err = arithmetic(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("mod コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newCommand("mod"))
			})
		})
		Context("TLを渡したとき", func() {
			Before(func() {
				expected = []byte{'M', 'R'}
				cmd, seek, err = arithmetic(expected)
			})
			It("コマンドが存在しないこと", func() {
				Expect(cmd).To(NotExist)
			})
			It("エラーが発生すること", func() {
				Expect(err).To(Exist)
			})
			It("読み込んだ文字数が0であること", func() {
				Expect(seek).To(Equal, 0)
			})
		})
	})
	Describe(t, "ヒープ領域を操作する", func() {
		Context("Sを渡したとき", func() {
			Before(func() {
				expected = []byte{'U'}
				cmd, seek, err = heapAccess(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("heap push コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newSubCommand("heap", "push"))
			})
		})
		Context("Tを渡したとき", func() {
			Before(func() {
				expected = []byte{'M'}
				cmd, seek, err = heapAccess(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("heap pop コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newSubCommand("heap", "pop"))
			})
		})
		Context("Lを渡したとき", func() {
			Before(func() {
				expected = []byte{'R'}
				cmd, seek, err = heapAccess(expected)
			})
			It("コマンドが存在しないこと", func() {
				Expect(cmd).To(NotExist)
			})
			It("エラーが発生すること", func() {
				Expect(err).To(Exist)
			})
			It("読み込んだ文字数が0であること", func() {
				Expect(seek).To(Equal, 0)
			})
		})
	})
	Describe(t, "I/O操作の関数", func() {
		Context("SSを渡したとき", func() {
			Before(func() {
				expected = []byte{'U', 'U'}
				cmd, seek, err = i_o(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("putc コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newCommand("putc"))
			})
		})
		Context("STを渡したとき", func() {
			Before(func() {
				expected = []byte{'U', 'M'}
				cmd, seek, err = i_o(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("putn コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newCommand("putn"))
			})
		})
		Context("TSを渡したとき", func() {
			Before(func() {
				expected = []byte{'M', 'U'}
				cmd, seek, err = i_o(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("gutc コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newCommand("getc"))
			})
		})
		Context("TTを渡したとき", func() {
			Before(func() {
				expected = []byte{'M', 'M'}
				cmd, seek, err = i_o(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("gutn コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newCommand("getn"))
			})
		})
		Context("TLを渡したとき", func() {
			Before(func() {
				expected = []byte{'M', 'R'}
				cmd, seek, err = i_o(expected)
			})
			It("コマンドが存在しないこと", func() {
				Expect(cmd).To(NotExist)
			})
			It("エラーが発生すること", func() {
				Expect(err).To(Exist)
			})
			It("読み込んだ文字数が0であること", func() {
				Expect(seek).To(Equal, 0)
			})
		})
	})
}
