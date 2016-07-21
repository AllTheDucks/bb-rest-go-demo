package main

import "golang.org/x/oauth2"
import "golang.org/x/oauth2/clientcredentials"
import "log"
import "flag"
import "os"
import "net/http"
import "encoding/json"
import "html/template"
import "fmt"
import "time"

var serverRoot string
var appKey string
var appSecret string
var tokenUrl string
var usersUrl string
var coursesUrl string

var client *http.Client
var courseService CourseService

func init() {
	flag.StringVar(&serverRoot, "serverRoot", "", "The base URL of the Bb Learn server. e.g. https://mybb.inst.edu.au")
	flag.StringVar(&appKey, "appKey", "", "The Application Key")
	flag.StringVar(&appSecret, "appSecret", "", "The Application Secret")

	flag.Parse()

	if serverRoot == "" || appKey == "" || appSecret == "" {
		flag.Usage()
		os.Exit(1)
	}

	tokenUrl = serverRoot + "/learn/api/public/v1/oauth2/token"
	usersUrl = serverRoot + "/learn/api/public/v1/users"
	coursesUrl = serverRoot + "/learn/api/public/v1/courses"

}

func main() {

	/*
	 * setup the clientcredentials Configuration data
	 */
	conf := &clientcredentials.Config{
		ClientID:     appKey,
		ClientSecret: appSecret,
		Scopes:       []string{},
		TokenURL:     tokenUrl,
	}

	/*
	 * Get a client based on the configuration
	 */
	client = conf.Client(oauth2.NoContext)

	courseService = CourseService{Client: *client}

	http.HandleFunc("/", courseListHandler)
	http.HandleFunc("/getenrolments", getCourseUsersHandler)
	http.ListenAndServe(":8080", nil)

}

func courseListHandler(w http.ResponseWriter, r *http.Request) {
	courses, err := courseService.getCourses()

	if err != nil {
		log.Fatal(err)
	}

	t, _ := template.ParseFiles("courselist.html")
	t.Execute(w, courses)

}

func getCourseUsersHandler(w http.ResponseWriter, r *http.Request) {

	courseId := r.URL.Query().Get("course_id")
	result, err := courseService.getCourseUsers(courseId)

	if err != nil {
		log.Fatal(err)
	}
	w.Header()["Content-Type"] = []string{"text/csv"}
	w.Header()["Content-Disposition"] = []string{"attachment; filename=\"enrollment-list-" +
		courseId + ".csv\""}
	fmt.Fprintf(w, "user_id, course_id, role, created\n")
	for _, c := range result.CourseUsers {
		// This assumes a lot.  Really should be escaping all these strings.
		fmt.Fprintf(w, "\"%s\",\"%s\",\"%s\",\"%s\"\n", c.UserId, c.CourseId, c.Role, c.CreatedDateTime)
	}

}

type CoursesResult struct {
	Courses []Course `json:"results"`
	Paging  Paging   `json:"paging"`
}

type CourseUsersResult struct {
	CourseUsers []CourseUser `json:"results"`
	Paging      Paging       `json:"paging"`
}

type Paging struct {
	NextPage string `json:"nextPage"`
}

type Course struct {
	Id           string `json:"id"`
	Uuid         string `json:"uuid"`
	ExternalId   string `json:"externalId"`
	DataSourceId string `json:"dataSourceId"`
	CourseId     string `json:"courseId"`
	Name         string `json:"name"`
}

type CourseUser struct {
	UserId          string       `json:"userId"`
	CourseId        string       `json:"courseId"`
	DataSourceId    string       `json:"dataSourceId"`
	CreatedDateTime time.Time    `json:"created"` // "2016-07-21T00:51:57.176Z"
	Availability    Availability `json:"availability"`
	Role            string       `json:"courseRoleId"`
}

type Availability struct {
	Available string `json:"available"`
}

type CourseService struct {
	Client http.Client
}

func (svc CourseService) getCourses() (result CoursesResult, err error) {

	resp, err := svc.Client.Get(coursesUrl)
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	return
}

func (svc CourseService) getCourseUsers(courseId string) (result CourseUsersResult, err error) {
	resp, err := svc.Client.Get(coursesUrl + "/" + courseId + "/users")
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	return
}
