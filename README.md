## 概要
Go言語のBDD用のフレームワークとしてGinkgoを使っているため、Ginkgoでのユニットテストの実行方法や実装方法についてまとめる。

## Ginkgoとは
- https://onsi.github.io/ginkgo/
- Go言語のBDD用テストフレームワークである
- Gomega というMatcherライブラリと組み合わせるのが最適である
    - 但し、Matcherライブラリには依存しない設計となっている

## Gomegaとは
- https://onsi.github.io/gomega/
- マッチャー/アサーションライブラリ
  - マッチャーとは
    - 期待値と実際の値を比較して、一致するかどうかの結果を返すオブジェクト
  - アサーションとは
    - コードが実行される時に満たされるべき条件を記述して、実行時にチェックする仕組みのこと
- 
## テストの実行
```
ginkgo
go test
```

## bootstrapの作成
```
ginkgo bootstrap
```

## Specの追加
```
ginkgo generate <ファイル名>
```
