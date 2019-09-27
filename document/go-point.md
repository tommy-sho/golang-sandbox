## Godoc

https://qiita.com/shibukawa/items/8c70fdd1972fad76a5ce
https://qiita.com/lufia/items/97acb391c26f967048f1


## Build Constraint
goは基本的に同一ディレクトリにあるコードはビルド対象になるが、「Build Constraints」というコメント的なものをつけることによって、環境ごとにビルド対象にするのかしないのかを設定することができる
以下のようなものを記述できる
```sh
OS名（$GOOSの値、linux, windows, darwinなど）
アーキテクチャ名（$GOARCHの値、arm, 386, amd64など）
cgo (cgoが有効なとき)
!cgo (cgoが無効なとき)
コンパイラ名（go, gccgo）
!コンパイラ名
tag (ctxt.BuildTags か ctxt.ReleaseTags に含まれるタグ)
!tag (ctxt.BuildTags と ctxt.ReleaseTags に含まれないタグ)
上記のカンマ区切りリスト（AND条件）

ReleaseTagsはGo 1.4だと次のような値になっています。これを使えばGoの特定バージョン以降・以前という表現ができます。

```


https://qiita.com/hnw/items/7dd4d10e837158d6156a