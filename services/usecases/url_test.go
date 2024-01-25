package usecases

import (
	"context"
	"fmt"
	"testing"
)

func Test_URLUseCase_GetHash_Successfull(t *testing.T) {
	const tcname string = "Test_URLUseCase_GetHash_Successfull"
	// BEGIN: have pre-existing hash in cache
	suite := getURLTestSuite(urlTestSuiteParam{withCache: true, cacheData: map[string]string{"https://tini.blezz.io": "abcdef"}, withLogger: true})
	hash, err := suite.uc.GetHash(context.Background(), "https://tini.blezz.io")
	if err != nil {
		t.Errorf("%s err=%v, expected nil", tcname, err)
	}
	if hash == "" {
		t.Errorf("%s hash=%s, expected non-empty", tcname, hash)
	}

	actualCacheGetCalled, _ := suite.cacheCallLog.Called(suite.uc.cache.GetContext)
	expectedCacheGetCalled := 1
	if actualCacheGetCalled != expectedCacheGetCalled {
		t.Errorf("%s cache_get_call=%d, expected %d", tcname, actualCacheGetCalled, expectedCacheGetCalled)
	}

	actualGetHashCalled, _ := suite.repoCallLog.Called(suite.uc.repo.GetHash)
	expectedGetHashCalled := 0
	if actualGetHashCalled != expectedGetHashCalled {
		t.Errorf("%s get_hash_call=%d, expected %d", tcname, actualGetHashCalled, expectedGetHashCalled)
	}

	actualCreateCalled, _ := suite.repoCallLog.Called(suite.uc.repo.Create)
	expectedCreateCalled := 0
	if actualCreateCalled != expectedCreateCalled {
		t.Errorf("%s create_call=%d, expected %d", tcname, actualCreateCalled, expectedCreateCalled)
	}
	// END: have pre-existing hash in cache

	// START: have no pre-existing hash
	suite = getURLTestSuite(urlTestSuiteParam{withCache: true, withLogger: true})
	hash, err = suite.uc.GetHash(context.Background(), "https://tini.blezz.io")
	if err != nil {
		t.Errorf("%s err=%v, expected nil", tcname, err)
	}
	if hash == "" {
		t.Errorf("%s hash=%s, expected non-empty", tcname, hash)
	}

	actualCacheGetCalled, _ = suite.cacheCallLog.Called(suite.uc.cache.GetContext)
	expectedCacheGetCalled = 1
	if actualCacheGetCalled != expectedCacheGetCalled {
		t.Errorf("%s cache_get_call=%d, expected %d", tcname, actualCacheGetCalled, expectedCacheGetCalled)
	}

	actualGetHashCalled, _ = suite.repoCallLog.Called(suite.uc.repo.GetHash)
	expectedGetHashCalled = 1
	if actualGetHashCalled != expectedGetHashCalled {
		t.Errorf("%s get_hash_call=%d, expected %d", tcname, actualGetHashCalled, expectedGetHashCalled)
	}

	actualCreateCalled, _ = suite.repoCallLog.Called(suite.uc.repo.Create)
	expectedCreateCalled = 1
	if actualCreateCalled != expectedCreateCalled {
		t.Errorf("%s create_call=%d, expected %d", tcname, actualCreateCalled, expectedCreateCalled)
	}
	// END: have no pre-existing hash

	// BEGIN: have no pre-existing and no cache
	suite = getURLTestSuite(urlTestSuiteParam{withCache: false, withLogger: true})
	hash, err = suite.uc.GetHash(context.Background(), "https://tini.blezz.io")
	if err != nil {
		t.Errorf("%s err=%v, expected nil", tcname, err)
	}
	if hash == "" {
		t.Errorf("%s hash=%s, expected non-empty", tcname, hash)
	}

	actualGetHashCalled, _ = suite.repoCallLog.Called(suite.uc.repo.GetHash)
	expectedGetHashCalled = 1
	if actualGetHashCalled != expectedGetHashCalled {
		t.Errorf("%s get_hash_call=%d, expected %d", tcname, actualGetHashCalled, expectedGetHashCalled)
	}

	actualCreateCalled, _ = suite.repoCallLog.Called(suite.uc.repo.Create)
	expectedCreateCalled = 1
	if actualCreateCalled != expectedCreateCalled {
		t.Errorf("%s create_call=%d, expected %d", tcname, actualCreateCalled, expectedCreateCalled)
	}
	// END: have pre-existing hash in db, no cache
}

