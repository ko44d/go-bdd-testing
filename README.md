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
## 参考リソース
- [Ginkgo公式ドキュメント](https://onsi.github.io/ginkgo/)
- [Gomega公式ドキュメント](https://onsi.github.io/gomega/)
