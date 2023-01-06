package app

import (
	"context"

	"github.com/joaocarvalhodb1/arch-poc/services/account_service/domain"
	"github.com/joaocarvalhodb1/arch-poc/shared/contracts/protog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *AccountAppService) CreateAccount(ctx context.Context, request *protog.AccountRequest) (*protog.AccountResponse, error) {
	account, err := domain.NewAccount(
		request.Data.Name,
		request.Data.Email,
		request.Data.CellPhone,
		request.Data.CreditLimit,
	)
	if err != nil {
		s.log.Error("Error on validate Account", err)
		return nil, status.Errorf(codes.InvalidArgument, "Error on validate Account  %s: ", err.Error())
	}
	if err = s.accountRepo.Save(account); err != nil {
		s.log.Error("Error on creating Account", err)
		return nil, status.Errorf(codes.Internal, "Internal error  %s: ", err.Error())
	}
	res := &protog.AccountResponse{
		Result: &protog.Account{
			Id:          account.ID,
			Name:        account.Name,
			Email:       account.Email,
			CellPhone:   account.CellPhone,
			AccountType: protog.AccountType(protog.AccountType_value[account.AccountType]),
			CreditLimit: account.CreditLimit,
			CreatedAt:   timestamppb.New(account.CreatedAt.Time),
		},
	}
	return res, nil
}

func (s *AccountAppService) UpdateAccount(ctx context.Context, request *protog.AccountRequest) (*protog.Empty, error) {
	account := &domain.Account{
		ID:          request.Data.Id,
		Name:        request.Data.Name,
		Email:       request.Data.Email,
		CellPhone:   request.Data.CellPhone,
		AccountType: request.Data.AccountType.String(),
		CreditLimit: request.Data.CreditLimit,
	}
	err := account.Validate()
	if err != nil {
		s.log.Error("Error on validate Account", err)
		return nil, status.Errorf(codes.InvalidArgument, "Error on validate Account  %s: ", err.Error())
	}
	if err = s.accountRepo.Update(account); err != nil {
		s.log.Error("Error on Update Account", err)
		return nil, status.Errorf(codes.Internal, "Internal error  %s: ", err.Error())
	}
	res := &protog.Empty{}
	return res, nil
}

func (s *AccountAppService) DeleteAccount(ctx context.Context, request *protog.AccountDeleteRequest) (*protog.Empty, error) {
	accountId := request.AccountId
	if accountId == "" {
		s.log.Error("Error, ID is required to delete account")
		return nil, status.Errorf(codes.InvalidArgument, "Error on validate Account id")
	}
	if err := s.accountRepo.Delete(accountId); err != nil {
		s.log.Error("Error on Delete Account", err)
		return nil, status.Errorf(codes.Internal, "Internal error  %s: ", err.Error())
	}
	res := &protog.Empty{}
	return res, nil
}

func (s *AccountAppService) FindOne(ctx context.Context, request *protog.AccountFilterRequest) (*protog.AccountResponse, error) {
	accountId := request.AccountId
	if accountId == "" {
		s.log.Error("Error, ID is required to Find one account")
		return nil, status.Errorf(codes.InvalidArgument, "Error on validate Account id")
	}
	account, err := s.accountRepo.GetOne(accountId)
	if err != nil {
		s.log.Error("Error on find Account", err)
		return nil, status.Errorf(codes.Internal, "Internal error  %s: ", err.Error())
	}
	if account.Id == "" {
		s.log.Error("not found account", err)
		return nil, status.Errorf(codes.NotFound, "Not found account  %s: ", err.Error())
	}
	res := &protog.AccountResponse{
		Result: &protog.Account{
			Id:          account.Id,
			Name:        account.Name,
			Email:       account.Email,
			CellPhone:   account.CellPhone,
			AccountType: protog.AccountType(protog.AccountType_value[account.AccountType]),
			CreditLimit: account.CreditLimit,
			CreatedAt:   timestamppb.New(account.CreatedAt.Time),
			UpdatedAt:   timestamppb.New(account.UpdatedAt.Time),
		},
	}
	return res, nil
}

func (s *AccountAppService) FindMany(ctx context.Context, request *protog.AccountFilterRequest) (*protog.AccountListResponse, error) {
	accountList, err := s.accountRepo.GetMany(request.AccountType)
	if err != nil {
		s.log.Error("Error on find Account", err)
		return nil, status.Errorf(codes.Internal, "Internal error  %s: ", err.Error())
	}
	resultList := &protog.AccountListResponse{
		Result: convertAccountToProtogAccount(accountList),
	}
	return resultList, nil
}

func (s *AccountAppService) CreditLimitApply(ctx context.Context, request *protog.CreditLimiteRequest) (*protog.Empty, error) {
	err := s.creditLimitService.Apply(request.AccountType, request.CreditLimit)
	if err != nil {
		s.log.Error("Error on apply Credit limit", err)
		return nil, status.Errorf(codes.Internal, "Internal error  %s: ", err.Error())
	}
	res := &protog.Empty{}
	return res, nil
}

func convertAccountToProtogAccount(list []*domain.AccountOutpuDTO) []*protog.Account {
	accountList := make([]*protog.Account, 0, len(list))
	for _, ac := range list {
		account := &protog.Account{
			Id:          ac.Id,
			Name:        ac.Name,
			Email:       ac.Email,
			CellPhone:   ac.CellPhone,
			AccountType: protog.AccountType(protog.AccountType_value[ac.AccountType]),
			CreditLimit: ac.CreditLimit,
			CreatedAt:   timestamppb.New(ac.CreatedAt.Time),
			UpdatedAt:   timestamppb.New(ac.UpdatedAt.Time),
		}
		accountList = append(accountList, account)
	}
	return accountList
}
