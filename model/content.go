package model

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/KL-Engineering/common-log/log"
	"github.com/golang-jwt/jwt"
	"io/ioutil"
	"kidsloop-stm-lambda/config"
	"kidsloop-stm-lambda/entity"
	"net/http"
	"strings"
	"sync"
	"time"
)

type IContent interface {
	MapContents(ctx context.Context, IDs []string) (map[string]*entity.LessonPlan, error)
}

type LocalContent struct{}

func (localContent *LocalContent) MapContents(ctx context.Context, IDs []string) (map[string]*entity.LessonPlan, error) {
	dir := "/Users/yanghui/kidsloop/kidsloop-stm-lambda/doc/json/lesson_plans"
	result := make(map[string]*entity.LessonPlan)
	for _, id := range IDs {
		data, err := ioutil.ReadFile(strings.Join([]string{dir, id + ".json"}, "/"))
		if err != nil {
			log.Error(ctx, "read lesson_plan", log.Err(err), log.String("id", id))
			return nil, err
		}
		var lessonPlan entity.LessonPlan
		err = json.Unmarshal(data, &lessonPlan)
		if err != nil {
			log.Error(ctx, "unmarshal lesson_plan", log.Err(err), log.String("data", string(data)))
			return nil, err
		}
		result[id] = &lessonPlan
	}
	return result, nil
}

type KidsloopProvider struct {
	httpClient *http.Client
	session    string
	refreshAt  int64
}

func (kidsloopProvider *KidsloopProvider) refreshToken(ctx context.Context) error {
	now := time.Now()
	if now.Unix() < kidsloopProvider.refreshAt {
		log.Info(ctx, "don't need refresh")
		return nil
	}
	expiresAt := now.Add(entity.TokenValidityPeriod).Unix()
	claims := &jwt.StandardClaims{
		Audience:  "kidsloop-cms",
		ExpiresAt: expiresAt,
		IssuedAt:  now.Unix(),
		Issuer:    "stm-lambda",
		Subject:   "authorization",
	}
	claim := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	token, err := claim.SignedString(config.Get().CmsAccessKey)
	if err != nil {
		log.Error(ctx, "sign token", log.Err(err), log.Any("claims", claim))
		return err
	}
	kidsloopProvider.session = token
	kidsloopProvider.refreshAt = now.Add(entity.TokenValidityPeriod - entity.TokenRefreshBefore).Unix()
	return nil
}

func (kidsloopProvider *KidsloopProvider) MapContents(ctx context.Context, IDs []string) (map[string]*entity.LessonPlan, error) {
	err := kidsloopProvider.refreshToken(ctx)
	if err != nil {
		log.Error(ctx, "refresh token", log.Err(err), log.Strings("ids", IDs))
		return nil, err
	}
	body, err := json.Marshal(IDs)
	if err != nil {
		log.Error(ctx, "marshal ids", log.Err(err), log.Strings("ids", IDs))
		return nil, err
	}
	requestUrl := config.Get().CmsEndpoint + "/v1/internal/stm/contents"
	request, err := http.NewRequest(http.MethodGet, requestUrl, bytes.NewBuffer(body))
	cookie := http.Cookie{
		Name:  "access",
		Value: kidsloopProvider.session,
	}
	request.AddCookie(&cookie)
	response, err := kidsloopProvider.httpClient.Do(request)
	if err != nil {
		log.Error(ctx, "do http", log.Err(err), log.Strings("ids", IDs))
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		log.Error(ctx, "http status is not ok", log.Int("status", response.StatusCode), log.Strings("ids", IDs))
		return nil, entity.ErrHttpStatusNotOk
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(ctx, "read body", log.Err(err), log.Strings("ids", IDs))
		return nil, err
	}
	defer response.Body.Close()
	var lessonPlans []*entity.LessonPlan
	err = json.Unmarshal(data, &lessonPlans)
	if err != nil {
		log.Error(ctx, "unmarshal lesson_plan", log.Err(err), log.String("data", string(data)))
		return nil, err
	}
	result := make(map[string]*entity.LessonPlan)
	for _, lp := range lessonPlans {
		result[lp.ID] = lp
	}
	return result, nil
}

var (
	_contentProvider IContent
	_contentOnce     sync.Once
)

func mustKidsloopProvider(ctx context.Context) *KidsloopProvider {
	now := time.Now()
	expiresAt := now.Add(entity.TokenValidityPeriod).Unix()
	claims := &jwt.StandardClaims{
		Audience:  "kidsloop-cms",
		ExpiresAt: expiresAt,
		IssuedAt:  now.Unix(),
		Issuer:    "stm-lambda",
		Subject:   "authorization",
	}
	claim := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	token, err := claim.SignedString(config.Get().CmsAccessKey)
	if err != nil {
		log.Panic(ctx, "sign token", log.Err(err), log.Any("claims", claim))
	}

	provider := &KidsloopProvider{
		httpClient: http.DefaultClient,
		session:    token,
		refreshAt:  now.Add(entity.TokenValidityPeriod - entity.TokenRefreshBefore).Unix(),
	}
	return provider
}

func GetContentProvider(ctx context.Context) IContent {
	_contentOnce.Do(func() {
		_contentProvider = mustKidsloopProvider(ctx)
	})
	return _contentProvider
}
