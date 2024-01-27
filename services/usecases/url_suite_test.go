package usecases

import (
	"context"
	"fmt"
	"testing"

	"github.com/blezzio/triniti/mocks"
	"github.com/blezzio/triniti/services/interfaces"
)

type urlTestSuiteParam struct {
	withCache         bool
	withLogger        bool
	repoData          map[string]string
	cacheData         map[string]string
	hasherResult      string
	repoErr, cacheErr error
}

type urlTestSuite struct {
	cacheCallLog  interfaces.CallLog
	loggerCallLog interfaces.CallLog
	repoCallLog   interfaces.CallLog
	hashCallLog   interfaces.CallLog
	uc            *URL
}

func getURLTestSuite(param urlTestSuiteParam) urlTestSuite {
	if param.cacheData == nil {
		param.cacheData = map[string]string{}
	}
	if param.repoData == nil {
		param.repoData = map[string]string{}
	}

	suite := urlTestSuite{}

	opts := []URLOpt{}
	if param.withCache {
		cache := mocks.NewTestCache(
			param.cacheData, param.cacheErr,
		)
		opts = append(
			opts,
			WithCache(cache),
		)
		suite.cacheCallLog = cache
	}

	if param.withLogger {
		logger := mocks.NewTestLogger()
		opts = append(
			opts,
			WithLogger(logger),
		)
		suite.loggerCallLog = logger
	}

	repo := mocks.NewUrlTestRepo(param.repoData, param.repoErr)

	hasherResult := "abcdef"
	if len(param.hasherResult) != 0 {
		hasherResult = param.hasherResult
	}
	hash := mocks.NewURLTestHash(hasherResult)

	uc := NewURL(repo, hash, opts...)

	suite.repoCallLog = repo
	suite.hashCallLog = hash
	suite.uc = uc

	return suite
}

func test_URLUseCase_GetHash_HaveHashInDB(tcname string, t *testing.T) {
	expectedHash := "123456"
	suite := getURLTestSuite(urlTestSuiteParam{withCache: false, withLogger: true, hasherResult: expectedHash, repoData: map[string]string{expectedHash: "https://triniti.blezz.io"}})
	hash, err := suite.uc.GetHash(context.Background(), "https://triniti.blezz.io")
	if err != nil {
		t.Errorf("%s err=%v, expected nil", tcname, err)
	}
	if hash != expectedHash {
		t.Errorf("%s hash=%s, expected %s", tcname, hash, expectedHash)
	}

	actualCreateCalled, _ := suite.repoCallLog.Called(suite.uc.repo.Create)
	expectedCreateCalled := 0
	if actualCreateCalled != expectedCreateCalled {
		t.Errorf("%s create_call=%d, expected %d", tcname, actualCreateCalled, expectedCreateCalled)
	}
}

func test_URLUseCase_GetHash_DBFailed(tcname string, t *testing.T) {
	suite := getURLTestSuite(
		urlTestSuiteParam{
			withCache:  true,
			withLogger: true,
			repoErr:    fmt.Errorf("test error"),
		},
	)
	_, err := suite.uc.GetHash(context.Background(), "123456")
	if err == nil {
		t.Errorf("%s err=nil, expected error", tcname)
	}
}

func test_URLUseCase_GetHash_NoPreExistingHash(tcname string, t *testing.T) {
	expectedHash := "123456"
	suite := getURLTestSuite(urlTestSuiteParam{withCache: true, withLogger: true, hasherResult: expectedHash})
	hash, err := suite.uc.GetHash(context.Background(), "https://triniti.blezz.io")
	if err != nil {
		t.Errorf("%s err=%v, expected nil", tcname, err)
	}
	if hash != expectedHash {
		t.Errorf("%s hash=%s, expected non-empty", tcname, hash)
	}

	actualCacheGetCalled, _ := suite.cacheCallLog.Called(suite.uc.cache.GetContext)
	expectedCacheGetCalled := 1
	if actualCacheGetCalled != expectedCacheGetCalled {
		t.Errorf("%s cache_get_call=%d, expected %d", tcname, actualCacheGetCalled, expectedCacheGetCalled)
	}

	actualCreateCalled, _ := suite.repoCallLog.Called(suite.uc.repo.Create)
	expectedCreateCalled := 1
	if actualCreateCalled != expectedCreateCalled {
		t.Errorf("%s create_call=%d, expected %d", tcname, actualCreateCalled, expectedCreateCalled)
	}
}

func test_URLUseCase_GetHash_HaveHashInCache(tcname string, t *testing.T) {
	expectedHash := "abcdef"
	suite := getURLTestSuite(urlTestSuiteParam{withCache: true, cacheData: map[string]string{expectedHash: "https://triniti.blezz.io"}, withLogger: true})
	hash, err := suite.uc.GetHash(context.Background(), "https://triniti.blezz.io")
	if err != nil {
		t.Errorf("%s err=%v, expected nil", tcname, err)
	}
	if hash != expectedHash {
		t.Errorf("%s hash=%s, expected %v", tcname, hash, expectedHash)
	}

	actualCacheGetCalled, _ := suite.cacheCallLog.Called(suite.uc.cache.GetContext)
	expectedCacheGetCalled := 1
	if actualCacheGetCalled != expectedCacheGetCalled {
		t.Errorf("%s cache_get_call=%d, expected %d", tcname, actualCacheGetCalled, expectedCacheGetCalled)
	}

	actualCreateCalled, _ := suite.repoCallLog.Called(suite.uc.repo.Create)
	expectedCreateCalled := 0
	if actualCreateCalled != expectedCreateCalled {
		t.Errorf("%s create_call=%d, expected %d", tcname, actualCreateCalled, expectedCreateCalled)
	}
}

