package main_test

import (
	"bytes"
	"catinator-backend/pkg/auth"
	"catinator-backend/pkg/cat"
	"catinator-backend/pkg/config"
	"catinator-backend/pkg/db/ent/enttest"
	"catinator-backend/pkg/log"
	"catinator-backend/pkg/model"
	"catinator-backend/pkg/user"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

type mockFileSvc struct {
	buf         bytes.Buffer
	deletedPath string
}

func (s *mockFileSvc) Create(path string, src io.Reader) error {
	s.buf = bytes.Buffer{}
	_, err := io.Copy(&s.buf, src)
	return err
}

func (s *mockFileSvc) Delete(path string) error {
	s.deletedPath = path
	return nil
}

func TestBasicAuth(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	l := log.New(true)

	tokenAuth := jwtauth.New("HS256", []byte("test"), nil)

	fileSvc := &mockFileSvc{}

	catSvc := cat.NewService(
		*client,
		tokenAuth,
		config.Config{
			Server: config.Server{
				PublicStorageFolder: "./test/",
			},
		},
		fileSvc,
		l,
	)

	userSvc := user.NewService(
		*client,
		tokenAuth,
		l,
	)

	r := chi.NewRouter()

	// Protected routes
	r.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(tokenAuth))

		// Handle valid / invalid tokens. In this example, we use
		// the provided authenticator middleware, but you can write your
		// own very easily, look at the Authenticator method in jwtauth.go
		// and tweak it, its not scary.
		r.Use(auth.Authenticator)

		catSvc.MountHandlers(r)
	})

	// Public routes
	r.Group(func(r chi.Router) {
		userSvc.MountHandlers(r)
	})

	user1 := model.LoginDetails{}
	user2 := model.LoginDetails{}

	t.Run("create user 1", func(t *testing.T) {
		d := model.Registration{
			Email:    "p1@mail.com",
			Password: "abc123",
			Name:     "User 1",
		}
		// Create a New Request
		req, err := newJsonRequest("POST", "/auth/register", d)
		require.Nil(t, err)

		// Execute Request
		response := executeRequest(req, r)

		// Check the response code
		checkResponseCode(t, http.StatusOK, response.Code)

		var resMsg model.SucessMessage

		err = json.Unmarshal(response.Body.Bytes(), &resMsg)
		require.Nil(t, err)

		// We can use testify/require to assert values, as it is more convenient
		require.Equal(t, model.SucessMessage{
			Message: "user created sucessfully",
		}, resMsg)
	})

	t.Run("create user 2", func(t *testing.T) {
		d := model.Registration{
			Email:    "p2@mail.com",
			Password: "xyz123",
			Name:     "User 2",
		}
		// Create a New Request
		req, err := newJsonRequest("POST", "/auth/register", d)
		require.Nil(t, err)

		// Execute Request
		response := executeRequest(req, r)

		// Check the response code
		checkResponseCode(t, http.StatusOK, response.Code)

		var resMsg model.SucessMessage

		err = json.Unmarshal(response.Body.Bytes(), &resMsg)
		require.Nil(t, err)

		// We can use testify/require to assert values, as it is more convenient
		require.Equal(t, model.SucessMessage{
			Message: "user created sucessfully",
		}, resMsg)
	})

	t.Run("login user 1", func(t *testing.T) {
		d := model.Login{
			Email:    "p1@mail.com",
			Password: "abc123",
		}
		// Create a New Request
		req, err := newJsonRequest("POST", "/auth/login", d)
		require.Nil(t, err)

		// Execute Request
		response := executeRequest(req, r)

		// Check the response code
		checkResponseCode(t, http.StatusOK, response.Code)

		err = json.Unmarshal(response.Body.Bytes(), &user1)
		require.Nil(t, err)

		// We can use testify/require to assert values, as it is more convenient
		require.NotEmpty(t, user1.Token)
	})

	t.Run("create user 2", func(t *testing.T) {
		d := model.Login{
			Email:    "p2@mail.com",
			Password: "xyz123",
		}
		// Create a New Request
		req, err := newJsonRequest("POST", "/auth/login", d)
		require.Nil(t, err)

		// Execute Request
		response := executeRequest(req, r)

		// Check the response code
		checkResponseCode(t, http.StatusOK, response.Code)

		err = json.Unmarshal(response.Body.Bytes(), &user2)
		require.Nil(t, err)

		// We can use testify/require to assert values, as it is more convenient
		require.NotEmpty(t, user1.Token)
	})

	var whiskey model.Cat

	t.Run("create pet1 for user 1", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		fw, err := writer.CreateFormField("name")
		require.Nil(t, err)
		_, err = io.Copy(fw, strings.NewReader("Whiskey"))
		require.Nil(t, err)

		fw, err = writer.CreateFormField("description")
		require.Nil(t, err)

		_, err = io.Copy(fw, strings.NewReader("The sweetest cat!"))
		require.Nil(t, err)

		fw, err = writer.CreateFormField("tags")
		require.Nil(t, err)

		_, err = io.Copy(fw, strings.NewReader("sweet,big,hungry"))
		require.Nil(t, err)

		fw, err = writer.CreateFormFile("image", "cat.jpeg")
		require.Nil(t, err)

		file, err := os.Open("cat.jpeg")
		require.Nil(t, err)

		_, err = io.Copy(fw, file)
		require.Nil(t, err)

		// Close multipart writer.
		writer.Close()
		// Create a New Request

		req, err := http.NewRequest("POST", "/cats", bytes.NewBuffer(body.Bytes()))
		require.Nil(t, err)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Add("Authorization", "Bearer "+user1.Token)

		// Execute Request
		response := executeRequest(req, r)

		// Check the response code
		checkResponseCode(t, http.StatusOK, response.Code)

		err = json.Unmarshal(response.Body.Bytes(), &whiskey)
		require.Nil(t, err)

		// We can use testify/require to assert values, as it is more convenient
		require.NotEmpty(t, whiskey)

		require.Equal(t, "Whiskey", whiskey.Name)
		require.Equal(t, "The sweetest cat!", whiskey.Description)
		require.Equal(t, []string{"sweet", "big", "hungry"}, whiskey.Tags)
		require.NotEmpty(t, whiskey.ImageID)
	})

	t.Run("create pet 1 user 2", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		fw, err := writer.CreateFormField("name")
		require.Nil(t, err)
		_, err = io.Copy(fw, strings.NewReader("Miko"))
		require.Nil(t, err)

		fw, err = writer.CreateFormField("description")
		require.Nil(t, err)

		_, err = io.Copy(fw, strings.NewReader("Boxer!"))
		require.Nil(t, err)

		fw, err = writer.CreateFormField("tags")
		require.Nil(t, err)

		_, err = io.Copy(fw, strings.NewReader("karate,judo,boxing"))
		require.Nil(t, err)

		fw, err = writer.CreateFormFile("image", "cat.jpeg")
		require.Nil(t, err)

		file, err := os.Open("cat.jpeg")
		require.Nil(t, err)

		_, err = io.Copy(fw, file)
		require.Nil(t, err)

		// Close multipart writer.
		writer.Close()
		// Create a New Request

		req, err := http.NewRequest("POST", "/cats", bytes.NewBuffer(body.Bytes()))
		require.Nil(t, err)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Add("Authorization", "Bearer "+user2.Token)

		// Execute Request
		response := executeRequest(req, r)

		// Check the response code
		checkResponseCode(t, http.StatusOK, response.Code)

		cat := model.Cat{}

		err = json.Unmarshal(response.Body.Bytes(), &cat)
		require.Nil(t, err)

		// We can use testify/require to assert values, as it is more convenient
		require.NotEmpty(t, cat)

		require.Equal(t, "Miko", cat.Name)
		require.Equal(t, "Boxer!", cat.Description)
		require.Equal(t, []string{"karate", "judo", "boxing"}, cat.Tags)
		require.NotEmpty(t, cat.ImageID)
	})

	t.Run("create pet 2 for user 1", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		fw, err := writer.CreateFormField("name")
		require.Nil(t, err)
		_, err = io.Copy(fw, strings.NewReader("Leo"))
		require.Nil(t, err)

		fw, err = writer.CreateFormField("description")
		require.Nil(t, err)

		_, err = io.Copy(fw, strings.NewReader("The lazy cat!"))
		require.Nil(t, err)

		fw, err = writer.CreateFormField("tags")
		require.Nil(t, err)

		_, err = io.Copy(fw, strings.NewReader("lazy,small,lazier"))
		require.Nil(t, err)

		fw, err = writer.CreateFormFile("image", "cat.jpeg")
		require.Nil(t, err)

		file, err := os.Open("cat.jpeg")
		require.Nil(t, err)

		_, err = io.Copy(fw, file)
		require.Nil(t, err)

		// Close multipart writer.
		writer.Close()
		// Create a New Request

		req, err := http.NewRequest("POST", "/cats", bytes.NewBuffer(body.Bytes()))
		require.Nil(t, err)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Add("Authorization", "Bearer "+user1.Token)

		// Execute Request
		response := executeRequest(req, r)

		// Check the response code
		checkResponseCode(t, http.StatusOK, response.Code)

		cat := model.Cat{}

		err = json.Unmarshal(response.Body.Bytes(), &cat)
		require.Nil(t, err)

		// We can use testify/require to assert values, as it is more convenient
		require.NotEmpty(t, cat)

		require.Equal(t, "Leo", cat.Name)
		require.Equal(t, "The lazy cat!", cat.Description)
		require.Equal(t, []string{"lazy", "small", "lazier"}, cat.Tags)
		require.NotEmpty(t, cat.ImageID)
	})

	t.Run("get user 1 cats", func(t *testing.T) {

		req, _ := http.NewRequest("GET", "/cats?order=asc", nil)
		req.Header.Add("Authorization", "Bearer "+user1.Token)
		// Execute Request
		response := executeRequest(req, r)

		cats := []model.Cat{}

		// Check the response code
		checkResponseCode(t, http.StatusOK, response.Code)

		err := json.Unmarshal(response.Body.Bytes(), &cats)
		require.Nil(t, err)

		// We can use testify/require to assert values, as it is more convenient
		require.NotEmpty(t, cats)

		require.Equal(t, 2, len(cats))

		require.Equal(t, "Whiskey", cats[0].Name)
		require.Equal(t, "The sweetest cat!", cats[0].Description)
		require.Equal(t, []string{"sweet", "big", "hungry"}, cats[0].Tags)
		require.NotEmpty(t, cats[0].ImageID)

		require.Equal(t, "Leo", cats[1].Name)
		require.Equal(t, "The lazy cat!", cats[1].Description)
		require.Equal(t, []string{"lazy", "small", "lazier"}, cats[1].Tags)
		require.NotEmpty(t, cats[1].ImageID)
	})

	t.Run("get user 1 cats in desc order", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/cats?sort=desc", nil)
		req.Header.Add("Authorization", "Bearer "+user1.Token)
		// Execute Request
		response := executeRequest(req, r)

		cats := []model.Cat{}

		// Check the response code
		checkResponseCode(t, http.StatusOK, response.Code)

		err := json.Unmarshal(response.Body.Bytes(), &cats)
		require.Nil(t, err)

		// We can use testify/require to assert values, as it is more convenient
		require.NotEmpty(t, cats)

		require.Equal(t, 2, len(cats))

		require.Equal(t, "Whiskey", cats[1].Name)
		require.Equal(t, "The sweetest cat!", cats[1].Description)
		require.Equal(t, []string{"sweet", "big", "hungry"}, cats[1].Tags)
		require.NotEmpty(t, cats[1].ImageID)

		require.Equal(t, "Leo", cats[0].Name)
		require.Equal(t, "The lazy cat!", cats[0].Description)
		require.Equal(t, []string{"lazy", "small", "lazier"}, cats[0].Tags)
		require.NotEmpty(t, cats[0].ImageID)
	})

	t.Run("get pet 1 of user 1", func(t *testing.T) {

		req, _ := http.NewRequest("GET", "/cat/"+whiskey.ID, nil)
		req.Header.Add("Authorization", "Bearer "+user1.Token)
		// Execute Request
		response := executeRequest(req, r)

		// Check the response code
		checkResponseCode(t, http.StatusOK, response.Code)

		cat := model.Cat{}

		err := json.Unmarshal(response.Body.Bytes(), &cat)
		require.Nil(t, err)

		// We can use testify/require to assert values, as it is more convenient
		require.NotEmpty(t, cat)

		require.Equal(t, "Whiskey", cat.Name)
		require.Equal(t, "The sweetest cat!", cat.Description)
		require.Equal(t, []string{"sweet", "big", "hungry"}, cat.Tags)
		require.NotEmpty(t, cat.ImageID)
	})

	t.Run("delte pet 1 of user 1", func(t *testing.T) {

		req, _ := http.NewRequest("DELETE", "/cat/"+whiskey.ID, nil)
		req.Header.Add("Authorization", "Bearer "+user1.Token)
		// Execute Request
		response := executeRequest(req, r)

		// Check the response code
		checkResponseCode(t, http.StatusOK, response.Code)

		msg := model.SucessMessage{}

		err := json.Unmarshal(response.Body.Bytes(), &msg)
		require.Nil(t, err)

		// We can use testify/require to assert values, as it is more convenient
		require.NotEmpty(t, msg)

		require.Equal(t, "cat deleted sucessfully", msg.Message)
	})

	t.Run("get user 1 pets", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/cats?sort=desc", nil)
		req.Header.Add("Authorization", "Bearer "+user1.Token)
		// Execute Request
		response := executeRequest(req, r)

		cats := []model.Cat{}

		// Check the response code
		checkResponseCode(t, http.StatusOK, response.Code)

		err := json.Unmarshal(response.Body.Bytes(), &cats)
		require.Nil(t, err)

		// We can use testify/require to assert values, as it is more convenient
		require.NotEmpty(t, cats)

		require.Equal(t, 1, len(cats))

		require.Equal(t, "Leo", cats[0].Name)
		require.Equal(t, "The lazy cat!", cats[0].Description)
		require.Equal(t, []string{"lazy", "small", "lazier"}, cats[0].Tags)
		require.NotEmpty(t, cats[0].ImageID)
	})

	t.Run("get user 2 cats", func(t *testing.T) {

		req, _ := http.NewRequest("GET", "/cats?sort=desc", nil)
		req.Header.Add("Authorization", "Bearer "+user2.Token)
		// Execute Request
		response := executeRequest(req, r)

		cats := []model.Cat{}

		// Check the response code
		checkResponseCode(t, http.StatusOK, response.Code)

		err := json.Unmarshal(response.Body.Bytes(), &cats)
		require.Nil(t, err)

		// We can use testify/require to assert values, as it is more convenient
		require.NotEmpty(t, cats)

		require.Equal(t, 1, len(cats))

		require.Equal(t, "Miko", cats[0].Name)
		require.Equal(t, "Boxer!", cats[0].Description)
		require.Equal(t, []string{"karate", "judo", "boxing"}, cats[0].Tags)
		require.NotEmpty(t, cats[0].ImageID)
	})
}

func newJsonRequest(method, url string, v any) (*http.Request, error) {
	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(v)
	if err != nil {
		return nil, err
	}
	req, _ := http.NewRequest(method, url, buf)
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

// executeRequest, creates a new ResponseRecorder
// then executes the request by calling ServeHTTP in the router
// after which the handler writes the response to the response recorder
// which we can then inspect.
func executeRequest(req *http.Request, r chi.Router) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	return rr
}

// checkResponseCode is a simple utility to check the response code
// of the response
func checkResponseCode(t *testing.T, expected, actual int) {
	require.Equal(t, expected, actual, "Expected response code %d. Got %d\n", expected, actual)
}
