// MinIO Browser WebRPC 对象标签接口实现
// 该文件用于集中管理 Web 相关对象标签操作，便于升级时不被覆盖
package cmd

import (
	"net/http"
	"context"
	"github.com/minio/minio-go/v7/pkg/tags"
)

// AddObjectTagsArgs WebRPC 请求参数
// Tags 字段为 XML 字符串
// VersionID 可选
type AddObjectTagsArgs struct {
	Bucket    string `json:"bucket"`
	Object    string `json:"object"`
	Tags      string `json:"tags"`
	VersionID string `json:"versionId,omitempty"`
}

type GetObjectTagsArgs struct {
	Bucket    string `json:"bucket"`
	Object    string `json:"object"`
	VersionID string `json:"versionId,omitempty"`
}

type DeleteObjectTagsArgs struct {
	Bucket    string `json:"bucket"`
	Object    string `json:"object"`
	VersionID string `json:"versionId,omitempty"`
}

// Web 通用 JSON 响应
type WebObjectTagsResponse struct {
	Status  string      `json:"status"` // "success" or "error"
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// func (web *webAPIHandlers) AddObjectTags(r *http.Request, args *AddObjectTagsArgs, reply *WebObjectTagsResponse) error {
// 	claims, owner, authErr := webRequestAuthenticate(r)
// 	if authErr != nil {
// 		reply.Status = "error"
// 		reply.Message = "认证失败: " + authErr.Error()
// 		return nil
// 	}
// 	objAPI := web.ObjectAPI()
// 	if objAPI == nil {
// 		reply.Status = "error"
// 		reply.Message = "服务未初始化"
// 		return nil
// 	}
// 	if !objAPI.IsTaggingSupported() {
// 		reply.Status = "error"
// 		reply.Message = "对象标签功能未启用"
// 		return nil
// 	}
// 	// 鉴权
// 	if s3Err := checkWebRequestAuthType(ctx, r, policy.PutObjectTaggingAction, args.Bucket, args.Object, claims, owner); s3Err != nil {
// 		reply.Status = "error"
// 		reply.Message = "无权限: " + s3Err.Error()
// 		return nil
// 	}
// 	// 校验 XML
// 	tagObj, err := tags.ParseObjectXML(io.NopCloser(strings.NewReader(args.Tags)))
// 	if err != nil {
// 		reply.Status = "error"
// 		reply.Message = "标签 XML 格式错误: " + err.Error()
// 		return nil
// 	}
// 	opts := cmd.ObjectOptions{VersionID: args.VersionID}
// 	objInfo, err := objAPI.PutObjectTags(ctx, args.Bucket, args.Object, tagObj.String(), opts)
// 	if err != nil {
// 		reply.Status = "error"
// 		reply.Message = "写入标签失败: " + err.Error()
// 		return nil
// 	}
// 	reply.Status = "success"
// 	reply.Data = map[string]interface{}{"versionId": objInfo.VersionID}
// 	return nil
// }
// // AddObjectTags WebRPC 方法
// func (web *WebAPIHandlers) AddObjectTags(r *http.Request, args *AddObjectTagsArgs, reply *WebObjectTagsResponse) error {
// }

// // GetObjectTags WebRPC 方法
// func (web *WebAPIHandlers) GetObjectTags(r *http.Request, args *GetObjectTagsArgs, reply *WebObjectTagsResponse) error {
// 	ctx := cmd.NewContext(r, nil, "WebGetObjectTags")
// 	claims, owner, authErr := webRequestAuthenticate(r)
// 	if authErr != nil {
// 		reply.Status = "error"
// 		reply.Message = "认证失败: " + authErr.Error()
// 		return nil
// 	}
// 	objAPI := web.ObjectAPI()
// 	if objAPI == nil {
// 		reply.Status = "error"
// 		reply.Message = "服务未初始化"
// 		return nil
// 	}
// 	if !objAPI.IsTaggingSupported() {
// 		reply.Status = "error"
// 		reply.Message = "对象标签功能未启用"
// 		return nil
// 	}
// 	if s3Err := checkWebRequestAuthType(ctx, r, policy.GetObjectTaggingAction, args.Bucket, args.Object, claims, owner); s3Err != nil {
// 		reply.Status = "error"
// 		reply.Message = "无权限: " + s3Err.Error()
// 		return nil
// 	}
// 	opts := cmd.ObjectOptions{VersionID: args.VersionID}
// 	tagObj, err := objAPI.GetObjectTags(ctx, args.Bucket, args.Object, opts)
// 	if err != nil {
// 		reply.Status = "error"
// 		reply.Message = "获取标签失败: " + err.Error()
// 		return nil
// 	}
// 	reply.Status = "success"
// 	reply.Data = tagObj.ToMap()
// 	return nil
// }

// // DeleteObjectTags WebRPC 方法
// func (web *WebAPIHandlers) DeleteObjectTags(r *http.Request, args *DeleteObjectTagsArgs, reply *WebObjectTagsResponse) error {
// 	ctx := cmd.NewContext(r, nil, "WebDeleteObjectTags")
// 	claims, owner, authErr := webRequestAuthenticate(r)
// 	if authErr != nil {
// 		reply.Status = "error"
// 		reply.Message = "认证失败: " + authErr.Error()
// 		return nil
// 	}
// 	objAPI := web.ObjectAPI()
// 	if objAPI == nil {
// 		reply.Status = "error"
// 		reply.Message = "服务未初始化"
// 		return nil
// 	}
// 	if !objAPI.IsTaggingSupported() {
// 		reply.Status = "error"
// 		reply.Message = "对象标签功能未启用"
// 		return nil
// 	}
// 	if s3Err := checkWebRequestAuthType(ctx, r, policy.DeleteObjectTaggingAction, args.Bucket, args.Object, claims, owner); s3Err != nil {
// 		reply.Status = "error"
// 		reply.Message = "无权限: " + s3Err.Error()
// 		return nil
// 	}
// 	opts := cmd.ObjectOptions{VersionID: args.VersionID}
// 	_, err := objAPI.DeleteObjectTags(ctx, args.Bucket, args.Object, opts)
// 	if err != nil {
// 		reply.Status = "error"
// 		reply.Message = "删除标签失败: " + err.Error()
// 		return nil
// 	}
// 	reply.Status = "success"
// 	return nil
// }

// // checkWebRequestAuthType 用于 WebRPC 鉴权
// func checkWebRequestAuthType(ctx context.Context, r *http.Request, action policy.Action, bucket, object string, claims map[string]interface{}, owner string) error {
// 	// 可根据项目实际情况调整
// 	return nil // TODO: 实现基于 claims/owner 的权限校验
// }

// // webRequestAuthenticate 获取当前用户 claims/owner
// func webRequestAuthenticate(r *http.Request) (claims map[string]interface{}, owner string, err error) {
// 	// 可根据项目实际情况调整
// 	return nil, "", nil // TODO: 实现 Web 认证信息提取
// }


type TagArgs struct {
	Bucket string `json:"bucket"`
	Object string `json:"object"`
	Tags   string `json:"tags,omitempty"` // 仅 PUT 使用
}



func (web *webAPIHandlers) WebPutObjectTags(r *http.Request, args *TagArgs, reply *ObjectInfo) error {
	objectAPI := web.ObjectAPI()
	if objectAPI == nil {
		return toJSONError(r.Context(), errServerNotInitialized)
	}
	opts := ObjectOptions{} // 可解析 token、versionId 等
	info, err := objectAPI.PutObjectTags(r.Context(), args.Bucket, args.Object, args.Tags, opts)
	if err != nil {
		return toJSONError(r.Context(), err)
	}
	*reply = info
	return nil
}


func (web *webAPIHandlers) WebGetObjectTags(r *http.Request, args *TagArgs, reply *tags.Tags) error {
	objectAPI := web.ObjectAPI()
	if objectAPI == nil {
		return toJSONError(r.Context(), errServerNotInitialized)
	}
	opts := ObjectOptions{}
	tg, err := objectAPI.GetObjectTags(r.Context(), args.Bucket, args.Object, opts)
	if err != nil {
		return toJSONError(r.Context(), err)
	}
	*reply = *tg
	return nil
}


func (web *webAPIHandlers) WebDeleteObjectTags(r *http.Request, args *TagArgs, reply *ObjectInfo) error {
	objectAPI := web.ObjectAPI()
	if objectAPI == nil {
		return toJSONError(r.Context(), errServerNotInitialized)
	}
	opts := ObjectOptions{}
	info, err := objectAPI.DeleteObjectTags(r.Context(), args.Bucket, args.Object, opts)
	if err != nil {
		return toJSONError(r.Context(), err)
	}
	*reply = info
	return nil
}


func (web *webAPIHandlers) WebUpdateObjectTags(r *http.Request, args *TagArgs, reply *bool) error {
	objectAPI := web.ObjectAPI()
	if objectAPI == nil {
		return toJSONError(r.Context(), errServerNotInitialized)
	}

	opts := ObjectOptions{}
	// 调用 JuiceFS 后端实现的方法
	objLayer, ok := objectAPI.(interface {
		UpdateObjectTags(ctx context.Context, bucket, object, tags string, opts ObjectOptions) (ObjectInfo, error)
	})
	if !ok {
		return toJSONError(r.Context(), errServerNotInitialized)
	}

	_, err := objLayer.UpdateObjectTags(r.Context(), args.Bucket, args.Object, args.Tags, opts)
	if err != nil {
		return toJSONError(r.Context(), err)
	}

	*reply = true
	return nil
}
