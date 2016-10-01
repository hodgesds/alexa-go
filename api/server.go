package api

func onPost(c echo.Context) error {
	var err error
	var res *Response
	req := &Request{}

	if err := c.Bind(req); err != nil {
		return err
	}

	switch req.AlexaRequest.Type {
	case Intent:
		if res, err = onIntent(req); err != nil {
			return err
		}

	case Launch:
		if res, err = onLaunch(req); err != nil {
			return err
		}

	case SessionEnd:
		if res, err = onSessionEnd(req); err != nil {
			return err
		}
	}

	return c.JSON(http.StatusOK, res)
}

func onIntent(req *Request) (*Response, error) {
	return &Response{}, nil
}

func onLaunch(req *Request) (*Response, error) {
	return &Response{}, nil
}

func onSessionEnd(req *Request) (*Response, error) {
	return &Response{}, nil
}
