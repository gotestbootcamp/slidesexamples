package dependency

func Parse(r io.Reader) (User, error) {
	res := User{}
	buf := &bytes.Buffer{}
	buf.ReadFrom(r)

	data := buf.Bytes()
	err := json.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

}
