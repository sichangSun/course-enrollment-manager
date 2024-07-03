# course-enrollment-manager
## 概要
学生が授業履修登録できるアプリケーションです。

## 機能紹介
ログイン機能<br>
（個人情報変更、パスワード変更）<br>
時間割で登録済みな授業を確認<br>
授業一覧から授業を選べ、詳細を確認<br>
授業の登録、削除。<br><br>
<img width="1728" alt="スクリーンショット 2024-07-03 18 23 54" src="https://github.com/sichangSun/course-enrollment-manager/assets/122949473/17c13682-ef23-4f3b-9086-a65116b4c163">
<br><br>
<img width="1728" alt="スクリーンショット 2024-07-03 18 24 50" src="https://github.com/sichangSun/course-enrollment-manager/assets/122949473/66c0a514-8574-461e-8de0-83532da7a67f">
<br><br>
<img width="1728" alt="スクリーンショット 2024-07-03 18 25 28" src="https://github.com/sichangSun/course-enrollment-manager/assets/122949473/389ff89a-ecc8-49b8-832c-ff54a56ab8fa">

## 使用技術・環境
### 【フロントエンド】
・HTML<br>
・CSS<br>
・Vite / Vue3<br>
・Javascript / Axios <br>
・Vuetify3 <br>
・vitejs<br>


### 【バックエンド】
・Go 1.2 / Echo 4.12.0<br>

### 【開発環境】
・MySQL<br>
・Docker / Docker-compose
<br>
## サーバーサイドアーキテクチャ
![109972929-26a7b680-7d3b-11eb-9a7f-41a2df491d57](https://github.com/sichangSun/course-enrollment-manager/assets/122949473/e3a1b66d-bf06-4aed-9afa-32d64094aff9)

## 工夫点
1. ユーザーの操作性を考え、授業詳細画面と授業一覧画面両方が授業を登録を削除できるように実装しました。
2. ユーザーが直感的に扱いやすいように、時間割で登録済みな授業が確認できるように実装しました。
3. フロントでは、繰り返す使えるようにコンポネートを分けるというところ意識して実装しました。


## ローカル実行手順
### 1.リポジトリのクローン<br>
### 2.env.localを追加<br>
フロントエンドのルートディレクトリに .env.local ファイルを作成し、次の内容を記入してください。<br>
```

APP_ENV=local
VITE_BASE_URL=http://localhost:8000

```


### 3. Docker コンテナの構築
以下のコマンドを使用して Docker コンテナをビルドします。<br>
```

docker compose build

```
コンテナを起動するためには、次のコマンドを実行します。<br>
```

docker compose up

```
<br>

### 注意点


> ・データベースの起動が遅れることがあります。そのため、バックエンドの起動に失敗する場合がDBが完全に起動していない可能性が高いです。<br>
<br>
・もしバックエンド起動が失敗した場合、DBが完全に起動するまで少し待ってから、バックエンドのビルドを再試行してください。<br><br>
```

docker compose build backend
docker compose up backend -d

```

### 4. アプリケーションへのアクセス
ブラウザを開いて以下のURLにアクセスします。<br>
http://localhost:5173/<br><br>


## 問題点＆これから解決ポイント
### ローカル
1. docker compose起動する際に、バックエンドがDB起動する前にbuild完成<br>
2. API側では、全授業を返す部分順番をASCで指定したけど、結果的にリフレッシュする際に、フロンドで順番が変わること<br>
3. ブラウザのリフレッシュで画面表示一部おかしくなること<br>
4. 授業を登録する際に、同じ時間を登録できないバリエーションチエックを追加<br>
5. パスワード変更機能が疎通していない<br>
6. API追加：　学生情報を取得 / Like検索<br>


