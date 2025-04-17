# Ginkgoを使ったテスト実装ガイド
## 概要
このリポジトリは、Go言語のBDD（振る舞い駆動開発）テストフレームワークGinkgoを使ったテスト実装方法をまとめたものです。Ginkgoの基本的な使い方からテストの実行方法までを解説します。
## Ginkgoとは
- Go言語で書かれたBDD（振る舞い駆動開発）スタイルのテストフレームワーク
- 階層的なテスト構造を記述できるため、複雑なテストシナリオを整理しやすい
- Gomega というMatcherライブラリと組み合わせて使うのが一般的
  - ただし、他のMatcherライブラリとも併用可能な設計になっている

## Gomegaとは
- Ginkgoと組み合わせて使うマッチャー/アサーションライブラリ
  - **マッチャー**: 期待値と実際の値を比較して、一致するかどうかの結果を返すオブジェクト
  - **アサーション**: コードが実行される時に満たされるべき条件を記述して、実行時にチェックする仕組み

## インストール方法
Ginkgoをインストールするには以下のコマンドを実行します：
``` bash
# ginkgoコマンドラインツールのインストール
go install github.com/onsi/ginkgo/v2/ginkgo@latest

# テストで使用するライブラリのインストール
go get github.com/onsi/ginkgo/v2
go get github.com/onsi/gomega
```
## Ginkgoの基本的な使い方
### 1. テストスイートの作成（Bootstrap）
プロジェクトディレクトリで以下のコマンドを実行してテストスイートを初期化します：
``` bash
ginkgo bootstrap
```
これにより、現在のパッケージ用のテストスイートファイル（`xxx_suite_test.go`）が生成されます。
### 2. テストファイルの生成
特定のファイルに対するテストを生成するには：
``` bash
ginkgo generate <ファイル名>
```
例えば、 に対するテストを生成する場合は `ginkgo generate book` を実行します。 `book.go`
### 3. テストの実行
テストを実行するには以下のコマンドを使用します：
``` bash
# ginkgoコマンドを使う場合
ginkgo

# または標準のgoコマンドを使う場合
go test
```
その他の便利なオプション：
``` bash
# 再帰的に全てのサブディレクトリのテストを実行
ginkgo -r

# テスト実行時にコードカバレッジを計測
ginkgo -cover

# 詳細な出力を表示
ginkgo -v
```
## テストの書き方の例
Ginkgoでは、、、などの関数を使ってテストを構造化します： `Describe``Context``It`
``` go
var _ = Describe("Book", func() {
    Describe("Categorizing book length", func() {
        var (
            longBook  Book
            shortBook Book
        )

        BeforeEach(func() {
            longBook = NewBook("title1", "author1", 1000)
            shortBook = NewBook("title2", "author2", 10)
        })

        Context("With more than 300 pages", func() {
            It("should be a novel", func() {
                actual := longBook.CategoryByLength()
                Expect(actual).To(Equal("Novel"))
            })
        })

        Context("With fewer than 300 pages", func() {
            It("should be a short story", func() {
                actual := shortBook.CategoryByLength()
                Expect(actual).To(Equal("Short story"))
            })
        })
    })
})
```
## mockgenを使ったモックファイル作成
Goのテストで外部依存関係をモック化するために、GoMockフレームワークとmockgenツールが非常に便利です。インターフェースに基づいてモックを自動生成できます。
### インストール方法
まず、mockgenツールをインストールします：
``` bash
go install go.uber.org/mock/mockgen@latest
```
### モックの生成方法
mockgenを使用するには主に2つの方法があります：
#### 1. ソースモード：ソースファイルからインターフェースを直接指定
``` bash
mockgen -source=<ソースファイルのパス> -destination=<出力先ファイルパス> -package=<パッケージ名>
```
例：
``` bash
mockgen -source=./repository/user_repository.go -destination=./mocks/mock_user_repository.go -package=mocks
```
#### 2. リフレクションモード：パッケージとインターフェース名を指定
``` bash
mockgen <パッケージパス> <インターフェース名1>,<インターフェース名2> -destination=<出力先ファイルパス> -package=<パッケージ名>
```
例：
``` bash
mockgen ko44d/go-bdd-testing/repository UserRepository -destination=./mocks/mock_user_repository.go -package=mocks
```
### go:generateを使った自動生成
コード内に特別なコメントを追加することで、`go generate`コマンドを使用してモックを自動生成できます：
``` go
//go:generate mockgen -source=$GOFILE -destination=../mocks/mock_$GOFILE -package=mocks
```
このコメントをインターフェースを含むファイルに追加し、以下のコマンドを実行：
``` bash
go generate ./...
```
### モックの使用例
モックを使用したテストの例：
``` go
func TestWithMock(t *testing.T) {
    // mockコントローラーの作成
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()
    
    // モックの作成
    mockRepo := mocks.NewMockUserRepository(ctrl)
    
    // モックの振る舞いを定義
    mockRepo.EXPECT().
        GetUserByID(gomock.Eq(1)).
        Return(&models.User{ID: 1, Name: "テストユーザー"}, nil)
    
    // テスト対象のサービスにモックを注入
    service := NewUserService(mockRepo)
    
    // テスト実行
    user, err := service.GetUser(1)
    
    // 結果の検証
    assert.NoError(t, err)
    assert.Equal(t, "テストユーザー", user.Name)
}
```
### Ginkgoとの組み合わせ
GoMockはGinkgoと組み合わせて使用することもできます：
``` go
var _ = Describe("UserService", func() {
    var (
        mockCtrl *gomock.Controller
        mockRepo *mocks.MockUserRepository
        service  *UserService
    )
    
    BeforeEach(func() {
        mockCtrl = gomock.NewController(GinkgoT())
        mockRepo = mocks.NewMockUserRepository(mockCtrl)
        service = NewUserService(mockRepo)
    })
    
    AfterEach(func() {
        mockCtrl.Finish()
    })
    
    Context("GetUser", func() {
        It("正常にユーザーを取得できること", func() {
            mockRepo.EXPECT().
                GetUserByID(gomock.Eq(1)).
                Return(&models.User{ID: 1, Name: "テストユーザー"}, nil)
                
            user, err := service.GetUser(1)
            
            Expect(err).NotTo(HaveOccurred())
            Expect(user.Name).To(Equal("テストユーザー"))
        })
        
        It("ユーザーが存在しない場合はエラーを返すこと", func() {
            mockRepo.EXPECT().
                GetUserByID(gomock.Any()).
                Return(nil, errors.New("user not found"))
                
            _, err := service.GetUser(999)
            
            Expect(err).To(HaveOccurred())
            Expect(err.Error()).To(Equal("user not found"))
        })
    })
})
```
### mockgenの便利なオプション
- `-mock_names`: モック名をカスタマイズ（例：`-mock_names=UserRepository=MockUserRepo`）
- `-copyright_file`: 著作権表示を含むファイルを指定
- `-imports`: 追加のインポートを指定
- `-aux_files`: 補助ファイルを指定（インターフェースが複数のファイルに分かれている場合）
- `-self_package`: 生成されたモックコードからオリジナルのパッケージを参照する場合に使用

モックを使用することで、外部依存関係（データベース、APIクライアントなど）を持つコードも確実にテストでき、テストの再現性と信頼性が向上します [[1]](https://github.com/uber-go/mock)[[2]](https://speedscale.com/blog/getting-started-gomock/)。

## 参考リソース
- [Ginkgo公式ドキュメント](https://onsi.github.io/ginkgo/)
- [Gomega公式ドキュメント](https://onsi.github.io/gomega/)
