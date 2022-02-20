package handlers

import (
	"bytes"
	"context"
	"errors"
	"go.uber.org/zap"
	"goImageStore/iternal/utils"
	"goImageStore/pb"
)

func GrpcHandlerNew(repo RepositoryInterface, log *zap.SugaredLogger, serverURL string) *GrpcHandler {
	return &GrpcHandler{
		repo:      repo,
		log:       log,
		serverURL: serverURL,
	}
}

type GrpcHandler struct {
	pb.UnimplementedFileServer
	repo      RepositoryInterface
	serverURL string
	log       *zap.SugaredLogger
}

func (gh *GrpcHandler) SaveFile(ctx context.Context, in *pb.FileRequest) (*pb.FileResponse, error) {
	file := in.GetFile()

	imageData := bytes.Buffer{}
	imageData.Write(file)

	fileName, err := utils.FormatFileName(in.GetTitle())
	if err != nil {
		gh.log.Error(err.Error())
		return nil, errors.New("error with open file")
	}
	if err := gh.repo.SaveImage(imageData, fileName); err != nil {
		gh.log.Error(err.Error())
		return nil, errors.New("error with save file")
	}

	message := "http://" + gh.serverURL + "/api/image/" + fileName
	return &pb.FileResponse{Status: "OK", Url: message}, nil
}
