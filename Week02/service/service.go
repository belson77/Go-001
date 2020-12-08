package service

import (
	"io"
	"bytes"
	"strconv"
	"net/http"
	"encoding/json"
	"github.com/pkg/errors"
	d "github.com/belson7/Go-000/Week02/dao"
)

func New(dao *d.Dao, err chan error) *Service {
	return &Service{dao, err}
}

type Service struct {
	dao *d.Dao
	errCh chan error
}

func (s *Service) GetUserHandle(resp http.ResponseWriter, req *http.Request) {
	uid, _ := strconv.ParseInt(req.FormValue("uid"), 10, 64)
	user, err := s.dao.GetUser(uid)
	if err != nil {
		s.error(err, resp)
		return
	}
	if user == nil {
		resp.WriteHeader(http.StatusNotFound)
		return
	}

	b, err := json.Marshal(user)
	if err != nil {
		s.error(errors.Wrap(err, "json decode error"), resp)
		return
	}

	if _, err = io.Copy(resp, bytes.NewReader(b)); err != nil {
		s.error(errors.Wrap(err, "response write error"), resp)
		return
	}
	return
}

func (s *Service) error(err error, resp http.ResponseWriter) {
	resp.WriteHeader(http.StatusInternalServerError)
	s.errCh <- err
}