package repository

import (
	"context"
	sql2 "database/sql"
	"diplomaProject/application/event"
	"diplomaProject/application/models"
	"diplomaProject/pkg/constants"
	"errors"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"strings"
)

type EventDatabase struct {
	conn *pgxpool.Pool
}

func NewEventDatabase(db *pgxpool.Pool) event.Repository {
	return &EventDatabase{conn: db}
}

func (e *EventDatabase) GetSoloEventUsers(evtID int) (*models.UserArr, error) {
	var usrArr models.UserArr
	sql := `select u1.* from event_users eu1 
			left join (
				select tu1.user_id from team_users tu1 
				join team t1 on t1.id=tu1.team_id where event=$1
			) as tu1 
			on eu1.user_id=tu1.user_id 
			join users u1 on eu1.user_id=u1.id
			where eu1.event_id=$1 and tu1.user_id is null`

	err := pgxscan.Select(context.Background(), e.conn, &usrArr, sql, evtID)

	if err != nil {
		return &models.UserArr{}, err
	}

	return &usrArr, nil
}

func (e *EventDatabase) CreateManyEventTeams(evtID int, usrArr *models.UserArr) error {
	sql := `insert into team values `
	sqlU := `insert into team_users values `
	var id []int
	teamName := ""
	for i := range *usrArr {
		teamName = fmt.Sprintf("team_%v_%v", (*usrArr)[i].LastName, (*usrArr)[i].FirstName)
		sql += fmt.Sprintf("(default,'%v',%v,%v),", teamName, evtID, (*usrArr)[i].Id)
	}
	sql = sql[:len(sql)-1] + "returning id"
	err := pgxscan.Select(context.Background(), e.conn, &id, sql)
	if err != nil {
		return err
	}

	for i := range id {
		sqlU += fmt.Sprintf("(%v,%v,0),", id[i], (*usrArr)[i].Id)
	}
	sqlU = sqlU[:len(sqlU)-1]
	queryResult, err := e.conn.Exec(context.Background(), sqlU)
	if err != nil {
		return err
	}
	affected := queryResult.RowsAffected()
	log.Println(affected)
	return nil
}

func (e *EventDatabase) GetTopEvents() (*models.EventDBArr, error) {
	var evtArr models.EventDBArr
	sql := `SELECT * from event where 
			state = 'Open' 
			and isVerified = true 
			and isPrivate = false 
			order by participants_count desc limit 10`

	err := pgxscan.Select(context.Background(), e.conn, &evtArr, sql)

	if err != nil {
		return &models.EventDBArr{}, err
	}

	return &evtArr, nil
}

func (e EventDatabase) GetPrize(prizeID int) (*models.Prize, error) {
	pr := models.Prize{}
	sql := `select * from prize where id = $1`

	queryResult := e.conn.QueryRow(context.Background(), sql, prizeID)
	err := queryResult.Scan(&pr.Id, &pr.EventID, &pr.Name, &pr.Place,
		&pr.Amount, &pr.Total, &pr.WinnerTeamIDs)
	if err != nil {
		return nil, err
	}
	return &pr, err
}

func (e EventDatabase) RemovePrize(evtId int, prArr *models.PrizeArr) error {
	sql := `delete from prize_users where `
	sqlPrize := `delete from prize where event_id = $1 AND (`
	for i := range *prArr {
		sql += fmt.Sprintf(" prize_id = %v OR ", (*prArr)[i].Id)
		sqlPrize += fmt.Sprintf(" id = %v OR ", (*prArr)[i].Id)
	}
	fmt.Println(sql[:len(sql)-3])
	fmt.Println(sqlPrize[:len(sqlPrize)-3] + ")")
	queryResult, err := e.conn.Exec(context.Background(), sql[:len(sql)-3])
	if err != nil {
		return err
	}
	affected := queryResult.RowsAffected()
	log.Println(affected)

	queryResult, err = e.conn.Exec(context.Background(), sqlPrize[:len(sqlPrize)-3]+")", evtId)
	if err != nil {
		return err
	}
	affected = queryResult.RowsAffected()
	log.Println(affected)

	return nil
}

