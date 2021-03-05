package repository

import (
	"diplomaProject/application/invite"
	"diplomaProject/application/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type InviteRepository struct {
	conn *pgxpool.Pool
}

func NewInviteRepository(db *pgxpool.Pool) invite.Repository {
	return &InviteRepository{conn: db}
}

func (i InviteRepository) InviteUser(invitation *models.Invitation) error {
	panic("implement me")
}

func (i InviteRepository) InviteTeam(invitation *models.Invitation) error {
	panic("implement me")
}

func (i InviteRepository) GetInvitedUser(invitation *models.Invitation) ([]models.FeedUser, error) {
	panic("implement me")
}

func (i InviteRepository) GetInvitedTeam(invitation *models.Invitation) (models.TeamArr, error) {
	panic("implement me")
}

func (i InviteRepository) GetInvitationUser(invitation *models.Invitation) ([]models.FeedUser, error) {
	panic("implement me")
}

func (i InviteRepository) GetInvitationTeam(invitation *models.Invitation) (models.TeamArr, error) {
	panic("implement me")
}