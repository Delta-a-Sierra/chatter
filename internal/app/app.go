package app

import restv1 "github.com/Delta-a-Sierra/chatter/internal/app/adapters/presentation/rest"

func Start() error {
	v1 := restv1.V1API{}
	v1.Start()
	return nil
}
