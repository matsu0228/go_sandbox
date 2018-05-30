# sample of JSON API

## requirement

* おみくじAPI
  * JSON形式でおみくじの結果を返す
  * 正月（1/1-1/3）だけ大吉にする

## design

* packageは、`lib`にまとめた
  * 大きく、`server: API server` と`fortune: おみくじを引く機能`がある。
  * それぞれにテストを書いた。UnitTestを書きやすくするため、`Drawer`のinterfaceをつくり、`Server`ではこのinterfaceにのみ依存するように実装
* endpoint
  * localhost:8080/
  * localhost:8080/twice