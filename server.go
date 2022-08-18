package website

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

var err error

func NewServer() *Server {
	srv := &Server{
		datas: Slice[Data]{},
	}
	return srv
}

func (srv *Server) SignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case "GET":
			path := r.URL.Path
			fmt.Println(path)
			http.ServeFile(w, r, htmlPagesPath+"signup/index.html")
		case "POST":
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}
			firstName = Name(r.FormValue("firstName"))
			lastName = Name(r.FormValue("lastName"))
			nation = Name(r.FormValue("nation"))
			gendre = Gendre(r.FormValue("gendre"))
			gmail = Gmail(r.FormValue("gmail"))
			password = Password(r.FormValue("password"))
			sbirthday := r.FormValue("birthday")
			verifyPassword = Password(r.FormValue("verifyPassword"))
			birthday, err = Sbirthday(sbirthday)

			var sliceOfErrors = []error{}
			sliceOfErrors = append(sliceOfErrors, firstName.Check(), lastName.Check(), nation.Check(), gendre.Check(), gmail.Check(), password.Check(), verifyPassword.Check(), err, birthday.Check(), password.Compare(verifyPassword))

			for _, v := range sliceOfErrors {
				if v != nil {
					fmt.Fprint(w, v.Error())
					return
				}
			}

			user := NewUser(firstName, lastName, nation, gendre, gmail, password, int(birthday.Day), int(birthday.Month), int(birthday.Year), uuid.New())
			path := string(gmail)

			err = os.Mkdir(dataBasePath+path+"dir", os.ModePerm)
			if err != nil {
				fmt.Fprint(w, errAlreadyOccupiedGmail.Error())
				return
			}

			jsonFile, err := os.Create(dataBasePath + path + "dir/" + "data.json")
			if err != nil {
				fmt.Fprint(w, errAlreadyOccupiedGmail.Error())
				return
			}

			err = json.NewEncoder(jsonFile).Encode(user)
			if err != nil {
				fmt.Fprint(w, err.Error())
				return
			}

			err = jsonFile.Close()
			if err != nil {
				log.Fatal(err)
			}
			redirectURL := fmt.Sprintf(`http://localhost%s/login/%s`, Port, gmail)
			http.Redirect(w, r, redirectURL, http.StatusSeeOther)

		default:
			fmt.Fprintf(w, "No other methods are supported so far.")
		}
	}
}

func (srv *Server) LogIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(htmlPagesPath + "logIn.html")
		if err != nil {
			log.Fatal(err)
		}
		switch r.Method {

		case "GET":
			t.Execute(w, logInRespond{})
		case "POST":

			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}

			gmail = Gmail(r.FormValue("gmail"))
			password = Password(r.FormValue("password"))

			for _, err := range []error{gmail.Check(), password.Check()} {
				if err != nil {
					t.Execute(w, logInRespond{
						Gmail:    gmail,
						Password: password,
						Error:    "Error: " + err.Error(),
					})
					return
				}
			}
			user := &User{}
			path := string(gmail)

			jsonFile, err := os.Open(dataBasePath + path + "dir/" + "data.json")
			if err != nil {
				t.Execute(w, logInRespond{
					Gmail:    gmail,
					Password: password,
					Error:    "Error: " + errNoRegisteredOrDeletedAccount.Error(),
				})
				return
			}

			err = json.NewDecoder(jsonFile).Decode(user)
			if err != nil {
				t.Execute(w, logInRespond{
					Gmail:    gmail,
					Password: password,
					Error:    "Error: " + errLostOrDeletedData.Error(),
				})
				return
			}

			err = password.Compare(user.Password)
			if err != nil {
				t.Execute(w, logInRespond{
					Gmail:    gmail,
					Password: password,
					Error:    "Error: " + err.Error(),
				})
				return
			}

			err = jsonFile.Close()
			if err != nil {
				t.Execute(w, logInRespond{
					Gmail:    gmail,
					Password: password,
					Error:    "Error: " + err.Error(),
				})
				log.Fatal(err)
			}

			redirectURL := fmt.Sprintf(`http://localhost%s/login/%s`, Port, gmail)
			http.Redirect(w, r, redirectURL, http.StatusSeeOther)

		default:
			fmt.Fprintf(w, "No other methods are supported so far.")
		}
	}
}

func (srv *Server) Welcome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		http.ServeFile(w, r, `C:\HtmlPages\html\welcome.html`)
	}
}
