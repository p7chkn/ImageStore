package handlers

import (
	"bytes"
	"context"
	"errors"
	"goImageStore/iternal/utils"
	"goImageStore/pb"

	"go.uber.org/zap"
)

func GrpcHandlerNew(repo RepositoryInterfaceGrpc, log *zap.SugaredLogger, serverURL string) *GrpcHandler {
	return &GrpcHandler{
		repo:      repo,
		log:       log,
		serverURL: serverURL,
	}
}

type GrpcHandler struct {
	pb.UnimplementedFileServer
	repo      RepositoryInterfaceGrpc
	serverURL string
	log       *zap.SugaredLogger
}

func (gh *GrpcHandler) SaveFile(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	file := in.GetFile()

	imageData := bytes.Buffer{}
	imageData.Write(file)

	fileName, err := utils.FormatFileName(in.GetTitle())
	if err != nil {
		gh.log.Error(err.Error())
		return nil, errors.New(err.Error())
	}
	if err := gh.repo.SaveImage(imageData, fileName); err != nil {
		gh.log.Error(err.Error())
		return nil, errors.New(err.Error())
	}

	message := "http://" + gh.serverURL + "/file/image/" + fileName
	return &pb.CreateResponse{Status: "OK", Url: message}, nil
}

func (gh *GrpcHandler) DeleteFile(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := gh.repo.DeleteImage(in.GetFilename())
	if err != nil {
		return nil, err
	}
	return &pb.DeleteResponse{Status: "OK"}, nil
}
