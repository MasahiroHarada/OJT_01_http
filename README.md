# HTTP基礎　サンプルアプリ
## 起動方法
### Go言語のインストール
[https://golang.org/dl/](https://golang.org/dl/)
### GOPATHを設定する
1. ホームフォルダ直下に ```go``` フォルダを作成
1. 環境変数 ```GOPATH``` に ```%HOME%\go``` を設定  
### サンプルアプリ展開
1. zipをダウンロード
1. ```GOPATH``` 配下に展開  

最終的に以下のようなフォルダ構成になってほしい。
```
%HOME%
  └ go
    └ OJT_01_http
      └ アプリケーションコード
```
### 依存パッケージをダウンロード
コマンドプロンプトで下記を実行。
```
go get -u github.com/kardianos/govendor
govendor sync
```
### アプリをビルドする
コマンドプロンプトで下記を実行。
```
cd
cd go\OJT_01_http
go build
```
### 起動
アプリケーションフォルダ ```OJT_01_http``` 直下に生成されるexeファイルを実行する。
