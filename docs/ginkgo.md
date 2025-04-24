## Describe句
テスト対象の「動作」や機能全体を表すブロック。トップレベルの説明単位として用いられる。

```go
Describe("対象の説明", func() {

})
```

## Context句
前提条件やシナリオの違いを明確にするためのブロック。Describe内にネストして使用する。

```go
Describe("機能", func() {
	Context("特定の条件下", func() {
	
	})
})
```

## It句
「期待する振る舞い（仕様）」を具体的に記述する。1つのテストケースを表す最小単位。

```go
Describe("機能", func() {
	Context("特定の条件下", func() {
		It("期待される動作を説明", func() {
		
		})
	})
})
```

## BeforeEach句
- 各テスト実行前に毎回実行される処理を定義する。
- 共通の初期化処理や前提データのセットアップに使用。
- アサーションを含めることで、準備段階でのエラーも検知可能。

```go
Describe("機能", func() {
	BeforeEach(func() {
	
	})
	
	Context("特定の条件下", func() {
		BeforeEach(func() {
		
		})
		It("期待される動作", func() {
		
		})
	})
})
```

## AfterEach句
各テスト実行後に呼び出されるクリーンアップ処理を記述する。

```go
Describe("機能", func() {
	BeforeEach(func() {
	
	})
	
	AfterEach(func() {
	
	})
	
	Context("特定の条件下", func() {
		BeforeEach(func() {
		
		})
		
		It("期待される動作", func() {
			
		})
	})
})
```
