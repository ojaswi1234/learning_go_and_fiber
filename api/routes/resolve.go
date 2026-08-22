package routes


import(
	"github.com/ojaswi1234/learning_go_and_fiber/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/v2"
)


func ResolveUrl(c *fiber.Ctx) error{
	url := c.Params("url")

	r := database.CreateClient(0)

	defer r.Close()



	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "short url not found in the database",
		})
	} else if err != nil {
		return s.Status(fiber.StatusInternalError).JSON(fiber.Map{
			"error": "cannot connect to database",
		})
	}


	rnr := database.CreateClient(1)
	defer rnr.Close()

	_ = rnr.Incr(database.ctx, "counter")


	return c.Ridirect(value, 301)
}
