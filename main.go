package main

import (
	"net/http"
	"html/template"
	"log"
	"github.com/satori/go.uuid"
)

var tpl *template.Template
var session SessionManager

var cookieName string = "DEMO_SESSION_ID"

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	session = NewFileSessionManager()
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set-cookie", set)
	http.HandleFunc("/reset-cookie", reset)

	// static contents
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	http.ListenAndServe(":9876", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	// Cookieがセットされているか調べる
	cookie, err := req.Cookie(cookieName)

	// Cookie発行前
	if err != nil {
		err = tpl.ExecuteTemplate(w, "index.html", nil)
		handleError(w, err)
		return
	}

	// Cookie発行済み
	// セッションデータの検索
	username, err := session.Fetch(cookie.Value)

	// Cookieの値に対応するセッションデータが見つからなかった場合
	if err != nil {
		// Cookieの破棄
		w.Header().Set("Location", "/reset")
		w.WriteHeader(http.StatusSeeOther)
	}

	// templateに渡すデータを生成
	templateData := struct {
		Name string
	}{Name: username}

	// templateにデータを渡してHTMLを生成
	err = tpl.ExecuteTemplate(w, "index.html", templateData)
	handleError(w, err)
}

func set(w http.ResponseWriter, req *http.Request) {
	// POSTメソッド以外は受け付けない
	if req.Method != http.MethodPost {
		// POST以外のメソッドだった場合は、METHOD NOT ALLOWED
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// リクエストボディの読み込み
	req.ParseForm()
	// usernameの値を取得
	username := req.PostFormValue("username")

	// Cookieを生成
	cookie := newCookie()

	// レスポンスヘッダーにCookie付加
	http.SetCookie(w, cookie)

	// セッションデータ生成
	session.Add(cookie.Value, username)

	// TOPページにリダイレクト
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusSeeOther)
}

func reset(w http.ResponseWriter, req *http.Request) {
	// Cookieがセットされているか調べる
	cookie, err := req.Cookie(cookieName)
	handleError(w, err)

	// セッションを破棄
	session.Destroy(cookie.Value)

	// Cookieを破棄
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)

	// TOPページにリダイレクト
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusSeeOther)
}

func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err.Error())
	}
}

func newCookie() *http.Cookie {
	id := uuid.NewV4()

	return &http.Cookie{
		Name: cookieName,
		Value: id.String(),
		Path: "/",
		HttpOnly: true,
	}
}