func (e EventDatabase) GetEventWinnerTeams(evtID int) (*models.TeamWinnerArr, error) {
	var tms models.TeamWinnerArr
	t := models.TeamWinner{}
	pr := models.Prize{}
	leadID := sql2.NullInt64{}
	sql := `select t1.*,p1.* from team t1
join prize p1 on t1.id = any(p1.winnerteamids)
where t1.event=$1`
	queryResult, err := e.conn.Query(context.Background(), sql, evtID)
	if err != nil {
		return nil, err
	}
	for queryResult.Next() {
		err = queryResult.Scan(&t.Id, &t.Name, &t.EventID, &leadID,
			&pr.Id, &pr.EventID, &pr.Name,
			&pr.Place, &pr.Amount, &pr.Total, &pr.WinnerTeamIDs)
		if err != nil {
			return nil, err
		}
		t.LeadID = int(leadID.Int64)
		t.Prize = pr
		tms = append(tms, t)
	}
	queryResult.Close()

	return &tms, nil
}

func (e EventDatabase) GetEventTeams(evtID int) (*models.TeamArr, error) {
	var tms models.TeamArr
	tm := models.Team{}
	leadID := sql2.NullInt64{}
	sql := `select * from team where event = $1`
	queryResult, err := e.conn.Query(context.Background(), sql, evtID)
	if err != nil {
		return nil, err
	}
	for queryResult.Next() {
		err = queryResult.Scan(&tm.Id, &tm.Name, &tm.EventID, &leadID)
		if err != nil {
			return nil, err
		}
		tm.LeadID = int(leadID.Int64)
		tms = append(tms, tm)
	}
	queryResult.Close()

	return &tms, nil
}

func (e EventDatabase) UpdateEvent(evt *models.Event) error {
	sql := `update event set  `
	if evt.Name != "" {
		sql += "name = '" + evt.Name + "', "
	}
	if evt.Description != "" {
		sql += "description = '" + evt.Description + "', "
	}
	if !evt.DateStart.IsZero() {
		sql += fmt.Sprintf("date_start = to_timestamp(%d), ", evt.DateStart.Unix())
	}
	if !evt.DateEnd.IsZero() {
		sql += fmt.Sprintf("date_end = to_timestamp(%d), ", evt.DateEnd.Unix())
	}
	if evt.Place != "" {
		sql += "place = '" + evt.Place + "', "
	}
	if evt.Site != "" {
		sql += "site = '" + evt.Site + "', "
	}
	if evt.TeamSize != 0 {
		sql += fmt.Sprintf("team_size = '%v', ", evt.TeamSize)
	}
	if len(sql) <= 18 {
		return nil
	}
	sql = sql[:len(sql)-2] + ` where id=$1`

	queryResult, err := e.conn.Exec(context.Background(), sql, evt.Id)
	if err != nil {
		return err
	}
	affected := queryResult.RowsAffected()
	log.Println(affected)
	//if affected != 1 {
	//	return errors.New("no event")
	//}
	return nil
}

func (e EventDatabase) UpdatePrize(pr *models.Prize) error {
	sql := `update prize set  `
	if pr.Name != "" {
		sql += "name = '" + pr.Name + "', "
	}
	if pr.Place != 0 {
		sql += fmt.Sprintf("place = '%v', ", pr.Place)
	}
	if pr.Amount != 0 {
		sql += fmt.Sprintf("amount = '%v', ", pr.Amount)
	}
	if pr.Total != 0 {
		sql += fmt.Sprintf("total = '%v', ", pr.Total)
	}

	if len(sql) <= 18 {
		return nil
	}
	sql = sql[:len(sql)-2] + ` where id=$1`

	queryResult, err := e.conn.Exec(context.Background(), sql, pr.Id)
	if err != nil {
		return err
	}
	affected := queryResult.RowsAffected()
	log.Println(affected)
	//if affected != 1 {
	//	return errors.New("no prize")
	//}
	return nil
}

