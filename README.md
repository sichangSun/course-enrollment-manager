# course-enrollment-manager

## 概要

学生が授業履修登録できるアプリケーションです。

## 機能紹介

- ログイン機能 （個人情報変更、パスワード変更）
- 時間割で登録済み授業を確認
- 授業一覧から授業を選択、詳細を確認
- 授業の登録、削除
<img width="1728" alt="スクリーンショット 2024-07-03 18 23 54" src="https://github.com/sichangSun/course-enrollment-manager/assets/122949473/17c13682-ef23-4f3b-9086-a65116b4c163">
<br><br>
<img width="1728" alt="スクリーンショット 2024-07-03 18 24 50" src="https://github.com/sichangSun/course-enrollment-manager/assets/122949473/66c0a514-8574-461e-8de0-83532da7a67f">
<br><br>
<img width="1728" alt="スクリーンショット 2024-07-03 18 25 28" src="https://github.com/sichangSun/course-enrollment-manager/assets/122949473/389ff89a-ecc8-49b8-832c-ff54a56ab8fa">

## 使用技術・環境

### 【フロントエンド】

- HTML
- CSS
- Vite / Vue3
- Javascript / Axios
- Vuetify3
- vitejs

### 【バックエンド】

- Go 1.22
- Echo 4.12.0

### 【開発環境】

- MySQL
- Docker / Docker-compose

## バックエンドアーキテクチャ

![109972929-26a7b680-7d3b-11eb-9a7f-41a2df491d57](https://github.com/sichangSun/course-enrollment-manager/assets/122949473/e3a1b66d-bf06-4aed-9afa-32d64094aff9)

## 工夫点

1. ユーザーの操作性を考え、授業詳細画面と授業一覧画面の両方で授業の登録と削除ができるように実装しました。
2. ユーザーが直感的に扱いやすいように、時間割で登録済み授業を確認できるように実装しました。
3. フロントエンドでは、繰り返し使えるようにコンポーネントを分けることを意識して実装しました。


## ローカル実行手順

### 1. リポジトリのクローン

### 2. .env.localを追加

フロントエンドのルートディレクトリに .env.local ファイルを作成し、次の内容を記入してください。<br>
```

APP_ENV=local
VITE_BASE_URL=http://localhost:8000

```


### 3. Docker コンテナの構築

以下のコマンドを使用して Docker コンテナをビルドします。
```

docker compose build

```

コンテナを起動するためには、次のコマンドを実行します。
```

docker compose up

```


### 注意点

- データベースの起動が遅れることがあります。そのため、バックエンドの起動に失敗する場合、DBが完全に起動していない可能性が高いです。
- もしバックエンドの起動が失敗した場合、DBが完全に起動するまで少し待ってから、バックエンドのビルドを再試行してください。


```

docker compose build backend

```

```

docker compose up backend -d

```

### 4. アプリケーションへのアクセス

ブラウザを開いて以下のURLにアクセスします。<br>
http://localhost:5173/<br>
user: `suzuki@scon.com`<br>
password:`KRhjh439jsc`

## 問題点とこれから解決すべきポイント

### ローカル

1. docker composeを起動する際に、DB起動後にバッククエントコンテナを起動
2. API側では全授業を返す部分の順番をASCで指定しているが、リフレッシュする際にフロンドエンドで順番が変わること
3. ブラウザのリフレッシュ時に画面表示の一部がおかしくなる
4. 授業を登録する際に、同じ時間の登録ができないバリデーションチェックを追加
5. パスワード変更機能が疎通していない
6. 学生情報を取得機能、授業検索機能
