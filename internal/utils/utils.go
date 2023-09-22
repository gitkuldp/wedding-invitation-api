package utils

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type ArrayList []string

func (j *ArrayList) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := []string{}
	err := json.Unmarshal(bytes, &result)
	*j = ArrayList(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (j ArrayList) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func Ref[T any](t T) *T {
	return &t
}

// Deref return a dereference of given input
func Deref[T any](t *T) T {
	if t == nil {
		return *new(T)
	}
	return *t
}

func SanitizeString(src string) string {
	return strings.Replace(src, "'", "''", -1)
}

type Pagination struct {
	Paginate bool
	Page     int
	Limit    int
}

func (p Pagination) Offset() int {
	return (p.Page - 1) * p.Limit
}

func (p Pagination) TotalPage(total int64) int {
	return int(math.Ceil(float64(total) / float64(p.Limit)))
}

func ParsePagination(ctx echo.Context) Pagination {
	pagination := Pagination{
		Page:     1,
		Limit:    20,
		Paginate: true,
	}

	paginate := ctx.QueryParam("pagination")

	if paginate == "false" {
		pagination.Paginate = false
		return pagination
	}

	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if err == nil && page > 0 {
		pagination.Page = page
	}

	limit, err := strconv.Atoi(ctx.QueryParam("limit"))
	if err == nil && limit > 0 {
		pagination.Limit = limit
	}

	return pagination
}

const MaxLoginAttempt = 3
const LockedTimeInMinutes = 30
const Hours = 24
const Minutes = 60
const DateFormatYMD = "2006-01-02"

func ConvertMinutes(min int) string {
	days := min / (Hours * Minutes)
	hours := (min % (Hours * Minutes)) / Minutes
	mins := (min % Minutes)

	result := ""

	if days > 0 {
		result += fmt.Sprintf("%dday ", days)
	}
	if hours > 0 {
		result += fmt.Sprintf("%dhr ", hours)
	}
	if mins > 0 {
		result += fmt.Sprintf("%dmin", mins)
	}

	return result
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// StringToUUID returns a new uuid
func StringToUUID(uuidStr string) uuid.UUID {
	val, err := uuid.Parse(uuidStr)
	if err != nil {
		logrus.Error(err)
		logrus.Info("creating new uuid...")
		return uuid.New()
	}
	return val
}