func (e EventDatabase) UnSelectWinner(prizeID, tId int, winnerArr []int) error {
	//	sql := `update prize set winnerteamids = '{}' ,
	//amount = amount +1
	//where id = $2`
	arr := fmt.Sprintf("%v", winnerArr)
	arr = arr[1 : len(arr)-1]
	arr = strings.Replace(arr, " ", ",", -1)
	sql := fmt.Sprintf(`update prize set winnerteamids = '{%v}' , 
amount = amount +1
where id = $1`, arr)

	queryResult, err := e.conn.Exec(context.Background(), sql, prizeID)
	if err != nil {
		return err
	}
	affected := queryResult.RowsAffected()
	if affected != 1 {
		return errors.New("no prize")
	}

	return e.UpdateUnWinUsers(prizeID, tId)
}

func (e EventDatabase) UpdateUnWinUsers(prizeID, tId int) error {
	var us []int
	u := 0
	sql := `select user_id from team_users where team_id=$1`
	queryResult, err := e.conn.Query(context.Background(), sql, tId)
	if err != nil {
		return err
	}
	for queryResult.Next() {
		err = queryResult.Scan(&u)
		if err != nil {
			return err
		}
		us = append(us, u)
	}
	queryResult.Close()

	if len(us) == 0 {
		return nil
	}

	sql = `delete from prize_users where `
	for i := range us {
		sql += fmt.Sprintf("(prize_id = %v AND user_id = %v) OR ", prizeID, us[i])
	}
	sql = sql[:len(sql)-3] + ``
	fmt.Println(sql)
	_, err = e.conn.Exec(context.Background(), sql)
	if err != nil {
		return err
	}
	//affected := qR.RowsAffected()
	//if affected != 1 {
	//	return errors.New("already finished")
	//}
	return nil
}

func (e EventDatabase) SelectWinner(prizeID, tId int) error {
	sql := `update prize set winnerteamids = array_append(winnerteamids,$1) , 
amount = amount -1
where id = $2`
	queryResult, err := e.conn.Exec(context.Background(), sql, tId, prizeID)
	if err != nil {
		return err
	}
	affected := queryResult.RowsAffected()
	if affected != 1 {
		return errors.New("no prize")
	}

	return e.UpdateWinUsers(prizeID, tId)
}

func (e EventDatabase) UpdateWinUsers(prizeID, tId int) error {
	var us []int
	u := 0
	sql := `select user_id from team_users where team_id=$1`
	queryResult, err := e.conn.Query(context.Background(), sql, tId)
	if err != nil {
		return err
	}
	for queryResult.Next() {
		err = queryResult.Scan(&u)
		if err != nil {
			return err
		}
		us = append(us, u)
	}
	queryResult.Close()

	if len(us) == 0 {
		return nil
	}

	sql = `insert into prize_users values `
	for i := range us {
		sql += fmt.Sprintf("(%v,%v),", prizeID, us[i])
	}
	sql = sql[:len(sql)-1] + ` on conflict on CONSTRAINT uniq_pair3 do nothing`
	fmt.Println(sql)
	_, err = e.conn.Exec(context.Background(), sql)
	if err != nil {
		return err
	}
	//affected := qR.RowsAffected()
	//if affected != 1 {
	//	return errors.New("already finished")
	//}
	return nil
}

func (e EventDatabase) GetEventPrize(evtID int) (*models.PrizeArr, error) {
	var prArr models.PrizeArr
	pr := models.Prize{}
	sql := `select * from prize
where event_id=$1`

	queryResult, err := e.conn.Query(context.Background(), sql, evtID)
	if err != nil {
		return nil, err
	}
	for queryResult.Next() {
		err = queryResult.Scan(&pr.Id, &pr.EventID, &pr.Name,
			&pr.Place, &pr.Amount, &pr.Total, &pr.WinnerTeamIDs)
		if err != nil {
			return nil, err
		}
		if len(pr.WinnerTeamIDs) < 1 {
			pr.WinnerTeamIDs = []int{}
		}
		prArr = append(prArr, pr)
	}
	queryResult.Close()

	return &prArr, nil
}

func (e EventDatabase) Finish(id int) error {
	sql := `update event set state = '` + constants.EventStatusClosed + `' where id = $1`
	queryResult, err := e.conn.Exec(context.Background(), sql, id)
	if err != nil {
		return err
	}
	affected := queryResult.RowsAffected()
	if affected != 1 {
		return errors.New("already finished")
	}

	return nil
}