func test_URLUseCase_GetFullURL_NoCache(tcname string, t *testing.T) {
	expectedFullURL := "https://triniti.blezz.io"
	suite := getURLTestSuite(urlTestSuiteParam{withCache: false, repoData: map[string]string{"abcdef": expectedFullURL}, withLogger: true})
	fullURL, err := suite.uc.GetFullURL(context.Background(), "abcdef")
	if err != nil {
		t.Errorf("%v err=%v, expected nil", tcname, err)
	}
	if fullURL == "" {
		t.Errorf("%v fullURL=%s, expected non-empty", tcname, fullURL)
	}

	actualGetFullURLCalled, _ := suite.repoCallLog.Called(suite.uc.repo.GetFullURL)
	expectedGetFullURLCalled := 1
	if actualGetFullURLCalled != expectedGetFullURLCalled {
		t.Errorf("%v get_fullURL_call=%d, expected %d", tcname, actualGetFullURLCalled, expectedGetFullURLCalled)
	}
}

func test_URLUseCase_GetFullURL_NoFullURLInCache(tcname string, t *testing.T) {
	expectedFullURL := "https://triniti.blezz.io"
	suite := getURLTestSuite(urlTestSuiteParam{withCache: true, repoData: map[string]string{"abcdef": expectedFullURL}, withLogger: true})
	fullURL, err := suite.uc.GetFullURL(context.Background(), "abcdef")
	if err != nil {
		t.Errorf("%s err=%v, expected nil", tcname, err)
	}
	if fullURL != expectedFullURL {
		t.Errorf("%s fullURL=%s, expected non-empty", tcname, fullURL)
	}

	actualCacheGetCalled, _ := suite.cacheCallLog.Called(suite.uc.cache.GetContext)
	expectedCacheGetCalled := 1
	if actualCacheGetCalled != expectedCacheGetCalled {
		t.Errorf("%s cache_get_call=%d, expected %d", tcname, actualCacheGetCalled, expectedCacheGetCalled)
	}

	actualGetFullURLCalled, _ := suite.repoCallLog.Called(suite.uc.repo.GetFullURL)
	expectedGetFullURLCalled := 1
	if actualGetFullURLCalled != expectedGetFullURLCalled {
		t.Errorf("%s get_fullURL_call=%d, expected %d", tcname, actualGetFullURLCalled, expectedGetFullURLCalled)
	}
}

func test_URLUseCase_GetFullURL_HaveFullURLInCache(tcname string, t *testing.T) {
	expectedFullURL := "https://triniti.blezz.io"
	suite := getURLTestSuite(urlTestSuiteParam{withCache: true, cacheData: map[string]string{"abcdef": expectedFullURL}, withLogger: true})
	fullURL, err := suite.uc.GetFullURL(context.Background(), "abcdef")
	if err != nil {
		t.Errorf("%s err=%v, expected nil", tcname, err)
	}
	if fullURL != expectedFullURL {
		t.Errorf("%s fullURL=%s, expected %s", tcname, fullURL, expectedFullURL)
	}

	actualCacheGetCalled, _ := suite.cacheCallLog.Called(suite.uc.cache.GetContext)
	expectedCacheGetCalled := 1
	if actualCacheGetCalled != expectedCacheGetCalled {
		t.Errorf("%s cache_get_call=%d, expected %d", tcname, actualCacheGetCalled, expectedCacheGetCalled)
	}

	actualGetFullURLCalled, _ := suite.repoCallLog.Called(suite.uc.repo.GetFullURL)
	expectedGetFullURLCalled := 0
	if actualGetFullURLCalled != expectedGetFullURLCalled {
		t.Errorf("%s get_fullURl_call=%d, expected %d", tcname, actualGetFullURLCalled, expectedGetFullURLCalled)
	}
}

func test_URLUseCase_GetFullURL_RepoError(tcname string, t *testing.T) {
	suite := getURLTestSuite(
		urlTestSuiteParam{
			withCache:  true,
			withLogger: true,
			repoErr:    fmt.Errorf("test error"),
		},
	)
	_, err := suite.uc.GetFullURL(context.Background(), "123456")
	if err == nil {
		t.Errorf("%s err=nil, expected error", tcname)
	}
}

func test_URLUseCase_GetFullURL_HashNotExisted(tcname string, t *testing.T) {
	suite := getURLTestSuite(urlTestSuiteParam{withCache: true, withLogger: true})
	_, err := suite.uc.GetFullURL(context.Background(), "abcdef")
	if err == nil {
		t.Errorf("%s err=%v, expected not nil", tcname, err)
	}

	actualCacheGetCalled, _ := suite.cacheCallLog.Called(suite.uc.cache.GetContext)
	expectedCacheGetCalled := 1
	if actualCacheGetCalled != expectedCacheGetCalled {
		t.Errorf("%s cache_get_call=%d, expected %d", tcname, actualCacheGetCalled, expectedCacheGetCalled)
	}

	actualGetFullURLCalled, _ := suite.repoCallLog.Called(suite.uc.repo.GetFullURL)
	expectedGetFullURLCalled := 1
	if actualGetFullURLCalled != expectedGetFullURLCalled {
		t.Errorf("%s get_fullURL_call=%d, expected %d", tcname, actualGetFullURLCalled, expectedGetFullURLCalled)
	}
}