func Test_URLUseCase_GetHash_Failed(t *testing.T) {
	const tcname string = "Test_URLUseCase_GetHash_Failed"
	suite := getURLTestSuite(
		urlTestSuiteParam{
			withCache:  true,
			withLogger: true,
			repoErr:    fmt.Errorf("test error"),
		},
	)
	hash, err := suite.uc.GetHash(context.Background(), "123456")
	if err == nil {
		t.Errorf("%s err=nil, expected error", tcname)
	}
	if hash == "" {
		t.Errorf("%s hash=%s, expected non-empty", tcname, hash)
	}
}

func Test_URLUseCase_GetFullURL_Successfull(t *testing.T) {
	const tcname string = "Test_URLUseCase_GetFullURL_Successfull"
	// BEGIN: have full url in cache
	suite := getURLTestSuite(urlTestSuiteParam{withCache: true, cacheData: map[string]string{"abcdef": "https://tini.blezz.io"}, withLogger: true})
	fullURL, err := suite.uc.GetFullURL(context.Background(), "abcdef")
	if err != nil {
		t.Errorf("%s err=%v, expected nil", tcname, err)
	}
	if fullURL == "" {
		t.Errorf("%s fullURL=%s, expected non-empty", tcname, fullURL)
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
	// END: have full url in cache

	// BEGIN: no full url in cache
	suite = getURLTestSuite(urlTestSuiteParam{withCache: true, repoData: map[string]string{"abcdef": "https://tini.blezz.io"}, withLogger: true})
	fullURL, err = suite.uc.GetFullURL(context.Background(), "abcdef")
	if err != nil {
		t.Errorf("%s err=%v, expected nil", tcname, err)
	}
	if fullURL == "" {
		t.Errorf("%s fullURL=%s, expected non-empty", tcname, fullURL)
	}

	actualCacheGetCalled, _ = suite.cacheCallLog.Called(suite.uc.cache.GetContext)
	expectedCacheGetCalled = 1
	if actualCacheGetCalled != expectedCacheGetCalled {
		t.Errorf("%s cache_get_call=%d, expected %d", tcname, actualCacheGetCalled, expectedCacheGetCalled)
	}

	actualGetFullURLCalled, _ = suite.repoCallLog.Called(suite.uc.repo.GetFullURL)
	expectedGetFullURLCalled = 1
	if actualGetFullURLCalled != expectedGetFullURLCalled {
		t.Errorf("%s get_fullURL_call=%d, expected %d", tcname, actualGetFullURLCalled, expectedGetFullURLCalled)
	}

	actualCacheSetCalled, _ := suite.cacheCallLog.Called(suite.uc.cache.SetContext)
	expectedCacheSetCalled := 1
	if actualCacheSetCalled != expectedCacheSetCalled {
		t.Errorf("%s cache_set_call=%d, expected %d", tcname, actualCacheSetCalled, expectedCacheSetCalled)
	}
	// END: no full url in cache

	// BEGIN: no  cache
	suite = getURLTestSuite(urlTestSuiteParam{withCache: false, repoData: map[string]string{"abcdef": "https://tini.blezz.io"}, withLogger: true})
	fullURL, err = suite.uc.GetFullURL(context.Background(), "abcdef")
	if err != nil {
		t.Errorf("%v err=%v, expected nil", tcname, err)
	}
	if fullURL == "" {
		t.Errorf("%v fullURL=%s, expected non-empty", tcname, fullURL)
	}

	actualGetFullURLCalled, _ = suite.repoCallLog.Called(suite.uc.repo.GetFullURL)
	expectedGetFullURLCalled = 1
	if actualGetFullURLCalled != expectedGetFullURLCalled {
		t.Errorf("%v get_fullURL_call=%d, expected %d", tcname, actualGetFullURLCalled, expectedGetFullURLCalled)
	}
	// END: no cache
}

func Test_URLUseCase_GetFullURL_Failed(t *testing.T) {
	const tcname string = "Test_URLUseCase_GetFullURL_Failed"
	// START: hash not existed
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
	// END: hash not existed

	// START: repo error
	suite = getURLTestSuite(
		urlTestSuiteParam{
			withCache:  true,
			withLogger: true,
			repoErr:    fmt.Errorf("test error"),
		},
	)
	_, err = suite.uc.GetFullURL(context.Background(), "123456")
	if err == nil {
		t.Errorf("%s err=nil, expected error", tcname)
	}
	// END: repo error
}
