package user

import (
	"ko44d/go-bdd-testing/howto-mockgen/book/mock"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	var (
		mockCtrl  *gomock.Controller
		mockBook  *mocks.MockBook
		testUser  User
		bookTitle string
	)

	BeforeEach(func() {
		// gomockコントローラーを初期化
		mockCtrl = gomock.NewController(GinkgoT())

		// Book インターフェースのモックを作成
		mockBook = mocks.NewMockBook(mockCtrl)

		// テスト用のタイトルを設定
		bookTitle = "テスト本のタイトル"

		// GetTitle()が呼ばれたらbookTitleを返すように設定
		mockBook.EXPECT().GetTitle().Return(bookTitle).AnyTimes()

		// テスト対象のユーザーを作成（モックを注入）
		testUser = NewUser(mockBook)
	})

	AfterEach(func() {
		// テスト後にコントローラーを終了
		mockCtrl.Finish()
	})

	Describe("UseBookTitle", func() {
		Context("本のタイトルを使用する場合", func() {
			It("Bookインターフェースから取得したタイトルを返すこと", func() {
				result := testUser.UseBookTitle()
				Expect(result).To(Equal(bookTitle))
			})
		})
	})
})
