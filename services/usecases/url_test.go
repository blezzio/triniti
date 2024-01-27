package usecases

import (
	"testing"
)

func Test_URLUseCase_GetHash_Successfull(t *testing.T) {
	const tcname string = "Test_URLUseCase_GetHash_Successfull"
	test_URLUseCase_GetHash_HaveHashInCache(tcname, t)
	test_URLUseCase_GetHash_NoPreExistingHash(tcname, t)
	test_URLUseCase_GetHash_HaveHashInDB(tcname, t)
}

func Test_URLUseCase_GetHash_Failed(t *testing.T) {
	const tcname string = "Test_URLUseCase_GetHash_Failed"
	test_URLUseCase_GetHash_DBFailed(tcname, t)
}

func Test_URLUseCase_GetFullURL_Successfull(t *testing.T) {
	const tcname string = "Test_URLUseCase_GetFullURL_Successfull"
	test_URLUseCase_GetFullURL_HaveFullURLInCache(tcname, t)
	test_URLUseCase_GetFullURL_NoFullURLInCache(tcname, t)
	test_URLUseCase_GetFullURL_NoCache(tcname, t)
}

func Test_URLUseCase_GetFullURL_Failed(t *testing.T) {
	const tcname string = "Test_URLUseCase_GetFullURL_Failed"
	test_URLUseCase_GetFullURL_HashNotExisted(tcname, t)
	test_URLUseCase_GetFullURL_RepoError(tcname, t)
}
