create-react-appの機能で、APIのリクエストURLをプロキシする機能がある。これを使うことで、ローカル環境でCORSエラーが起こらなくなる

https://create-react-app.dev/docs/proxying-api-requests-in-development/

Goコンテナではbashでログインすることができない。alpineにしていたから？

JWTにユーザ情報が含まれているのか？だとしたら、JS側でデコードしないといけないのか？
 - そんなわけはなく、フロントエンドにJWTのデコードはできないはず
 - react query authはlogin apiの戻り値にユーザ情報が返ってくることを期待しているが、gin-jwtの実装では、login apiはjwtを返すのみとなっている
 - gin-jwtの実装に合わせる
 - と思ったが、loginfnがUserオブジェクトを返さないといけないようになっているので、me apiを呼ぶようにした

react-queryについて
 - Reactアプリケーションのサーバー状態のフェッチ、キャッシュ、同期、更新が簡単になります。
 - https://react-query.tanstack.com/overview

react-query-authについて
react-queryのおかげで、サーバーの状態をキャッシュすることでコードベースを大幅に削減することができました。
ただし、ユーザーデータをどこに保存するかを考える必要があります。
ユーザーデータは、アプリケーション内のさまざまな場所からアクセスする必要があるため、グローバルなアプリケーション状態と見なすことができます。
一方、すべてのユーザーデータはサーバーから到着すると予想されるため、サーバーの状態でもあります。
このライブラリを使用すると、ユーザー認証を簡単に管理できます。
これは、アプリケーションの認証に使用している方法に依存せず、使用されているAPIに応じて調整できます。
構成を提供するだけで、残りは自動的に設定されます。

loadUserをリトライしている。リトライをいい感じにやってくれるのかも
websocketのプロセスが動いていて、そこでサーバとやりとりしてる？ => そんなことはない。サーバがWebSocket対応じゃないので。

何をするのか
 - プロバイダとフックを提供
 - これにより、ログインユーザ情報と、login, logoutのメソッドを任意のコンポーネントからアクセスできるようになる

向いているユースケース
 - 認証APIを自前で提供している場合
 - API側で特別な実装は必要なし
 - Cookie / JWTも問わない(そこの実装に関するAPIは提供していない)
   - 逆に言えば、JWTをローカルストレージに保存するコードは自前実装する必要あり(サンプルコードがあるので真似すればいい)

向いていないユースケース
 - Firebase Authenticationなど、フロントエンド側のライブラリが提供されている認証プロバイダを使う場合（敢えて採用する必要はなさそう）
   - Firebase SDKは、ログイン状態が変わったときのコールバックを定義できる

挙動
 - ローカルストレージからtokenを消したりサーバを落としたりトークンの期限が切れるとログアウト状態になる
 - ログイン状態であればキャッシュする

gin-jwt
 - セッションが期限切れになった場合の挙動を考えたい。
 - トークンの有効期限を延長するAPIを叩く必要が出てきそうな雰囲気
   - トークンをリフレッシュすることがはできるが、いわゆるリフレッシュトークンとは異なるっぽい。
      - それって意味あるの？ JWTが流出したらアウトになってしまうのでは。
      - 勘違いしているかもしれないので、調査した方が良さそう。

ローダーとエラーコンポーネントを定義した。
 - エラーコンポーネントに副作用でローカルストレージをフラッシュする機能を追加した。
   - ちょっと気持ち悪い感じはするが、このライブラリを使うならこうなるかなという感じ。
