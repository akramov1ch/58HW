package loan

import (
    "context"

    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

type Loan struct {
    ID     string
    BookID string
    UserID string
}

var loans = make(map[string]Loan)

type LoanServiceServer struct {
    UnimplementedLoanServiceServer
}

func (s *LoanServiceServer) IssueLoan(ctx context.Context, req *IssueLoanRequest) (*IssueLoanResponse, error) {
    loan := Loan{
        ID:     "1", // This should be generated uniquely
        BookID: req.GetBookId(),
        UserID: req.GetUserId(),
    }
    loans[loan.ID] = loan
    return &IssueLoanResponse{LoanId: loan.ID}, nil
}

func (s *LoanServiceServer) ReturnLoan(ctx context.Context, req *ReturnLoanRequest) (*ReturnLoanResponse, error) {
    _, exists := loans[req.GetLoanId()]
    if !exists {
        return nil, status.Errorf(codes.NotFound, "Loan not found")
    }
    delete(loans, req.GetLoanId())
    return &ReturnLoanResponse{Success: true}, nil
}
