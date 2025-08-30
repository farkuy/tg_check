package targets

import (
	"log/slog"
	"net/http"
	"tg_check/internal/httpModel"
	"time"

	"github.com/go-chi/render"
)

type requestTarget struct {
	Sum          int       `json:"sum"`
	Accumulated  int       `json:"accumulated"`
	DeadLineDate time.Time `json:"deadLineDate"`
}

type responseTarget struct {
	Id int `json:"id"`
	httpModel.Response
}

type targetCreacte interface {
	post(sum, accumulated int, deadLineDate time.Time) (*Target, error)
}

func PostTarget(tCreate targetCreacte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := slog.With("service", "target", "create")

		var req requestTarget
		err := render.DecodeJSON(r.Body, &req)

		if err != nil {
			log.Error("Ошибка получения данных из json", err.Error())
			render.JSON(w, r, httpModel.Error("ошибка при проверки данных"))
			return
		}

		if req.Sum >= req.Accumulated {
			log.Error("Сумма не может быть меньше накоплений",
				slog.Int("sum", req.Sum),
				slog.Int("accumulated", req.Accumulated),
			)
			render.JSON(w, r, httpModel.Error("Сумма не может быть меньше накоплений"))
			return
		}

		if req.DeadLineDate.Before(time.Now()) {
			log.Error("Пользователь задал дату меньше текущей", slog.Time("deadLineDate", req.DeadLineDate))
			render.JSON(w, r, httpModel.Error("Дата не может быть меньше текущей"))
			return
		}

		res, err := tCreate.post(req.Sum, req.Accumulated, req.DeadLineDate)
		if err != nil {
			log.Error("Ошибка создания targets", err.Error())
			render.JSON(w, r, httpModel.Error("Что-то пошло не так при создании цели"))
			return
		}

		if res != nil {
			log.Error("Ошибка получение данных targets", err.Error())
			render.JSON(w, r, httpModel.Error("Что-то пошло не так при создании цели"))
			return
		}

		render.JSON(w, r, responseTarget{
			Id:       res.Id,
			Response: httpModel.OK(),
		})
	}
}
