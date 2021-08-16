package server

import (
	"context"
	"github.com/DataWorkbench/filemanager/executor"
	"github.com/DataWorkbench/gproto/pkg/fmpb"
	"github.com/DataWorkbench/gproto/pkg/model"
)

type FileManagerServer struct {
	fmpb.UnimplementedFileManagerServer
	executor   *executor.FileManagerExecutor
	emptyReply *model.EmptyStruct
}

func NewFileManagerServer(executor *executor.FileManagerExecutor) *FileManagerServer {
	return &FileManagerServer{
		executor:   executor,
		emptyReply: &model.EmptyStruct{},
	}
}

// UploadFile @title  文件上传
// @description   上传文件到hadoop文件系统
// @auth      gx             时间（2021/7/28   10:57 ）
// @param     SpaceID        string         "工作空间id"
// @param     FileName       string         "文件名"
// @param     FilePath       string         "文件路径"
// @param     FileType       int32          "文件类型 1 jar包文件 2 udf文件"
// @param     SpaceID        string         "工作空间"
// @param     SpaceID        string         "工作空间"
// @return    ID             string         "文件id"
// @return    SpaceID        string         "工作空间id"
// @return    FileName       string         "文件名"
// @return    FilePath       string         "文件路径"
// @return    FileType       int32          "文件类型 1 jar包文件 2 udf文件"
// @return    HdfsAddress    string         "hadoop地址"
// @return    URL            string         "文件的hadoop地址"
func (fs *FileManagerServer) UploadFile(fu fmpb.FileManager_UploadFileServer) error {
	return fs.executor.UploadFile(fu)
}

// DownloadFile @title  文件下载
// @description   下载hadoop文件到本地
// @auth      gx             时间（2021/7/28   10:57 ）
// @param     ID             string         "文件id"
func (fs *FileManagerServer) DownloadFile(req *fmpb.DownloadRequest, res fmpb.FileManager_DownloadFileServer) error {
	return fs.executor.DownloadFile(req.Id, res)
}

// ListFiles @title  文件下载
// @description   下载hadoop文件到本地
// @auth      gx             时间（2021/7/28   10:57 ）
// @param     SpaceId       string         "根据spaceId查询列表"
// @param     Limit          int32         "分页限制"
// @param     Offset         int32         "页码"
// @return    Total          int32         "文件总数"
// @return    HasMode        int32         "是否有更多"
// @return    ID             string        "文件id"
// @return    SpaceID        string        "工作空间id"
// @return    FileName       string        "文件名"
// @return    FilePath       string        "文件路径"
// @return    FileType       int32         "文件类型 1 jar包文件 2 udf文件"
// @return    Url            string        "文件的hadoop路径"
func (fs *FileManagerServer) ListFiles(ctx context.Context, req *fmpb.ListRequest) (*fmpb.ListResponse, error) {
	infos, count, err := fs.executor.ListFiles(ctx, req.SpaceId, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}
	reply := &fmpb.ListResponse{
		Infos:   infos,
		HasMore: len(infos) >= int(req.Limit),
		Total:   count,
	}
	return reply, nil
}

// UpdateFile @title  更新文件
// @description   更新文件信息
// @auth      gx             时间（2021/7/28   10:57 ）
// @param     Id            string        "文件id"
// @param     FileName      string        "文件名"
// @param     FilePath      string        "文件路径"
// @param     FileType      int32         "文件类型"
func (fs *FileManagerServer) UpdateFile(ctx context.Context, req *fmpb.UpdateFileRequest) (*model.EmptyStruct, error) {
	return fs.executor.UpdateFile(ctx, req.Id, req.FileName, req.FilePath, req.FileType)
}

// DeleteFiles @title  根据id批量删除文件
// @description   删除文件
// @auth      gx             时间（2021/7/28   10:57 ）
// @param     Ids           string         "文件id"
func (fs *FileManagerServer) DeleteFiles(ctx context.Context, req *fmpb.DeleteFilesRequest) (*model.EmptyStruct, error) {
	return fs.executor.DeleteFiles(ctx, req.Ids)
}

// DeleteAllFiles @title  根据space_id批量删除文件
// @description   删除文件
// @auth      gx             时间（2021/7/28   10:57 ）
// @param     SpaceIds           string    "工作空间id"
func (fs *FileManagerServer) DeleteAllFiles(ctx context.Context, req *fmpb.DeleteAllFilesRequest) (*model.EmptyStruct, error) {
	return fs.executor.DeleteAllFiles(ctx, req.SpaceIds)
}

// DescribeFile @title
// @description   根据id查询文件信息
// @auth      gx             时间（2021/7/28   10:57 ）
// @param     id           string    "文件id"
func (fs *FileManagerServer) DescribeFile(ctx context.Context,req *fmpb.DescribeRequest) (*fmpb.FileInfoResponse, error) {
	return fs.executor.DescribeFile(ctx,req.Id)
}
