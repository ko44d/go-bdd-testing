## Describe句
テストケースの「動作」説明を記述する
```go
Describe("", func(){
	
})
```

## Context句
テストケースの「条件」説明を記述する
```go
Describe("", func(){
    Context("", func(){
	
    })
})
```
## IT句
テストケースの「仕様」説明を記述する
```go
Describe("", func(){
    Context("", func(){
        It("",func(){

        })
    })
})

```
## BeforeEach句
- テストコード間での共通設定を記述する
- テストメソッド実行前の準備について記述する
- ブロック内でのアサーションを配置することは一般的である。なぜなら、テスト実行前の準備中にエラーが発生しないかチェックするため。

```go
Describe("", func(){
    BeforeEach(func(){

    })

    Context("", func(){
        BeforeEach(func(){

        })       
		
		It("",func(){

        })
    })
})
```
## AfterEach句
テストの最後にクリーンアップするものを記述する
```go
Describe("", func(){
    BeforeEach(func(){

    })
	
    AfterEach(func(){
	
    })

    Context("", func(){
        BeforeEach(func(){

        })

        It("",func(){

        })
	})
})
```
