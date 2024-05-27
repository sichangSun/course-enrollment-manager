# course-enrollment-manager
## 概要
学生が授業履修登録できるアプリケーションです。

## 機能紹介
ログイン、パスワード変更。<br>
時間割で登録済みな授業を確認できる。<br>
授業一覧から授業を選べ、詳細を確認できる。<br>
授業の登録、削除。<br>

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

## ローカル実行手順
### 1.リポジトリのクローン<br>
### 2.env.localを追加<br>
フロントエンドのルートディレクトリに .env.local ファイルを作成し、次の内容を記入してください。<br>
APP_ENV=local<br>
VITE_BASE_URL=http://localhost:8000/<br>


### 3. Docker コンテナの構築
以下のコマンドを使用して Docker コンテナをビルドします。<br>
docker compose build<br>
コンテナを起動するためには、次のコマンドを実行します。<br>
docker compose up<br>
### 注意点
・データベースの起動が遅れることがあります。そのため、バックエンドの起動に失敗する場合がDBが完全に起動していない可能性が高いです。
・もしバックエンド起動が失敗した場合、DBが完全に起動するまで少し待ってから、バックエンドのビルドを再試行してください。<br>
docker compose build backend<br>
docker compose up backend -d<br>

### 4. アプリケーションへのアクセス
ブラウザを開いて以下のURLにアクセスします。<br>
http://localhost:5173/<br>

## 現状　
フロントは完成していない状況。ログインのあとにある学生ホーム画面以外は、ずべて仮のfakeデータで画面表示のみになっている。<br>

## 問題点＆これから解決必要
### ローカル
1.docker compose起動する際に、バックエンドがDB起動する前にbuild完成<br>
2.API側では、全授業を返す部分、limitとindexでの検索を検討必要。それに応じてフロントでの対応<br>
3.API追加：　学生情報を取得 / Like検索<br>


