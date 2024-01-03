package services

import "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/ports"

type serviceMisc struct {
	RepositoryMisc ports.RepositoryMisc
}

func NewServiceMisc(rMisc ports.RepositoryMisc) *serviceMisc {
	return &serviceMisc{
		RepositoryMisc: rMisc,
	}
}