package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"crypto/hmac"
	"golang.org/x/crypto/bcrypt"
)

var tpl *template.Template

var key = []byte("astronaut juggernaut flouride")

// secretKey DO NOT SHARE
const secretKey = "home florpus magic astronaut"

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

// User is exported so that it can be an embedded type in pageData
type User struct {
	First    string
	Email    string
	Password []byte
	LoggedIn bool
}

type pageData struct {
	User
	Title   string
	Heading string
}

// create variable to hold value of TYPE map[string]user
// DOES NOT INITIALIZE variable with VALUE in memory
// var db map[string]user

// create variable to hold value of TYPE map[string]user
// INITIALIZES variable with VALUE in memory
// email is the key
// user is the value
var db = map[string]User{}

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/signupprocess", processSignUp)
	http.HandleFunc("/login", login)
	http.HandleFunc("/loginprocess", processLogin)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe(":3000", nil)
}

func getUser(r *http.Request) (User, error) {
	u := User{}

	c, err := r.Cookie("user-session")
	if err != nil {
		c = &http.Cookie{}
	}

	// make sure client did not change the email stored
	// eg, we don't want the client altering the email in an attempt
	// to access another user's account
	// the email in the cookie needs to be verified
	// as the email we set in the cookie
	xs := strings.Split(c.Value, "|")
	if len(xs) == 2 {
		e := xs[0]
		cookieHash := xs[1]

		sum := sha256.Sum256([]byte(e + secretKey))
		hash := fmt.Sprintf("%x", sum)

		if cookieHash != hash {
			return u, fmt.Errorf("hashes weren't equal")
		}

		ok := false
		if u, ok = db[e]; ok {
			u.LoggedIn = true
		}
	}
	
	return u, nil
}

func index(w http.ResponseWriter, r *http.Request) {
	u, err := getUser(r)
	if err != nil {
		http.Error(w, "are you hacking us? "+err.Error(), http.StatusForbidden)
		return
	}

	pd := pageData{
		User:    u,
		Title:   "Acme Inc.",
		Heading: "Welcome To Acme Inc.",
	}

	err = tpl.ExecuteTemplate(w, "index.gohtml", pd)
	if err != nil {
		log.Println("index.gohtml couldn't ExecuteTemplate", err)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	u, err := getUser(r)
	if err != nil {
		http.Error(w, "are you hacking us?"+err.Error(), http.StatusForbidden)
		return
	}

	pd := pageData{
		User:    u,
		Title:   "About Acme Inc.",
		Heading: "All About Acme",
	}

	err = tpl.ExecuteTemplate(w, "about.gohtml", pd)
	if err != nil {
		log.Println("about.gohtml couldn't ExecuteTemplate", err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	u, err := getUser(r)
	if err != nil {
		http.Error(w, "are you hacking us?"+err.Error(), http.StatusForbidden)
		return
	}

	pd := pageData{
		User:    u,
		Title:   "Contact Acme Inc.",
		Heading: "Contact Acme Incorporated",
	}

	err = tpl.ExecuteTemplate(w, "contact.gohtml", pd)
	if err != nil {
		log.Println("contact.gohtml couldn't ExecuteTemplate", err)
	}
}

func signup(w http.ResponseWriter, r *http.Request) {
	u, err := getUser(r)
	if err != nil {
		http.Error(w, "are you hacking us?"+err.Error(), http.StatusForbidden)
		return
	}

	if u.Email != "" {
		// already logged in, so redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	pd := pageData{
		User:    u,
		Title:   "SIGN UP",
		Heading: "SIGN UP AT ACME INC.",
	}

	err = tpl.ExecuteTemplate(w, "signup.gohtml", pd)
	if err != nil {
		log.Println("signup.gohtml couldn't ExecuteTemplate", err)
	}
}

func processSignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusNotFound)
		return
	}

	f := r.FormValue("fn")
	if f == "" {
		http.Error(w, "first name cannot be empty", http.StatusBadRequest)
		return
	}

	e := r.FormValue("em")
	if e == "" {
		http.Error(w, "email cannot be empty", http.StatusBadRequest)
		return
	}

	p := r.FormValue("pw")
	if p == "" {
		http.Error(w, "password cannot be empty", http.StatusBadRequest)
		return
	}

	bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "password could not be bcrypt'd", http.StatusBadRequest)
		return
	}

	// never directly store a password
	// store the bcrypt or scrypt of the password
	u := User{
		First:    f,
		Email:    e,
		Password: bs,
	}

	db[e] = u

	// never directly store an important value without checking to see if ...
	// the user has tampered with the value
	sum := sha256.Sum256([]byte(e + secretKey))
	hash := fmt.Sprintf("%x", sum)
	v := e + "|" + hash
	c := &http.Cookie{
		Name:  "user-session",
		Value: v,
		Path:  "/",
	}

	http.SetCookie(w, c)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func login(w http.ResponseWriter, r *http.Request) {
	u, err := getUser(r)
	if err != nil {
		http.Error(w, "are you hacking us?"+err.Error(), http.StatusForbidden)
		return
	}

	if u.Email != "" {
		// already logged in, so redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	pd := pageData{
		User:    u,
		Title:   "LOGIN",
		Heading: "LOGIN TO ACME INC.",
	}

	tpl.ExecuteTemplate(w, "login.gohtml", pd)
}

func processLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusNotFound)
		return
	}

	e := r.FormValue("em")
	if e == "" {
		http.Error(w, "email cannot be empty", http.StatusBadRequest)
		return
	}

	p := r.FormValue("pw")
	if p == "" {
		http.Error(w, "password cannot be empty", http.StatusBadRequest)
		return
	}

	u := User{}
	ok := false
	if u, ok = db[e]; !ok {
		http.Error(w, "username or password not correct", http.StatusBadRequest)
		return
	}

	err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
	if err != nil {
		http.Error(w, "username or password not correct BCRYPT'D", http.StatusBadRequest)
		return
	}

	// never directly store an important value without checking to see if ...
	// the user has tampered with the value
	sum := sha256.Sum256([]byte(e + secretKey))
	hash := fmt.Sprintf("%x", sum)
	v := e + "|" + hash
	c := &http.Cookie{
		Name:  "user-session",
		Value: v,
		Path:  "/",
	}

	http.SetCookie(w, c)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func logout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("user-session")
	if err != nil {
		c = &http.Cookie{
			Name: "user-session",
		}
	}

	c.MaxAge = -1

	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func signMessage(msg []byte) ([]byte, error) {
	h := hmac.New(sha512.New, key)
	_, err := h.Write(msg)
	if err != nil {
		return nil, fmt.Errorf("error in signMessage writing to hash %w", err)
	}
	signature := h.Sum(nil)
	return signature, nil
}

func checkSig(msg, sig []byte) (bool, error) {
	newSig, err := signMessage(msg)
	if err != nil {
		return false, fmt.Errorf("Error in checkSig while getting signature of message: %w", err)
	}

	same := hmac.Equal(newSig, sig)
	return same, nil
}
