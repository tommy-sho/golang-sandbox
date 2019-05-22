## JWT 認証
まずは/authにログインしてJWTをゲットする。
ゲットしたJWTをAuthorizationって名前のヘッダーにつけて、Bearerもつけて送信する

`curl -i -H "Authorization:Bearer {さっきのJWT}" localhost:8080/private`