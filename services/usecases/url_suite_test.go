package usecases

import (
	"github.com/blezzio/triniti/mocks"
	"github.com/blezzio/triniti/services/interfaces"
)

type urlTestSuiteParam struct {
	withCache         bool
	withLogger        bool
	repoData          map[string]string
	cacheData         map[string]string
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
	hash := mocks.NewURLTestHash()
	uc := NewUrl(repo, hash, opts...)

	suite.repoCallLog = repo
	suite.hashCallLog = hash
	suite.uc = uc

	return suite
}
