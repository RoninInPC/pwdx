package pwdx

type Output struct {
	error error
	dir   string
	pid   string
}

func initOutput(dir string, pid string, err error) Output {
	return Output{err, dir, pid}
}

func (o Output) Error() string {
	if o.error == nil {
		return ""
	}
	return o.error.Error()
}

func (o Output) Pid() string {
	return o.pid
}

func (o Output) Dir() string {
	if o.error != nil {
		return ""
	}
	return o.dir
}

func (o Output) Ans() (string, error) {
	return o.dir, o.error
}

type Outputs []Output

func (o Outputs) GetMap() map[string]string {
	answer := map[string]string{}
	for _, v := range o {
		if v.error == nil {
			answer[v.pid] = v.dir
		}
	}
	return answer
}
