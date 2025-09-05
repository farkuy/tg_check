package storages

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

func validateRequestPost(req requestStorage, log *slog.Logger, w http.ResponseWriter, r *http.Request) error {
	validate := validator.New()
	err := validate.Struct(req)

	if err != nil {
		errors := err.(validator.ValidationErrors)
		log.Info("Переданны не все данные", errors)
		return fmt.Errorf("Перенданы не все необходимые поля")
	}

	if req.Sum < req.Accumulated {
		log.Error("Сумма не может быть меньше накоплений",
			slog.Int("sum", req.Sum),
			slog.Int("accumulated", req.Accumulated),
		)
		return fmt.Errorf("Сумма не может быть меньше накоплений")

	}

	if req.DeadLineDate.Before(time.Now()) {
		log.Error("Пользователь задал дату меньше текущей", slog.Time("deadLineDate", req.DeadLineDate))
		return fmt.Errorf("Дата не может быть меньше текущей")
	}

	return nil
}
