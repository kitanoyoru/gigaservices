package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/guregu/null"
	"github.com/kitanoyoru/kita/apps/emailservice/internal/config"
	"github.com/kitanoyoru/kita/apps/emailservice/internal/models"
	"github.com/kitanoyoru/kita/apps/emailservice/internal/services"
	"github.com/kitanoyoru/kita/apps/emailservice/pkg/database"
	"github.com/kitanoyoru/kita/apps/emailservice/pkg/events"
	pb "github.com/kitanoyoru/kita/apps/emailservice/pkg/proto"
	uuid "github.com/satori/go.uuid"
	"go-micro.dev/v4/util/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Email struct {
	emailService *services.Email
	db           *database.Postgres
	p            *events.KafkaProducer
}

func NewEmail() *Email {
	emailService, err := initEmailService()
	if err != nil {
		log.Error(err)
	}

	db, err := initDb()
	if err != nil {
		log.Error(err)
	}

	p, err := initEventProducer()
	if err != nil {
		log.Error(err)
	}

	return &Email{
		emailService,
		db,
		p,
	}
}

func (e *Email) SendOrderConfirmation(ctx context.Context, in *pb.SendOrderConfirmationRequest, out *pb.Empty) error {
	err := e.emailService.SendConfirmationMail(in)
	if err != nil {
		e.log(in.Email, err)
		return status.Errorf(codes.Internal, fmt.Sprintf("failed to send message: %v", err))
	}

	e.log(in.Email, nil)
	return nil
}

func initEmailService() (*services.Email, error) {
	emailService := services.NewEmail()

	if err := emailService.Init(); err != nil {
		return nil, err
	}

	return emailService, nil
}

func initDb() (*database.Postgres, error) {
	cfg := config.Database()

	dbConn, err := database.NewPostgres(cfg.URL)
	if err != nil {
		return nil, err
	}

	if err := dbConn.Init(); err != nil {
		return nil, err
	}

	return dbConn, err
}

func initEventProducer() (*events.KafkaProducer, error) {
	cfg := config.MessageBroker()

	p, err := events.NewKafkaProducer(cfg.BrokersUrl)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (e *Email) log(email string, err error) {
	var (
		isSended  bool
		errReason *string
	)

	if err != nil {
		isSended = true
		*errReason = err.Error()
	}

	logModel := &models.LogModel{
		ID:        uuid.NewV4().String(),
		Email:     email,
		IsSended:  isSended,
		ErrReason: null.StringFromPtr(errReason),
		CreatedAt: time.Now().UTC(),
	}

	if err := e.db.Save(logModel); err != nil {
		log.Error(err)
	}
}
