## What's this?
GoogleのProgrammable Search Engineのお試し
https://developers.google.com/custom-search

## How to run?
1. clone this repo.
2. set environment variables.
```bash
export CUSTOMSEARCH_API_TOKEN="{API_TOKEN}"
export SEARCH_ID="{CustomSearchEngine_ID}"
```
3. build and run
```bash
$ go build webserver.go
$ ./webserver
```
You can access the test page at [`localhost:3030`](localhost:3030).

## ページ内容
[`localhost:3030`](localhost:3030)にアクセスして表示される内容．各番号が対応している．
1. Programmable Search Element Control API: [https://developers.google.com/custom-search/docs/element](https://developers.google.com/custom-search/docs/element)
    1. JavaScriptで埋め込むやつ
    2. 通常のGoogle検索みたいなUIをカスタマイズして埋め込める．
    3. ↑のカスタマイズはGoogle Programmable Search Engineのコントロールパネルでグラフィカルに設定できる．
    4. JSにコールバック関数を登録すれば多分検索結果も取得できる
2. カスタム検索 JSON API: [https://developers.google.com/custom-search/v1/overview](https://developers.google.com/custom-search/v1/overview)
    1. 次のURLにパラメータを渡して叩くとJSONで検索結果が返ってくる．[https://customsearch.googleapis.com/customsearch/v1](https://customsearch.googleapis.com/customsearch/v1)
    2. これが返ってくる： [https://developers.google.com/custom-search/v1/reference/rest/v1/Search](https://developers.google.com/custom-search/v1/reference/rest/v1/Search)
    3. UIは自前で用意．
    4. Google検索を利用している旨の表記がどこかに必要っぽい？（要確認）
    5. このURLも叩ける．違いはまだ分からない： [https://www.googleapis.com/customsearch/v1](https://www.googleapis.com/customsearch/v1)
3. 各言語で実装されたAPI（Golang, Python等）
    1. 基本的にJSON APIと同じ．
    2. golangで試した．
    3. 結果がオブジェクトで返ってくるので扱いやすいかも



## (Archived)雑作業ログ
https://programmablesearchengine.google.com/controlpanel/all
でプログラム可能検索エンジンを作成

プログラム可能検索エンジンには以下の２つが含まれる
> * Programmable Search Element Control API を使用すると、プログラム可能な検索要素を JavaScript でウェブページやその他のウェブ アプリケーションに埋め込めます。
> * Custom Search JSON API を使用すると、プログラム可能検索エンジンから検索結果をプログラムで取得して表示するためのウェブサイトやプログラムを作成できます。この API では、RESTful リクエストを使用して JSON 形式の検索結果を取得できます。

### Custom Search JSON APIのkeyを発行
https://developers.google.com/custom-search/v1/overview
キーの取得を押下で発行できる．


### Custom Search JSON APIの利用
このAPIはOpenSearch1.1↓を満たすらしい．
https://github.com/dewitt/opensearch/blob/master/opensearch-1-1-draft-6.md


### 高速化
GolangのAPIや直接エンドポイントを叩くAPIはレスポンスが遅い．（8秒くらいかかる）
GolangのAPIで以下の高速化を試した
* レスポンスのgzip圧縮→デフォルトで有効
* `Start` オプションの指定をやめる→効果なし
* レスポンスの要素の限定→効果なし
* 画像検索の無効化（GCP上のUIで変更可能）→It works!
数百msecでレスポンスが返ってくるように．（サーバ上での実行時間と数百msec差）
