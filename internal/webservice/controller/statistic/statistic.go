package statistic

import (
	"github.com/gin-gonic/gin"

	"github.com/forgocode/family/internal/pkg/newlog"
	"github.com/forgocode/family/internal/pkg/response"
	categoryService "github.com/forgocode/family/internal/webservice/service/category"
	commentService "github.com/forgocode/family/internal/webservice/service/comemnt"
	tagService "github.com/forgocode/family/internal/webservice/service/tag"
	topicService "github.com/forgocode/family/internal/webservice/service/topic"
)

type StaticCountInfo struct {
	ArticleTotal      int64 `json:"articleTotal"`
	TagTotal          int64 `json:"tagTotal"`
	CategoryTotal     int64 `json:"categoryTotal"`
	TopicTotal        int64 `json:"topicTotal"`
	ShortCommentTotal int64 `json:"shortCommentTotal"`
}

func StatisticCounts(ctx *gin.Context) {

	tagCount, err := tagService.AdminGetTagCount()
	if err != nil {
		newlog.Logger.Errorf("failed to statistic tag count, err: %+v\n", err)
	}
	categoryCount, err := categoryService.AdminGetCategoryCount()
	if err != nil {
		newlog.Logger.Errorf("failed to statistic category count, err: %+v\n", err)
	}
	topicCount, err := topicService.AdminGetTopicCount()
	if err != nil {
		newlog.Logger.Errorf("failed to statistic topic count, err: %+v\n", err)
	}
	articleCount, err := 0, nil
	if err != nil {
		newlog.Logger.Errorf("failed to statistic article count, err: %+v\n", err)
	}
	shortCommentCount, err := commentService.AdminGetShortCommentCount()
	if err != nil {
		newlog.Logger.Errorf("failed to statistic short comment count, err: %+v\n", err)
	}
	result := &StaticCountInfo{
		ArticleTotal:      int64(articleCount),
		TagTotal:          tagCount,
		CategoryTotal:     categoryCount,
		TopicTotal:        topicCount,
		ShortCommentTotal: shortCommentCount,
	}
	response.Success(ctx, *result, 1)

}
