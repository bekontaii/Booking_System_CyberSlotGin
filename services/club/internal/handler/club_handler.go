package handler

import "github.com/bekontaii/Booking_System_CyberSlotGin/services/internal/usecase"
type ClubHandler struct{
	authUC usecase.ClubUseCase
}
func NewClubHander(authUC usecase.ClubUseCase) *ClubHandler{
	return &ClubHandler{
		authUC: authUC,
	}
}