func (e EventDatabase) CreatePrize(evtID int, prizeArr models.PrizeArr) error {
	sql := `INSERT INTO prize VALUES`
	for i := range prizeArr {
		prizeArr[i].EventID = evtID
		prizeArr[i].WinnerTeamIDs = []int{}
		sql += fmt.Sprintf("(default,$1,'%v',%v,%v,%v,null),",
			prizeArr[i].Name, prizeArr[i].Place, prizeArr[i].Amount, prizeArr[i].Total)
	}
	println(sql[:len(sql)-1])
	queryResult, err := e.conn.Exec(context.Background(), sql[:len(sql)-1], evtID)
	if err != nil {
		return err
	}
	affected := queryResult.RowsAffected()
	if affected != int64(len(prizeArr)) {
		return errors.New("prize wasn't created")
	}

	return nil
}

func (e EventDatabase) GetEventUsers(evtID int) (*models.UserArr, error) {
	var us models.UserArr
	u := models.User{}
	sql := `select u1.* from event_users eu1
join users u1 on eu1.user_id=u1.id where eu1.event_id=$1
order by u1.lastname`
	queryResult, err := e.conn.Query(context.Background(), sql, evtID)
	if err != nil {
		return nil, err
	}
	for queryResult.Next() {
		err = queryResult.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.Bio, &u.Description,
			&u.WorkPlace, &u.Vk, &u.Tg, &u.Git, &u.Avatar)
		if err != nil {
			return nil, err
		}
		us = append(us, u)
	}
	queryResult.Close()

	return &us, err
}

func (e EventDatabase) Get(id int) (*models.EventDB, error) {
	evt := models.EventDB{}
	sql := `select * from event where id = $1`

	queryResult := e.conn.QueryRow(context.Background(), sql, id)
	err := queryResult.Scan(&evt.Id, &evt.Name, &evt.Description, &evt.Founder,
		&evt.DateStart, &evt.DateEnd, &evt.State, &evt.Place, &evt.ParticipantsCount,
		&evt.Logo, &evt.Background, &evt.Site, &evt.TeamSize)
	if err != nil {
		return nil, err
	}
	return &evt, err
}

func (e EventDatabase) Create(newEvent *models.Event) (*models.EventDB, error) {
	sql := `INSERT INTO event 
			(id,name, description, founder, date_start, date_end, state, place,
				participants_count, site, team_size, is_private)
			VALUES(default,$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)  RETURNING id`
	id := 0
	fmt.Println(sql)
	err := e.conn.QueryRow(context.Background(), sql, newEvent.Name, newEvent.Description,
		newEvent.Founder, newEvent.DateStart, newEvent.DateEnd,
		newEvent.State, newEvent.Place, newEvent.ParticipantsCount,
		newEvent.Site, newEvent.TeamSize, newEvent.IsPrivate).Scan(&id)
	if err != nil {
		return nil, err
	}
	return e.Get(id)
}

func (e EventDatabase) CheckUser(evtID, uid int) bool {
	sql := `select * from event_users where event_id = $1 AND user_id = $2`

	queryResult := e.conn.QueryRow(context.Background(), sql, evtID, uid)
	err := queryResult.Scan(&evtID, &uid)

	return err != nil
}

func (e *EventDatabase) SetLogo(uid, eid int, link string) error {
	sql := `update event set logo=$1 where id=$2 and founder=$3`

	queryResult, err := e.conn.Exec(context.Background(), sql, link, eid, uid)
	if err != nil {
		return err
	}
	affected := queryResult.RowsAffected()
	if affected != 1 {
		return errors.New("no such event")
	}
	return nil
}

func (e *EventDatabase) SetBackground(uid, eid int, link string) error {
	sql := `update event set background=$1 where id=$2 and founder=$3`

	queryResult, err := e.conn.Exec(context.Background(), sql, link, eid, uid)
	if err != nil {
		return err
	}
	affected := queryResult.RowsAffected()
	if affected != 1 {
		return errors.New("no such event")
	}
	return nil
}
