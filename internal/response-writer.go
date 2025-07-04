package internal

import "io"

type ResponseWriter struct {
	Dest io.Writer
}

func (r *ResponseWriter) Write(data Resp) error {
	d, err := data.Encode()
	if err != nil {
		return err
	}
	_, err = r.Dest.Write(d)
	if err != nil {
		return err
	}

	return nil
}
