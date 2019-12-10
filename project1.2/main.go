package main
import (
	"html/template"
	"net/http"
	_"sql/mysql"
	"database/sql"
	"fmt"
	"strconv"
	"github.com/satori/go.uuid"
	
)
var db *sql.DB
var err error
var i,it,itq,j,t,s int64
var no,k int64
var tpl *template.Template
var uID,sem string
var dbUsers = map[string]string{} 
var dbSessions = map[string]string{}
func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
	
}
type ds struct{
	Username string
	Passw string
	Userid string
	Per string
}
type dss struct{
	Username string
	Sem string
	Passw string
	Userid string
}
type ques struct{
	Qid ,Qno int64
	Admin,Com,Q,Sem,Sub,Code,Uids string

}
type resu struct{
	Qid,Uidr,Ans,Res,Sem string
}
type question struct{
	Qid,Qno,Qs,O1,O2,O3,O4,Ans string
}
func main(){
	itq = 0 
	i = 0
	it = 0
	t = 0
	s =0
	fs := http.FileServer(http.Dir("templates"))
  http.Handle("/", fs)
  
	http.HandleFunc("/login",login)
	http.HandleFunc("/page",page)
	http.HandleFunc("/paget",paget)
	http.HandleFunc("/addteacher",teacher)
	http.HandleFunc("/view",view)
	http.HandleFunc("/create",create)
	http.HandleFunc("/update",update)
	http.HandleFunc("/update1",update1)
	http.HandleFunc("/updated",updated)
	http.HandleFunc("/delete",delete)
	http.HandleFunc("/deleted",deleted)
	http.HandleFunc("/delete1",delete1)
	http.HandleFunc("/aas",aas)
	http.HandleFunc("/asc",asc)
	http.HandleFunc("/ads",ads)
	http.HandleFunc("/aes",aes)
	http.HandleFunc("/aes1",aes1)
	http.HandleFunc("/aesf",aesf)
	http.HandleFunc("/vs",vs)	
	http.HandleFunc("/ads1",ads1)
	http.HandleFunc("/adsf",adsf)
	http.HandleFunc("/taq",taq)
	http.HandleFunc("/tas",tas)
	http.HandleFunc("/tsc",tsc)
	http.HandleFunc("/tvs",tvs)
	http.HandleFunc("/tes",tes)
	http.HandleFunc("/tes1",tes1)
	http.HandleFunc("/tesf",tesf)
	http.HandleFunc("/tds",tds)
	http.HandleFunc("/tds1",tds1)
	http.HandleFunc("/tdsf",tdsf)
	http.HandleFunc("/gq",gq)
	http.HandleFunc("/vq",viewquestion)
	http.HandleFunc("/tevq",tevq)
	http.HandleFunc("/teq",teq)
	http.HandleFunc("/teq1",teq1)
	http.HandleFunc("/teq2",teq2)
	http.HandleFunc("/teq3",teq3)
	http.HandleFunc("/tdq",tdq)
	http.HandleFunc("/tdq1",tdq1)
	http.HandleFunc("/tdq2",tdq2)
	http.HandleFunc("/app",app)
	http.HandleFunc("/aap",aap)
	http.HandleFunc("/aap2",aap2)
	http.HandleFunc("/avq",avq)
	http.HandleFunc("/agq",agq)
	http.HandleFunc("/agq2",agq2)
	http.HandleFunc("/pages",pages)
	http.HandleFunc("/gt",gt)
	http.HandleFunc("/gt1",gt1)
	http.HandleFunc("/result",result)
	http.HandleFunc("/logout",logout)
	http.HandleFunc("/aresult",aresult)
	http.HandleFunc("/tresult",tresult)
	http.HandleFunc("/sresult",sresult)
	http.ListenAndServe(":8080", nil)

}
func tes(w http.ResponseWriter,r *http.Request){
	Uids := r.FormValue("id")
	check(Uids,r,w)
	tpl.ExecuteTemplate(w, "tes.html",Uids)
}
func tes1(w http.ResponseWriter,r *http.Request){
	Uids := r.FormValue("id")
	check(Uids,r,w)
		
		uid := r.FormValue("userid")
		uID = uid
		db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Database created successfully")
		}
		_,err = db.Exec("USE test_1")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("DB selected successfully..")
		}
		var un string
		var id string
		var se string

		
		err = db.QueryRow("select username,userid,sem from student WHERE userid =?",uid).Scan(&un,&id,&se)
		dxs := struct{
			Username string
			Sem string
			Userid string
			Uids string
						
		}{
			Username:un,
			Sem :se,
			Userid :id,
			Uids :Uids,
					
		} 
		fmt.Println(dxs)
		tpl.ExecuteTemplate(w, "tes1.html", dxs)

	
	
}
func tesf(w http.ResponseWriter, r *http.Request){
	Uids := r.FormValue("id")
	check(Uids,r,w)
	uname:=r.FormValue("username")
	sem:=r.FormValue("sem")
	passwd:=r.FormValue("passwd")
	
				db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("Database created successfully")
				}
				_,err = db.Exec("USE test_1")
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("DB selected successfully..")
				}
			INS,err := db.Prepare("update `student` set username =?,sem=?,password=? where userid = ?;")
			fmt.Println(uID)
			_,err = INS.Exec(uname,sem,passwd,uID)
			if err != nil {
				fmt.Println("error")
			}
			INSERT,err := db.Prepare("update `users` set password=? where username = ?;")
			_,err = INSERT.Exec(passwd,uID)
			if err != nil {
				tpl.ExecuteTemplate(w, "tese.html",Uids)
				return
			}
				tpl.ExecuteTemplate(w, "tesf.html",Uids)
			
			
	
	
}
func tas(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
	tpl.ExecuteTemplate(w,"tas.html",Uids)
}
func tsc(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
	uname:=r.FormValue("username")
	uid:=r.FormValue("userid")
	sem:=r.FormValue("sem")
	passwd:=r.FormValue("passwd")
	
	
		db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Database created successfully")
			}
			_,err = db.Exec("USE test_1")
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("DB selected successfully..")
			}
		INS,err := db.Prepare("INSERT  INTO `users` VALUES (?,?,'s');")
		_,err = INS.Exec(uid,passwd)
		INSERT,err := db.Prepare("INSERT  INTO `student` VALUES (?,?,?,?);")
		_,err = INSERT.Exec(uname,uid,sem,passwd)
		if err != nil {
			tpl.ExecuteTemplate(w, "tase.html",Uids)
			return
		}
		tpl.ExecuteTemplate(w, "tac.html",Uids)
	
}
func tvs(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
		var dc []dss
		db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
		if err != nil {
		fmt.Println(err.Error())
		} else {
		fmt.Println("Database created successfully")
		}
		_,err = db.Exec("USE test_1")
		if err != nil {
		fmt.Println(err.Error())
		} else {
		fmt.Println("DB selected successfully..")
		}
		rows,err := db.Query("SELECT * FROM student")
		if err == nil {
			for rows.Next() {
				var un string
				var id string
				var se string
				var pass string
				err = rows.Scan(&un,&id,&se,&pass)
				dc = append(dc,dss{Username:un,Userid:id,Sem:se,Passw:pass})
			}
		}
		d := struct{
			Uids string
			D []dss
		}{
			D:dc,
			Uids:Uids,
		}
		
		
		fmt.Println(d)
	tpl.ExecuteTemplate(w, "tvs.html",d)
		
	
}
func login(w http.ResponseWriter, r *http.Request){
		
	
	tpl.ExecuteTemplate(w, "log.html",nil)
	
	
}
func create(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
	uname:=r.FormValue("username")
	uid:=r.FormValue("userid")
	p:=r.FormValue("permission")
	passwd:=r.FormValue("passwd")
	passwd1:=r.FormValue("passwd1")
	fmt.Println(passwd1)
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Database created successfully")
			}
			_,err = db.Exec("USE test_1")
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("DB selected successfully..")
			}
		INS,err := db.Prepare("INSERT  INTO `users` VALUES (?,?,?);")
		_,err = INS.Exec(uid,passwd,p)
		if err != nil {
			tpl.ExecuteTemplate(w, "ateach1.html",Uids)
			return
		}
		INSERT,err := db.Prepare("INSERT  INTO `teacher` VALUES (?,?,?,?);")
		_,err = INSERT.Exec(uname,uid,passwd,p)
		if err != nil {
			tpl.ExecuteTemplate(w, "ateach1.html",Uids)
			return
		}
		tpl.ExecuteTemplate(w, "created.html",Uids)	
	
	
}
func view(w http.ResponseWriter, r *http.Request){
		id:=r.FormValue("id")
		check(id,r,w)
		var dz []ds
		db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
		if err != nil {
		fmt.Println(err.Error())
		} else {
		fmt.Println("Database created successfully")
		}
		_,err = db.Exec("USE test_1")
		if err != nil {
		fmt.Println(err.Error())
		} else {
		fmt.Println("DB selected successfully..")
		}
		rows,err := db.Query("SELECT * FROM teacher")
		if err == nil {
			for rows.Next() {
				var un string
				var id string
				var pass string
				var p string
				err = rows.Scan(&un,&id,&pass,&p)
				dz = append(dz,ds{Username:un,Userid:id,Passw:pass,Per:p})
			}
		}
		d := struct{
			Df []ds
			Uids string

		}{
			Df : dz ,
			Uids : id,
		}
		fmt.Println(d)
		tpl.ExecuteTemplate(w, "view.html", d)	
		

}
func teacher(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
	
	check(Uids,r,w)
			tpl.ExecuteTemplate(w, "ateach.html",Uids)	

}
func page(w http.ResponseWriter,r *http.Request){
	


	if i == 1 {
		Uids := r.FormValue("id")
		if Uids == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		fmt.Println(Uids)
		c, _ := r.Cookie(Uids)
		if c.Name!= Uids{
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		
		tpl.ExecuteTemplate(w, "page.html", Uids)	
		return
	}
	
		Uids := r.FormValue("userid")
		if Uids == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		passwd := r.FormValue("password")
		var pass string
		var p string
		db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Database created successfully")
		}
		_,err = db.Exec("USE test_1")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("DB selected successfully..")
		}
		err = db.QueryRow("select password,p from users WHERE username =?",Uids).Scan(&pass,&p)

		if err != nil {
			fmt.Println("fail1")
			tpl.ExecuteTemplate(w, "log1.html",nil)
			return
		}
		fmt.Println(pass)
		fmt.Println(p)
		

		if pass == passwd {
			sID, _ := uuid.NewV4()
				c := &http.Cookie{
					Name:  Uids,
					Value: sID.String(),
				}
				http.SetCookie(w, c)
		dbSessions[c.Value] = Uids
		
			if p == "a"{
				i = 1
				
				tpl.ExecuteTemplate(w, "page.html", Uids)
				return
			}else if p == "t"{
				it=1
				
				tpl.ExecuteTemplate(w,"paget.html",Uids)
				return
			}else if p == "s"{
				s =1
				
				tpl.ExecuteTemplate(w,"pages.html",Uids)
				return
			}
		}
		fmt.Println("fail")
		tpl.ExecuteTemplate(w, "log1.html",nil)	
		
	
}
func paget(w http.ResponseWriter, r *http.Request ){
	
		Uids := r.FormValue("id")
		
		check(Uids,r,w)

	tpl.ExecuteTemplate(w,"paget.html",Uids)
}
func update(w http.ResponseWriter,r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
	tpl.ExecuteTemplate(w, "update.html",Uids)
}
func update1(w http.ResponseWriter,r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
		uid := r.FormValue("userid")
		uID = uid
		db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Database created successfully")
		}
		_,err = db.Exec("USE test_1")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("DB selected successfully..")
		}
		var un string
		var id string

		
		err = db.QueryRow("select username,userid from teacher WHERE userid =?",uid).Scan(&un,&id,)
		if err!=nil {
			tpl.ExecuteTemplate(w, "update.html", Uids)
		}
		dx := struct{
			Username string
			Userid string
			Uids string
						
		}{
			Username:un,
			Userid :id,
			Uids:Uids,
					
		} 
		fmt.Println(dx)
		tpl.ExecuteTemplate(w, "update1.html", dx)
		return


}
func updated(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
	uname:=r.FormValue("username")
	p:=r.FormValue("permission")
	passwd:=r.FormValue("passwd")
	
	
				db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("Database created successfully")
				}
				_,err = db.Exec("USE test_1")
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("DB selected successfully..")
				}
			INS,err := db.Prepare("update `teacher` set username =?,password=?,p=? where userid = ?;")
			fmt.Println(uID)
			_,err = INS.Exec(uname,passwd,p,uID)
			if err != nil {
				tpl.ExecuteTemplate(w, "update2.html",Uids)
			}
			INSERT,err := db.Prepare("update `users` set password=?,p=? where username = ?;")
			_,err = INSERT.Exec(passwd,p,uID)
			if err == nil {
				tpl.ExecuteTemplate(w, "updated.html",Uids)
			} else {	
	tpl.ExecuteTemplate(w, "update2.html",Uids)
	}
	
}
func delete(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
	tpl.ExecuteTemplate(w, "delete.html",Uids)	
}
func delete1(w http.ResponseWriter,r *http.Request){
	
	Uids:=r.FormValue("id")
	check(Uids,r,w)
		uid := r.FormValue("userid")
		uID = uid
		db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Database created successfully")
		}
		_,err = db.Exec("USE test_1")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("DB selected successfully..")
		}
		var un string
		var id string
		var p string
		err = db.QueryRow("select username,userid,p from teacher WHERE userid =?",uid).Scan(&un,&id,&p)
		dx := struct{
			Username string
			Passw string
			Userid string
			Uids string
						
		}{
			Username:un,
			Userid :id,
			Passw :	p,
			Uids:Uids,
		} 
		fmt.Println(dx)
		tpl.ExecuteTemplate(w, "delete1.html", dx)

	
}
func deleted(w http.ResponseWriter, r *http.Request){
		Uids:=r.FormValue("id")
		check(Uids,r,w)
				db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("Database created successfully")
				}
				_,err = db.Exec("USE test_1")
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("DB selected successfully..")
				}
			
			INSERT,err := db.Prepare("DELETE FROM `teacher` WHERE userid = ?;")
			
			_,err = INSERT.Exec(uID)
			if err != nil {
				tpl.ExecuteTemplate(w, "delete2.html",Uids)	
			}
			INS,err := db.Prepare("DELETE FROM `users` WHERE username = ?;")
			fmt.Println(uID)
			_,err = INS.Exec(uID)
			if err != nil {
				tpl.ExecuteTemplate(w, "delete2.html",Uids)	
			}
			tpl.ExecuteTemplate(w, "deleted.html",Uids)
	
}
func aas(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
		check(Uids,r,w)
	tpl.ExecuteTemplate(w,"aastudent.html",Uids)
}
func asc(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
	uname:=r.FormValue("username")
	uid:=r.FormValue("userid")
	sem:=r.FormValue("sem")
	passwd:=r.FormValue("passwd")
	passwd1:=r.FormValue("passwd1")
	if passwd == passwd1 {
		db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Database created successfully")
			}
			_,err = db.Exec("USE test_1")
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("DB selected successfully..")
			}
		INS,err := db.Prepare("INSERT  INTO `users` VALUES (?,?,'s');")
		_,err = INS.Exec(uid,passwd)
		INSERT,err := db.Prepare("INSERT  INTO `student` VALUES (?,?,?,?);")
		_,err = INSERT.Exec(uname,uid,sem,passwd)
		if err == nil {
			tpl.ExecuteTemplate(w, "asc.html",Uids)
			}
	} else {
	tpl.ExecuteTemplate(w, "aastudent1.html",Uids)
	}	
	
}
func vs(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
		check(Uids,r,w)
		var stud []dss
		db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
		if err != nil {
		fmt.Println(err.Error())
		} else {
		fmt.Println("Database created successfully")
		}
		_,err = db.Exec("USE test_1")
		if err != nil {
		fmt.Println(err.Error())
		} else {
		fmt.Println("DB selected successfully..")
		}
		rows,err := db.Query("SELECT * FROM student")
		if err == nil {
			for rows.Next() {
				var un string
				var id string
				var se string
				var pass string
				err = rows.Scan(&un,&id,&se,&pass)
				stud = append(stud,dss{Username:un,Userid:id,Sem:se,Passw:pass})
			}
		}
		d := struct{
			Uids string
			D []dss
		}{
			Uids :Uids,
			D:stud,
		}
		fmt.Println(d)
		tpl.ExecuteTemplate(w, "vs.html",d)
			
		
}
func aes(w http.ResponseWriter,r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
	tpl.ExecuteTemplate(w, "aes.html",Uids)
}
func ads(w http.ResponseWriter,r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
	tpl.ExecuteTemplate(w, "ads.html",Uids)
}
func aes1(w http.ResponseWriter,r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
		
		uid := r.FormValue("userid")
		uID = uid
		db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Database created successfully")
		}
		_,err = db.Exec("USE test_1")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("DB selected successfully..")
		}
		var un string
		var id string
		var se string

		
		err = db.QueryRow("select username,userid,sem from student WHERE userid =?",uid).Scan(&un,&id,&se)
		dxs := struct{
			Username string
			Sem string
			Userid string
			Uids string
						
		}{
			Username:un,
			Sem :se,
			Userid :id,
			Uids:Uids,		
		} 
		fmt.Println(dxs)
		tpl.ExecuteTemplate(w, "aes1.html", dxs)

	
}
func aesf(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
	uname:=r.FormValue("username")
	sem:=r.FormValue("sem")
	passwd:=r.FormValue("passwd")
			db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("Database created successfully")
				}
				_,err = db.Exec("USE test_1")
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("DB selected successfully..")
				}
			INS,err := db.Prepare("update `student` set username =?,sem=?,password=? where userid = ?;")
			fmt.Println(uID)
			_,err = INS.Exec(uname,sem,passwd,uID)
			if err != nil {
				tpl.ExecuteTemplate(w, "aese.html",Uids)
			}
			INSERT,err := db.Prepare("update `users` set password=? where username = ?;")
			_,err = INSERT.Exec(passwd,uID)
			if err == nil {
				tpl.ExecuteTemplate(w, "aesf.html",Uids)
		} else {	
	tpl.ExecuteTemplate(w, "aese.html",Uids)
	}
	
}
func ads1(w http.ResponseWriter,r *http.Request){		
	Uids:=r.FormValue("id")
	check(Uids,r,w)	
	uid := r.FormValue("userid")
		uID = uid
		db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Database created successfully")
		}
		_,err = db.Exec("USE test_1")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("DB selected successfully..")
		}
		var un string
		var id string
		var se string
		err = db.QueryRow("select username,userid,sem from student WHERE userid =?",uid).Scan(&un,&id,&se)
		dx := struct{
			Username string
			Sem string
			Userid string
			Uids string			
		}{
			Username:un,
			Sem :se,
			Userid :id,
			Uids : Uids,
		} 
		fmt.Println(dx)
		tpl.ExecuteTemplate(w, "ads1.html", dx)

	
}
func adsf(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
				db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("Database created successfully")
				}
				_,err = db.Exec("USE test_1")
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("DB selected successfully..")
				}
			
			INSERT,err := db.Prepare("DELETE FROM `student` WHERE userid = ?;")
			
			_,err = INSERT.Exec(uID)
			if err != nil {
				tpl.ExecuteTemplate(w, "ads2.html",Uids)	
			}
			_,err = db.Exec("DELETE FROM `users` WHERE username = ? ;",uID)
			fmt.Println(uID)
			if err != nil {
				tpl.ExecuteTemplate(w, "ads2.html",Uids)
			}
			_,err = db.Exec("DELETE FROM `result` WHERE uid = ? ;",uID)
			fmt.Println(uID)
			if err != nil {
				tpl.ExecuteTemplate(w, "ads2.html",Uids)
			}
			tpl.ExecuteTemplate(w, "adsf.html",Uids)
		
	
}
func taq(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)	
	if itq == 0 {
			qp:=r.FormValue("qp")
			sem:=r.FormValue("sem")
			sub:=r.FormValue("sub")
			sc:=r.FormValue("sc")
			nq:=r.FormValue("nq")
			fmt.Println(nq)
			no,err = strconv.ParseInt(nq, 10, 64)
			db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Database created successfully")
			}
			_,err = db.Exec("USE test_1")
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("DB selected successfully..")
			}
			var i string
			var x int64
			var f string 
			err = db.QueryRow("select q_id from question_paper where comp=0;").Scan(&f)
			if f !="null"	{
						INSERTf,err := db.Prepare(" Call deleteincomp(?);")
			if err != nil{
				fmt.Println(err)
			}
			
			_,err = INSERTf.Exec(f)
			}
			err = db.QueryRow("select max(q_id) from question_paper").Scan(&i)
			if i == "null" {
				x=0
			}else {
				x,err=strconv.ParseInt(i, 10, 64)
				x++
				fmt.Println(x)
			}
			INS,err := db.Prepare("INSERT  INTO `question_paper` VALUES (?,?,?,?,?,?,false,false,false);")
			_,err = INS.Exec(x,qp,nq,sem,sub,sc)
			if err != nil {
				fmt.Println("error")
				tpl.ExecuteTemplate(w,"generatepaper.html",nil)
			}
			k=1
			q := struct {
				Qp, Nq, Sem, Sub, Sc string
				Z,K int64
				Uids string
			}{
				Qp: qp,
				Nq:  nq,
				Sem: sem,
				Sub:  sub,
				Sc: sc,
				Z: x,
				K:k,
				Uids:Uids,
			}
			itq =1
			fmt.Println("error")
			tpl.ExecuteTemplate(w, "taq.html",q)
			return
		}else if itq ==1 {
				if k != (no+1)  {
				ques:=r.FormValue("question")
				o1:=r.FormValue("o1")
				o2:=r.FormValue("o2")
				o3:=r.FormValue("o3")
				o4:=r.FormValue("o4")
				ans:=r.FormValue("ans")
				db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("Database created successfully")
				}
				_,err = db.Exec("USE test_1;")
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("DB selected successfully..")
				}
				var i,qp,sem,sub,sc string
				var nq int64
				var x int64
				err = db.QueryRow("select max(q_id) from question_paper;").Scan(&i)
					x,err=strconv.ParseInt(i, 10, 64)
					fmt.Println(x)
					err = db.QueryRow("select q_paper,no_ques,sem,sub,sub_code from question_paper where q_id = ?",x).Scan(&qp,&nq,&sem,&sub,&sc)	
					INS,err := db.Prepare("INSERT  INTO `question` VALUES (?,?,?,?,?,?,?,?);")
					_,err = INS.Exec(x,k,ques,o1,o2,o3,o4,ans)
					k++		
					q := struct {
						Qp, Sem, Sub, Sc,Uids string
						Nq int64
						Z,K int64
					}{
						Qp: qp,
						Nq:  nq,
						Sem: sem,
						Sub:  sub,
						Sc: sc,
						Z: x,
						K:k,
						Uids:Uids,
					}
					fmt.Println(q)
					if k == no {
						tpl.ExecuteTemplate(w, "taq.html",q)
						return	
					} else if k > no{
						INS,err = db.Prepare("update `question_paper` set comp=true where q_id = ?;")
						_,err = INS.Exec(x)
						itq=0
						if err == nil{
							tpl.ExecuteTemplate(w,"tpg.html",Uids)
							return
						}
					}
					tpl.ExecuteTemplate(w, "taq.html",q)
					return	
				}
				
			}
}
func tds(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
	tpl.ExecuteTemplate(w, "tds.html",Uids)
}
func tds1(w http.ResponseWriter,r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
		uid := r.FormValue("userid")
		uID = uid
		db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Database created successfully")
		}
		_,err = db.Exec("USE test_1")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("DB selected successfully..")
		}
		var un string
		var id string
		var se string
		err = db.QueryRow("select username,userid,sem from student WHERE userid =?",uid).Scan(&un,&id,&se)
		dx := struct{
			Username string
			Sem string
			Userid string
			Uids string
						
		}{
			Username:un,
			Sem :se,
			Userid :id,
			Uids:Uids,
		} 
		fmt.Println(dx)
		tpl.ExecuteTemplate(w, "tds1.html", dx)
	
}
func tdsf(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
				db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("Database created successfully")
				}
				_,err = db.Exec("USE test_1")
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("DB selected successfully..")
				}
			
			INSERT,err := db.Prepare("DELETE FROM `student` WHERE userid = ?;")
			
			_,err = INSERT.Exec(uID)
			if err != nil {
				fmt.Println("err1")
				tpl.ExecuteTemplate(w, "tdse.html",Uids)
			}
			_,err = db.Exec("DELETE FROM `users` WHERE username = ? ;",uID)
			fmt.Println(uID)
			if err != nil {
				fmt.Println("err2")
				tpl.ExecuteTemplate(w, "tdse.html",Uids)
			}
			_,err = db.Exec("DELETE FROM `result` WHERE uid = ? ;",uID)
			fmt.Println(uID)
			if err != nil {
				fmt.Println("err3")
				tpl.ExecuteTemplate(w, "tdse.html",Uids)
			}
			tpl.ExecuteTemplate(w, "tdsf.html",Uids)
		
	
}
func gq(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
	itq =0
	tpl.ExecuteTemplate(w, "generatepaper.html",Uids)
	return
}
func viewquestion(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
		
	
		var dc []ques
		db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
		if err != nil {
		fmt.Println(err.Error())
		} else {
		fmt.Println("Database created successfully")
		}
		_,err = db.Exec("USE test_1")
		if err != nil {
		fmt.Println(err.Error())
		} else {
		fmt.Println("DB selected successfully..")
		}
		rows,err := db.Query("SELECT q_id,q_paper,no_ques,sem,sub,sub_code,admin,comp FROM question_paper")
		if err == nil {
			for rows.Next() {
				var paper,sem,sub,code,ad,comp string
				var id,no int64
				err = rows.Scan(&id,&paper,&no,&sem,&sub,&code,&ad,&comp)
				if ad == "1"{
					ad = "approved"
				}else {
					ad = "pending"
				}
				dc = append(dc,ques{Qid:id,Qno:no,Admin:ad,Com:comp,Sem:sem,Sub:sub,Code:code,Q:paper,Uids:Uids})
			}
		}
		d:=struct{
			D []ques
			Uids string
		}{
			Uids :Uids,
			D:dc,
		}
		fmt.Println(d)
		tpl.ExecuteTemplate(w, "viewq.html",d)
		
		
}
func tevq(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
	name := r.FormValue("name")
	fmt.Println(name)
	var dc []question
		db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
		if err != nil {
		fmt.Println(err.Error())
		} else {
		fmt.Println("Database created successfully")
		}
		_,err = db.Exec("USE test_1")
		if err != nil {
		fmt.Println(err.Error())
		} else {
		fmt.Println("DB selected successfully..")
		}
		rows,err := db.Query("SELECT q_id,nq,question,op1,op2,op3,op4,ans FROM question where q_id = ?",name)
		if err == nil {
			for rows.Next() {
				var id,no,qu,o1,o2,o3,o4,ans string
				
				err = rows.Scan(&id,&no,&qu,&o1,&o2,&o3,&o4,&ans)
				
				dc = append(dc,question{Qid:id,Qno:no,Qs:qu,O1:o1,O2:o2,O3:o3,O4:o4,Ans:ans})
			}
		}
		d:=struct{
			D []question
			Uids string
		}{
			D:dc,
			Uids:Uids,
		}
		fmt.Println(d)
		tpl.ExecuteTemplate(w,"tvq.html",d)
}
func teq(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
	tpl.ExecuteTemplate(w,"teq.html",Uids)
}
func teq1(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
	filter:=r.FormValue("userid")
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println(err.Error())
	} else {
	fmt.Println("Database created successfully")
	}
	_,err = db.Exec("USE test_1")
	if err != nil {
	fmt.Println(err.Error())
	} else {
	fmt.Println("DB selected successfully..")
	}
	fmt.Println(filter)
	var id,paper,no,sem,sub,code string
	err = db.QueryRow("SELECT q_id,q_paper,no_ques,sem,sub,sub_code FROM question_paper where q_id = ?;", filter).Scan(&id,&paper,&no,&sem,&sub,&code)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)
	d := struct {
		Qid , Paper, No, Sem, Sub, Code,Uids string
	}{
		Qid:id,
		Paper:paper,
		No:no ,
		Sem:sem,
		Sub:sub,
		Code:code,
		Uids:Uids,
	}
	fmt.Println(d)
	tpl.ExecuteTemplate(w,"teq1.html",d)
}
func teq2(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
	filter:=r.FormValue("name")
	fil:=r.FormValue("nq")
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println(err.Error())
	} else {
	fmt.Println("Database created successfully")
	}
	_,err = db.Exec("USE test_1")
	if err != nil {
	fmt.Println(err.Error())
	} else {
	fmt.Println("DB selected successfully..")
	}
	fmt.Println(filter)
	var id,no,ques,op1,op2,op3,op4,ans string
	err = db.QueryRow("SELECT q_id,nq,question,op1,op2,op3,op4,ans FROM question where q_id = ? and nq = ?;", filter,fil).Scan(&id,&no,&ques,&op1,&op2,&op3,&op4,&ans)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)
	d := struct {
		Qid,No,Ques,Op1,Op2,Op3,Op4,Ans,Uids string
	}{
		Qid:id,
		No:no ,
		Ques:ques,
		Op1:op1,
		Op2:op2,
		Op3:op3,
		Op4:op4,
		Ans:ans,
		Uids:Uids,
	}
	fmt.Println(d)
	tpl.ExecuteTemplate(w,"teq2.html",d)
}
func teq3(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
	id:=r.FormValue("name")
	fmt.Println(id)
	no:=r.FormValue("num")
	fmt.Println(no)
	ques:=r.FormValue("question")
	o1:=r.FormValue("o1")
	o2:=r.FormValue("o2")
	o3:=r.FormValue("o3")
	o4:=r.FormValue("o4")
	ans:=r.FormValue("ans")
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Database created successfully")
	}
	_,err = db.Exec("USE test_1;")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("DB selected successfully..")
	}
	INS,err := db.Prepare("update `question` set question=?,op1=?,op2=?,op3=?,op4=?,ans=? where q_id=? and nq=?;")
	if err != nil{
		fmt.Println(err)
	}
	_,err = INS.Exec(ques,o1,o2,o3,o4,ans,id,no)
	INSa,err := db.Prepare("update `question_paper` set admin=false where q_id=?;")	
	_,err = INSa.Exec(id)
	tpl.ExecuteTemplate(w,"teq3.html",Uids)
}
func tdq(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
	tpl.ExecuteTemplate(w,"tdq.html",Uids)
}
func tdq1(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
	filter:=r.FormValue("userid")
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println(err.Error())
	} else {
	fmt.Println("Database created successfully")
	}
	_,err = db.Exec("USE test_1")
	if err != nil {
	fmt.Println(err.Error())
	} else {
	fmt.Println("DB selected successfully..")
	}
	fmt.Println(filter)
	var id,paper,no,sem,sub,code string
	err = db.QueryRow("SELECT q_id,q_paper,no_ques,sem,sub,sub_code FROM question_paper where q_id = ?;", filter).Scan(&id,&paper,&no,&sem,&sub,&code)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)
	d := struct {
		Qid , Paper, No, Sem, Sub, Code,Uids string
	}{
		Qid:id,
		Paper:paper,
		No:no ,
		Sem:sem,
		Sub:sub,
		Code:code,
		Uids:Uids,
	}
	fmt.Println(d)
	tpl.ExecuteTemplate(w,"tdq1.html",d)
}
func tdq2(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
	id:=r.FormValue("name")
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println(err.Error())
	} else {
	fmt.Println("Database created successfully")
	}
	_,err = db.Exec("USE test_1")
	if err != nil {
	fmt.Println(err.Error())
	} else {
	fmt.Println("DB selected successfully..")
	}
	INSERTf,err := db.Prepare(" Call deleteincomp(?);")
			if err != nil{
				fmt.Println(err)
			}
			
			_,err = INSERTf.Exec(id)
	tpl.ExecuteTemplate(w, "tdq2.html",Uids)
	
}
func app(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
		
		var dq []ques
		db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
		if err != nil {
		fmt.Println(err.Error())
		} else {
		fmt.Println("Database created successfully")
		}
		_,err = db.Exec("USE test_1")
		if err != nil {
		fmt.Println(err.Error())
		} else {
		fmt.Println("DB selected successfully..")
		}
		rows,err := db.Query("SELECT q_id,q_paper,no_ques,sem,sub,sub_code,admin,comp FROM question_paper")
		if err == nil {
			for rows.Next() {
				var paper,sem,sub,code,ad,comp string
				var id,no int64
				err = rows.Scan(&id,&paper,&no,&sem,&sub,&code,&ad,&comp)
				if ad == "1"{
					ad = "approved"
				}else {
					ad = "pending"
				}
				dq = append(dq,ques{Qid:id,Qno:no,Admin:ad,Com:comp,Sem:sem,Sub:sub,Code:code,Q:paper,Uids:Uids})
			}
		}
		d := struct{
			D []ques
			Uids string
		}{
			D : dq,
			Uids:Uids,
		}
		fmt.Println(d)
		tpl.ExecuteTemplate(w, "app.html",d)
		

}
func aap(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
	qid:=r.FormValue("name")
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println(err.Error())
	} else {
	fmt.Println("Database created successfully")
	}
	_,err = db.Exec("USE test_1")
	if err != nil {
	fmt.Println(err.Error())
	} else {
	fmt.Println("DB selected successfully..")
	}
	
	var id,paper,no,sem,sub,code string
	err = db.QueryRow("SELECT q_id,q_paper,no_ques,sem,sub,sub_code FROM question_paper where q_id = ?;", qid).Scan(&id,&paper,&no,&sem,&sub,&code)
	if err != nil {
		fmt.Println(err)
	}
	
	d := struct {
		Qid , Paper, No, Sem, Sub, Code,Uids string
	}{
		Qid:id,
		Paper:paper,
		No:no ,
		Sem:sem,
		Sub:sub,
		Code:code,
		Uids:Uids,
	}
	fmt.Println(d)
	tpl.ExecuteTemplate(w,"aap.html",d)
}
func aap2(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
	qid :=r.FormValue("name")
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Database created successfully")
	}
	_,err = db.Exec("USE test_1")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("DB selected successfully..")
	}
	INSERT,err := db.Prepare("update `question_paper` set admin=true where q_id = ?;")
	_,err = INSERT.Exec(qid)
	
	tpl.ExecuteTemplate(w, "aap2.html",Uids)
	
}
func avq(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
	name := r.FormValue("name")
	fmt.Println(name)
	var dc []question
		db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
		if err != nil {
		fmt.Println(err.Error())
		} else {
		fmt.Println("Database created successfully")
		}
		_,err = db.Exec("USE test_1")
		if err != nil {
		fmt.Println(err.Error())
		} else {
		fmt.Println("DB selected successfully..")
		}
		rows,err := db.Query("SELECT q_id,nq,question,op1,op2,op3,op4,ans FROM question where q_id = ?",name)
		if err == nil {
			for rows.Next() {
				var id,no,qu,o1,o2,o3,o4,ans string
				
				err = rows.Scan(&id,&no,&qu,&o1,&o2,&o3,&o4,&ans)
				
				dc = append(dc,question{Qid:id,Qno:no,Qs:qu,O1:o1,O2:o2,O3:o3,O4:o4,Ans:ans})
			}
		}
		d:= struct{
			Uids string
			D []question
		}{
			D:dc,
			Uids:Uids,
		}
		fmt.Println(d)
		tpl.ExecuteTemplate(w,"avq.html",d)
}
func agq(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
	qid:=r.FormValue("name")
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println(err.Error())
	} else {
	fmt.Println("Database created successfully")
	}
	_,err = db.Exec("USE test_1")
	if err != nil {
	fmt.Println(err.Error())
	} else {
	fmt.Println("DB selected successfully..")
	}
	
	var id,paper,no,sem,sub,code,ad string
	err = db.QueryRow("SELECT q_id,q_paper,no_ques,sem,sub,sub_code,student FROM question_paper where q_id = ?;", qid).Scan(&id,&paper,&no,&sem,&sub,&code,&ad)
	if err != nil {
		fmt.Println(err)
	}
	if ad == "1"{
		ad = "approved"
	}else {
		ad = "pending"
	}
	d := struct {
		Qid , Paper, No, Sem, Sub, Code,Ad,Uids string
	}{
		Qid:id,
		Paper:paper,
		No:no ,
		Sem:sem,
		Sub:sub,
		Code:code,
		Ad:ad,
		Uids:Uids,
	}
	fmt.Println(d)
	tpl.ExecuteTemplate(w,"agq.html",d)
}
func agq2(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
	qid :=r.FormValue("name")
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Database created successfully")
	}
	_,err = db.Exec("USE test_1")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("DB selected successfully..")
	}
	var ad string
	err = db.QueryRow("SELECT student FROM question_paper where q_id = ?;", qid).Scan(&ad)
	if ad == "1"{
		INSERTs,_ := db.Prepare("update `question_paper` set student=false where q_id = ?;")
		_,err = INSERTs.Exec(qid)
	}else {
		INSERT,_ := db.Prepare("update `question_paper` set student=true where q_id = ?;")
		_,err = INSERT.Exec(qid)
	} 
	
	tpl.ExecuteTemplate(w, "agp2.html",Uids)
	
}
func pages(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
	tpl.ExecuteTemplate(w,"pages.html",Uids)
}
func gt(w http.ResponseWriter, r *http.Request){
	Uids:=r.FormValue("id")
	check(Uids,r,w)
	var sem string
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Database created successfully")
	}
	_,err = db.Exec("USE test_1")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("DB selected successfully..")
	}
	err = db.QueryRow("select sem from student WHERE userid =?",Uids).Scan(&sem)
	fmt.Println(sem)
	var qid,qid2,paper,no,sem1,sub string
	err = db.QueryRow("SELECT q_id,q_paper,no_ques,sem,sub FROM question_paper where sem = ? and comp = true and admin = true and student = true;",sem).Scan(&qid,&paper,&no,&sem1,&sub)
	if qid == "" {
		tpl.ExecuteTemplate(w,"nulla.html",Uids)
		return
	}
	err = db.QueryRow("SELECT q_id FROM result where uid=?;",Uids).Scan(&qid2)
	if qid2 == qid {
		tpl.ExecuteTemplate(w,"nulla.html",Uids)
		return
	}
	
	d := struct {
		Qid, Uids, Paper, No, Sem, Sub string
	}{
		Qid:qid,
		Uids:Uids,
		Paper:paper,
		No:no ,
		Sem:sem,
		Sub:sub,
		
	}
	fmt.Println(d)
	tpl.ExecuteTemplate(w,"gt1.html",d)
}
func gt1(w http.ResponseWriter, r *http.Request){
	
	id := r.FormValue("id")
	qid := r.FormValue("qid")
	if t == 1 {
		
		tpl.ExecuteTemplate(w,"pages.html",id)
		return
	}
	var d []question
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
	if err != nil {
	fmt.Println(err.Error())
	} else {
	fmt.Println("Database created successfully")
	}
	_,err = db.Exec("USE test_1")
	if err != nil {
	fmt.Println(err.Error())
	} else {
	fmt.Println("DB selected successfully..")
	}
	rows,err := db.Query("SELECT q_id,nq,question,op1,op2,op3,op4,ans FROM question where q_id = ?",qid)
	if err == nil {
		for rows.Next() {
			var id,no,qu,o1,o2,o3,o4,ans string
			
			err = rows.Scan(&id,&no,&qu,&o1,&o2,&o3,&o4,&ans)
			
			d = append(d,question{Qid:id,Qno:no,Qs:qu,O1:o1,O2:o2,O3:o3,O4:o4,Ans:ans})
		}
	}
	err = db.QueryRow("SELECT no_ques FROM question_paper where q_id = ?;",qid).Scan(&k)
	
	df := struct {
		Dl []question
		Uids,Qid string
		Time int64
	}{
		Dl:d,
		Time:(k*60), 
		Uids:id,
		Qid:qid,
	}
	fmt.Println(df)
	tpl.ExecuteTemplate(w,"test.html",df)

}
func result(w http.ResponseWriter, r *http.Request){
	t=1
	Uids :=r.FormValue("id")
	qid :=r.FormValue("qid")

	var f string
	f = ""
	fmt.Println(f)
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
	if err != nil {
	fmt.Println(err.Error())
	} else {
	fmt.Println("Database created successfully")
	}
	_,err = db.Exec("USE test_1")
	if err != nil {
	fmt.Println(err.Error())
	} else {
	fmt.Println("DB selected successfully..")
	}
	err = db.QueryRow("SELECT no_ques,sem FROM question_paper where q_id = ?;",qid).Scan(&k,&sem)
	var i int64
	for i = 1; i <= k; i++ {
		
		x:=r.FormValue(strconv.FormatInt(i, 10))
		if x == ""{
			f=f+"0"
		} else {
			f=f+x
		}

	}
	
	var ans,ab string
	ans = ""
	i = 0
	rows,err := db.Query("SELECT ans FROM question where q_id = ? ",qid)
	if err == nil {
		for rows.Next() {
			err = rows.Scan(&ab)
			ans=ans+ab	
		}
	}
	fmt.Println(ans)
	fmt.Println(f)
	i=0
	for j:=range f {
		if f[j]==ans[j]{
		i++
		}
	}
	fmt.Println(i)
	INS,err := db.Prepare("INSERT  INTO `result` VALUES (?,?,?,?,?);")
	_,err = INS.Exec(qid,f,i,Uids,sem)
	d := struct{
		Uids string
		I int64
	}{
		Uids:Uids,
		I:i,
	}
	tpl.ExecuteTemplate(w,"result.html",d)
}
func aresult(w http.ResponseWriter, r *http.Request){
	Uids := r.FormValue("id")
		
	check(Uids,r,w)

	var dc []resu
		db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
		if err != nil {
		fmt.Println(err.Error())
		} else {
		fmt.Println("Database created successfully")
		}
		_,err = db.Exec("USE test_1")
		if err != nil {
		fmt.Println(err.Error())
		} else {
		fmt.Println("DB selected successfully..")
		}
		rows,err := db.Query("SELECT * FROM result")
		if err == nil {
			for rows.Next() {
				var uid,qid,res,ans,sem string
				err = rows.Scan(&qid,&ans,&res,&uid,&sem)
				dc = append(dc,resu{Qid:qid,Uidr:uid,Ans:ans,Res:res,Sem:sem })
				
			}
		}
		d:=struct{
			D []resu
			Uids string
		}{
			Uids:Uids,
			D:dc,
		}
		fmt.Println(d)
		tpl.ExecuteTemplate(w, "aresult.html",d)
			
		
}
func logout(w http.ResponseWriter, req *http.Request) {
	Uids:=req.FormValue("id")
	
		c, _ := req.Cookie(Uids)
		fmt.Println(c.Name)
	// delete the session
	dbSessions[c.Value]=""
	// remove the cookie
	c = &http.Cookie{
		Name:   Uids,
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	i =0
	it = 0
	http.Redirect(w, req, "/login", http.StatusSeeOther)
	return
	
	
}
func check(Uids string, r *http.Request,w http.ResponseWriter) bool{
	fmt.Println(Uids)
		
		c, err := r.Cookie(Uids)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return	true
		}
		if c.Name != Uids{
			fmt.Println("error")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		return true
}
func tresult(w http.ResponseWriter, r *http.Request){
	Uids := r.FormValue("id")
		
	check(Uids,r,w)

		var dc []resu
		db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
		if err != nil {
		fmt.Println(err.Error())
		} else {
		fmt.Println("Database created successfully")
		}
		_,err = db.Exec("USE test_1")
		if err != nil {
		fmt.Println(err.Error())
		} else {
		fmt.Println("DB selected successfully..")
		}
		rows,err := db.Query("SELECT * FROM result")
		if err == nil {
			for rows.Next() {
				var uid,qid,res,ans,sem string
				err = rows.Scan(&qid,&ans,&res,&uid,&sem)
				dc = append(dc,resu{Qid:qid,Uidr:uid,Ans:ans,Res:res,Sem:sem })
				
			}
		}
		d:=struct{
			Uids string
			D []resu
		}{
			Uids:Uids,
			D:dc,
		}
		fmt.Println(d)
		tpl.ExecuteTemplate(w, "tresult.html",d)
}
func sresult(w http.ResponseWriter, r *http.Request){
	id := r.FormValue("id")
	check(id,r,w)
	var df []resu
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
		if err != nil {
		fmt.Println(err.Error())
		} else {
		fmt.Println("Database created successfully")
		}
		_,err = db.Exec("USE test_1")
		if err != nil {
		fmt.Println(err.Error())
		} else {
		fmt.Println("DB selected successfully..")
		}
		rows,err := db.Query("SELECT * FROM result where uid = ? ",id)
		if err == nil {
			for rows.Next() {
				var uid,qid,res,ans,sem string
				err = rows.Scan(&qid,&ans,&res,&uid,&sem)
				df = append(df,resu{Qid:qid,Uidr:uid,Ans:ans,Res:res,Sem:sem })
				
			}
		}
		fmt.Println(df)
		d := struct{
			Dl []resu
			Uids string
		}{
			Uids:id,
			Dl :df,
		}
		tpl.ExecuteTemplate(w,"sresult.html",d)

}